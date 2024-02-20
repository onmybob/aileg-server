package controllers

import "net/http"

func Echo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
}