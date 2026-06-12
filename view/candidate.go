package view

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

func ManageCandidates(candidateSetting *model.DataSetting) {
	var choose string
	fmt.Print(normalLine())
	fmt.Println("selamat datang di Candidate Management")
	fmt.Println("Please choose your menu")
	fmt.Print(normalLine())
	fmt.Println("1. View All Candidate")
	fmt.Println("2. Add Candidate")
	fmt.Println("3. Delete Candidate")
	fmt.Println()
	fmt.Println("b, to back to home")
	fmt.Println("q, to quit the app")
	fmt.Print(normalLine())

	fmt.Print("Choice: ")
	fmt.Scan(&choose)
	switch choose {
		case "1" : 
		ViewAllCandidate(candidateSetting)
	case "2" :
		AddCandidate(candidateSetting)
	case "3" :
		DeleteCandidate(candidateSetting)
	case "b" :
		Home()
		
	case "q" :
		exit()
		default : 
		wrong(choose)
		ManageCandidates(candidateSetting)
	}
}

func displayTableCandidate(){
	var i int = 0
	fmt.Print(tableLine())
	fmt.Printf("%-30s%-44s%-40s\n", "CANDIDATENUMBER","NAME","VISIONMISSION")
	fmt.Print(tableLine())

	for i <model.Candidate.Length {
		fmt.Printf("|%-30d|%-44s|%-40s\n",model.Candidate.Data[i].CandidateNumber,model.Candidate.Data[i].Name,model.Candidate.Data[i].VisionMission)
		i++
	}
	fmt.Print(tableLine())
		
}

func ViewAllCandidate(candidateSetting *model.DataSetting) {
	var choice string
	fmt.Println("kamu sedang berada di ViewAllCandidate")
	displayTableCandidate()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()

	fmt.Print("Choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1" :
		SearchCandidate(candidateSetting)
		
	case "2" :
		SortCandidate(candidateSetting)
		
	case "b" : 
		ManageCandidates(candidateSetting)
	case "q" :
		exit()
		default : 
		wrong(choice)
		inputChoice(&choice)
	}
}

func AddCandidate(candidateSetting *model.DataSetting) {
	var choice string
	var candidate model.CandidateData

	fmt.Printf("%-16s%s", "CandidateNumber", ": ")
	fmt.Scan(&candidate.CandidateNumber)
	fmt.Printf("%-16s%s", "Name", ": ")
	fmt.Scan(&candidate.Name)
	confirmation(&choice)
	
	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		// [Controller] -> Adding Voter
		var res controller.Response = controller.CreateCandidate(candidate, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%s has suuccesfully added to the data!\n",candidate.Name)
		}else {
			fmt.Printf("%s", res.Message)
			fmt.Println(normalLine())
			AddCandidate(candidateSetting)
		}
	case "n", "N":
		ManageCandidates(candidateSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		AddCandidate(candidateSetting)
	}

	ManageCandidates(candidateSetting)
}

func DeleteCandidate(candidateSetting *model.DataSetting) {
	var choice string
	var id int

	displayTableCandidate()
	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&id)
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "y","Y" : 
		var res controller.Response = controller.DeleteCandidate(id, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", id)
		} else {
			fmt.Println("ID not found! try again,")
			DeleteCandidate(candidateSetting)
		}
	case "n","N" : 
		ManageCandidates(candidateSetting)
	default :
		fmt.Println("something went wrong! try again,")
		DeleteCandidate(candidateSetting)
	}
	ManageCandidates(candidateSetting)

}

func SearchCandidate(candidateSetting *model.DataSetting) {
	var choice string
	var res controller.Response


	fmt.Print("[i] ID ")
	fmt.Print("[n] Name ")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i":
		var id int 
		fmt.Scan(&id)

		res = controller.ShowCandidateById(id, *candidateSetting)
	case "n":
		var name string
		fmt.Scan(&name)

		res = controller.ShowCandidateByName(name, *candidateSetting)
	default :
		fmt.Println("wrong choice! try again")
		SearchCandidate(candidateSetting)
	}

	if !res.Success {
		fmt.Println(res.Message)
		SearchCandidate(candidateSetting)
	} else {
		var data model.CandidateData
		data = res.Data.(model.CandidateData)
		displayCandidate(data, candidateSetting)
	}

	ViewAllCandidate(candidateSetting)

}

func displayCandidate(data model.CandidateData, candidateSetting *model.DataSetting) {
	fmt.Println()
	fmt.Printf("%-10s%-18s\n", "Candidate Number", "Name")
	fmt.Printf("%-10d%-18s\n", data.CandidateNumber, data.Name)
	fmt.Println()
	fmt.Println("1. Change name\n2. Delete candidate\n3. Return to candidate list")

	fmt.Println()
	var choice string
	inputChoice(&choice)
	switch choice {
	case "1":
		changeNameCandidate(data, candidateSetting)
	case "2":
		deleteDataCandidateDetail(data, candidateSetting)
	case "3":
		ViewAllCandidate(candidateSetting)
	}
}

func changeNameCandidate(data model.CandidateData, candidateSetting *model.DataSetting) {
	var id int = data.CandidateNumber
	fmt.Print("Enter new name: ")
	fmt.Scan(&data.Name)

	var res controller.Response = controller.UpdateCandidate(id, data, *candidateSetting)
	fmt.Println(res.Message)
	displayCandidate(data, candidateSetting)
}

func deleteDataCandidateDetail(data model.CandidateData, candidateSetting *model.DataSetting) {
	var choice string
	confirmation(&choice)
	switch choice {
	case "Y", "y":
		var res controller.Response = controller.DeleteVoter(data.CandidateNumber, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", data.CandidateNumber)
		} else {
			fmt.Println("ID not found! try again,")
			deleteVoter(candidateSetting)
		}
	case "n", "N":
		viewAllVoters(candidateSetting)
	}
}

func SortCandidate(candidateSetting *model.DataSetting) {
	var choice string

	fmt.Println("[i] ID ")
	fmt.Println("[n] Name ")
	fmt.Print("Sort By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i" :
		candidateSetting.SortBy ="id"
	case "n" :
		candidateSetting.SortBy ="name"
	default :
		fmt.Println("wrong choice! try again")
		SortCandidate(candidateSetting)
	}
	fmt.Print("Sorting [a] Ascending [d] Descending: ")
	fmt.Scan(&choice)

	switch choice {
	case "a":
		candidateSetting.SortOrder = "asc"
	case "d":
		candidateSetting.SortOrder = "desc"
	default:
		fmt.Println("Wrong choice! try again")
		SortCandidate(candidateSetting)
	}

	fmt.Println()

	controller.SelectionCandidateSorting(candidateSetting.SortOrder, candidateSetting.SortBy)

	ViewAllCandidate(candidateSetting)
}




