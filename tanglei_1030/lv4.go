package main

import (
	"fmt"
	"math"
	"strings"
)
type Painter interface {
}
type Heart struct {

}
type Line struct {

}
func Paint(painter Painter)  {
	switch painter.(type) {
	case Heart:
		hev();
	case Line:
		line()

	}
}
func main()  {
	var heart Heart
	Paint(heart)

	var li Line
	Paint(li)


}
func hev()  {
	var MYWORD string
	MYWORD = "*"
	var sep string
	sep =" "
	var zoom float64
	zoom = 1.0
	chars := strings.Split(MYWORD, sep)
	for _, char := range chars{
		allChar := make([]string, 0)

		for y := 12 * zoom; y > -12*zoom; y-- {
			lst := make([]string, 0)
			lstCon := ""
			for x := -30 * zoom; x < 30*zoom; x++ {
				x2 := float64(x)
				y2 := float64(y)
				formula := math.Pow(math.Pow(x2*0.04/zoom, 2)+math.Pow(y2*0.1/zoom, 2)-1, 3) - math.Pow(x2*0.04/zoom, 2)*math.Pow(y2*0.1/zoom, 3)
				if formula <= 0 {
					index := int(x) % len(char)
					if index >= 0 {
						lstCon += string(char[index])
					} else {
						lstCon += string(char[int(float64(len(char))-math.Abs(float64(index)))])
					}

				} else {
					lstCon += " "
				}
			}
			lst = append(lst, lstCon)
			allChar = append(allChar, lst...)
		}

		for _, text := range allChar {
			fmt.Printf("%s\n", text)
		}
	}

}
func line(){
	fmt.Println("--------------------------------------------------------------------------")
}