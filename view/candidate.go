package view

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

/*
 * manageCandidates displays the candidate management submenu.
 * Purpose: Let the user view, add, or delete candidates.
 * Flow: Show submenu -> read choice -> route to the selected action or go back home.
 */
func manageCandidates(candidateSetting, voterSetting *model.DataSetting) {
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
	case "1":
		viewAllCandidate(candidateSetting, voterSetting)
	case "2":
		addCandidate(candidateSetting, voterSetting)
	case "3":
		deleteCandidate(candidateSetting, voterSetting)
	case "b":
		Home(candidateSetting, voterSetting)
	case "q":
		exit()
	default:
		wrong(choose)
		manageCandidates(candidateSetting, voterSetting)
	}
}

/*
 * displayTableCandidate prints all candidates in a formatted table.
 * Purpose: Show candidate number, name, and vision mission in the console.
 * Flow: Print table header -> loop each candidate and join vision mission words -> print rows.
 */
func displayTableCandidate() {
	var i int = 0
	fmt.Print(tableLine())
	fmt.Printf("|%-10s|%-54s|%-50s|\n", "Number", "Name", "Vision Mission")
	fmt.Print(tableLine())

	for i < model.Candidate.Length {
		var j int = 0
		var text string = ""
		for j < model.Candidate.Data[i].VMLength {
			text += model.Candidate.Data[i].VisionMission[j] + " "
			j++
		}
		fmt.Printf("|%-10d|%-54s|%-50s|\n", model.Candidate.Data[i].CandidateNumber, model.Candidate.Data[i].Name, text)
		i++
	}
	fmt.Print(tableLine())

}

/*
 * viewAllCandidate shows the candidate list with search and sort options.
 * Purpose: Display all candidates and let the user search, sort, go back, or quit.
 * Flow: Print candidate table -> show action menu -> route to search, sort, or navigation.
 */
func viewAllCandidate(candidateSetting, voterSetting *model.DataSetting) {
	var choice string
	fmt.Println("kamu sedang berada di ViewAllCandidate")
	displayTableCandidate()
	fmt.Printf("%-30s%-30s%-30s%-30s\n", "1: Searching", "2: Sorting", "b: Back", "q: Quit")
	fmt.Println()

	fmt.Print("Choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		searchCandidate(candidateSetting, voterSetting)
	case "2":
		sortCandidate(candidateSetting, voterSetting)
	case "b":
		manageCandidates(candidateSetting, voterSetting)
	case "q":
		exit()
	default:
		wrong(choice)
		inputChoice(&choice)
	}
}

/*
 * addCandidate handles the form to create a new candidate.
 * Purpose: Collect candidate details from the user and save them via the controller.
 * Flow: Read number, name, and vision mission words -> confirm -> call CreateCandidate -> return to menu.
 */
func addCandidate(candidateSetting, voterSetting *model.DataSetting) {
	var choice string
	var candidate model.CandidateData

	fmt.Printf("%-16s%s", "Number", ": ")
	fmt.Scan(&candidate.CandidateNumber)
	fmt.Printf("%-16s%s", "Name", ": ")
	fmt.Scan(&candidate.Name)
	fmt.Printf("%-16s%s", "VisionMission", ": ")
	fmt.Println("use word 'STOP' to finish typing.")

	var word string
	fmt.Scan(&word)

	for word != "STOP" {
		candidate.VisionMission[candidate.VMLength] = word
		candidate.VMLength++
		fmt.Scan(&word)
	}

	confirmation(&choice)

	fmt.Print("\n", normalLine())
	switch choice {
	case "Y", "y":
		var res controller.Response = controller.CreateCandidate(candidate, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%s has suuccesfully added to the data!\n", candidate.Name)
		} else {
			fmt.Printf("%s", res.Message)
			fmt.Println(normalLine())
			addCandidate(candidateSetting, voterSetting)
		}
	case "n", "N":
		manageCandidates(candidateSetting, voterSetting)
	default:
		fmt.Println("Something went wrong! try again,")
		addCandidate(candidateSetting, voterSetting)
	}

	manageCandidates(candidateSetting, voterSetting)
}

/*
 * deleteCandidate handles removing a candidate by ID.
 * Purpose: Let the user pick a candidate to delete with confirmation.
 * Flow: Show table -> read ID -> confirm -> call DeleteCandidate -> return to menu.
 */
func deleteCandidate(candidateSetting, voterSetting *model.DataSetting) {
	var choice string
	var id int

	displayTableCandidate()
	fmt.Printf("%-16s%s", "ID", ": ")
	fmt.Scan(&id)
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(&choice)
	fmt.Print("\n", normalLine())
	switch choice {
	case "y", "Y":
		var res controller.Response = controller.DeleteCandidate(id, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", id)
		} else {
			fmt.Println("ID not found! try again,")
			deleteCandidate(candidateSetting, voterSetting)
		}
	case "n", "N":
		manageCandidates(candidateSetting, voterSetting)
	default:
		fmt.Println("something went wrong! try again,")
		deleteCandidate(candidateSetting, voterSetting)
	}
	manageCandidates(candidateSetting, voterSetting)

}

/*
 * searchCandidate finds and displays a single candidate.
 * Purpose: Let the user search by ID or name and view the result.
 * Flow: Choose search field -> read value -> call ShowCandidateById or ShowCandidateByName -> display result or retry.
 */
