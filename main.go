package main

import (
	"fmt"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
	Handler "github.com/sridharu346/loginPractice/handlers"
)

func main() {

	DB, err := DBconfig.DBConnection()
	if err != nil {
		fmt.Print("Error Occured")
	}
	defer DB.Close()
	fmt.Println("DB connected Successfully")

	//handling the routes
	http.HandleFunc("/Signup", Handler.SignupHandler)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
