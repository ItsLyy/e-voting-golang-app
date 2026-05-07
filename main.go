package main

import (
	"e-voting/model"
	"e-voting/controller"
	"fmt"
)


func main() {
	 var voters model.Voters = model.GenerateData()

	fmt.Println(controller.SelectionSortingVoters(voters, 20,"decreasing","pusing"))
	fmt.Println(controller.SelectionSortingVoters(voters, 20,"increasing","singlet"))
}
	