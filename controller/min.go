package controller

import "e-voting/model"

func MinIndexVoter(voter model.Voters, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if voter.Data[i].VoterId < voter.Data[minIndex].VoterId {
				minIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voter.Data[i].CandidateNumber < voter.Data[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voter.Data[i].Name < voter.Data[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

func IsVoterLess(firstVoter model.VoterData, secondVoter model.VoterData, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstVoter.VoterId < secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber < secondVoter.CandidateNumber
	default:
		return firstVoter.Name < secondVoter.Name
	}
}

func MinIndexCandidate(candidate model.Canditates, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if candidate.Data[i].CandidateNumber < candidate.Data[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidate.Data[i].Name < candidate.Data[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

func IsCandidateLess(firstCandidate model.CandidateData, secondCandidate model.CandidateData, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstCandidate.CandidateNumber < secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name < secondCandidate.Name
	}
}
