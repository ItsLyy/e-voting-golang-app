package controller

import "e-voting/model"

/*
 * CreateVoter adds a new voter to the list.
 * Purpose: Handle voter creation from the UI with validation and sorted insertion.
 * Flow: Validate input -> check duplicate via search -> insert at correct position based on sort settings -> increment length -> return response.
 */
func CreateVoter(voter model.VoterData, voterSetting model.DataSetting) Response {
	if voter.VoterId == 0 || voter.Name == "" {
		return Response{Success: false, Message: "invalid voter data"}
	}

	var index = optimizationVoterIdSearchIndex(voter.VoterId, voterSetting)

	if index != -1 {
		return Response{Success: false, Message: "voter already exists"}
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
	return Response{Success: true, Message: "voter created successfully"}
}

/*
 * DeleteVoter removes a voter by voter ID.
 * Purpose: Remove an existing voter from the data.
 * Flow: Find index by ID -> return error if not found -> shift remaining entries left -> decrement length -> return response.
 */
func DeleteVoter(id int, voterSetting model.DataSetting) Response {
	var index = optimizationVoterIdSearchIndex(id, voterSetting)

	if index == -1 {
		return Response{Success: false, Message: "voter not found"}
	}

	model.Voter.Data[index] = model.VoterData{}
	for index < model.Voter.Length-1 {
		model.Voter.Data[index] = model.Voter.Data[index+1]
		index++
	}

	model.Voter.Length--

	return Response{Success: true, Message: "voter deleted successfully"}
}

/*
 * UpdateVoter replaces an existing voter's data.
 * Purpose: Allow editing a voter's name or other fields.
 * Flow: Find index by ID -> return error if not found -> overwrite data at index -> return response.
 */
func UpdateVoter(id int, voter model.VoterData, voterSetting model.DataSetting) Response {
	var index = optimizationVoterIdSearchIndex(id, voterSetting)

	if index == -1 {
		return Response{Success: false, Message: "voter not found"}
	}

	model.Voter.Data[index] = voter
	return Response{Success: true, Message: "voter updated successfully"}
}

/*
 * CastingVoteVoter records a voter's choice for a candidate.
 * Purpose: Assign a candidate number to a voter during voting.
 * Flow: Find voter by ID -> return error if not found -> set CandidateNumber -> return response.
 */
func CastingVoteVoter(id int, candidateId int, voterSetting model.DataSetting) Response {
	var index = optimizationVoterIdSearchIndex(id, voterSetting)

	if index == -1 {
		return Response{Success: false, Message: "voter not found"}
	}

	model.Voter.Data[index].CandidateNumber = candidateId
	return Response{Success: true, Message: "Voter voted successfully"}
}

/*
 * ShowVoterById retrieves a voter by voter ID.
 * Purpose: Return a single voter record for display or editing.
 * Flow: Find index by ID -> return error if not found -> return voter data in response.
 */
func ShowVoterById(id int, voterSetting model.DataSetting) Response {
	var index = optimizationVoterIdSearchIndex(id, voterSetting)
	if index == -1 {
		return Response{Success: false, Message: "voter not found"}
	}

	return Response{Success: true, Message: "voter found", Data: model.Voter.Data[index]}
}

/*
 * ShowVoterByName retrieves a voter by name.
 * Purpose: Return a single voter record when searching by name.
 * Flow: Find index by name -> return error if not found -> return voter data in response.
 */
func ShowVoterByName(name string, voterSetting model.DataSetting) Response {
	var index = optimizationVoterNameSearchIndex(name, voterSetting)
	if index == -1 {
		return Response{Success: false, Message: "voter not found"}
	}

	return Response{Success: true, Message: "voter found", Data: model.Voter.Data[index]}
}

/*
 * optimizationVoterIdSearchIndex finds a voter's array index by ID.
 * Purpose: Pick the fastest search method (binary or sequential) based on current settings.
 * Flow: If sorted by ID with binary search enabled, use binary search -> otherwise use sequential search -> return index or -1.
 */
func optimizationVoterIdSearchIndex(id int, voterSetting model.DataSetting) int {
	var index = -1
	if (voterSetting.SortOrder == "asc" || voterSetting.SortOrder == "desc") && voterSetting.SortBy == "id" && voterSetting.SearchSetting == "binary" {
		index = SearchBinaryVoterById(id, voterSetting.SortOrder)
	} else {
		index = SearchSequentialVoterById(id)
	}

	return index
}

/*
 * optimizationVoterNameSearchIndex finds a voter's array index by name.
 * Purpose: Pick the fastest search method (binary or sequential) based on current settings.
 * Flow: If sorted by name with binary search enabled, use binary search -> otherwise use sequential search -> return index or -1.
 */
func optimizationVoterNameSearchIndex(name string, voterSetting model.DataSetting) int {
	var index = -1
	if (voterSetting.SortOrder == "asc" || voterSetting.SortOrder == "desc") && voterSetting.SortBy == "name" && voterSetting.SearchSetting == "binary" {
		index = SearchBinaryVoterByName(name, voterSetting.SortOrder)
	} else {
		index = SearchSequentialVoterByName(name)
	}

	return index
}
