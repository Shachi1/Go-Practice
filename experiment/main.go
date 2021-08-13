package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Int int
	Float float32
	Slice []string
	Map map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := User{
		Name: "Mehjabeen Shachi",
		Int: 123,
		Float: 100.3,
		Slice :[]string{"abc","def"},
		Map: map[string]string{
			"keyl": "valuel",
			"key2": "value2",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}