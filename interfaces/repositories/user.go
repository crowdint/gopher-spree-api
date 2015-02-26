package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
)

func (this *DbRepository) UserRoles(user *json.User) {
	this.dbHandler.Model(user).Related(&user.Roles, "Roles")
}
