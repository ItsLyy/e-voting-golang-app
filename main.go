package main

import (
	"e-voting/model"
	"e-voting/view"
)

func main() {
	model.Candidate = model.GenerateCandidatesData()
	model.Voter = model.GenerateVotersData()
	view.Home()
}
