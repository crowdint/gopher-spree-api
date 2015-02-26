package domain

type Role struct {
	Id   int64
	Name string
}

func (this Role) TableName() string {
	return "spree_roles"
}
