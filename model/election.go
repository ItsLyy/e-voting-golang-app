package model

const NMAXELECTION = 10

type ElectionData struct {
	CandidateIndex int
	TotalVote      int
}

type Elections struct {
	Data [NMAXELECTION]ElectionData
}
