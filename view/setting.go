package view

import (
	"e-voting/model"
	"fmt"
)

func settingMenu(candidateSettings, voterSettings *model.DataSetting) {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Println("Please choose your menu!")
	fmt.Print(normalLine())
	fmt.Println("1. Candidate Sorting")
	fmt.Println("2. Candidate Searching")
	fmt.Println("3. Voter Sorting")
	fmt.Println("4. Voter Searching")
	fmt.Println()
	fmt.Println("b. Back to Main Menu")
	fmt.Println("q. Quit App")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		dataSortingMenu(candidateSettings)
	case "2":
		dataSearchingMenu(candidateSettings)
	case "3":
		dataSortingMenu(voterSettings)
	case "4":
		dataSearchingMenu(voterSettings)
	case "b":
		Home(candidateSettings, voterSettings)
	case "q":
		exit()
	default:
		settingMenu(candidateSettings, voterSettings)
	}

	Home(candidateSettings, voterSettings)
}

func dataSortingMenu(dataSetting *model.DataSetting) {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Print("Currently sorting using: ")
	switch dataSetting.SortSetting {
	case "insertion":
		fmt.Println("Insertion Sorting")
	default:
		fmt.Println("Selection Sorting")
	}
	fmt.Println("Please choose a sorting option:")
	fmt.Print(normalLine())
	fmt.Println("1. Selection Sorting")
	fmt.Println("2. Insertion Sorting")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		dataSetting.SortSetting = "selection"
	case "2":
		dataSetting.SortSetting = "insertion"
	default:
		dataSortingMenu(dataSetting)
	}

	fmt.Println()
	fmt.Println("Sorting setting updated:", dataSetting.SortSetting)
	fmt.Println()
}

func dataSearchingMenu(dataSetting *model.DataSetting) {
	var choice string

	fmt.Println()
	fmt.Print(normalLine())
	fmt.Print("Currently searching using: ")
	switch dataSetting.SearchSetting {
	case "sequential":
		fmt.Println("Sequential Searching")
	default:
		fmt.Println("Binary Searching")
	}
	fmt.Println("Please choose a searching option:")
	fmt.Print(normalLine())
	fmt.Println("1. Sequential Searching")
	fmt.Println("2. Binary Searching")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		dataSetting.SearchSetting = "sequential"
	case "2":
		dataSetting.SearchSetting = "binary"
	default:
		dataSearchingMenu(dataSetting)
	}

	fmt.Println()
	fmt.Println("Searching setting updated:", dataSetting.SearchSetting)
	fmt.Println()
}
