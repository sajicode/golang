package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

//* notify implements a method that can be called via
//* a value of type user
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

//* admin represents an admin user with privileges
type admin struct {
	user
	level string
}

//* notify implements a method that can be called
//* via a value of type admin
//* by this, the inner type is not promoted since the
//* outer type is also implementing notify
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

//* main is the entry point for the app
func main() {
	//* create an admin user
	ad := admin{
		user: user{
			name:  "Icarus Daedalion",
			email: "icarus@wings.com",
		},
		level: "super",
	}

	//* send the admin a user notification
	//* the embedded inner type is NOT promoted to the outer type

	sendNotification(&ad)

	//* we can still access the inner type's method directly
	ad.user.notify()

	//* the inner type's method is not promoted
	ad.notify()
}

//* sendNotification accepts values that implement the notifier
//* interface sends notifications
func sendNotification(n notifier) {
	n.notify()
}
