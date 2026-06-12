package main

import (
	"e-voting/model"
	"e-voting/view"
)

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
