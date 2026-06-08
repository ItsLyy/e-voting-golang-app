package view

import (
	"e-voting/model"
	"fmt"
)

func manageElection() {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Println("MENU: Election Managements")
	fmt.Println("Please choose your menu!")
	fmt.Print(normalLine())
	fmt.Println("1. View Election Results")
	fmt.Println("2. Reset Election")
	fmt.Println()
	fmt.Println("b. Back to Home")
	fmt.Println("q. Quit App")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		viewElectionResults()
	case "2":
		resetElection()
	case "b":
		//Home()
	case "q":
		exit()
	}
}

func displayTableElectionResults() {
	var i int = 0

	fmt.Print(tableLine())
	fmt.Printf("|%-80s|%-20s|%-13s|\n", "Name", "Persentage(%)", "Total Voted")
	fmt.Print(tableLine())

	for i < model.Candidate.Length {
		fmt.Printf("|%-80s|%-20.2f|%-13d|\n", model.Candidate.Data[i].Name, 20.02, 10)
		i++
	}
	fmt.Print(tableLine())
}

func viewElectionResults() {
	var choice string
	displayTableElectionResults()
	fmt.Printf("%-30s%-30s%-30s\n", "1: Sorting", "b: Back", "q: Quit")
	fmt.Println() // Upcoming

	inputChoice(&choice)
	fmt.Println()

	switch choice {
	case "1":
		sortElectionResults()
	case "b":
		//manageVoter()
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

func sortElectionResults() {
	var choice, sortBy, sorting string

	fmt.Print("[i] ID  ")
	fmt.Print("[n] Name  ")
	fmt.Print("[c] Choosen Candidate  \n")
	switch choice {
	case "i":
		sortBy = "created"
	case "n":
		sortBy = "name"
	case "c":
		sortBy = "candidate"
	default:
		fmt.Println("Wrong choice! try again")
		//sortVoters()
	}
	fmt.Print("Sort By: ")
	fmt.Scan(&sortBy)
	fmt.Print("Sorting [a] Ascending [d] Descending: ")
	fmt.Scan(&sorting)

	fmt.Println()

	//viewAllVoters()
}

func resetElection() {
	var choice string

	fmt.Print("Are you sure you want to reset the election? (Y/n) ")
	fmt.Scan(&choice)
	fmt.Println()

	switch choice {
	case "Y", "y":
		// TODO: Reset election from controller
		manageElection()
	case "n", "N":
		manageElection()
	default:
		fmt.Println("Invalid choice. Please enter 'Y' or 'n'.")
		resetElection()
	}
}
