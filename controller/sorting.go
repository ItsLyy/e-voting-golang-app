package controller

import (
	"e-voting/model"
	"fmt"
)

func InsertionCandidateSorting(sort string, sortBy string) {
	var candidate model.Candidates = model.Candidate
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
