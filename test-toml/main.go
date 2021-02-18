package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type File struct {
	Group `toml:"group"`
}

type Group struct {
	Name    string   `toml:"name"`
	Company string   `toml:"company"`
	Member  []Person `toml:"member"`
}

type Person struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

func main() {
	var group File
	_, err := toml.DecodeFile("group.toml", &group)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Enter group: %v(%v)\n", group.Name, group.Company)
	for k, v := range group.Member {
		fmt.Printf("Member %v: %v(%v)\n", k, v.Name, v.Age)
	}
}
