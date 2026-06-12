package controller

import "e-voting/model"

/*
 * CreateCandidate adds a new candidate to the list.
 * Purpose: Handle candidate creation from the UI with validation and sorted insertion.
 * Flow: Validate input -> check duplicate via search -> insert at correct position based on sort settings -> increment length -> return response.
 */
func CreateCandidate(candidate model.CandidateData, candidateSetting model.DataSetting) Response {
	if candidate.CandidateNumber == 0 || candidate.Name == "" {
		return Response{Success: false, Message: "invalid candidate data"}
	}

	var index = optimizationCandidateSearchIndex(candidate.CandidateNumber, candidateSetting)

	if index != -1 {
		return Response{Success: false, Message: "candidate already exists"}
	}

	var i int = model.Candidate.Length - 1

	switch candidateSetting.SortOrder {
	case "asc":
		switch candidateSetting.SortBy {
		case "id":
			for i >= 0 && IsCandidateLess(candidate, model.Candidate.Data[i], "id") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		case "candidate":
			for i >= 0 && IsCandidateLess(candidate, model.Candidate.Data[i], "candidate") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		case "name":
			for i >= 0 && IsCandidateLess(candidate, model.Candidate.Data[i], "name") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		}
	case "desc":
		switch candidateSetting.SortBy {
		case "id":
			for i >= 0 && IsCandidateGreater(candidate, model.Candidate.Data[i], "id") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		case "candidate":
			for i >= 0 && IsCandidateGreater(candidate, model.Candidate.Data[i], "candidate") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		case "name":
			for i >= 0 && IsCandidateGreater(candidate, model.Candidate.Data[i], "name") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			model.Candidate.Data[i+1] = candidate
		}
	default:
		model.Candidate.Data[model.Candidate.Length] = candidate
	}

	model.Candidate.Length++
	return Response{Success: true, Message: "candidate created successfully"}
}

/*
 * DeleteCandidate removes a candidate by candidate number.
 * Purpose: Remove an existing candidate from the data.
 * Flow: Find index by ID -> return error if not found -> shift remaining entries left -> decrement length -> return response.
 */
func DeleteCandidate(id int, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateSearchIndex(id, candidateSetting)

	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	model.Candidate.Data[index] = model.CandidateData{}
	for index < model.Candidate.Length-1 {
		model.Candidate.Data[index] = model.Candidate.Data[index+1]
		index++
	}

	model.Candidate.Length--

	return Response{Success: true, Message: "candidate deleted successfully"}
}

/*
 * UpdateCandidate replaces an existing candidate's data.
 * Purpose: Allow editing a candidate's name or vision mission.
 * Flow: Find index by ID -> return error if not found -> overwrite data at index -> return response.
 */
func UpdateCandidate(id int, candidate model.CandidateData, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateSearchIndex(id, candidateSetting)

	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	model.Candidate.Data[index] = candidate
	return Response{Success: true, Message: "candidate updated successfully"}
}

/*
 * optimizationCandidateSearchIndex finds a candidate's array index by number.
 * Purpose: Pick the fastest search method (binary or sequential) based on current settings.
 * Flow: If sorted by ID, use binary search -> otherwise use sequential search -> return index or -1.
 */
func optimizationCandidateSearchIndex(candidateNumber int, candidateSetting model.DataSetting) int {
	var index = -1
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "id" {
		index = SearchBinaryCandidateByCandidateNumber(candidateNumber, candidateSetting.SortOrder)
	} else {
		index = SearchSequentialCandidateByCandidateNumber(candidateNumber)
	}

	return index
}

/*
 * ShowCandidateById retrieves a candidate by candidate number.
 * Purpose: Return a single candidate record for display or editing.
 * Flow: Find index by ID -> return error if not found -> return candidate data in response.
 */
func ShowCandidateById(id int, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateSearchIndex(id, candidateSetting)
	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	return Response{Success: true, Message: "candidate found", Data: model.Candidate.Data[index]}
}

/*
 * ShowCandidateByName retrieves a candidate by name.
 * Purpose: Return a single candidate record when searching by name.
 * Flow: Find index by name -> return error if not found -> return candidate data in response.
 */
func ShowCandidateByName(name string, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateNameSearchIndex(name, candidateSetting)
	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	return Response{Success: true, Message: "candidate found", Data: model.Candidate.Data[index]}
}

/*
 * optimizationCandidateNameSearchIndex finds a candidate's array index by name.
 * Purpose: Pick the fastest search method (binary or sequential) based on current settings.
 * Flow: If sorted by name with binary search enabled, use binary search -> otherwise use sequential search -> return index or -1.
 */
func optimizationCandidateNameSearchIndex(name string, candidateSetting model.DataSetting) int {
	var index = -1
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "name" && candidateSetting.SearchSetting == "binary" {
		index = SearchBinaryCandidateByName(name, candidateSetting.SortOrder)
	} else {
		index = SearchSequentialCandidateByName(name)
	}

	return index
}
