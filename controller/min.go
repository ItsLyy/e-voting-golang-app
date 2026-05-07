package controller

import "e-voting/model"

func MinVoterIdx(voters model.Voters, N int, comparedBy string)int{
	var i int = 1
	var MinIdx int = 0

	switch comparedBy {
	case "created" :
		for i <= N {
			if voters[MinIdx].VoterId > voters[i].VoterId {
				MinIdx = i
			}
			i++
		}
	case "candidate" :
		for i <= N {
			if voters[MinIdx].CandidateNumber > voters[i].CandidateNumber {
				MinIdx = i
			}
			i++
		}
	default : 
		for i <= N {
			if voters[MinIdx].Name > voters[i].Name {
				MinIdx = i
			}
			i++
		}
	}
	return MinIdx
}