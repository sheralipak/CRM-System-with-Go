package main

import (
	"fmt"
	"os"
)

type User struct {
	username string
	password string
}

type Search struct {
	user    User
	query   string
	results []string
}

var searches []Search

func main() {

	var mode string

	fmt.Println("Welcome! Please log in or sign up.")
	fmt.Println("1. Log in")
	fmt.Println("2. Sign up")
	fmt.Print("Enter choice: ")
	fmt.Scan(&mode)

	switch mode {
	case "1":

		login()
		break
	case "2":
		signup()
		break
	default:
		fmt.Println("Invalid choice. Exiting.")
		os.Exit(1)

	}

	for {

		fmt.Println("1. Search")
		fmt.Println("2. Log out")
		fmt.Print("Enter choice: ")
		fmt.Scan(&mode)

		switch mode {
		case "1":

			search()
			break
		case "2":

			logout()
			break
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func login() {
	var username, password string
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

}

func signup() {
	var username, password string
	fmt.Print("Enter a username: ")
	fmt.Scan(&username)
	fmt.Print("Enter a password: ")
	fmt.Scan(&password)

}

func search() {
	var query string
	fmt.Print("Enter your search query: ")
	fmt.Scan(&query)

}

func logout() {
	fmt.Println("Entering admin mode.")
	// handle admin mode
	for {
		var mode string
		// prompt for admin options
		fmt.Println("1. View all searches")
		fmt.Println("2. View searches by user")
		fmt.Println("3. Log out")
		fmt.Print("Enter choice: ")
		fmt.Scan(&mode)

		switch mode {
		case "1":
			// handle view all searches
			viewSearches()
			break
		case "2":
			// handle view searches by user
			viewUserSearches()
			break
		case "3":
			// handle admin logout
			fmt.Println("Exiting.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func viewSearches() {
	// handle view all searches
	for _, search := range searches {
		fmt.Println("User:", search.user.username)
		fmt.Println("Query:", search.query)
		fmt.Println("Results:", search.results)
	}
}

func viewUserSearches() {
	var username string
	fmt.Print("Enter a username: ")
	fmt.Scan(&username)
	// handle view searches by user
	for _, search := range searches {
		if search.user.username == username {
			fmt.Println("User:", search.user.username)
			fmt.Println("Query:", search.query)
			fmt.Println("Results:", search.results)
		}
	}
}
