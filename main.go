package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"authsecure_go/authsecure"
)

func main() {

	AuthsecureApp := authsecure.AuthSecure{}
	AuthsecureApp.Api(
		"XD",
		"3ezshCmkXrn",
		"7a8bfeb28afcd690812ee5de010a6860",
		"1.0",
	)
	AuthsecureApp.Init()

	fmt.Println("\n[1] Login\n[2] Register\n[3] License Login\n[4] Exit")
	fmt.Print("Choose option: ")

	var choice string
	fmt.Scan(&choice)

	reader := bufio.NewReader(os.Stdin)
	// 🧹 FIX: Clear leftover newline from fmt.Scan()
	reader.ReadString('\n')

	switch choice {
	case "1":
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		AuthsecureApp.Login(strings.TrimSpace(username), strings.TrimSpace(password))

	case "2":
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		fmt.Print("License: ")
		license, _ := reader.ReadString('\n')
		AuthsecureApp.Register(strings.TrimSpace(username), strings.TrimSpace(password), strings.TrimSpace(license))

	case "3":
		fmt.Print("License: ")
		license, _ := reader.ReadString('\n')
		AuthsecureApp.License(strings.TrimSpace(license))

	default:
		fmt.Println("Goodbye!")
	}
}
