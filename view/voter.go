package view

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

func manageVoter(voterSetting *model.DataSetting) {
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
		viewAllVoters(voterSetting)
	case "2":
		addVoter(voterSetting)
	case "3":
		deleteVoter(voterSetting)
	case "b":
		Home()
	case "q":
		exit()
	default:
		wrong(choice)
		manageVoter(voterSetting)
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

func viewAllVoters(voterSetting *model.DataSetting) {
	var choice string

	displayTableVoter()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()
	inputChoice(&choice)
	fmt.Println()

	switch choice {
	case "1":
		searchVoters(voterSetting)
	case "2":
		sortVoters(voterSetting)
	case "b":
		manageVoter(voterSetting)
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

func addVoter(voterSetting *model.DataSetting) {
	var choice string
	var voter model.VoterData

	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&voter.VoterId)
	fmt.Printf("%-16s%s", "Name", ": ")
	fmt.Scan(&voter.Name)

	confirmation(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		var res controller.Response = controller.CreateVoter(voter, *voterSetting)
		if res.Success {
			fmt.Printf("\n%s has successfully added to the data!\n", voter.Name)
		} else {
			fmt.Printf("%s", res.Message)
			fmt.Println(normalLine())
			addVoter(voterSetting)
		}
	case "n", "N":
		manageVoter(voterSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		addVoter(voterSetting)
	}

	manageVoter(voterSetting)
}

func deleteVoter(voterSetting *model.DataSetting) {
	var choice string
	var id int

	displayTableVoter()
	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&id)
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		// [Controller] -> Delete Voters
		var res controller.Response = controller.DeleteVoter(id, *voterSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", id)
		} else {
			fmt.Println("ID not found! try again,")
			deleteVoter(voterSetting)
		}
	case "n", "N":
		manageVoter(voterSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		deleteVoter(voterSetting)
	}

	manageVoter(voterSetting)
}

func searchVoters(voterSetting *model.DataSetting) {
	var choice string
	var res controller.Response

	fmt.Print("[i] ID  ")
	fmt.Print("[n] Name  ")
	fmt.Print("[c] Choosen Candidate  \n")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	fmt.Print("Search: ")
	switch choice {
	case "i":
		var id int
		fmt.Scan(&id)

		res = controller.ShowVoter(id, *voterSetting)
	case "n":
	case "c":
	default:
		fmt.Println("Wrong choice! try again")
		searchVoters(voterSetting)
	}

	if !res.Success {
		fmt.Println(res.Message)
		searchVoters(voterSetting)
	} else {
		var data model.VoterData
		data = res.Data.(model.VoterData)
		displayVoter(data)
	}

	viewAllVoters(voterSetting)
}

func displayVoter(data model.VoterData) {
	fmt.Println()
	fmt.Printf("%-10s %-10s %-10s\n", "Voter ID", "Candidate Number", "Name")
	fmt.Printf("%-10d %-10d %-10s\n", data.VoterId, data.CandidateNumber, data.Name)
	fmt.Println()
	fmt.Println("1. Change name\n2. Delete voter\n3. Return to voter list")
}

func sortVoters(voterSetting *model.DataSetting) {
	var choice string

	fmt.Print("[i] ID  ")
	fmt.Print("[n] Name  ")
	fmt.Print("[c] Choosen Candidate  \n")
	fmt.Print("Sort By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i":
		voterSetting.SortBy = "id"
	case "n":
		voterSetting.SortBy = "name"
	case "c":
		voterSetting.SortBy = "candidate"
	default:
		fmt.Println("Wrong choice! try again")
		sortVoters(voterSetting)
	}
	fmt.Print("Sorting [a] Ascending [d] Descending: ")
	fmt.Scan(&choice)

	switch choice {
	case "a":
		voterSetting.SortOrder = "asc"
	case "d":
		voterSetting.SortOrder = "desc"
	default:
		fmt.Println("Wrong choice! try again")
		sortVoters(voterSetting)
	}

	fmt.Println()

	controller.SelectionVotersSorting(voterSetting.SortOrder, voterSetting.SortBy)

	viewAllVoters(voterSetting)
}
