package gophpserialize

import "testing"

func TestUnmarshal(t *testing.T) {
	data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`

	obj := Unmarshal([]byte(data)).(map[string]interface{})

	if obj["apple"] != 1 {
		t.Error("Unmarshal failed")
	}
	if obj["orange"] != 2 {
		t.Error("Unmarshal failed")
	}
	if obj["grape"] != 3 {
		t.Error("Unmarshal failed")
	}
}

func TestPhpToJson(t *testing.T) {
	data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`

	obj, err := PhpToJson([]byte(data))

	jsonStr := `{"apple":1,"grape":3,"orange":2}`

	if err != nil {
		t.Error(err)
	}

	if string(obj) != jsonStr {
		t.Error("convert to json error")
	}
}

func TestPhpToJson2(t *testing.T) {
	data := `a:2:{i:2;a:4:{s:11:"id_language";s:1:"2";s:9:"iso2_code";s:2:"zh";s:4:"name";s:7:"Chinese";s:10:"is_default";s:1:"0";}i:1;a:4:{s:11:"id_language";s:1:"1";s:9:"iso2_code";s:2:"en";s:4:"name";s:7:"English";s:10:"is_default";s:1:"1";}}`

	obj, _ := PhpToJson([]byte(data))

	jsonStr := `{"0":{"id_language":"1","is_default":"1","iso2_code":"en","name":"English"},"2":{"id_language":"2","is_default":"0","iso2_code":"zh","name":"Chinese"}}`

	if string(obj) != jsonStr {
		t.Error("convert to json error")
	}
}

