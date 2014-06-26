package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
			return "Hello world!"
		})
	m.Get("/hello/:name", func(param martini.Params) string {
			return "Hello " + param["name"]
		})
	m.Run()
}
