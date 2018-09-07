package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Generate a UUID for each Id before the entity is saved.
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
