package main

import "fmt"

func main()  {
	person1 := NewChinese()
	age := person1.AgeChinese
	if age<23 {
		fmt.Println("终于找到你")
	}else {
		fmt.Println("不符合要求")
	}
}

type Person interface {
	Sex() bool
	Age() int
	Height() int
	Weight() int
	Salary() int
}



type Chinese struct {
	Name string
	SexChinese bool
	AgeChinese int
	Hei int
	Wei int
	AiHao bool
	Money int
}

func NewChinese() *Chinese {
	c := &Chinese{
		Name:       "上官燕雀",
		SexChinese: false,
		AgeChinese: 20,
		Hei:        165,
		Wei:        100,
		AiHao:      true,
		Money:      20000,
	}
	return c
}

func (c *Chinese) Sex() bool {
	return c.SexChinese
}
func (c *Chinese) Age() int {
	return c.AgeChinese
}
func (c *Chinese) Height() int {
	return c.Hei
}
func (c *Chinese) Weight() int {
	return c.Wei
}
func (c *Chinese) Salary() int {
	return c.Money
}
func (c *Chinese) Like()  {
	if c.AiHao == true {
		fmt.Println("他也喜欢打游戏")
	}else {
		fmt.Println("他不喜欢打游戏")
	}
}