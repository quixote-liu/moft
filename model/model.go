package model

type H map[string]interface{}

const (
	DirFile  = "./static_file"
	DirPhoto = "./static_photo"
)

type Role string

const (
	RoleAdmin  Role = "role_admin"
	RoleMember Role = "role_member"
)
