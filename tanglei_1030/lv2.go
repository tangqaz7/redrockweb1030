package main

import (
	"fmt"
)

func main()  {
	receiver(123)
	receiver("ikdfhdis")
	receiver(true)
}
func receiver(v interface{})  {
	var str string
	fmt.Println(v)
	switch v.(type) {
	case string:
		str="这个是string";
	case int:
		str="这个是int";
	case bool:
		str="这个是bool";
	}
	fmt.Println(str)

}