func TestPhpToJsonComplex(t *testing.T) {
	data := `a:9:{s:4:"meta";a:38:{s:3:"sku";s:14:"JO945AC20SULMY";s:17:"id_catalog_config";s:6:"100579";s:16:"attribute_set_id";s:2:"10";s:4:"name";s:21:"(fit Over) Sunglass S";s:12:"activated_at";s:19:"2012-09-10 23:49:18";s:12:"caching_hash";s:32:"acc3e0404646c57502b480dc052c4fe1";s:11:"fk_vertical";s:1:"2";s:10:"categories";s:17:"1|191|257|346|726";s:5:"brand";s:14:"JO JO NEW YORK";s:15:"sub_cat_type_id";s:2:"26";s:9:"gender_id";s:1:"1";s:6:"gender";s:4:"Male";s:15:"gender_position";s:1:"0";s:6:"season";s:11:"Fall-Winter";s:15:"season_position";s:1:"6";s:16:"color_name_brand";s:5:"Brown";s:12:"color_family";s:5:"brown";s:17:"short_description";s:345:"The JO JO NEW YORK (Fit over) Sunglass S features lightweight frame, polarized lens to eliminate harsh glare, side lens for improved peripheral vision, UV 400 protection block 100% of harmful UV rays. It is designed to be worn over your prescription glasses RX Eyewear to provide comfort from glare. Alternatively, it can also be worn by itself.";s:8:"ponumber";s:5:"M1050";s:14:"sunglass_width";s:7:"0.00000";s:19:"sunglass_arm_length";s:7:"0.00000";s:12:"strap_length";s:3:"255";s:13:"sourcing_type";s:13:"branded_local";s:22:"sourcing_type_position";s:1:"3";s:9:"live_date";s:10:"2012-12-19";s:12:"nonsale_item";s:1:"0";s:10:"brand_type";s:7:"branded";s:19:"brand_type_position";s:1:"1";s:10:"show_chart";s:1:"0";s:9:"config_id";s:6:"100579";s:9:"max_price";s:6:"139.00";s:5:"price";s:6:"139.00";s:18:"max_original_price";N;s:14:"original_price";N;s:17:"max_special_price";s:5:"97.30";s:13:"special_price";s:5:"97.30";s:21:"max_saving_percentage";s:2:"30";s:13:"original_name";s:21:"(fit Over) Sunglass S";}s:10:"attributes";a:1:{s:5:"color";s:5:"Brown";}s:17:"isMerchantProduct";b:0;s:7:"simples";a:1:{s:21:"JO945AC20SULMY-277946";a:2:{s:4:"meta";a:12:{s:3:"sku";s:21:"JO945AC20SULMY-277946";s:5:"price";s:6:"139.00";s:12:"caching_hash";s:32:"366137f1627c907aaae0160f5ce4db62";s:18:"shipment_cost_item";s:4:"0.00";s:19:"shipment_cost_order";s:4:"0.00";s:13:"special_price";s:5:"97.30";s:11:"tax_percent";s:4:"0.00";s:8:"quantity";s:1:"1";s:4:"size";s:8:"One Size";s:13:"size_position";s:2:"15";s:18:"estimated_delivery";s:42:"3 business days (4 days for East Malaysia)";s:27:"estimated_delivery_position";s:1:"4";}s:10:"attributes";a:0:{}}}s:6:"images";a:3:{i:0;a:16:{s:24:"id_catalog_product_image";s:6:"297248";s:17:"fk_catalog_config";i:100579;s:17:"fk_catalog_simple";N;s:5:"image";s:1:"1";s:4:"main";s:1:"1";s:10:"updated_at";s:19:"2012-09-05 10:36:58";s:13:"updated_at_ts";s:4:"2618";s:3:"sku";s:14:"JO945AC20SULMY";s:16:"fk_catalog_brand";s:3:"945";s:10:"sku_simple";N;s:16:"id_catalog_brand";s:3:"945";s:10:"brand_name";s:14:"JO JO NEW YORK";s:13:"brand_url_key";s:14:"jo-jo-new-york";s:3:"url";s:60:"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-1";s:4:"path";s:47:"%catalog.image_directory%/product/97/5001/1.jpg";s:6:"sprite";s:69:"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-sprite.jpg";}i:1;a:15:{s:24:"id_catalog_product_image";s:6:"297249";s:17:"fk_catalog_config";d:100579;s:17:"fk_catalog_simple";N;s:5:"image";s:1:"2";s:4:"main";s:1:"0";s:10:"updated_at";s:19:"2012-09-05 10:37:00";s:13:"updated_at_ts";s:4:"2620";s:3:"sku";s:14:"JO945AC20SULMY";s:16:"fk_catalog_brand";s:3:"945";s:10:"sku_simple";N;s:16:"id_catalog_brand";s:3:"945";s:10:"brand_name";s:14:"JO JO NEW YORK";s:13:"brand_url_key";s:14:"jo-jo-new-york";s:3:"url";s:60:"http://static03-my.zalora.com/p/jo-jo-new-york-2620-975001-2";s:4:"path";s:47:"%catalog.image_directory%/product/97/5001/2.jpg";}i:2;a:15:{s:24:"id_catalog_product_image";s:6:"297250";s:17:"fk_catalog_config";i:100579;s:17:"fk_catalog_simple";N;s:5:"image";s:1:"3";s:4:"main";s:1:"0";s:10:"updated_at";s:19:"2012-09-05 10:37:01";s:13:"updated_at_ts";s:4:"2621";s:3:"sku";s:14:"JO945AC20SULMY";s:16:"fk_catalog_brand";s:3:"945";s:10:"sku_simple";N;s:16:"id_catalog_brand";s:3:"945";s:10:"brand_name";s:14:"JO JO NEW YORK";s:13:"brand_url_key";s:14:"jo-jo-new-york";s:3:"url";s:60:"http://static03-my.zalora.com/p/jo-jo-new-york-2621-975001-3";s:4:"path";s:47:"%catalog.image_directory%/product/97/5001/3.jpg";}}s:14:"rating_classes";a:0:{}s:5:"image";s:60:"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-1";s:6:"sprite";s:69:"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-sprite.jpg";s:4:"link";s:37:"%28fit-Over%29-Sunglass-S-100579.html";}`

	obj, err := PhpToJson([]byte(data))

	jsonStr := `{"attributes":{"color":"Brown"},"image":"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-1","images":[{"brand_name":"JO JO NEW YORK","brand_url_key":"jo-jo-new-york","fk_catalog_brand":"945","fk_catalog_config":100579,"fk_catalog_simple":null,"id_catalog_brand":"945","id_catalog_product_image":"297248","image":"1","main":"1","path":"%catalog.image_directory%/product/97/5001/1.jpg","sku":"JO945AC20SULMY","sku_simple":null,"sprite":"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-sprite.jpg","updated_at":"2012-09-05 10:36:58","updated_at_ts":"2618","url":"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-1"},{"brand_name":"JO JO NEW YORK","brand_url_key":"jo-jo-new-york","fk_catalog_brand":"945","fk_catalog_config":100579,"fk_catalog_simple":null,"id_catalog_brand":"945","id_catalog_product_image":"297249","image":"2","main":"0","path":"%catalog.image_directory%/product/97/5001/2.jpg","sku":"JO945AC20SULMY","sku_simple":null,"updated_at":"2012-09-05 10:37:00","updated_at_ts":"2620","url":"http://static03-my.zalora.com/p/jo-jo-new-york-2620-975001-2"},{"brand_name":"JO JO NEW YORK","brand_url_key":"jo-jo-new-york","fk_catalog_brand":"945","fk_catalog_config":100579,"fk_catalog_simple":null,"id_catalog_brand":"945","id_catalog_product_image":"297250","image":"3","main":"0","path":"%catalog.image_directory%/product/97/5001/3.jpg","sku":"JO945AC20SULMY","sku_simple":null,"updated_at":"2012-09-05 10:37:01","updated_at_ts":"2621","url":"http://static03-my.zalora.com/p/jo-jo-new-york-2621-975001-3"}],"isMerchantProduct":false,"link":"%28fit-Over%29-Sunglass-S-100579.html","meta":{"activated_at":"2012-09-10 23:49:18","attribute_set_id":"10","brand":"JO JO NEW YORK","brand_type":"branded","brand_type_position":"1","caching_hash":"acc3e0404646c57502b480dc052c4fe1","categories":"1|191|257|346|726","color_family":"brown","color_name_brand":"Brown","config_id":"100579","fk_vertical":"2","gender":"Male","gender_id":"1","gender_position":"0","id_catalog_config":"100579","live_date":"2012-12-19","max_original_price":null,"max_price":"139.00","max_saving_percentage":"30","max_special_price":"97.30","name":"(fit Over) Sunglass S","nonsale_item":"0","original_name":"(fit Over) Sunglass S","original_price":null,"ponumber":"M1050","price":"139.00","season":"Fall-Winter","season_position":"6","short_description":"The JO JO NEW YORK (Fit over) Sunglass S features lightweight frame, polarized lens to eliminate harsh glare, side lens for improved peripheral vision, UV 400 protection block 100% of harmful UV rays. It is designed to be worn over your prescription glasses RX Eyewear to provide comfort from glare. Alternatively, it can also be worn by itself.","show_chart":"0","sku":"JO945AC20SULMY","sourcing_type":"branded_local","sourcing_type_position":"3","special_price":"97.30","strap_length":"255","sub_cat_type_id":"26","sunglass_arm_length":"0.00000","sunglass_width":"0.00000"},"rating_classes":[],"simples":{"JO945AC20SULMY-277946":{"attributes":[],"meta":{"caching_hash":"366137f1627c907aaae0160f5ce4db62","estimated_delivery":"3 business days (4 days for East Malaysia)","estimated_delivery_position":"4","price":"139.00","quantity":"1","shipment_cost_item":"0.00","shipment_cost_order":"0.00","size":"One Size","size_position":"15","sku":"JO945AC20SULMY-277946","special_price":"97.30","tax_percent":"0.00"}}},"sprite":"http://static03-my.zalora.com/p/jo-jo-new-york-2618-975001-sprite.jpg"}`

	if err != nil {
		t.Error(err)
	}

	if string(obj) != jsonStr {
		t.Error("convert to json complex error")
		println(string(obj))
	}
}
