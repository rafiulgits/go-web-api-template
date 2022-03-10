package domains

type Demo struct {
	ID int `gorm:"primarykey:size:30"`
}

const DemoTableName = "Demos"

func (Demo) TableName() string {
	return DemoTableName
}
