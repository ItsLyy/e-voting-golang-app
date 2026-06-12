package view

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

func manageElection(candidateSetting, voterSetting *model.DataSetting) {
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
		viewElectionResults(candidateSetting, voterSetting)
	case "2":
		resetElection(candidateSetting, voterSetting)
	case "b":
		Home(candidateSetting, voterSetting)
	case "q":
		exit()
	}
}

func displayTableElectionResults() {
	var i int = 0
	var election model.Elections = controller.InputElection()
	var totalVotes int = controller.SumVotes(election)

	fmt.Print(tableLine())
	fmt.Printf("|%-80s|%-20s|%-13s|\n", "Name", "Persentage(%)", "Total Voted")
	fmt.Print(tableLine())

	for i < model.Candidate.Length {
		var candidateIndex = election.Data[i].CandidateIndex
		fmt.Printf("|%-80s|%-20.2f|%-13d|\n", model.Candidate.Data[candidateIndex].Name, controller.CalculatePercentage(totalVotes, election.Data[i].TotalVote), election.Data[i].TotalVote)
		i++
	}
	fmt.Print(tableLine())
}

func viewElectionResults(candidateSetting, voterSetting *model.DataSetting) {
	var choice string

	displayTableElectionResults()
	fmt.Printf("%-30s%-30s\n", "b: Back", "q: Quit")
	fmt.Println() // Upcoming

	inputChoice(&choice)
	fmt.Println()

	switch choice {
	case "b":
		manageElection(candidateSetting, voterSetting)
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

func resetElection(candidateSetting, voterSetting *model.DataSetting) {
	var choice string

	fmt.Print("Are you sure you want to reset the election? (Y/n) ")
	fmt.Scan(&choice)
	fmt.Println()

	switch choice {
	case "Y", "y":
		controller.ResetElection()
		manageElection(candidateSetting, voterSetting)
	case "n", "N":
		manageElection(candidateSetting, voterSetting)
	default:
		fmt.Println("Invalid choice. Please enter 'Y' or 'n'.")
		resetElection(candidateSetting, voterSetting)
	}
}
