package main

import (
	// "fmt"
	"net/http"

	"github.com/RivierGrullon/GoWebService/controllers"
	// "github.com/RivierGrullon/GoWebService/models"
)

func main() {
	controllers.RegisterControllers()
	// panic("uy quieto")
	
	http.ListenAndServe(":8000", nil)
}
