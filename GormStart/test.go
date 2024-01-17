package main

import "fmt"

func main() {
	test01 := map[string]interface{}{"Name": "jinzhu_1", "Age": 18}
	fmt.Printf("%#v", test01)
	test02 := []map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	}
	fmt.Println("\r")
	fmt.Println(test02)
	// test03 := map[]interface{}{
	// 	"sdd", "sdsdsd",
	// }
	// fmt.Println(test03)
}
