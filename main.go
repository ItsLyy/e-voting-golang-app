package main

import (
	"e-voting/controller"
	"e-voting/model"
)

func main() {
	model.Candidate = model.GenerateCandidatesData()
	model.Voter = model.GenerateVotersData()
	controller.InsertionCandidateSorting("desc", "name")
}
