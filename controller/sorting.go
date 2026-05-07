package controller 

import  "e-voting/model"

func SelectionSortingVoters(voters model.Voters, N int, sort string, sortBy string)model.Voters {
	var apajalah = voters
	var temp model.Voter 
	var i int = N-1

	switch sort {
	case "increasing" :
		for i > 0 {
			var j int = MaxVoteridx(apajalah, i, sortBy)
			temp = apajalah[i]
			apajalah[j] = apajalah[i]
			apajalah[i] = temp 

			i--
		}
	default : 
		for i > 0 {
			var j int = MinVoterIdx(apajalah,i,sortBy)
			temp = apajalah[j]
			apajalah[j] = apajalah[i]
			apajalah[i] = temp

			i--
		}
	}
	return apajalah
}