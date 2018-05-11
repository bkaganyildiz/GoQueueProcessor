package main

import (
	"github.com/bkaganyildiz/GoQueueProcessor/token"
	"fmt"
)

func main() {
	var opsGenieToken = token.InitialTokenGenerator()
	fmt.Print(opsGenieToken.Token())
	// start long polling
	// before receiving from queue generateNewToken if necessary
	var generatedToken, err = token.IsValidClient(opsGenieToken)

	if err != nil {
		panic(err)
	} else {
		opsGenieToken = generatedToken
		fmt.Print(opsGenieToken.Token())
		// process queue
	}
}
