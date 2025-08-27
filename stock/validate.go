package stock

import (
	"errors"

	"github.com/flipkart/machine-coding-practice/utils"
)

func (s *Stock) Validate() error {
	if s == nil {
		return errors.New("invalid user create requests")
	}

	if s.Name == utils.EmptyString {
		return errors.New("error while user creation request with invalid name")
	}

	return nil
}
