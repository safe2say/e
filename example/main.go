package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", withError(handlerExample))
	fmt.Println("serving on :8080")
	fmt.Println("http://localhost:8080?username=nii236\nhttp://localhost:8080?username=nii237")
	http.ListenAndServe(":8080", m)
}
