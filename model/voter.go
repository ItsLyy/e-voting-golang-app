package model

const NMAXVOTER = 999

type Voter struct {
	VoterId int
	Name, CandidateNumber string
}

type Voters [NMAXVOTER]Voter
