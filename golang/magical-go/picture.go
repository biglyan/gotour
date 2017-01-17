package main

import (
	"io"
	"log"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, h, _ := r.FormFile("image")
		filename := h.Filename
		defer f.Close()

		
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}