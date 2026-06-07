package view

import (
	"e-voting/model"
	"fmt"
)

func ManageCandidates() {
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
		ViewAllCandidate()
	case "2" :
		AddCandidate()
	case "3" :
		DeleteCandidate()
	case "b" :
		
	case "q" :
		fmt.Print("GOOD BYE AND THANKS :)")
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

func ViewAllCandidate() {
	var choice string
	fmt.Println("kamu sedang berada di ViewAllCandidate")
	displayTableCandidate()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()

	fmt.Print("Choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1" :
		SearchCandidate()
		
	case "2" :
		SortCandidate()
		
	case "b" : 
		ManageCandidates()
	case "q" :
		fmt.Print("GOOD BYE AND THANKS :)")
	}
}

func AddCandidate() {
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
		ManageCandidates()
	default:
		fmt.Println("Something went wrong! try again,")
		AddCandidate()
	}

	ManageCandidates()
}

func DeleteCandidate() {
	var choice, id string

	displayTableCandidate()
	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&id)
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "y","Y" : 
		var index int = -1 
		if index != -1 {
			fmt.Printf("\n%s has succesfully deleted to data!\n", id)
		}else {
			fmt.Println("ID not found! try again,")
			DeleteCandidate()
		}
	case "n","N" : 
		ManageCandidates()
	default :
		fmt.Println("something went wrong! try again,")
		DeleteCandidate()
	}
	ManageCandidates()

}

func SearchCandidate() {
	var choice ,searchBY,target string

	fmt.Print("[i] ID ")
	fmt.Print("[n] Name ")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i":
		searchBY = "id"
	case "n":
		searchBY = "name"
	default :
		fmt.Println("wrong choice! try again")
		SearchCandidate()
	}

	fmt.Print("Search: ")
	fmt.Scan(&target)
	fmt.Println(searchBY)

	ViewAllCandidate()

}

func SortCandidate() {
	var choice, sortBy, sorting string

	fmt.Println("[i] ID ")
	fmt.Println("[n] Name ")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i" :
		sortBy = "id"
	case "n" :
		sortBy = "name"
	default :
		fmt.Println("wrong choice! try again")
		SortCandidate()
	}
	fmt.Print("Sort By: ")
	fmt.Scan(&sortBy)
	fmt.Print("sorting [a] Ascending [d] Descending")
	fmt.Scan(&sorting)

	fmt.Println()

	ViewAllCandidate()
}


