package main

import (
	"e-voting/controller"
	"e-voting/model"
	"fmt"
)

func main() {
	var candidates model.Canditates = model.GenerateCandidatesData()
	var voters model.Voters = model.GenerateVotersData()

	fmt.Println(controller.InsertionCandidateSorting(candidates, 10, "asc", "name"))
	fmt.Println(controller.SelectionVotersSorting(voters, 20, "desc", "name"))
}
