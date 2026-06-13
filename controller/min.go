package controller

import "e-voting/model"

/*
 * MinIndexVoter finds the index of the smallest voter in range [0..N].
 * Purpose: Used by selection sort to place the smallest element at the end when sorting descending.
 * Flow: Start at index 0 -> loop up to N -> compare by ID, candidate, or name -> return index of minimum.
 */
func MinIndexVoter(voter model.Voters, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if voter.Data[i].VoterId < voter.Data[minIndex].VoterId {
				minIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voter.Data[i].CandidateNumber < voter.Data[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voter.Data[i].Name < voter.Data[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

/*
 * IsVoterLess checks if the first voter is smaller than the second.
 * Purpose: Comparison helper for insertion sort and sorted insertion when adding voters.
 * Flow: Compare by ID, candidate number, or name -> return true if first is less than second.
 */
func IsVoterLess(firstVoter model.VoterData, secondVoter model.VoterData, comparedBy string) bool {
	switch comparedBy {
	case "id":
		return firstVoter.VoterId < secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber < secondVoter.CandidateNumber
	default:
		return firstVoter.Name < secondVoter.Name
	}
}

/*
 * MinIndexCandidate finds the index of the smallest candidate in range [0..N].
 * Purpose: Used by selection sort to place the smallest element at the end when sorting descending.
 * Flow: Start at index 0 -> loop up to N -> compare by ID or name -> return index of minimum.
 */
func MinIndexCandidate(candidate model.Candidates, N int, comparedBy string) int {
	var i int = 1
	var minIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if candidate.Data[i].CandidateNumber < candidate.Data[minIndex].CandidateNumber {
				minIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidate.Data[i].Name < candidate.Data[minIndex].Name {
				minIndex = i
			}
			i++
		}
	}

	return minIndex
}

/*
 * IsCandidateLess checks if the first candidate is smaller than the second.
 * Purpose: Comparison helper for insertion sort and sorted insertion when adding candidates.
 * Flow: Compare by ID or name -> return true if first is less than second.
 */
func IsCandidateLess(firstCandidate model.CandidateData, secondCandidate model.CandidateData, comparedBy string) bool {
	switch comparedBy {
	case "created":
		return firstCandidate.CandidateNumber < secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name < secondCandidate.Name
	}
}
