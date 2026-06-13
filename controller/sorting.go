package controller

import (
	"e-voting/model"
	"fmt"
)

/*
 * InsertionCandidateSorting sorts candidates using insertion sort.
 * Purpose: Reorder the candidate list by ID or name in ascending or descending order.
 * Flow: Copy candidate data -> for each element, shift larger/smaller items and insert -> save back to model.
 */
func InsertionCandidateSorting(sort string, sortBy string) {
	var candidate model.Candidates = model.Candidate
	var temp model.CandidateData
	// Start from the second element because the first element
	// is considered a sorted subarray by itself.
	var i int = 1

	switch sort {
	case "desc":
		for i < candidate.Length {
			var j int = i - 1
			// Store the current candidate that will be inserted
			// into its correct position in the sorted portion.
			temp = candidate.Data[i]

			// Shift elements one position to the right while
			// they are greater than the current element (descending order).
			for j >= 0 && IsCandidateGreater(temp, candidate.Data[j], sortBy) {
				candidate.Data[j+1] = candidate.Data[j]
				j--
			}

			candidate.Data[j+1] = temp
			i++
		}
	default:
		for i < candidate.Length {
			var j int = i - 1
			temp = candidate.Data[i]

			// Shift elements one position to the right while
			// they are smaller than the current element (ascending order).
			for j >= 0 && IsCandidateLess(temp, candidate.Data[j], sortBy) {
				candidate.Data[j+1] = candidate.Data[j]
				j--
			}

			// Insert the candidate into its correct sorted position.
			candidate.Data[j+1] = temp
			i++
		}
	}

	fmt.Println(candidate)
	model.Candidate = candidate
}

/*
 * SelectionVotersSorting sorts voters using selection sort.
 * Purpose: Reorder the voter list by ID, name, or chosen candidate in ascending or descending order.
 * Flow: Copy voter data -> find min/max index in unsorted portion -> swap with end position -> repeat -> save back to model.
 */
func SelectionVotersSorting(sort string, sortBy string) {
	var voter model.Voters = model.Voter
	var temp model.VoterData
	// Start from the last index and gradually build
	// the sorted portion of the array from right to left.
	var i int = voter.Length - 1

	switch sort {
	case "desc":
		for i > 0 {
			// Find the smallest element within the unsorted portion.
			var j int = MinIndexVoter(voter, i, sortBy)
			
			// Swap the selected element with the current end position.
			temp = voter.Data[j]
			voter.Data[j] = voter.Data[i]
			voter.Data[i] = temp

			i--
		}
	default:
		for i > 0 {
			// Find the largest element within the unsorted portion.
			var j int = MaxIndexVoter(voter, i, sortBy)

			temp = voter.Data[j]
			voter.Data[j] = voter.Data[i]
			voter.Data[i] = temp

			i--
		}
	}
	// Store the sorted voter data back into the Model layer.
	model.Voter = voter
}

/*
 * InsertionVotersSorting sorts voters using insertion sort.
 * Purpose: Reorder the voter list by ID, name, or chosen candidate in ascending or descending order.
 * Flow: Copy voter data -> for each element, shift larger/smaller items and insert -> save back to model.
 */
func InsertionVotersSorting(sort string, sortBy string) {
	var voter model.Voters = model.Voter
	var temp model.VoterData
	// Begin insertion sort from the second voter record.
	var i int = 1

	switch sort {
	case "desc":
		for i < voter.Length {
			var j int = i - 1
			// Save the current voter temporarily while
			// searching for the correct insertion position.
			temp = voter.Data[i]

			// Shift larger-priority elements to the right
			// until the correct descending position is found.
			for j >= 0 && IsVoterGreater(temp, voter.Data[j], sortBy) {
				voter.Data[j+1] = voter.Data[j]
				j--
			}

			// Place the voter record in its correct sorted position.
			voter.Data[j+1] = temp
			i++
		}
	default:
		for i < voter.Length {
			var j int = i - 1
			temp = voter.Data[i]
			
			// Shift smaller-priority elements to the right
			// until the correct ascending position is found.
			for j >= 0 && IsVoterLess(temp, voter.Data[j], sortBy) {
				voter.Data[j+1] = voter.Data[j]
				j--
			}

			// Place the voter record in its correct sorted position.
			voter.Data[j+1] = temp
			i++
		}
	}
	fmt.Println(voter)
	// Update the Model with the newly sorted voter list.
	model.Voter = voter

}

/*
 * SelectionCandidateSorting sorts candidates using selection sort.
 * Purpose: Reorder the candidate list by ID or name in ascending or descending order.
 * Flow: Copy candidate data -> find min/max index in unsorted portion -> swap with end position -> repeat -> save back to model.
 */
func SelectionCandidateSorting(sort string, sortBy string) {
	var candidate model.Candidates = model.Candidate
	var temp model.CandidateData
	// Start from the last candidate and repeatedly place
	// the correct element into its final sorted position.
	var i int = candidate.Length - 1

	switch sort {
	case "desc":
		for i > 0 {
			// Locate the smallest candidate in the unsorted section.
			var j int = MinIndexCandidate(candidate, i, sortBy)

			// Exchange the selected candidate with the candidate
			// currently positioned at the end of the unsorted section.
			temp = candidate.Data[j]
			candidate.Data[j] = candidate.Data[i]
			candidate.Data[i] = temp

			i--
		}
	default:
		for i > 0 {
			// Locate the largest candidate in the unsorted section.
			var j int = MaxIndexCandidate(candidate, i, sortBy)

			// Exchange the selected candidate with the candidate
			// currently positioned at the end of the unsorted section.
			temp = candidate.Data[j]
			candidate.Data[j] = candidate.Data[i]
			candidate.Data[i] = temp

			i--
		}
	}
	// Save the sorted candidate list back into the Model layer.
	model.Candidate = candidate
}

// Controller processes the sorting logic and then
// updates the Model with the sorted result.

// Controller is responsible for data manipulation,
// while the View only displays the sorted data.