package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/api", rest)

	http.HandleFunc("/marco", hello)

	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {

	name := req.URL.Query().Get("name")
	value2 := req.URL.Query().Get("address")
	value3 := req.URL.Query().Get("id")
	age := req.URL.Query().Get("age")
	f_name := req.URL.Query().Get("f_name")
	if age == "" {
		fmt.Fprint(w, "please pass name age  address id query string")
		return
	}
	if name == "" {
		fmt.Fprint(w, "please pass name")
		return
	}
	//fmt.Fprintf(w, "hello %s, age %s,address %s,id %s\n", name, age, value2, value3)

	f, err := os.Create(f_name)
	if err != nil {
		log.Fatalln(err)
	}

	f.WriteString(name + value2 + value3 + age)
}

func rest(w http.ResponseWriter, req *http.Request) {

	p := Person{Name: "nawaz", Age: 47}

	data, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)

}
