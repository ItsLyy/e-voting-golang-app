package controller

import "e-voting/model"

func MaxIndexVoter(voter model.Voters, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if voter.Data[i].VoterId > voter.Data[maxIndex].VoterId {
				maxIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voter.Data[i].CandidateNumber > voter.Data[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voter.Data[i].Name > voter.Data[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

func IsVoterGreater(firstVoter model.VoterData, secondVoter model.VoterData, comparedBy string) bool {
	switch comparedBy {
	case "id":
		return firstVoter.VoterId > secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber > secondVoter.CandidateNumber
	default:
		return firstVoter.Name > secondVoter.Name
	}
}

func MaxIndexCandidate(candidate model.Candidates, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if candidate.Data[i].CandidateNumber > candidate.Data[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidate.Data[i].Name > candidate.Data[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

func IsCandidateGreater(firstCandidate model.CandidateData, secondCandidate model.CandidateData, comparedBy string) bool {
	switch comparedBy {
	case "id":
		return firstCandidate.CandidateNumber > secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name > secondCandidate.Name
	}
}
