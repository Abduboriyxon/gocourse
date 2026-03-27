package basic

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[keyType]valueType
	// mapvariable = make(map[keyType]valueType)
	
	// using map literal
	// mapvariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap) 
	myMap["key1"] = 7
	myMap["key2"] = 9
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("key1 is present in the map")
	} else {
		fmt.Println("key1 is not present in the map")
	}
	//fmt.Println(value)
	fmt.Println("is a value associated with key1",ok)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}


	if maps.Equal(myMap3, myMap2) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

	for k, v := range myMap3{
		fmt.Println(k, v)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println("map is not nil")
	}

	val := myMap4["nonExistingKey"]
	fmt.Println("value:", val)

	myMap4 = make(map[string]string)
	myMap4["name"] = "Alice"
	fmt.Println("myMap4:", myMap4)

	fmt.Println("Length of myMap4:", len(myMap))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println("myMap5:", myMap5)

}