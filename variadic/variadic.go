//Pass different types of arguments in variadic function
package main

import (
	"fmt"
	"reflect"
)

func main() {
	variadicExample(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})

	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	arr := reflect.ValueOf(strArray)
	fmt.Println(arr.Kind())
	fmt.Println(reflect.Array)
}

func variadicExample(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
