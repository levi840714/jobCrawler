package crawler

import (
	"fmt"
	"time"
)

const (
	Crawler_Init       = "initial"
	Crawler_Shutdown   = "shutdown"
	Crawler_104        = "New104"
	Crawler_CakeResume = "NewCakeresume"
)

type jobInfo struct {
	Id       string
	Company  string
	Location string
	Title    string
	Salary   string
	Content  string
	Link     string
}

func (j *jobInfo) String() string {
	return fmt.Sprintf("公司: %s\n地區: %s\n職缺: %s\n薪資: %s\n内容: \n%s\n連結: %s", j.Company, j.Location, j.Title, j.Salary, j.Content, j.Link)
}

type IAction interface {
	Entry()
	Crawler() string
	Exit()
}

type JobCrawler struct {
	Initial string
	Final   string
	Action  map[string]IAction
}

// start running crawler action loop
func (j *JobCrawler) Run() {
	current := j.Initial
	for {
		if instance, ok := j.Action[current]; ok {
			instance.Entry()
			current = instance.Crawler()
			instance.Exit()
		}
	}
}

// setup jobCrawler action running process
func Run(keyword string) {
	jobCrawler := JobCrawler{
		Initial: Crawler_Init,
		Final:   Crawler_Shutdown,
		Action: map[string]IAction{
			Crawler_Init:       NewInit(keyword),
			Crawler_104:        New104(keyword),
			Crawler_CakeResume: NewCakeresume(keyword),
		},
	}

	jobCrawler.Run()
}

type Initial struct {
	Keyword string
	Name    string
	Next    string
}

func NewInit(keyword string) IAction {
	return Initial{
		Keyword: keyword,
		Name:    Crawler_Init,
		Next:    Crawler_104,
	}
}

func (i Initial) Entry() {
	fmt.Println("Entry initial")
}

func (i Initial) Crawler() string {
	return i.Next
}

func (i Initial) Exit() {
	fmt.Println("Exit initial")
	time.Sleep(2 * time.Second)
}
