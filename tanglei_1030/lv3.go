package main

import "fmt"
type newau struct {
	dianzan int
	col bool
	coin bool
	three bool
	name string
	author string
}


func (c *newau) dianzannum() int {
	return c.dianzan
}
func (c *newau) colis() bool {
	return c.col
}
func (c *newau) iscoin() bool {
	return c.coin
}
func (c *newau) isthree() bool {
	return c.three
}
func main() {
video("au","name")
}
func video(a string,b string)  {
	c := &newau{
		123,
		true,
		true,
		true,
		a,
		b,
	}
	fmt.Println(c)
	return
}