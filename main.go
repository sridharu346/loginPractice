package main

import (
	"fmt"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
)

func main() {

	DB, err := DBconfig.DBConection()
	if err != nil {
		fmt.Print("Error Occured")
	}
	defer DB.Close()
	fmt.Print("DB connected Successfully")
}
