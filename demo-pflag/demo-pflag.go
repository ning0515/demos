package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	flag "github.com/spf13/pflag"
)

// 定义命令行参数对应的变量
var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = flag.IntP("age", "a", 22, "Input Your Age")
var cliGender = flag.StringP("gender", "g", "male", "Input Your Gender")
var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

// 这个是标准化函数，目的是将前面设置的所有标志的长名称更改,凡是出现"-"或者"_"的情况,都会
// 被替换成"."流程如下：
// 1.标准化程序员设置的名称:
// 程序员在StringP等函数中第一个参数里设置的长名称会被替换为指定字符
// 执行flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)时，会遍历一遍所有的flag
// 此时已经将des-detail替换成了des.detail,实际上运行程序的时候可以使用des.detail来设置参数
// 2.标准化用户输入的名称,比如用户输入了
func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	fmt.Println("start:" + name)
	fmt.Printf("stack = %s\n", debug.Stack())
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	fmt.Println("stop:" + name)
	return flag.NormalizedName(name)
}

func main() {
	fmt.Println(1)
	fmt.Println(os.Args[1:])
	// 设置标准化参数名称的函数
	fmt.Println(2)
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	//NoOptDefVal的意思是如果使用了 --age或者-a但是没有用等号赋值,
	//那默认值就是NoOptDefVal,没有使用--age或者-a,那默认值就是
	//一开始设置的默认值,增加了更多灵活性
	//Lookup接受一个flag的全名，返回这个flag对应的结构体
	flag.Lookup("age").NoOptDefVal = "25"

	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 gender
	flag.CommandLine.MarkHidden("badflag")
	fmt.Println(3)
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	fmt.Println(4)
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}
