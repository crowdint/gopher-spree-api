package domain

import (
	"time"
)

type User struct {
	Id int64

	AuthenticationToken string
	BillAddressId       int64
	ConfirmationSentAt  time.Time
	ConfirmationToken   string
	ConfirmedAt         time.Time
	CreatedAt           time.Time
	CurrentSignInAt     time.Time
	CurrentSignInIp     string
	//DeletedAt           time.Time
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
	Roles               []Role `gorm:"many2many:spree_roles_users;"`
	ShipAddressId       int64
	SignInCount         int64
	SpreeApiKey         string
	UnlockToken         string
	UpdatedAt           time.Time
}

func (this User) TableName() string {
	return "spree_users"
}

func (this *User) HasRole(role string) bool {
	for _, r := range this.Roles {
		if r.Name == role {
			return true
		}
	}

	return false
}
