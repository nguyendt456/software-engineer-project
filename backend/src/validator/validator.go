package validator

import (
	"fmt"
	"regexp"

	pb "github.com/nguyendt456/software-engineer-project/pb"
)

func ValidateUser(user *pb.User) error {
	if user.Username == "" {
		return fmt.Errorf("Username field is required")
	}
	if user.Password == "" {
		return fmt.Errorf("Password field is required")
	}
	if user.Usertype == "" {
		return fmt.Errorf("User type field is required")
	}
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}
	if err := ValidateUserType(user.Usertype); err != nil {
		return err
	}
	return nil
}

func ValidateUsername(username string) error {
	minLength := 5
	maxLength := 100

	if len(username) <= minLength || len(username) >= maxLength {
		return fmt.Errorf("Username must contain from %d to %d", minLength, maxLength)
	}
	isValid := regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(username)
	if !isValid {
		return fmt.Errorf("Username must not contain capital or special letter !")
	}
	return nil
}

func ValidateUserType(usertype string) error {
	if usertype != "collector" && usertype != "backofficer" && usertype != "janitor" {
		return fmt.Errorf("Unknown user type: %s", usertype)
	}
	return nil
}

func ValidateLoginForm(login_form *pb.LoginForm) error {
	err := ValidateUsername(login_form.Username)
	if err != nil {
		return fmt.Errorf("Login failed")
	}
	return nil
}
