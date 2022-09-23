package sjson

import (
	"fmt"

	"github.com/tidwall/sjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const json = `{"name":{"first":"li","last":"dj"},"age":18}`

func SjsonTest() {
	value, _ := sjson.Set(json, "name.last", "dajun")
	fmt.Println(value)

	nilJSON, _ := sjson.Set("", "key", nil)
	fmt.Println(nilJSON)

	boolJSON, _ := sjson.Set("", "key", false)
	fmt.Println(boolJSON)

	intJSON, _ := sjson.Set("", "key", 1)
	fmt.Println(intJSON)

	floatJSON, _ := sjson.Set("", "key", 10.5)
	fmt.Println(floatJSON)

	strJSON, _ := sjson.Set("", "key", "hello")
	fmt.Println(strJSON)

	mapJSON, _ := sjson.Set("", "key", map[string]interface{}{"hello": "world"})
	fmt.Println(mapJSON)

	u := User{Name: "dj", Age: 18}
	structJSON, _ := sjson.Set("", "key", u)
	fmt.Println(structJSON)

	fruits := `{"fruits":["apple", "orange", "banana"]}`

	var newValue string
	newValue, _ = sjson.Delete(fruits, "fruits.1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.-1")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.5","")
	fmt.Println(newValue)

	user := `{"name":{"first":"li","last":"dj"},"age":18}`

	newValue, _ = sjson.Delete(user, "name.first")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(user, "name.full")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.-1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.5")
	fmt.Println(newValue)

	userErr := `{"name":dj,age:18}`
	newValue, err := sjson.Set(userErr, "name", "dajun")
	fmt.Println(err, newValue)

	userErrPath := `{"name":"dj","age":18}`
	newErrValue, err := sjson.Set(userErrPath, "na?e", "dajun")
	fmt.Println(err, newErrValue)
}
