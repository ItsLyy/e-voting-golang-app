package view

import (
	"e-voting/model"
	"fmt"
)

/*
 * Home displays the main application menu.
 * Purpose: Let the user navigate to candidate, voter, voting, election, or settings screens.
 * Flow: Show menu options -> read user choice -> route to the selected screen or quit.
 */
func Home(candidateSetting, voterSetting *model.DataSetting) {
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
	fmt.Println("s. Settings")
	fmt.Println("q. Quit App")
	fmt.Print(normalLine())
	inputChoice(&choice)
	fmt.Print(normalLine())
	fmt.Println()

	switch choice {
	case "1":
		manageCandidates(candidateSetting, voterSetting)
	case "2":
		manageVoter(candidateSetting, voterSetting)
	case "3":
		castingVote(candidateSetting, voterSetting)
	case "4":
		manageElection(candidateSetting, voterSetting)
	case "q":
		exit()
	case "s":
		settingMenu(candidateSetting, voterSetting)
	default:
		wrong(choice)
		Home(candidateSetting, voterSetting)
	}
}

/*
 * normalLine returns a decorative separator for menu headers.
 * Purpose: Format console output with consistent visual dividers.
 * Flow: Return a fixed string of dashes and equals signs.
 */
func normalLine() string {
	return "-----------------------------\n=============================\n"
}

/*
 * tableLine returns a wide decorative separator for table headers.
 * Purpose: Format table output with a longer visual divider.
 * Flow: Return a fixed wide string of dashes and equals signs.
 */
func tableLine() string {
	return "--------------------------------------------------------------------------------------------------------------------\n====================================================================================================================\n"
}

/*
 * wrong prints an error message for an invalid menu choice.
 * Purpose: Inform the user that their input was not recognized.
 * Flow: Print separator -> show error with the invalid choice -> print separator.
 */
func wrong(choice string) {
	fmt.Print(normalLine())
	fmt.Println("Wrong Choice there is no", choice)
	fmt.Print(normalLine())
}

/*
 * exit displays a goodbye message when the user quits.
 * Purpose: End the application session gracefully.
 * Flow: Print separator -> show goodbye message -> print separator.
 */
func exit() {
	fmt.Print(normalLine())
	fmt.Println("GOOD BYE THANKS!")
	fmt.Print(normalLine())
}

/*
 * inputChoice reads a single menu choice from the user.
 * Purpose: Reusable prompt for any screen that needs user input.
 * Flow: Print "Choice: " prompt -> scan input into the choice pointer.
 */
func inputChoice(choice *string) {
	fmt.Print("Choice: ")
	fmt.Scan(choice)
}

/*
 * confirmation asks the user to confirm an action.
 * Purpose: Prevent accidental create, update, or delete operations.
 * Flow: Print "Are you sure? (Y/n)" prompt -> scan input into the choice pointer.
 */
func confirmation(choice *string) {
	fmt.Print("\nAre you sure? (Y/n) : ")
	fmt.Scan(choice)
}
