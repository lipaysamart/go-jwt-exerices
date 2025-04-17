package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lipaysamart/go-jwt-exerices/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"not null" json:"email"`
	Username  string         `json:"username" `
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserLoginReq struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UserRegisterReq struct {
	Email    string `json:"email" `
	Username string `json:"username" `
	Password string `json:"password" `
}

type UserLoginResp struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.Password = utils.HashAndSalt([]byte(u.Password))
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.Password = utils.HashAndSalt([]byte(u.Password))
	return nil
}
