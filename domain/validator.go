package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotBlank = errors.New("can't be blank")
	ErrNotValid = errors.New("Invalid resource. Please fix errors and try again.")
)

type Validatorable interface {
	IsValid() bool
	GetErrors() *ValidatorErrors
}

type ValidatorErrors map[string]string

func (this *ValidatorErrors) Add(key string, errStr string) {
	(*this)[key] = errStr
}

func (this *ValidatorErrors) IsEmpty() bool {
	return len(*this) == 0
}

func (this *ValidatorErrors) Size() int {
	return len(*this)
}

func ErrTooShort(minLen int) error {
	errStr := fmt.Sprintf("is too short (minimum is %d characters)", minLen)
	return errors.New(errStr)
}

func ErrGreaterThanOrEqual(num int) error {
	errStr := fmt.Sprintf("must be greater than or equal to %d", num)
	return errors.New(errStr)
}
