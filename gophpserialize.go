package gophpserialize

import (
	"encoding/json"
	"strconv"
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

func (s *Serializer) readType() string {
	result := string(s.raw[s.pos])
	s.move()
	return result
}

func (s *Serializer) readBool() bool {
	result := string(s.raw[s.pos])
	s.move()
	if result == "0" {
		return false
	}
	return true
}

func (s *Serializer) readInt() int {
	result := string(s.raw[s.pos])
	for string(s.raw[s.pos+1]) != ":" && string(s.raw[s.pos+1]) != ";" {
		s.move()
		result = result + string(s.raw[s.pos])
	}
	i, _ := strconv.ParseInt(result, 10, 32)
	s.move()
	return int(i)
}

func (s *Serializer) readFloat() float64 {
	result := string(s.raw[s.pos])
	for string(s.raw[s.pos+1]) != ":" && string(s.raw[s.pos+1]) != ";" {
		s.move()
		result = result + string(s.raw[s.pos])
	}
	d, _ := strconv.ParseFloat(result, 64)
	s.move()
	return d
}

func (s *Serializer) readString(size int) string {
	s.move()
	result := string([]rune(string(s.raw[s.pos : s.pos+size])))
	for i := 0; i <= size; i++ {
		s.move()
	}
	return result
}

func (s *Serializer) readValue() interface{} {
	objType := s.readType()
	if objType == "N" {
		s.move()
		return nil
	}

	if objType == "i" {
		s.move()
		val := s.readInt()
		s.move()
		return val
	}
	if objType == "d" {
		s.move()
		val := s.readFloat()
		s.move()
		return val
	}

	if objType == "b" {
		s.move()
		val := s.readBool()
		s.move()
		return val
	}

	if objType == "s" {
		s.move()
		size := s.readInt()
		s.move()
		val := s.readString(size)
		s.move()
		return val
	}

	if objType == "a" {
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
				if hasStringKey {
					r[strconv.Itoa(v2)] = val
				} else {
					l = append(l, val)
				}
			}
		}

		// array close }
		s.move()
		if len(r) == 0 {
			return l
		}
		return r
	}
	println(string(s.raw))
	panic("Unknown objType: " + objType)
}

func (s *Serializer) move() {
	s.pos += 1
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
