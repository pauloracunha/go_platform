package models

type Tenant struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"type:varchar(100);unique"`
}

type User struct {
	ID      uint     `gorm:"primaryKey"`
	Nome    string   `gorm:"type:varchar(100)"`
	Email   string   `gorm:"type:varchar(100);unique"`
	Senha   string   `gorm:"type:varchar(255)"`
	Tenants []Tenant `gorm:"many2many:user_tenants"`
}

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"type:varchar(100);unique"`
}

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Nome        string       `gorm:"type:varchar(100);unique"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type UserTenant struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	TenantID    uint
	RoleID      uint
	Permissions []Permission `gorm:"many2many:user_tenant_permissions"`
}
