package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(){

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput(text)

		if len(cleaned_text) == 0{
			continue
		}

		commandName := cleaned_text[0];		
		
		availableCommands := getCommands();

		command , ok := availableCommands[commandName];

		if !ok{
			fmt.Println("Invalid response");
			continue
		}


		command.callback()
		
	}

	

}

type cliCommand struct {
	name string
	description string
	callback func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" : {
			name: "help",
			description: "look for all available options -",
			callback: callbackHelp,
		},
		"exit" : {
			name: "exit",
			description: "Exit from application",
			callback: callbackExit ,
		},
	}
}


func cleanInput(str string ) []string {

	loweredString := strings.ToLower(str);
	convertedToArray := strings.Fields(loweredString);

	return convertedToArray

}

