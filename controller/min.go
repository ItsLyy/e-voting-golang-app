package controller

import "e-voting/model"

func MinIndexVoter(voters model.Voters, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if voters[i].VoterId < voters[minIndex].VoterId {
				minIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voters[i].CandidateNumber < voters[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voters[i].Name < voters[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

func IsVoterLess(firstVoter model.Voter, secondVoter model.Voter, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstVoter.VoterId < secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber < secondVoter.CandidateNumber
	default:
		return firstVoter.Name < secondVoter.Name
	}
}

func MinIndexCandidate(candidates model.Canditates, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "created":
		for i <= N {
			if candidates[i].CandidateNumber < candidates[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidates[i].Name < candidates[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

func IsCandidateLess(firstCandidate model.Canditate, secondCandidate model.Canditate, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstCandidate.CandidateNumber < secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name < secondCandidate.Name
	}
}
