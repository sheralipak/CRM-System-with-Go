package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var searches []string // array to store all search queries

func main() {
	var username, password string
	var choice int
	var searchTerm string

	fmt.Println("Welcome to Search App")
	fmt.Println("1. Login")
	fmt.Println("2. Sign up")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		// validate login credentials here

		if username == "admin" {
			adminMenu() // call admin menu if the user is admin
		} else {
			fmt.Println("Login successful.")
			searchMenu(username)
		}

	} else if choice == 2 {
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		// create new account here

		fmt.Println("Sign up successful.")
		searchMenu(username)
	} else {
		fmt.Println("Invalid choice.")
		return
	}
}

func searchMenu(username string) {
	var searchTerm string

	fmt.Print("Enter search term: ")
	fmt.Scan(&searchTerm)

	// perform search here
	resp, _ := http.Get("https://duckduckgo.com/html/?q=" + searchTerm)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	doc, _ := html.Parse(strings.NewReader(string(body)))

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "result__url" {
					fmt.Println(n.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	searches = append(searches, username+": "+searchTerm) // store the search query with username
}

func adminMenu() {
	var adminChoice int

	fmt.Println("Welcome admin.")
	fmt.Println("1. View all searches")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&adminChoice)

	if adminChoice == 1 {
		for _, search := range searches {
			fmt.Println(search)
		}
	} else {
		fmt.Println("Invalid choice.")
	}
}
