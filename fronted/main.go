package main

import (
	. "config"
	"controller"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(View)))
	http.Handle("/search", controller.CreatSearchHandler(TemplateFilePath))
	err := http.ListenAndServe(fmt.Sprintf(":%d", Fronted), nil)
	if err != nil {
		panic(err)
	}
}
