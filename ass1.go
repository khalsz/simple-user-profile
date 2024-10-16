package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// generating unique user ID
func generateUserId() string {
	userId := uuid.New()
	return userId.String()
}

// userName takes in username stdinput and outputs it
func userName(reader *bufio.Reader) (string, error) {
	fmt.Println("Please enter your name")
	name, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("error reading username input: %s", err)
	}

	// trim posssible trailing and leading space from stdinput
	name = strings.TrimSpace(name)
	return name, nil
}

// userAge takes in user age stdinput and outputs it
func userAge(reader *bufio.Reader) (int, error) {
	fmt.Println("Please enter your age")
	age, err := reader.ReadString('\n')

	if err != nil {
		return 0, fmt.Errorf("error reading age input: %s", err)
	}

	// trim possible leading and trailing space from input
	age = strings.TrimSpace(age)

	// convert user input to integer
	userage, err := strconv.Atoi(age)
	if err != nil {
		return 0, fmt.Errorf("error converting user age to numer, please enter a valid user age")
	}
	return userage, nil
}

// initializing country code constant
const countryCode = "United Kingdom"

// definining location type
type Location int

// defining user location constant as enum
const (
	leicester Location = iota
	london
	lincoln
	unknown
)

// creating map for defined user location
var locationNames = map[Location]string{
	leicester: "Leicester",
	london:    "London",
	lincoln:   "Lincoln",
}

// userName takes in username stdinput and outputs it
func userLocation(reader *bufio.Reader) (string, error) {
	fmt.Println("Please enter your location (must be either of leicester, london, lincoln):")
	location, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("error reading username input: %s", err)
	}

	// trim possible leading and trailing space from input
	location = strings.TrimSpace(location)

	// checking if user location input matches already defined location
	for _, name := range locationNames {
		if strings.EqualFold(name, location) {
			return name, nil
		}

	}

	return "", fmt.Errorf("unknown location entered: %s", location)

}

// userName takes in username stdinput and outputs it
func userFavNumber(reader *bufio.Reader) (float64, error) {

	fmt.Println("Please enter your favorite single decimal number (e.g., 1.2, 5):")
	for {
		favNumber, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("error converting to float: %s", err)
		}

		// trim possible leading and trailing space from input
		favNumber = strings.TrimSpace(favNumber)

		// // ensuring user fav number input is a single decimal value
		// if len(favNumber) > 3 {
		// 	fmt.Println("Enter a single digit number")
		// 	continue
		// }

		// converting user input to float
		favNum, err := strconv.ParseFloat(favNumber, 64)

		if err != nil {
			fmt.Println("Error converting fav. Num. to float")
			continue
		}

		// ensuring user input is a single unit decimal number
		if favNum > 0 && favNum < 10 {
			return favNum, nil
		} else {
			fmt.Println("Please enter a number between 0 and 9.")
		}
	}
}

// userName takes in username stdinput and outputs it
func userHubby(reader *bufio.Reader) (string, error) {
	fmt.Println("Enter your hubby:")
	hubby, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading user hubby: %s", err)
	}

	// trim possible leading and trailing space from input
	hubby = strings.TrimSpace(hubby)
	return hubby, err
}

func main() {

	var loginStatus bool = true

	reader := bufio.NewReader(os.Stdin)

	newUserId := generateUserId()

	newUsername, err := userName(reader)

	if err != nil {
		fmt.Println("Error creating new user name", err)
	}

	newUserAge, err := userAge(reader)

	if err != nil {
		fmt.Println("Error creating new user name", err)
	}

	newUserLocation, err := userLocation(reader)

	if err != nil {
		fmt.Println("Error creating new user name", err)
	}

	newUserFavNumber, err := userFavNumber(reader)

	if err != nil {
		fmt.Println("Error creating new user name", err)
	}

	newUserHubby, err := userHubby(reader)

	if err != nil {
		fmt.Println("Error creating new user name", err)
	}

	// compile user input information into map
	userProfile := map[string]interface{}{
		"Id":               newUserId,
		"Name":             newUsername,
		"loginStatus":      loginStatus,
		"Age":              newUserAge,
		"Location":         newUserLocation,
		"Favourite Number": newUserFavNumber,
		"Hubby":            newUserHubby,
		"Country":          countryCode,
	}

	fmt.Println("\nThis is your profile:")

	for key, value := range userProfile {
		fmt.Printf("%s: %v\n", key, value)
	}

}
