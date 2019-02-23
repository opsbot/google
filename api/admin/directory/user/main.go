package user

import (
	"fmt"
	"log"

	"github.com/opsbot/google/api"
)

// https://developers.google.com/admin-sdk/directory/v1/guides/manage-users

// List -
func List() {
	srv := api.Service()
	r, err := srv.Users.List().Customer("my_customer").MaxResults(10).OrderBy("email").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain: %v", err)
	}

	if len(r.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		fmt.Print("Users:\n")
		for _, u := range r.Users {
			fmt.Printf("%s (%s)\n", u.PrimaryEmail, u.Name.FullName)
		}
	}
}
