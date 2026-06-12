package controller

import (
	"e-voting/model"
)

func CalculatePercentage(total int, currentVote int) float64 {
	if currentVote > 0 || total > 0 {
		return float64(currentVote) / float64(total) * 100
	} else {
		return 0
	}
}

func SumVotes(election model.Elections) int {
	var total int = 0
	var i int
	for i = 0; i < model.Candidate.Length; i++ {
		total += election.Data[i].TotalVote
	}
	return total
}

func ResetElection() {
	var i int
	for i = 0; i < model.Voter.Length; i++ {
		model.Voter.Data[i].CandidateNumber = 0
	}
}

func InputElection() model.Elections {
	var i int
	var election model.Elections

	for i = 0; i < model.Candidate.Length; i++ {
		election.Data[i] = model.ElectionData{
			CandidateIndex: i,
			TotalVote:      0,
		}
	}

	for i = 0; i < model.Voter.Length; i++ {
		var j int
		for j = 0; j < model.Candidate.Length; j++ {
			if model.Voter.Data[i].CandidateNumber == model.Candidate.Data[j].CandidateNumber {
				election.Data[j].TotalVote++
			}
		}
	}

	return election
}
