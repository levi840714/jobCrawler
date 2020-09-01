package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	JobOk   = "1"
	JobFail = "0"
)

var (
	DB *gorm.DB
)

type Job struct {
	ID       int       `gorm:"column:id"`
	JobId    string    `gorm:"column:JobId"`
	Keyword  string    `gorm:"column:Keyword"`
	Company  string    `gorm:"column:Company"`
	Location string    `gorm:"column:Location"`
	Title    string    `gorm:"column:Title"`
	Salary   string    `gorm:"column:Salary"`
	Content  string    `gorm:"column:Content"`
	Link     string    `gorm:"column:Link"`
	Website  string    `gorm:"column:Website"`
	Status   string    `gorm:"column:status"`
	CreateAt time.Time `gorm:"column:createAt"`
}

func (Job) TableName() string {
	return "job"
}

func InsertJob(jobid, keyword, company, location, title, salary, content, link, website string) bool {
	insert := Job{JobId: jobid, Keyword: keyword, Company: company, Location: location, Title: title, Salary: salary, Content: content, Link: link, Website: website, Status: JobOk, CreateAt: time.Now()}
	db := DB.Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE Location = VALUES(Location), Title = VALUES(Title), Salary = VALUES(Salary), Content = VALUES(Content), Link = VALUES(Link)").
		Create(&insert)
	if db.Error != nil {
		log.Println(db.Error)
	}

	result := db.RowsAffected == 1
	return result
}

func CloseDB() {
	for DB.DB().Stats().InUse > 0 {
		log.Println("DB not disconnected count: ", DB.DB().Stats().InUse)
		time.Sleep(time.Second * 1)
	}

	DB.DB().Close()
}
