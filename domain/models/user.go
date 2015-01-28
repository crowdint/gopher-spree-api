package models

import (
	"time"
)

type User struct {
	Id int64

	AuthenticationToken string
	BillAddress         int64
	ConfirmationSentAt  time.Time
	ConfirmationToken   string
	ConfirmedAt         time.Time
	CreatedAt           time.Time
	CurrentSignInAt     time.Time
	CurrentSignInIp     string
	DeletedAt           time.Time
	Email               string
	EncryptedPassword   string
	FailedAttempts      int64
	LastRequestAt       time.Time
	LastSignInAt        time.Time
	LastSignInIp        string
	LockedAt            time.Time
	Login               string
	PasswordSalt        string
	PerishableToken     string
	PersistenceToken    string
	RememberCreatedAt   time.Time
	RememberToken       string
	ResetPasswordSentAt time.Time
	ResetPasswordToken  string
	ShipAddress         int64
	SignInCount         int64
	SpreeApiKey         string
	UnlockToken         string
	UpdatedAt           time.Time
}

func (this User) TableName() string {
	return "spree_users"
}
