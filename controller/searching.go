package controller

import "e-voting/model"

/* Sequential Search Section */

/*
 * SearchSequentialCandidateByName finds a candidate by name using linear search.
 * Purpose: Look up a candidate when data is not sorted for binary search.
 * Flow: Loop through each candidate -> compare name -> return index when found, or -1 if not found.
 */
func SearchSequentialCandidateByName(name string) int {
	var index int = -1
	var i int = 0

	for i < model.Candidate.Length && index == -1 {
		if model.Candidate.Data[i].Name == name {
			index = i
		}

		i++
	}

	return index
}

/*
 * SearchSequentialCandidateByCandidateNumber finds a candidate by number using linear search.
 * Purpose: Look up a candidate by ID when data is not sorted for binary search.
 * Flow: Loop through each candidate -> compare number -> return index when found, or -1 if not found.
 */
func SearchSequentialCandidateByCandidateNumber(candidateNumber int) int {
	var index int = -1
	var i int = 0

	for i < model.Candidate.Length && index == -1 {
		if model.Candidate.Data[i].CandidateNumber == candidateNumber {
			index = i
		}

		i++
	}

	return index
}

/*
 * SearchSequentialVoterByName finds a voter by name using linear search.
 * Purpose: Look up a voter when data is not sorted for binary search.
 * Flow: Loop through each voter -> compare name -> return index when found, or -1 if not found.
 */
func SearchSequentialVoterByName(name string) int {
	var index int = -1
	var i int = 0

	for i < model.Voter.Length && index == -1 {
		if model.Voter.Data[i].Name == name {
			index = i
		}
		i++
	}

	return index
}

/*
 * SearchSequentialVoterById finds a voter by ID using linear search.
 * Purpose: Look up a voter by ID when data is not sorted for binary search.
 * Flow: Loop through each voter -> compare ID -> return index when found, or -1 if not found.
 */
func SearchSequentialVoterById(id int) int {
	var index int = -1
	var i int = 0

	for i < model.Voter.Length && index == -1 {
		if model.Voter.Data[i].VoterId == id {
			index = i
		}
		i++
	}

	return index
}

/* Binary Search Section (data must be sorted first) */

/*
 * SearchBinaryCandidateByName finds a candidate by name using binary search.
 * Purpose: Faster lookup when candidates are sorted by name.
 * Flow: Set left/right bounds -> compare middle element -> narrow search range -> return index when found, or -1 if not found.
 */
func SearchBinaryCandidateByName(name string, sort string) int {
	var left int = 0
	var right int = model.Candidate.Length - 1
	var index int = -1

	switch sort {
	case "desc":
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Candidate.Data[mid].Name > name {
				left = mid + 1
			} else if model.Candidate.Data[mid].Name < name {
				right = mid - 1
			} else {
				index = mid
			}
		}
	default:
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Candidate.Data[mid].Name < name {
				left = mid + 1
			} else if model.Candidate.Data[mid].Name > name {
				right = mid - 1
			} else {
				index = mid
			}
		}
	}

	return index
}

/*
 * SearchBinaryCandidateByCandidateNumber finds a candidate by number using binary search.
 * Purpose: Faster lookup when candidates are sorted by candidate number.
 * Flow: Set left/right bounds -> compare middle element -> narrow search range -> return index when found, or -1 if not found.
 */
func SearchBinaryCandidateByCandidateNumber(candidateNumber int, sort string) int {
	var left int = 0
	var right int = model.Candidate.Length - 1
	var index int = -1

	switch sort {
	case "desc":
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Candidate.Data[mid].CandidateNumber > candidateNumber {
				left = mid + 1
			} else if model.Candidate.Data[mid].CandidateNumber < candidateNumber {
				right = mid - 1
			} else {
				index = mid
			}
		}
	default:
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Candidate.Data[mid].CandidateNumber < candidateNumber {
				left = mid + 1
			} else if model.Candidate.Data[mid].CandidateNumber > candidateNumber {
				right = mid - 1
			} else {
				index = mid
			}
		}
	}

	return index
}

/*
 * SearchBinaryVoterByName finds a voter by name using binary search.
 * Purpose: Faster lookup when voters are sorted by name.
 * Flow: Set left/right bounds -> compare middle element -> narrow search range -> return index when found, or -1 if not found.
 */
func SearchBinaryVoterByName(name string, sort string) int {
	var left int = 0
	var right int = model.Voter.Length - 1
	var index int = -1

	switch sort {
	case "desc":
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Voter.Data[mid].Name > name {
				left = mid + 1
			} else if model.Voter.Data[mid].Name < name {
				right = mid - 1
			} else {
				index = mid
			}
		}
	default:
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Voter.Data[mid].Name < name {
				left = mid + 1
			} else if model.Voter.Data[mid].Name > name {
				right = mid - 1
			} else {
				index = mid
			}
		}
	}

	return index
}

/*
 * SearchBinaryVoterById finds a voter by ID using binary search.
 * Purpose: Faster lookup when voters are sorted by voter ID.
 * Flow: Set left/right bounds -> compare middle element -> narrow search range -> return index when found, or -1 if not found.
 */
func SearchBinaryVoterById(id int, sort string) int {
	var left int = 0
	var right int = model.Voter.Length - 1
	var index int = -1

	switch sort {
	case "desc":
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Voter.Data[mid].VoterId > id {
				left = mid + 1
			} else if model.Voter.Data[mid].VoterId < id {
				right = mid - 1
			} else {
				index = mid
			}
		}
	default:
		for left <= right && index == -1 {
			var mid int = (left + right) / 2
			if model.Voter.Data[mid].VoterId < id {
				left = mid + 1
			} else if model.Voter.Data[mid].VoterId > id {
				right = mid - 1
			} else {
				index = mid
			}
		}
	}

	return index
}
