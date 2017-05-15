package main

import (
	"controller"
	"fmt"
	"net/http"
)

func main() {
	routeHandler()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动失败: ", err.Error())
	}
}

func routeHandler() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/Index/Search", controller.Search)
}
