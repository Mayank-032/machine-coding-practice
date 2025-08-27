package user

import (
	"errors"

	"github.com/flipkart/machine-coding-practice/utils"
)

func (u *User) Validate() error {
	if u == nil {
		return errors.New("invalid user create requests")
	}

	if u.Name == utils.EmptyString {
		return errors.New("error while user creation request with invalid name")
	}

	return nil
}
