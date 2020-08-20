package crawler

import (
	"fmt"
	"jobCrawler/model"
	"jobCrawler/telegram"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	URL_CR = "https://www.cakeresume.com/jobs?q=golang&refinementList%5Blocation_list%5D%5B0%5D=%E5%8F%B0%E4%B8%AD&refinementList%5Blocation_list%5D%5B1%5D=%E5%8F%B0%E5%8C%97&page="
)

var chCR = make(chan bool, 1)

type cakeresume struct {
	Name string
	Next string
}

func NewCakeresume() Action {
	return cakeresume{
		Name: Crawler_CakeResume,
		Next: Crawler_Init,
	}
}

func (c cakeresume) Entry() {
	fmt.Println("Entry cakeresume")
}

func (c cakeresume) Crawler() string {
	var page int = 1
	for {
		select {
		case <-chCR:
			fmt.Println("stop cakeresume crawler")
			return c.Next
		default:
			crawlerCakeresume(page)
			page++
			time.Sleep(time.Second)
		}
	}
}

func crawlerCakeresume(page int) {
	url := fmt.Sprintf("%s%d", URL_CR, page)
	fmt.Println(url)
	stories := []jobInfo{}

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: www.cakeresume.com
		colly.AllowedDomains("www.cakeresume.com"),
		// Parallelism
		colly.Async(true),
	)

	c.OnHTML(".no-result", func(e *colly.HTMLElement) {
		chCR <- true
		return
	})

	// On every a element which has .top-matter attribute call callback
	// This class is unique to the div that holds all information about a story
	c.OnHTML(".is-condensed", func(e *colly.HTMLElement) {
		tmp := jobInfo{}
		link := e.ChildAttr(".job-link", "href")
		split := strings.Split(link, "/")
		id := split[len(split)-1:][0]

		salary := e.ChildText(".job-salary")
		if salary == "" {
			salary = "待遇面議"
		}

		tmp.Id = id
		tmp.Company = e.ChildText(".page-name")
		tmp.Title = e.ChildText(".job-link")
		tmp.Salary = salary
		tmp.Content = e.ChildText(".job-desc")
		tmp.Link = link
		stories = append(stories, tmp)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())

	// })

	c.Visit(url)

	c.Wait()

	for _, v := range stories {
		// fmt.Println("ID: ", v.Id)
		// fmt.Println("公司: ", v.Company)
		// fmt.Println("職缺: ", v.Title)
		// fmt.Println("薪資: ", v.Salary)
		// fmt.Println("内容: ", v.Content)
		// fmt.Println("連結: ", v.Link)
		result := model.InsertJob(v.Id, "golang", v.Company, v.Title, v.Salary, v.Content, v.Link, "CakeResume")
		if result == true {
			telegram.Send(v.String())
		}
	}
	fmt.Println(len(stories))
}

func (c cakeresume) Exit() {
	fmt.Println("Exit cakeresume")
	time.Sleep(4 * time.Hour)
}
