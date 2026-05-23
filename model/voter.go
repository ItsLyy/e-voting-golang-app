package model

const NMAXVOTER = 999

type VoterData struct {
	VoterId, CandidateNumber int
	Name                     string
}

type Voters struct {
	Data   [NMAXVOTER]VoterData
	Length int
}

var Voter Voters

func GenerateVotersData() Voters {
	var data Voters = Voters{
		Data: [NMAXVOTER]VoterData{
			{VoterId: 1, Name: "Hammam", CandidateNumber: 02},
			{VoterId: 2, Name: "Ali", CandidateNumber: 01},
			{VoterId: 3, Name: "Budi", CandidateNumber: 02},
			{VoterId: 4, Name: "Citra", CandidateNumber: 03},
			{VoterId: 5, Name: "Dewi", CandidateNumber: 01},
			{VoterId: 6, Name: "Eka", CandidateNumber: 02},
			{VoterId: 7, Name: "Fajar", CandidateNumber: 03},
			{VoterId: 8, Name: "Gina", CandidateNumber: 01},
			{VoterId: 9, Name: "Hadi", CandidateNumber: 02},
			{VoterId: 10, Name: "Indra", CandidateNumber: 03},
			{VoterId: 11, Name: "Joko", CandidateNumber: 01},
			{VoterId: 12, Name: "Kiki", CandidateNumber: 02},
			{VoterId: 13, Name: "Lina", CandidateNumber: 03},
			{VoterId: 14, Name: "Maya", CandidateNumber: 01},
			{VoterId: 15, Name: "Nanda", CandidateNumber: 02},
			{VoterId: 16, Name: "Oki", CandidateNumber: 03},
			{VoterId: 17, Name: "Putri", CandidateNumber: 01},
			{VoterId: 18, Name: "Rian", CandidateNumber: 02},
			{VoterId: 19, Name: "Sari", CandidateNumber: 03},
			{VoterId: 20, Name: "Tono", CandidateNumber: 01},
		},
		Length: 20,
	}
	return data
}
