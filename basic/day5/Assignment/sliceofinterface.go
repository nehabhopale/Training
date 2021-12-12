package main

import "fmt"

func main() {
	var values []interface{}
	values = append(values, "name")
	values = append(values, "second")
	values = append(values, 12)
	values = append(values, 13)
	values = append(values, false)
	values = append(values, 12.5)

	fmt.Println(values)
	fmt.Println(values[1:3])
	operation(values)
	fmt.Println(values)

}
func operation(slice []interface{}) {
	for i := 0; i < len(slice); i++ {
		switch slice[i].(type) {
		case int:
			slice[i] = slice[i].(int) + 12
		case float64:
			slice[i] = slice[i].(float64) * 0.1
		case bool:
			slice[i] = !slice[i].(bool)
		default:
			fmt.Println("we are good with this operatin ")
		}
	}
}
