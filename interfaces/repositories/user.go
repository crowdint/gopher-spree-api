package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

func (this *DbRepository) UserRoles(user *domain.User) {
	this.dbHandler.Model(user).Related(&user.Roles, "Roles")
}
