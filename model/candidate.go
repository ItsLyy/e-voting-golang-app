package model

const NMAXCANDIDATE = 10

type Canditate struct {
	CandidateNumber int
	Name, VisionMission string
}

type Canditates [NMAXCANDIDATE]Canditate