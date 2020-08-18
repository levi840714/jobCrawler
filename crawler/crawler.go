package crawler

import (
	"fmt"
	"time"
)

type jobInfo struct {
	Id      string
	Company string
	Title   string
	Salary  string
	Content string
	Link    string
}

const (
	Crawler_Init = "initial"
	Crawler_104  = "New104"
)

func (j *jobInfo) String() string {
	return fmt.Sprintf("公司: %s\n職缺: %s\n薪資: %s\n内容: \n%s\n連結: %s", j.Company, j.Title, j.Salary, j.Content, j.Link)
}

type Action interface {
	Entry()
	Crawler() string
	Exit()
}

var action = map[string]func() Action{
	Crawler_Init: NewInit,
	Crawler_104:  New104,
}

var current string = "initial"

func Run() {
	for {
		if fn, ok := action[current]; ok {
			instance := fn()
			instance.Entry()
			current = instance.Crawler()
			instance.Exit()
		}
	}
}

type Initial struct{}

func NewInit() Action {
	return Initial{}
}

func (Initial) Entry() {
	fmt.Println("Entry initial")
}

func (Initial) Crawler() string {
	return Crawler_104
}

func (Initial) Exit() {
	fmt.Println("Exit initial")
	time.Sleep(2 * time.Second)
}