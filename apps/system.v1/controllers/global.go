package controllers

import (
	"fmt"
	"net/http"
	//"../../../render"
	//"../../../system"
)

var guess string

func TMP(w http.ResponseWriter, req *http.Request) {
	guess = "SET"
}

func TMP2(w http.ResponseWriter, req *http.Request) {
	fmt.Println(guess)
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HERE")
}
