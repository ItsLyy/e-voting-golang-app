package controller

import (
	"e-voting/model"
)

/*
 * CalculatePercentage computes the vote share as a percentage.
 * Purpose: Display each candidate's portion of total votes in election results.
 * Flow: If votes exist, divide currentVote by total and multiply by 100 -> otherwise return 0.
 */
func CalculatePercentage(total int, currentVote int) float64 {
	if currentVote > 0 || total > 0 {
		return float64(currentVote) / float64(total) * 100
	} else {
		return 0
	}
}

/*
 * SumVotes counts all votes across every candidate.
 * Purpose: Get the total vote count needed for percentage calculation.
 * Flow: Loop through each candidate's vote count -> add them up -> return total.
 */
func SumVotes(election model.Elections) int {
	var total int = 0
	var i int
	for i = 0; i < model.Candidate.Length; i++ {
		total += election.Data[i].TotalVote
	}
	return total
}

/*
 * ResetElection clears all voter choices.
 * Purpose: Start a new election by removing every voter's selected candidate.
 * Flow: Loop through all voters -> set each CandidateNumber to 0.
 */
func ResetElection() {
	var i int
	for i = 0; i < model.Voter.Length; i++ {
		model.Voter.Data[i].CandidateNumber = 0
	}
}

/*
 * InputElection tallies votes from voter data into election results.
 * Purpose: Build the vote count per candidate for displaying results.
 * Flow: Initialize vote slots per candidate -> loop voters -> match voter choice to candidate -> increment vote count -> return Elections.
 */
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
