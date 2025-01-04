package db

const TableNameKeystoneUser = "keystone_users"

// KeystoneUser mapped from table <keystone_users>
type KeystoneUser struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	Email    string `gorm:"column:email;not null" json:"email"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Password string `gorm:"column:password;not null" json:"password"`
	Username string `gorm:"column:username;not null" json:"username"`
	Deleted  bool   `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName KeystoneUser's table name
func (*KeystoneUser) TableName() string {
	return TableNameKeystoneUser
}
