package controller

import "e-voting/model"

func CreateVoter(voter model.VoterData, voterSetting model.DataSetting) response {
	if voter.VoterId == 0 || voter.Name == "" {
		return response{success: false, message: "invalid voter data"}
	}

	var index = optimazitaionVoterIdSearchIndex(voter.VoterId, voterSetting)

	if index != -1 {
		return response{success: false, message: "voter already exists"}
	}

	var i int = model.Voter.Length - 1

	switch voterSetting.SortOrder {
	case "asc":
		switch voterSetting.SortBy {
		case "id":
			for i >= 0 && IsVoterLess(voter, model.Voter.Data[i], "id") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		case "candidate":
			for i >= 0 && IsVoterLess(voter, model.Voter.Data[i], "candidate") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		case "name":
			for i >= 0 && IsVoterLess(voter, model.Voter.Data[i], "name") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		}
	case "desc":
		switch voterSetting.SortBy {
		case "id":
			for i >= 0 && IsVoterGreater(voter, model.Voter.Data[i], "id") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		case "candidate":
			for i >= 0 && IsVoterGreater(voter, model.Voter.Data[i], "candidate") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		case "name":
			for i >= 0 && IsVoterGreater(voter, model.Voter.Data[i], "name") {
				model.Voter.Data[i] = model.Voter.Data[i-1]
				i--
			}
			model.Voter.Data[i+1] = voter
		}
	default:
		model.Voter.Data[model.Voter.Length] = voter
	}

	model.Voter.Length++
	return response{success: true, message: "voter created successfully"}
}

func DeleteVoter(id int, voterSetting model.DataSetting) response {
	var index = optimazitaionVoterIdSearchIndex(id, voterSetting)

	if index == -1 {
		return response{success: false, message: "voter not found"}
	}

	model.Voter.Data[index] = model.VoterData{}
	for index < model.Voter.Length-1 {
		model.Voter.Data[index] = model.Voter.Data[index+1]
		index++
	}

	model.Voter.Length--

	return response{success: true, message: "voter deleted successfully"}
}

func UpdateVoter(id int, voter model.VoterData, voterSetting model.DataSetting) response {
	var index = optimazitaionVoterIdSearchIndex(id, voterSetting)

	if index == -1 {
		return response{success: false, message: "voter not found"}
	}

	model.Voter.Data[index] = voter
	return response{success: true, message: "voter updated successfully"}
}

func optimazitaionVoterIdSearchIndex(id int, voterSetting model.DataSetting) int {
	var index = -1
	if (voterSetting.SortOrder == "asc" || voterSetting.SortOrder == "desc") && voterSetting.SortBy == "id" {
		index = SearchBinaryVoterById(id, voterSetting.SortOrder)
	} else {
		index = SearchSequentialVoterById(id)
	}

	return index
}
