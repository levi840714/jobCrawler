package main

import (
	"flag"
	"fmt"

	"jobCrawler/config"
	"jobCrawler/crawler"
	"jobCrawler/model"
	"jobCrawler/schema"
	"jobCrawler/telegram"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	keyword := flag.String("keyword", "", "Search keywords")
	flag.Parse()
	if *keyword == "" {
		panic("Please enter search keywords!!")
	}
	crawler.Keyword = *keyword
}

func main() {
	var err error
	//DB connect
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Ip, config.Config.Mysql.Port, config.Config.Mysql.Db)
	model.DB, err = gorm.Open("mysql", connectStr)
	if err != nil {
		panic("DB connection failed!")
	}

	model.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(schema.AllSchema...)
	defer model.DB.Close()

	// telegram bot setup
	telegram.Init()

	//start job crawler run!
	crawler.Run()
}
