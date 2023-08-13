package dao

type Admin struct {
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:passwd"`
	Name     string `gorm:"column:ad_name"`
	Gender   string `gorm:"column:gender"`
}

func (Admin) TableName() string {
	return "admin"
}

func AdminLogin(account, password string) bool {
	var admin Admin
	res := DB.Where("account=? AND passwd=?", account, password).First(&admin)
	if res.RowsAffected != 0 {
		return true
	}
	return false
}
