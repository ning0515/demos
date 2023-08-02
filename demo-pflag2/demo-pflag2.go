package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

func main() {
	//pflag提供了默认的CommandLine作为FlagSet
	//如果使用pflag.NewFlagSet,会新建一个类似CommandLine的FlagSet
	//这里的test会出现在Usage里面
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)
	var ip = flagset.IntP("ip", "i", 1234, "help message for ip")
	flagset2 := pflag.NewFlagSet("test2", pflag.ExitOnError)
	num := flagset2.Int("num", 1, "111")
	//AddFlagSet函数会把两个FlagSet合成一个
	flagset.AddFlagSet(flagset2)
	pflag.String("global", "", "")
	global := pflag.CommandLine
	f := global.Lookup("global")
	flagset.AddFlag(f)
	flagset.Parse(os.Args[1:])
	fmt.Println(*num)
	fmt.Println(*ip)
	//var ip = flagset.IntP("ip", "i", 1234, "help message for ip")
	//
	//var boolVar bool
	//flagset.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")
	//
	//var h host
	//flagset.VarP(&h, "host", "H", "help message for host")
	//
	//flagset.SortFlags = false
	//
	//flagset.Parse(os.Args[1:])
	//
	//fmt.Printf("ip: %d\n", *ip)
	//fmt.Printf("boolVar: %t\n", boolVar)
	//fmt.Printf("host: %+v\n", h)
	//
	//i, err := flagset.GetInt("ip")
	//fmt.Printf("i: %d, err: %v\n", i, err)
}
