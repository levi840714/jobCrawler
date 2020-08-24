package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"jobCrawler/config"
	"jobCrawler/crawler"
	"jobCrawler/model"
	"jobCrawler/schema"
	"jobCrawler/telegram"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	var wg sync.WaitGroup
	var keywords []string

	// Get search keywords separated by ","
	keyword := flag.String("keyword", "", "Search keywords")
	flag.Parse()
	if *keyword == "" {
		panic("Please enter search keywords!!")
	}
	keywords = strings.Split(*keyword, ",")
	fmt.Println(keywords)

	//DB connect
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Ip, config.Config.Mysql.Port, config.Config.Mysql.Db)
	model.DB, err = gorm.Open("mysql", connectStr)
	if err != nil {
		panic("DB connection failed!")
	}

	model.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(schema.AllSchema...)
	defer model.DB.Close()

	// telegram bot setup
	telegram.Init()

	//start job crawler run!

	wg.Add(1)
	for _, keyword := range keywords {
		if keyword != "" {
			go crawler.Run(keyword)
		}
	}
	wg.Wait()
}
