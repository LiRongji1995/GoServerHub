package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"` // 权限列表，如 "read,write,delete"
}

type Report struct {
	TotalUsers    int `json:"total_users"`
	ActiveUsers   int `json:"active_users"`
	NewUsersToday int `json:"new_users_today"`
}
