package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Msg struct {
	Code    int
	Message string
}

func main() {
	s := &Msg{0, "hello there"}
	j, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1 %v\n", j)
	fmt.Printf("2 %v\n", string(j))

	r := "{\"Code\":1,\"Message\":\"yo bro\"}"
	j = []byte(r)
	err = json.Unmarshal(j, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("3 %v\n", *s)

	r = "{\"Code\":1,\"Msg\":\"yo bro\"}"
	j = []byte(r)
	err = json.Unmarshal(j, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("4 %v\n", *s)

	r = "{\"Msg\":\"yo bro\",\"X\":1}"
	j = []byte(r)
	err = json.Unmarshal(j, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("5 %v\n", *s)

	r = "{\"Code\":\"ACK\",\"Msg\":\"yo bro\"}"
	j = []byte(r)
	err = json.Unmarshal(j, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("6 %v\n", *s)
}
