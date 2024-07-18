package main

import (
	"fmt"
	"yaro/auth"
)

//TODO: Argument parsing, config handling,

func main() {
	loginCode := auth.RunLoginWindow()

	authDetails := auth.GetSessionToken(loginCode, auth.Verifier)

	fmt.Println("Access token: ", authDetails.AccessToken)
	fmt.Println("Refresh token: ", authDetails.RefreshToken)
}
