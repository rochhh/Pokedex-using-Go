package main

import (
	"fmt"
)

func callbackHelp() error{

	fmt.Println("Welcome to the Pokedex");

	availableCommands := getCommands();

	for _ , values := range availableCommands {
		fmt.Printf("- %s:%s\n" , values.name , values.description)
	}
	return nil
}