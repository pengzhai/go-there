package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"log"
	"reflect"
)

func do(i int) int {
	fmt.Println(i)
	return 23
}

func main() {
	in := inject.New()
	in.Map(12)
	val, err := in.Invoke(do)
	if err != nil {
		log.Fatal("调用出错:" + err.Error())
		return
	}
	for _, v := range val {
		if v.Kind() == reflect.Int {
			fmt.Println(v.Int())
		}
	}
}