func searchCandidate(candidateSetting, voterSetting *model.DataSetting) {
	var choice string
	var res controller.Response

	fmt.Print("[i] ID ")
	fmt.Println("[n] Name ")
	fmt.Print("Search By: ")
	fmt.Scan(&choice)
	fmt.Print("Search: ")
	switch choice {
	case "i":
		var id int
		fmt.Scan(&id)

		res = controller.ShowCandidateById(id, *candidateSetting)
	case "n":
		var name string
		fmt.Scan(&name)

		res = controller.ShowCandidateByName(name, *candidateSetting)
	default:
		fmt.Println("wrong choice! try again")
		searchCandidate(candidateSetting, voterSetting)
	}

	if !res.Success {
		fmt.Println(res.Message)
		searchCandidate(candidateSetting, voterSetting)
	} else {
		var data model.CandidateData
		data = res.Data.(model.CandidateData)
		displayCandidate(data, candidateSetting, voterSetting)
	}

	viewAllCandidate(candidateSetting, voterSetting)

}

/*
 * displayCandidate shows one candidate's details with edit options.
 * Purpose: Let the user view and modify a single candidate after searching.
 * Flow: Print candidate info -> show edit menu -> route to change name, vision mission, delete, or go back.
 */
func displayCandidate(data model.CandidateData, candidateSetting, voterSetting *model.DataSetting) {
	fmt.Println()
	fmt.Printf("%-10s%-18s%-60s\n", "Number", "Name", "Vision Mission")
	fmt.Printf("%-10d%-18s", data.CandidateNumber, data.Name)
	var i int
	for i = 0; i < data.VMLength; i++ {
		fmt.Print(data.VisionMission[i], " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("1. Change name\n2. Change vision mission\n3. Delete candidate\n4. Return to candidate list")

	fmt.Println()
	var choice string
	inputChoice(&choice)
	switch choice {
	case "1":
		changeNameCandidate(data, candidateSetting, voterSetting)
	case "2":
		changeVisionMissionCandidate(data, candidateSetting, voterSetting)
	case "3":
		deleteDataCandidateDetail(data, candidateSetting, voterSetting)
	case "4":
		viewAllCandidate(candidateSetting, voterSetting)
	}
}

/*
 * changeNameCandidate updates a candidate's name.
 * Purpose: Let the user rename a candidate from the detail view.
 * Flow: Read new name -> call UpdateCandidate -> show updated detail view.
 */
func changeNameCandidate(data model.CandidateData, candidateSetting, voterSetting *model.DataSetting) {
	var id int = data.CandidateNumber
	fmt.Print("Enter new name: ")
	fmt.Scan(&data.Name)

	var res controller.Response = controller.UpdateCandidate(id, data, *candidateSetting)
	fmt.Println(res.Message)
	displayCandidate(data, candidateSetting, voterSetting)
}

/*
 * changeVisionMissionCandidate updates a candidate's vision mission.
 * Purpose: Let the user replace the vision mission words from the detail view.
 * Flow: Read new words until STOP -> call UpdateCandidate -> show updated detail view.
 */
func changeVisionMissionCandidate(data model.CandidateData, candidateSetting, voterSetting *model.DataSetting) {
	var word string
	fmt.Print("Enter new vision mission: ")
	fmt.Scan(&word)

	data.VisionMission = [999]string{}
	data.VMLength = 0

	for word != "STOP" {
		data.VisionMission[data.VMLength] = word
		data.VMLength++
		fmt.Scan(&word)
	}

	var res controller.Response = controller.UpdateCandidate(data.CandidateNumber, data, *candidateSetting)
	fmt.Println(res.Message)
	displayCandidate(data, candidateSetting, voterSetting)
}

/*
 * deleteDataCandidateDetail removes a candidate from the detail view.
 * Purpose: Let the user delete a candidate after viewing their details.
 * Flow: Confirm action -> call DeleteCandidate -> show result or return to list.
 */
func deleteDataCandidateDetail(data model.CandidateData, candidateSetting, voterSetting *model.DataSetting) {
	var choice string
	confirmation(&choice)
	switch choice {
	case "Y", "y":
		var res controller.Response = controller.DeleteCandidate(data.CandidateNumber, *candidateSetting)
		if res.Success {
			fmt.Printf("\n%d has successfully deleted to the data!\n", data.CandidateNumber)
		} else {
			fmt.Println("ID not found! try again,")
			deleteCandidate(candidateSetting, voterSetting)
		}
	case "n", "N":
		manageCandidates(candidateSetting, voterSetting)
	}
}

/*
 * sortCandidate configures and runs candidate sorting.
 * Purpose: Let the user choose sort field and order, then reorder the candidate list.
 * Flow: Choose sort by (ID/name) -> choose order (asc/desc) -> call SelectionCandidateSorting -> return to list.
 */
func sortCandidate(candidateSetting, voterSetting *model.DataSetting) {
	var choice string

	fmt.Println("[i] ID ")
	fmt.Println("[n] Name ")
	fmt.Print("Sort By: ")
	fmt.Scan(&choice)
	switch choice {
	case "i":
		candidateSetting.SortBy = "id"
	case "n":
		candidateSetting.SortBy = "name"
	default:
		fmt.Println("wrong choice! try again")
		sortCandidate(candidateSetting, voterSetting)
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
		sortCandidate(candidateSetting, voterSetting)
	}

	fmt.Println()

	controller.SelectionCandidateSorting(candidateSetting.SortOrder, candidateSetting.SortBy)

	viewAllCandidate(candidateSetting, voterSetting)
}
