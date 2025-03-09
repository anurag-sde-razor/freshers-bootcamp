package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Eroor occured", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll((res.Body))

	if err != nil {
		fmt.Println("Error occured", err)
		return
	}

	fmt.Println("Response",string(body))

}
