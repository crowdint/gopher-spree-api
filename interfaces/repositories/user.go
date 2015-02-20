package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

func (this *DbRepository) UserRoles(user *models.User) {
	this.dbHandler.Model(user).Related(&user.Roles, "Roles")
}
