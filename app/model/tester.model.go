package model
//
type Tester struct {
	Id    uint `gorm:"primarykey"`
	Key   string
	Value string
	UserId uint
}

func (Tester) TableName() string {
	return "testers"
}