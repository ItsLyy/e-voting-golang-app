package controller

import "e-voting/model"

func MaxIndexVoter(voters model.Voters, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if voters[i].VoterId > voters[maxIndex].VoterId {
				maxIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voters[i].CandidateNumber > voters[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voters[i].Name > voters[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

func IsVoterGreater(firstVoter model.Voter, secondVoter model.Voter, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstVoter.VoterId > secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber > secondVoter.CandidateNumber
	default:
		return firstVoter.Name > secondVoter.Name
	}
}

func MaxIndexCandidate(candidates model.Canditates, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if candidates[i].CandidateNumber > candidates[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidates[i].Name > candidates[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

func IsCandidateGreater(firstCandidate model.Canditate, secondCandidate model.Canditate, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstCandidate.CandidateNumber > secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name > secondCandidate.Name
	}
}