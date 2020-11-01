package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string  `json:"name"`
	Parent *Person `json:"-"`
	Child  *Person `json:"child"`
}

func main() {
	m := Person{}
	n := Person{}
	m.Name = "Momo Asakura"
	m.Child = &n
	n.Name = "Shiina Natsukawa"
	n.Parent = &m

	j, err := json.Marshal([]Person{m, n})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
