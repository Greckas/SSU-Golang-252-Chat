package main

import (
	"net/http"

	"fmt"

	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/Greckas/SSU-Golang-252-Chat/server/config"
	"github.com/Greckas/SSU-Golang-252-Chat/server/core"
	"github.com/gorilla/mux"
)

func main() {

	port := "3006"
	loger.Log.Infof("Server run with port: " + port)
	r := mux.NewRouter()
	fmt.Println(config.GetConfig())
	r.HandleFunc("/message", core.MessageHandler) // listen message
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		loger.Log.Panicf("Cannot run server %s", err.Error())
	}
}
