package view

import (
	"e-voting/model"
	"fmt"
)

func manageVoter() {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Println("MENU: Manage Voter")
	fmt.Println("Please choose your menu!")
	fmt.Print(normalLine())
	fmt.Println("1. View All Voters")
	fmt.Println("2. Add Voters")
	fmt.Println("3. Delete Voters")
	fmt.Println()
	fmt.Println("b. Back to Home")
	fmt.Println("q. Quit App")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		viewAllVoters()
	case "2":
		addVoter()
	case "3":
		deleteVoter()
	case "b":
		Home()
	case "q":
		exit()
	default:
		wrong(choice)
		manageVoter()
	}
}

func displayTableVoter() {
	var i int = 0

	fmt.Print(tableLine())
	fmt.Printf("|%-10s|%-74s|%-30s|\n", "ID", "Name", "Selected")
	fmt.Print(tableLine())

	for i < model.Voter.Length {
		fmt.Printf("|%-10d|%-74s|%-30d|\n", model.Voter.Data[i].VoterId, model.Voter.Data[i].Name, model.Voter.Data[i].CandidateNumber)
		i++
	}
	fmt.Print(tableLine())
}

func viewAllVoters() {
	var choice string

	displayTableVoter()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()
	inputChoice(&choice)
	fmt.Println()

	switch choice {
	case "1":
		searchVoters()
	case "2":
	case "b":
		manageVoter()
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

func addVoter() {
	var choice, name string

	fmt.Printf("%-16s%s", "Name", ": ")
	fmt.Scan(&name)
	confirmation(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		// [Controller] -> Adding Voter
		fmt.Printf("\n%s has successfully added to the data!\n", name)
	case "n", "N":
		manageVoter()
	default:
		fmt.Println("Something went wrong! try again,")
		addVoter()
	}

	manageVoter()
}

func deleteVoter() {
	var choice, id string

	displayTableVoter()
	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&id)
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		var index int = -1
		// [Controller] -> Delete Voters
		if index != -1 {
			fmt.Printf("\n%s has successfully deleted to the data!\n", id)
		} else {
			fmt.Println("ID not found! try again,")
			deleteVoter()
		}
	case "n", "N":
		manageVoter()
	default:
		fmt.Println("Something went wrong! try again,")
		deleteVoter()
	}

	manageVoter()
}

func searchVoters() {
	var choice, searchBy, target string

	fmt.Print("[i] ID  ")
	fmt.Print("[n] Name  ")
	fmt.Print("[c] Choosen Candidate  \n")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i":
		searchBy = "id"
	case "n":
		searchBy = "name"
	case "c":
		searchBy = "candidate"
	default:
		fmt.Println("Wrong choice! try again")
		searchVoters()
	}

	fmt.Print("Search: ")
	fmt.Scan(&target)
	fmt.Println(searchBy)

	// [Controller] -> data = SearchData(searchBy, target)

	viewAllVoters()
}

func sortVoters() {
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
		sortVoters()
	}
	fmt.Print("Sort By: ")
	fmt.Scan(&sortBy)
	fmt.Print("Sorting [a] Ascending [d] Descending: ")
	fmt.Scan(&sorting)

	fmt.Println()

	viewAllVoters()
}
