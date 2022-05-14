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

type TicketStatus string

const (
	TicketStatusPending    TicketStatus = "pending"
	TicketStatusProcessing TicketStatus = "processing"
	TicketStatusCompleted  TicketStatus = "completed"
)

type FileType string

const (
	FileTypeFile  FileType = "file"
	FileTypePhoto FileType = "photo"
)
