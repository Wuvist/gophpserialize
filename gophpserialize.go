package gophpserialize

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

type Serializer struct {
	raw []byte
	pos int
}

func (s *Serializer) SetRaw(msg []byte) {
	s.raw = msg
}

func (s *Serializer) read() map[string]interface{} {
	r := s.readValue().(map[string]interface{})
	return r
}

func (s *Serializer) readType() byte {
	result := s.raw[s.pos]
	s.move()
	return result
}

func (s *Serializer) readBool() bool {
	result := s.raw[s.pos]
	s.move()
	if result == '0' {
		return false
	}
	return true
}

func (s *Serializer) readInt() int {
	start := s.pos
	end := start + 1

	c := s.raw[end]
	for c != ':' && c != ';' {
		end = end + 1
		c = s.raw[end]
	}
	s.pos = end
	i, _ := strconv.ParseInt(string(s.raw[start:end]), 10, 32)
	return int(i)
}

func (s *Serializer) readFloat() float64 {
	start := s.pos
	end := start + 1

	c := s.raw[end]
	for c != ':' && c != ';' {
		end = end + 1
		c = s.raw[end]
	}
	s.pos = end
	d, _ := strconv.ParseFloat(string(s.raw[start:end]), 64)
	return d
}

func (s *Serializer) readString(size int) string {
	s.move()
	result := string(s.raw[s.pos : s.pos+size])
	s.pos += size + 1
	return result
}

func (s *Serializer) readValue() interface{} {
	objType := s.readType()
	if objType == 'N' {
		s.move()
		return nil
	}

	if objType == 'i' {
		s.move()
		val := s.readInt()
		s.move()
		return val
	}
	if objType == 'd' {
		s.move()
		val := s.readFloat()
		s.move()
		return val
	}

	if objType == 'b' {
		s.move()
		val := s.readBool()
		s.move()
		return val
	}

	if objType == 's' {
		s.move()
		size := s.readInt()
		s.move()
		val := s.readString(size)
		s.move()
		return val
	}

	if objType == 'a' {
		s.move()
		size := s.readInt()
		s.move()

		// array open {
		s.move()

		r := make(map[string]interface{})
		l := make([]interface{}, 0)

		//hack to handle array that has both string/int as key
		//convert int key to string key
		hasStringKey := false

		for i := 0; i < size; i++ {
			key := s.readValue()
			val := s.readValue()
			switch v2 := key.(type) {
			case string:
				hasStringKey = true
				r[v2] = val
			case int:
				if hasStringKey || v2 != i {
					r[strconv.Itoa(v2)] = val
				} else {
					l = append(l, val)
				}
			}
		}

		if len(r) > 0 && len(l) > 0 {
			for i, val := range l {
				r[strconv.Itoa(i)] = val
			}
		}

		// array close }
		s.move()
		if len(r) == 0 {
			return l
		}
		return r
	}
	panic("Unknown objType: " + string(objType) + "\n" + strconv.Itoa(s.pos) + "\n\n" + string(s.raw))
}

func (s *Serializer) slice(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	objType := s.readType()

	if objType != 'a' {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	s.move()
	size := s.readInt()
	s.move()

	// array open {
	s.move()

	l := reflect.MakeSlice(t, size, size)

	for i := 0; i < size; i++ {
		s.readValue() //skip the key
		elem := l.Index(i)
		elem.Set(reflect.New(elem.Type().Elem()))
		s.value(elem)
	}

	// array close }
	s.move()

	v.Set(l)
	return nil
}

func (s *Serializer) obj(v reflect.Value) error {
	e := v.Elem()
	t := e.Type()

	fields := make(map[string]reflect.Value)
	for i := 0; i < e.NumField(); i++ {
		thriftInfo := t.Field(i).Tag.Get("thrift")
		pos := strings.Index(thriftInfo, ",")
		field := e.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		fields[thriftInfo[0:pos]] = field
	}

	objType := s.readType()

	if objType != 'a' {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	s.move()
	size := s.readInt()
	s.move()

	// array open {
	s.move()

	for i := 0; i < size; i++ {
		key := s.readValue()

		field, found := fields[key.(string)]
		if found == false {
			s.readValue()
			continue
		}
		s.value(field)
	}

	// array close }
	s.move()

	return nil
}

func (s *Serializer) dict(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	objType := s.readType()

	if objType != 'a' {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	s.move()
	size := s.readInt()
	s.move()

	// array open {
	s.move()

	m := reflect.MakeMap(t)

	elem := m.Type().Elem()

	for i := 0; i < size; i++ {
		key := s.readValue().(string)
		val := reflect.New(elem).Elem()
		s.value(val)
		m.SetMapIndex(reflect.ValueOf(key), val)
	}

	// array close }
	s.move()

	v.Set(m)
	return nil
}

func (s *Serializer) value(v reflect.Value) error {
	kind := v.Kind()

	switch kind {
	case reflect.Ptr:
		v2 := v.Elem()
		kind = v2.Kind()
		switch kind {
		case reflect.String:
			v2.SetString(s.readValue().(string))
			return nil
		case reflect.Int64:
			v2.SetInt(s.readValue().(int64))
			return nil
		case reflect.Map:
			return s.dict(v)
		case reflect.Slice:
			return s.slice(v)
		case reflect.Struct:
			return s.obj(v)
		}
	case reflect.Map:
		return s.dict(v)
	case reflect.String:
		v.SetString(s.readValue().(string))
		return nil
	case reflect.Float64:
		val := s.readValue().(float64)
		v.SetFloat(val)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := s.readValue().(int)
		v.SetInt(int64(val))
		return nil
	}

	return &InvalidUnmarshalError{reflect.TypeOf(v)}
}

func (s *Serializer) move() {
	s.pos += 1
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

func UnmarshalThrift(data []byte, v interface{}) error {
	s := new(Serializer)
	s.SetRaw(data)
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	return s.value(rv)
}

func Unmarshal(data []byte) interface{} {
	s := new(Serializer)
	s.SetRaw(data)
	return s.readValue()
}

func PhpToJson(phpData []byte) (jsonData []byte, err error) {
	r := Unmarshal(phpData)
	jsonData, err = json.Marshal(r)
	return
}
