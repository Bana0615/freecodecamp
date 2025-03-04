package main

import "fmt"

func main() {
	//Error handling instead of try catch blocks
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	profile, err := getUserProfile(user.id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(profile)
}

// type userError struct {
// 	name string
// }

// func sendSMS(msg, userName string) error {
// 	if !canSendToUser(userName) {
// 		return userError{name: userName}
// 	}

// 	...
// }

// func (e userError) Error() string {
// 	return fmt.Sprintf("%v has a problem with their account", e.name)
// }

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

func getUserProfile(id int) (User, error) {
	// do some get user logic here

	user := User{
		id:   1,
		name: "name",
	}

	return user, nil
}
