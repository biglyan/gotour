package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
	        Title string
	    }{
	        Title: "golang html template demo",
	    }

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}