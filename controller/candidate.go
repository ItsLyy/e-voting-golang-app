package controller

import "e-voting/model"

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

func UpdateCandidate(id int, candidate model.CandidateData, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateSearchIndex(id, candidateSetting)

	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	model.Candidate.Data[index] = candidate
	return Response{Success: true, Message: "candidate updated successfully"}
}

func optimizationCandidateSearchIndex(candidateNumber int, candidateSetting model.DataSetting) int {
	var index = -1
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "id" {
		index = SearchBinaryCandidateByCandidateNumber(candidateNumber, candidateSetting.SortOrder)
	} else {
		index = SearchSequentialCandidateByCandidateNumber(candidateNumber)
	}

	return index
}

func ShowCandidateById(id int, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateSearchIndex(id, candidateSetting)
	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	return Response{Success: true, Message: "candidate found", Data: model.Candidate.Data[index]}
}

func ShowCandidateByName(name string, candidateSetting model.DataSetting) Response {
	var index = optimizationCandidateNameSearchIndex(name, candidateSetting)
	if index == -1 {
		return Response{Success: false, Message: "candidate not found"}
	}

	return Response{Success: true, Message: "candidate found", Data: model.Candidate.Data[index]}
}

func optimizationCandidateNameSearchIndex(name string, candidateSetting model.DataSetting) int {
	var index = -1
	if (candidateSetting.SortOrder == "asc" || candidateSetting.SortOrder == "desc") && candidateSetting.SortBy == "name" {
		index = SearchBinaryVoterByName(name, candidateSetting.SortOrder)
	} else {
		index = SearchSequentialVoterByName(name)
	}

	return index
}