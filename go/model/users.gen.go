// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email               string    `gorm:"column:email;not null" json:"email"`
	EncryptedPassword   string    `gorm:"column:encrypted_password;not null" json:"encrypted_password"`
	ResetPasswordToken  string    `gorm:"column:reset_password_token" json:"reset_password_token"`
	ResetPasswordSentAt time.Time `gorm:"column:reset_password_sent_at" json:"reset_password_sent_at"`
	RememberCreatedAt   time.Time `gorm:"column:remember_created_at" json:"remember_created_at"`
	CreatedAt           time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Roles               string    `gorm:"column:Roles;default:user" json:"Roles"`
	Jti                 string    `gorm:"column:jti;not null" json:"jti"`
	Name                string    `gorm:"column:name" json:"name"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
