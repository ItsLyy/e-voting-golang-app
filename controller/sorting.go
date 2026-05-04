package controller

import (
	"e-voting/model"
)

func InsertingCandidateSortingByPointAsc(candidates model.Canditates, N int, sort string, sortBy string) model.Canditates {
	var sortedCandidates model.Canditates = candidates
	var temp model.Canditate
	var i int = 1

	switch sort {
	case "desc":
		for i < N {
			var j int = i
			for j >= 1 && IsCandidateGreater(candidates[j], candidates[j-1], sortBy) {
				temp=sortedCandidates[j]
				sortedCandidates[j] = sortedCandidates[j-1]
				sortedCandidates[j-1] = temp
				j--
			}
			i++
		}
	default:
		for i < N {
			var j int = i
			for j >= 1 && IsCandidateLess(candidates[j], candidates[j-1], sortBy) {
				temp=sortedCandidates[j]
				sortedCandidates[j] = sortedCandidates[j-1]
				sortedCandidates[j-1] = temp
				j--
			}
			i++
		}
	}

	return sortedCandidates
}

func SelectionVotersSortingByName(voters model.Voters, N int, sort string, sortBy string) model.Voters {
	var sortedVoters model.Voters = voters
	var temp model.Voter
	var i int = N - 1

	switch sort {
	case "desc":
		for i > 0 {
			var j int = MinIndexVoter(sortedVoters, i, sortBy)
	
			temp = sortedVoters[j]
			sortedVoters[j] = sortedVoters[i]
			sortedVoters[i] = temp
	
			i--
		}
	default:
		for i > 0 {
			var j int = MaxIndexVoter(sortedVoters, i, sortBy)
	
			temp = sortedVoters[j]
			sortedVoters[j] = sortedVoters[i]
			sortedVoters[i] = temp
	
			i--
		}
	}

	return sortedVoters
}