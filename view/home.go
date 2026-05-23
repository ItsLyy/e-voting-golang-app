package view

import (
	"fmt"
)

func Home() {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Println("Please choose your menu!")
	fmt.Print(normalLine())
	fmt.Println("1. Candidate Management")
	fmt.Println("2. Voter Management")
	fmt.Println("3. Voting")
	fmt.Println("4. Election Management")
	fmt.Println()
	fmt.Println("q. Quit App")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
	case "2":
		manageVoter()
	case "3":
	case "4":
		manageElection()
	case "q":
		exit()
	default:
		wrong(choice)
		Home()
	}
}

func normalLine() string {
	return "-----------------------------\n=============================\n"
}

func tableLine() string {
	return "--------------------------------------------------------------------------------------------------------------------\n====================================================================================================================\n"
}

func wrong(choice string) {
	fmt.Print(normalLine())
	fmt.Println("Wrong Choice there is no", choice)
	fmt.Print(normalLine())
}

func exit() {
	fmt.Print(normalLine())
	fmt.Println("GOOD BYE THANKS!")
	fmt.Print(normalLine())
}

func inputChoice(choice *string) {
	fmt.Print("Choice: ")
	fmt.Scan(choice)
}

func confirmation(choice *string) {
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(choice)
}
