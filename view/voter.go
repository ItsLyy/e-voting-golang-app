package view

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

func manageVoter(candidateSettings, voterSetting *model.DataSetting) {
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
		viewAllVoters(candidateSettings, voterSetting)
	case "2":
		addVoter(candidateSettings, voterSetting)
	case "3":
		deleteVoter(candidateSettings, voterSetting)
	case "b":
		Home(candidateSettings, voterSetting)
	case "q":
		exit()
	default:
		wrong(choice)
		manageVoter(candidateSettings, voterSetting)
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

func viewAllVoters(candidateSettings, voterSetting *model.DataSetting) {
	var choice string

	displayTableVoter()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()
	inputChoice(&choice)
	fmt.Println()

	switch choice {
	case "1":
		searchVoters(candidateSettings, voterSetting)
	case "2":
		sortVoters(candidateSettings, voterSetting)
	case "b":
		manageVoter(candidateSettings, voterSetting)
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

func addVoter(candidateSettings, voterSetting *model.DataSetting) {
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
			addVoter(candidateSettings, voterSetting)
		}
	case "n", "N":
		manageVoter(candidateSettings, voterSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		addVoter(candidateSettings, voterSetting)
	}

	manageVoter(candidateSettings, voterSetting)
}

func deleteVoter(candidateSettings, voterSetting *model.DataSetting) {
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
		var res controller.Response = controller.DeleteVoter(id, *voterSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", id)
		} else {
			fmt.Println("ID not found! try again,")
			deleteVoter(candidateSettings, voterSetting)
		}
	case "n", "N":
		manageVoter(candidateSettings, voterSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		deleteVoter(candidateSettings, voterSetting)
	}

	manageVoter(candidateSettings, voterSetting)
}

func searchVoters(candidateSettings, voterSetting *model.DataSetting) {
	var choice string
	var res controller.Response

	fmt.Print("[i] ID  ")
	fmt.Print("[n] Name  ")
	fmt.Println()
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	fmt.Print("Search: ")
	switch choice {
	case "i":
		var id int
		fmt.Scan(&id)

		res = controller.ShowVoterById(id, *voterSetting)
	case "n":
		var name string
		fmt.Scan(&name)

		res = controller.ShowVoterByName(name, *voterSetting)
	default:
		fmt.Println("Wrong choice! try again")
		searchVoters(candidateSettings, voterSetting)
	}

	if !res.Success {
		fmt.Println(res.Message)
		searchVoters(candidateSettings, voterSetting)
	} else {
		var data model.VoterData
		data = res.Data.(model.VoterData)
		displayVoter(data, candidateSettings, voterSetting)
	}

	viewAllVoters(candidateSettings, voterSetting)
}

func displayVoter(data model.VoterData, candidateSettings, voterSetting *model.DataSetting) {
	fmt.Println()
	fmt.Printf("%-10s%-18s%-18s\n", "Voter ID", "Candidate Number", "Name")
	fmt.Printf("%-10d%-18d%-18s\n", data.VoterId, data.CandidateNumber, data.Name)
	fmt.Println()
	fmt.Println("1. Change name\n2. Delete voter\n3. Return to voter list")

	fmt.Println()
	var choice string
	inputChoice(&choice)
	switch choice {
	case "1":
		changeNameVoter(data, candidateSettings, voterSetting)
	case "2":
		deleteDataVoterDetail(data, candidateSettings, voterSetting)
	case "3":
		viewAllVoters(candidateSettings, voterSetting)
	}
}

func deleteDataVoterDetail(data model.VoterData, candidateSettings, voterSetting *model.DataSetting) {
	var choice string
	confirmation(&choice)
	switch choice {
	case "Y", "y":
		var res controller.Response = controller.DeleteVoter(data.VoterId, *voterSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", data.VoterId)
		} else {
			fmt.Println("ID not found! try again,")
			deleteVoter(candidateSettings, voterSetting)
		}
	case "n", "N":
		viewAllVoters(candidateSettings, voterSetting)
	}
}

func changeNameVoter(data model.VoterData, candidateSettings, voterSetting *model.DataSetting) {
	var id int = data.VoterId
	fmt.Print("Enter new name: ")
	fmt.Scan(&data.Name)

	var res controller.Response = controller.UpdateVoter(id, data, *voterSetting)
	fmt.Println(res.Message)
	displayVoter(data, candidateSettings, voterSetting)
}

func sortVoters(candidateSettings, voterSetting *model.DataSetting) {
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
		sortVoters(candidateSettings, voterSetting)
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
		sortVoters(candidateSettings, voterSetting)
	}

	fmt.Println()

	if voterSetting.SortSetting == "selection" {
		controller.SelectionVotersSorting(voterSetting.SortOrder, voterSetting.SortBy)
	} else if voterSetting.SortSetting == "insertion" {
		controller.InsertionVotersSorting(voterSetting.SortOrder, voterSetting.SortBy)
	}

	viewAllVoters(candidateSettings, voterSetting)
}

func castingVote(candidateSettings, voterSetting *model.DataSetting) {
	displayTableVoter()
	fmt.Println()

	var voterId, candidateId int
	fmt.Print("Select Voter's ID: ")
	fmt.Scan(&voterId)

	fmt.Println()
	displayTableCandidate()
	fmt.Println()
	fmt.Print("Select Candidate's ID: ")
	fmt.Scan(&candidateId)

	var res controller.Response = controller.CastingVoteVoter(voterId, candidateId, *voterSetting)
	fmt.Println(res.Message)
	fmt.Println()
	manageVoter(candidateSettings, voterSetting)
}
