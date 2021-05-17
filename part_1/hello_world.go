 package main
 import "fmt"
 func main() {
 	fmt.Println("Hello,世界")
 	/*
 	此处要注意单引号和双引号的区分；
 	在go语法中，双引号是常用的来表达字符串；
 	如果使用了单引号，编辑器会提示出错：invalid character literal (more than one character)
 	单引号只能包含一个字符，例如’a’ ,程序会输出97表示字符a的ascii码。

	如果非要使用单引号输出必须使用string函数转换：fmt.Println(string('a')) 

	-- fmt包中的函数用来格式化输出和扫描输入。Println是fmt中一个基本的输出函数，
	它输出一个或多个用空格分隔的值，结尾使用一个换行符
 	 */
 }