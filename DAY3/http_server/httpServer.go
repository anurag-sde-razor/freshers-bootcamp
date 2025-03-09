package main

import (
	"fmt"
	"net/http"
)

func main(){
    http.HandleFunc("/anurag",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Anurag"))
	})
	fmt.Println("server started at port 8080")

    http.ListenAndServe(":8080",nil)

}