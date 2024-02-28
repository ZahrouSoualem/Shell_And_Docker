package main

import (
	"fmt"

	"github.com/dockerintro/util"
)

func main() {
	//soualem
	//Eddine
	//Eddine
	//changes in haled branch
	//squash changes 2
	//changes in haled branch
	//squash changes 2
	//changes in khaled branch
	//we are doing changes in main branch
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
