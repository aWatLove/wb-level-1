package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/
func main() {
	isType(12)
	isType("12")
	isType(12.0)
}

func isType(value interface{}) {
	// используя switch и type, определяем тип value
	switch v := value.(type) {
	// далее перебираем каждый случай
	case int:
		fmt.Printf("is int type %%T=%T and value = %v\n", v, v)
	case string:
		fmt.Printf("is string type %%T=%T and value = %v\n", v, v)
	case bool:
		fmt.Printf("is bool type %%T=%T and value = %v\n", v, v)
	case chan int:
		fmt.Printf("is chan int type %%T=%T and value = %v\n", v, v)
	default:
		fmt.Printf("idk about this type %%T=%T\n", v)
	}
}
