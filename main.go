package main

import (
	"fmt"

	"github.com/dockerintro/util"
)

func main() {
	//soualem
	//Eddine
	//Eddine
	//do changes
	/* var a int
	fmt.Print("enter a number :")
	fmt.Scanf("%d", &a)
	fmt.Println("the variable is :", a) */

	conf, err := util.ConfigurationFile()

	if err != nil {
		fmt.Println("the err ", err)
	}

	fmt.Println("our config file as Name : ", conf.Name)
}
