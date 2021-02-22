package webserver

import (
	"net/http"
	"fmt"
)

func RunWebServer() {
	fmt.Println("Started webserver at port 8080")

	static := http.FileServer(http.Dir("./templates"))
	http.Handle("/",static)



	http.ListenAndServe(":8080", nil)
}