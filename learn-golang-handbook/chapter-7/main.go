package main

import "fmt"

func main() {
	//Error handling instead of try catch blocks
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	profile, err = getUserProfile(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type User struct {
	id   int
	name string
}

func getUser() (User, error) {
	// do some get user logic here

	user := User{
		id:   1,
		name: "name",
	}

	return user, nil
}
