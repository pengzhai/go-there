package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type MyApp struct {
	Name string `inject`
	Age  int    `inject`
}

func main() {
	app := MyApp{}
	in := inject.New()
	in.Map("zhengk")
	in.Map(25)
	in.Apply(&app)
	fmt.Println(app.Name)
	fmt.Println(app.Age)
}