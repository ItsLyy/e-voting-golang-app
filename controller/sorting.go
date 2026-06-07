package controller

import (
	"e-voting/model"
	"fmt"
)

func InsertionCandidateSorting(sort string, sortBy string) {
	var candidate model.Canditates = model.Candidate
	var temp model.CandidateData
	var i int = 1

	switch sort {
	case "desc":
		for i < candidate.Length {
			var j int = i - 1
			temp = candidate.Data[i]

			for j >= 0 && IsCandidateGreater(temp, candidate.Data[j], sortBy) {
				candidate.Data[j+1] = candidate.Data[j]
				j--
			}

			candidate.Data[j+1] = temp
			i++
		}
	default:
		for i < candidate.Length {
			var j int = i - 1
			temp = candidate.Data[i]

			for j >= 0 && IsCandidateLess(temp, candidate.Data[j], sortBy) {
				candidate.Data[j+1] = candidate.Data[j]
				j--
			}

			candidate.Data[j+1] = temp
			i++
		}
	}

	fmt.Println(candidate)
	model.Candidate = candidate
}

func SelectionVotersSorting(sort string, sortBy string) {
	var voter model.Voters = model.Voter
	var temp model.VoterData
	var i int = voter.Length - 1

	switch sort {
	case "desc":
		for i > 0 {
			var j int = MinIndexVoter(voter, i, sortBy)

			temp = voter.Data[j]
			voter.Data[j] = voter.Data[i]
			voter.Data[i] = temp

			i--
		}
	default:
		for i > 0 {
			var j int = MaxIndexVoter(voter, i, sortBy)

			temp = voter.Data[j]
			voter.Data[j] = voter.Data[i]
			voter.Data[i] = temp

			i--
		}
	}

	model.Voter = voter
}

func InsertionVotersSorting(sort string, sortBy string) {
	var voter model.Voters = model.Voter
	var temp model.VoterData
	var i int = 1

	switch sort {
	case "desc" :
		for i < voter.Length {
			var j int = i - 1
			temp = voter.Data[i]

			for j >= 0 && IsVoterGreater(temp, voter.Data[j], sortBy) {
				voter.Data[j+1] = voter.Data[j]
				j--
			}

			voter.Data[j+1] = temp
			i++ 
		}
	default :
		for i < voter.Length {
			var j int = i - 1
			temp = voter.Data[i]

			for j >= 0 && IsVoterLess(temp, voter.Data[j], sortBy) {
				voter.Data[j+1] = voter.Data[j]
				j--
			}

			voter.Data[j+1] = temp
			i++
		}
	}
	fmt.Println(voter)
	model.Voter = voter

}

func SelectionCandidateSorting(sort string, sortBy string) {
	var candidate model.Canditates = model.Candidate
	var temp model.CandidateData
	var i int = candidate.Length - 1

	switch sort {
	case "desc":
		for i > 0 {
			var j int = MinIndexCandidate(candidate, i, sortBy)

			temp = candidate.Data[j]
			candidate.Data[j] = candidate.Data[i]
			candidate.Data[i] = temp

			i--
		}
	default:
		for i > 0 {
			var j int = MaxIndexCandidate(candidate, i, sortBy)

			temp = candidate.Data[j]
			candidate.Data[j] = candidate.Data[i]
			candidate.Data[i] = temp

			i--
		}
	}

	model.Candidate = candidate
}
