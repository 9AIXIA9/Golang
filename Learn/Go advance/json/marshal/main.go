package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name   string
	Age    int
	weight int
}

func main() {
	p1 := Person{
		Name:   "构式",
		Age:    10,
		weight: 38,	//小写不可访问，所以他的值为零
	}
	log.Print("p1", p1)

	by, err := json.Marshal(p1)
	if err != nil {
		log.Print("json marshal failed", err)
		return
	}
	log.Print("byte", by)
	log.Printf("%s", by)

	var p2 Person
	err = json.Unmarshal(by, &p2)
	if err != nil {
		log.Print("json unmarshal failed", err)
		return
	}
	log.Print("p2", p2)

}
