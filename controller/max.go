package controller

import "e-voting/model"

/*
 * MaxIndexVoter finds the index of the largest voter in range [0..N].
 * Purpose: Used by selection sort to place the largest element at the end when sorting ascending.
 * Flow: Start at index 0 -> loop up to N -> compare by ID, candidate, or name -> return index of maximum.
 */
func MaxIndexVoter(voter model.Voters, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if voter.Data[i].VoterId > voter.Data[maxIndex].VoterId {
				maxIndex = i
			}
			i++
		}
	case "candidate":
		for i <= N {
			if voter.Data[i].CandidateNumber > voter.Data[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if voter.Data[i].Name > voter.Data[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

/*
 * IsVoterGreater checks if the first voter is larger than the second.
 * Purpose: Comparison helper for insertion sort and sorted insertion when adding voters.
 * Flow: Compare by ID, candidate number, or name -> return true if first is greater than second.
 */
func IsVoterGreater(firstVoter model.VoterData, secondVoter model.VoterData, comparedBy string) bool {
	switch comparedBy {
	case "id":
		return firstVoter.VoterId > secondVoter.VoterId
	case "candidate":
		return firstVoter.CandidateNumber > secondVoter.CandidateNumber
	default:
		return firstVoter.Name > secondVoter.Name
	}
}

/*
 * MaxIndexCandidate finds the index of the largest candidate in range [0..N].
 * Purpose: Used by selection sort to place the largest element at the end when sorting ascending.
 * Flow: Start at index 0 -> loop up to N -> compare by ID or name -> return index of maximum.
 */
func MaxIndexCandidate(candidate model.Candidates, N int, comparedBy string) int {
	var i int = 1
	var maxIndex int = 0

	switch comparedBy {
	case "id":
		for i <= N {
			if candidate.Data[i].CandidateNumber > candidate.Data[maxIndex].CandidateNumber {
				maxIndex = i
			}
			i++
		}
	default:
		for i <= N {
			if candidate.Data[i].Name > candidate.Data[maxIndex].Name {
				maxIndex = i
			}
			i++
		}
	}

	return maxIndex
}

/*
 * IsCandidateGreater checks if the first candidate is larger than the second.
 * Purpose: Comparison helper for insertion sort and sorted insertion when adding candidates.
 * Flow: Compare by ID or name -> return true if first is greater than second.
 */
func IsCandidateGreater(firstCandidate model.CandidateData, secondCandidate model.CandidateData, comparedBy string) bool {
	switch comparedBy {
	case "id":
		return firstCandidate.CandidateNumber > secondCandidate.CandidateNumber
	default:
		return firstCandidate.Name > secondCandidate.Name
	}
}
