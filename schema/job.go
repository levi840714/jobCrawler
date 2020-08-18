package schema

import "time"

var AllSchema = []interface{}{
	&Job{},
}

type Job struct {
	ID       int       `gorm:"column:id;AUTO_INCREMENT"`
	Company  string    `gorm:"column:Company;type:varchar(80);NOT NULL"`
	JobId    string    `gorm:"column:JobId;type:varchar(40);UNIQUE;NOT NULL"`
	Keyword  string    `gorm:"column:Keyword;type:varchar(40);NOT NULL"`
	Title    string    `gorm:"column:Title;type:varchar(128);NOT NULL"`
	Salary   string    `gorm:"column:Salary;type:varchar(40)"`
	Content  string    `gorm:"column:Content;type:text"`
	Link     string    `gorm:"column:Link;type:varchar(256);NOT NULL"`
	Website  string    `gorm:"column:Website;type:varchar(20);NOT NULL"`
	Status   string    `gorm:"column:status;type:enum('0', '1');default:'1'"`
	CreateAt time.Time `gorm:"column:createAt;type:datetime;NOT NULL"`
}

func (Job) TableName() string {
	return "job"
}
