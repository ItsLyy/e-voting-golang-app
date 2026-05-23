package main

import (
	// "e-voting/controller"
	// "e-voting/model"
	"e-voting/model"
	"e-voting/view"
	// "fmt"
)

func main() {
	// var candidates model.Canditates = model.GenerateCandidatesData()
	// var voters model.Voters = model.GenerateVotersData()

	// fmt.Println(controller.InsertionCandidateSorting(candidates, 10, "asc", "name"))
	// fmt.Println(controller.SelectionVotersSorting(voters, 20, "desc", "name"))
	model.Candidate = model.GenerateCandidatesData()
	view.ManageCandidates()
}
