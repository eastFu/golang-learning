package main

import "fmt"

func test1(){
	/*
	所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
	你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。
	内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。因为每台机器可能有不同的存储器布局，并且位置分配也可能不同。
	更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
	 */
	var cc = 456
	dd := &cc
	fmt.Println(cc,dd)
	cc = 789
	fmt.Println(cc,dd)
}




func main(){
	test1()
}