package routes

import (
	"fmt"
	"net/http"
)

func Home(req *http.Request, res http.ResponseWriter) {
	fmt.Fprintf(res, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
