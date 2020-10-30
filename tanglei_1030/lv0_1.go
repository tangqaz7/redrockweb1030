package main

import "fmt"
type Author struct {
Name string             //名字
VIP bool                //是否是高贵的带会员
Icon string             //头像
Signature string        //签名
Focus int               //关注人数
}

func main()  {
	var dagongren Author
	dagongren.VIP = false
	dagongren.Name = "三Lu有毒"
	dagongren.Icon = "thunder"
	dagongren.Signature = "哇塞，是谁在看我。群：829581926~"
	dagongren.Focus = 145000
	fmt.Println(dagongren)
}