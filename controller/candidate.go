package controller

import "e-voting/model"

/*
 * CreateCandidate adds a new candidate to the list.
 * Purpose: Handle candidate creation from the UI with validation and sorted insertion.
 * Flow: Validate input -> check duplicate via search -> insert at correct position based on sort settings -> increment length -> return response.
 */
func CreateCandidate(candidate model.CandidateData, candidateSetting model.DataSetting) Response {
	// Ensure the candidate has a valid ID and name
	// before processing the creation request.
	if candidate.CandidateNumber == 0 || candidate.Name == "" {
		return Response{Success: false, Message: "invalid candidate data"}
	}

	// Check whether a candidate with the same ID
	// already exists in the system.
	var index = optimizationCandidateSearchIndex(candidate.CandidateNumber, candidateSetting)

	if index != -1 {
		return Response{Success: false, Message: "candidate already exists"}
	}

	// Start from the last occupied position and
	// locate the correct insertion point.
	var i int = model.Candidate.Length - 1

	switch candidateSetting.SortOrder {
	case "asc":
		switch candidateSetting.SortBy {
		case "id":
			// Shift existing candidates one position to the right
			// until the correct ascending position is found.
			for i >= 0 && IsCandidateLess(candidate, model.Candidate.Data[i], "id") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			// Insert the new candidate into its sorted position.
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
			// Shift candidates to maintain descending order
			// before inserting the new candidate.
			for i >= 0 && IsCandidateGreater(candidate, model.Candidate.Data[i], "id") {
				model.Candidate.Data[i] = model.Candidate.Data[i-1]
				i--
			}
			// Insert the new candidate into its sorted position.
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
		// If no sorting configuration is applied,
		// append the candidate to the end of the array.
		model.Candidate.Data[model.Candidate.Length] = candidate
	}

	// Increase the total number of stored candidates.
	model.Candidate.Length++
	// Controller successfully inserts the candidate
	// and updates the Model layer.
	return Response{Success: true, Message: "candidate created successfully"}
}

/*
 * DeleteCandidate removes a candidate by candidate number.
 * Purpose: Remove an existing candidate from the data.
 * Flow: Find index by ID -> return error if not found -> shift remaining entries left -> decrement length -> return response.
 */
func DeleteCandidate(id int, candidateSetting model.DataSetting) Response {
	// Locate the candidate's index before attempting deletion.
	var index = optimizationCandidateSearchIndex(id, candidateSetting)

	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	// Clear the candidate data at the target index.
	model.Candidate.Data[index] = model.CandidateData{}

	// Shift all subsequent candidates one position left
	// to fill the gap created by the deletion.
	for index < model.Candidate.Length-1 {
		model.Candidate.Data[index] = model.Candidate.Data[index+1]
		index++
	}

	// Decrease the number of active candidates.
	model.Candidate.Length--

	return Response{Success: true, Message: "candidate deleted successfully"}
}

/*
 * UpdateCandidate replaces an existing candidate's data.
 * Purpose: Allow editing a candidate's name or vision mission.
 * Flow: Find index by ID -> return error if not found -> overwrite data at index -> return response.
 */
func UpdateCandidate(id int, candidate model.CandidateData, candidateSetting model.DataSetting) Response {
	// Find the candidate record that will be updated.
	var index = optimizationCandidateSearchIndex(id, candidateSetting)

	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	// Replace the existing candidate record
	// with the updated data provided by the user.
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
	// Select the most efficient search algorithm
	// based on the current sorting configuration.
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "id" {
		// Binary Search is used when the data
		// is sorted by candidate ID.
		index = SearchBinaryCandidateByCandidateNumber(candidateNumber, candidateSetting.SortOrder)
	} else {
		// Sequential Search is used when the data
		// is not sorted by candidate ID.
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
	// Retrieve the candidate's position in the array
	// before returning the record.
	var index = optimizationCandidateSearchIndex(id, candidateSetting)
	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}
	// Return the requested candidate data
	// to the View layer for display.
	return Response{Success: true, Message: "candidate found", Data: model.Candidate.Data[index]}
}

/*
 * ShowCandidateByName retrieves a candidate by name.
 * Purpose: Return a single candidate record when searching by name.
 * Flow: Find index by name -> return error if not found -> return candidate data in response.
 */
func ShowCandidateByName(name string, candidateSetting model.DataSetting) Response {
	// Search for the candidate using their name.
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
	// Choose between Binary Search and Sequential Search
	// depending on the current sort and search settings.
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "name" && candidateSetting.SearchSetting == "binary" {
		// Binary Search can only be applied when
		// the data is sorted by candidate name.
		index = SearchBinaryCandidateByName(name, candidateSetting.SortOrder)
	} else {
		// Fall back to Sequential Search when
		// Binary Search requirements are not met.
		index = SearchSequentialCandidateByName(name)
	}

	return index
}


// The Controller receives requests from the View,
// processes the required business logic, and
// updates the Model accordingly.