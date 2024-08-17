package dtos

import "errors"

type AddUserDTO struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
}

func (dto AddUserDTO) Validate() error {
	if dto.Username == "" {
		return errors.New("username is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password != dto.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}
	return nil
}
