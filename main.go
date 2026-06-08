package main

import (
	"e-voting/model"
)

func main() {
	model.Candidate = model.GenerateCandidatesData()
	model.Voter = model.GenerateVotersData()

	var voterSetting model.DataSetting
	var candidateSetting model.DataSetting

	voterSetting.SortBy = ""
	voterSetting.SortOrder = ""
	candidateSetting.SortBy = ""
	candidateSetting.SortOrder = ""
}
