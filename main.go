package main

import (
	"e-voting/model"
	"e-voting/view"
)

/*
 * main is the entry point of the e-voting application.
 * Purpose: Initialize sample data and default settings, then start the home menu.
 * Flow: Load candidate and voter data -> set default sort/search settings -> call Home menu.
 */
func main() {
	model.Candidate = model.GenerateCandidatesData()
	model.Voter = model.GenerateVotersData()

	var voterSetting model.DataSetting
	var candidateSetting model.DataSetting

	voterSetting.SortBy = ""
	voterSetting.SortOrder = ""
	voterSetting.SortSetting = "selection"
	voterSetting.SearchSetting = "binary"
	candidateSetting.SortBy = ""
	candidateSetting.SortOrder = ""
	candidateSetting.SortSetting = "selection"
	candidateSetting.SearchSetting = "binary"

	view.Home(&candidateSetting, &voterSetting)
}
