package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // Embedding User struct, anonymous field
}

// New is a constructor function that creates a new User instance.
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, fmt.Errorf("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// NewAdmin is a constructor function that creates a new Admin instance.
func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, fmt.Errorf("email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "--/--/----",
			createdAt: time.Now(),
		},
	}, nil
}

// ClearUserName clears the firstName and lastName fields of the User struct.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// OutputUserDetails prints the details of the User struct.
// It displays different tim.Time format too.
func (u User) OutputUserDetails() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("birthdate:", u.birthdate)
	fmt.Println("Created At (Local):", u.createdAt.Local())
	fmt.Println("Created At (UTC):", u.createdAt.UTC())
	fmt.Println("Created At (Unix):", u.createdAt.Unix())
	fmt.Println("Created At (UnixNano):", u.createdAt.UnixNano())
	fmt.Println("Created At (UnixMicro):", u.createdAt.UnixMicro())
	fmt.Println("Created At (UnixMilli):", u.createdAt.UnixMilli())
	fmt.Println("Created At (ISO):", u.createdAt.Format(time.RFC3339))
	fmt.Println("Created At (RFC1123):", u.createdAt.Format(time.RFC1123))
	fmt.Println("Created At (RFC850):", u.createdAt.Format(time.RFC850))
	fmt.Println("Created At (ANSIC):", u.createdAt.Format(time.ANSIC))
	fmt.Println("Created At (Kitchen):", u.createdAt.Format(time.Kitchen))
	fmt.Println("Created At (Stamp):", u.createdAt.Format(time.Stamp))
	fmt.Println("Created At (StampMilli):", u.createdAt.Format(time.StampMilli))
	fmt.Println("Created At (StampMicro):", u.createdAt.Format(time.StampMicro))
	fmt.Println("Created At (StampNano):", u.createdAt.Format(time.StampNano))
	fmt.Println("----------------------------------------")
}

// OutputUserDetails2 prints the details of the User struct in a different format.
// This function is a placeholder for a different output format.
func (u *User) OutputUserDetails2() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt.Local())
	fmt.Println("----------------------------------------")
}
