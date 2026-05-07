package controller

import "e-voting/model"

func MaxVoteridx(voters model.Voters, N int, comparedBy string)int {
	var i int = 1
	var MaxIdx int = 0
	
	switch comparedBy {
	case "created" :
		for i <= N {
			if voters[MaxIdx].VoterId < voters[i].VoterId {
				MaxIdx = i
			}
			i++
		}
	case "candidate" : 
		for i <= N {
			if voters[MaxIdx].CandidateNumber < voters[i].CandidateNumber{
				MaxIdx = i
			}
			i++
		}
	default :
		for i <= N {
			if voters[MaxIdx].Name < voters[i].Name {
				MaxIdx = i
			}
		}
	 
		
	}
	return  MaxIdx
}

