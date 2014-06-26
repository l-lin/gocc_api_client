package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"github.com/go-martini/martini"
)

const API_URL string = "http://occapi.herokuapp.com/api"
const BASE_TPL string = "public/base.html"
const ROOT_TPL string = "public/root.html"

type Context struct {
	Content string
}

func Root(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(API_URL)
	if err != nil {
		log.Print("Error fetching content of "+API_URL, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading content of "+API_URL, err)
	}
	context := Context{Content: string(body)}
	tmpl_list := []string{BASE_TPL, ROOT_TPL}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("Template error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
        log.Print("template executing error: ", err)
    }
}

func main() {
	m := martini.Classic()
	m.Get("/", Root)
	m.Get("/hello/:name", func(param martini.Params) string {
			return "Hello " + param["name"]
		})

	m.Run()
}
