package main

import (
	"fmt"
	"unsafe"
)

func testIota1(){
	/*
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	iota 可以被用作枚举值
	iota 表示从 0 开始自动加 1
	 */
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}

func testIota2(){
	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	//1<<0,3<<1,3<<2,3<<3
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}


func main() {
	testIota1()
	testIota2()

	// const 常量定义 , 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型, 不可修改
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH

	fmt.Println(a, b, c)
	fmt.Printf("面积为 : %d", area)

	const (
		UNKNOWN = 0
		FEMAIE  = 1
		MALE    = 2
	)

	const (
		d = "abc"
		e = len(d)
		f= unsafe.Sizeof(d)
	)

	/*
		1. golang的命名需要使用驼峰命名法，且不能出现下划线
		2、golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，
		如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
	  可以简单的理解成，首字母大写是公有的，首字母小写是私有的
	 */

	 //变量
	 var name1 string = "niuniu1"
	 var name2 ="niuniu2"
	 name3:="niuniu3"
	 var name4,name5="niuniu4","niuniu5"

	 fmt.Println(name1,name2,name3,name4,name5)


}
