package main

import (
	"fmt"
	"net/http"

	"github.com/eqiruz22/go-server-json/handler"
	"github.com/eqiruz22/go-server-json/utils"
)

func init() {
	err := utils.LoadDB()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Database loaded successfully")
	}
}


func main() {
	http.HandleFunc("/posts", handler.PostHandler)
	http.HandleFunc("/posts/", handler.HandleIdWithPath)
	http.ListenAndServe(":8080",nil)
}