package controller

import "e-voting/model"

/**
 * Sequential Search Section
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

/**
 * Binary Search Section
 * Note: Must be sorted before searching
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
