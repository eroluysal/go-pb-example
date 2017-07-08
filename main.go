package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	pb "github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"gopkg.in/kyokomi/emoji.v1"
)

var steps = []string{
	"connecting to server",
	"downloading sources",
	"installing dependencies",
	"seeding databases",
	"collecting assets",
	"restarting services",
	"publishing deploy",
}

var servers = []string{
	"aws-en",
	"aws-us",
	"aws-uk",
	"aws-jp",
	"aws-tr",
	"aws-eu",
	"aws-as",
	"aws-eur",
	"aws-kin",
	"aws-s3",
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Starting deploy ...")

	pb.Start()
	for _, v := range servers {
		wg.Add(1)
		go deploy(v, &wg)
	}
	wg.Wait()
	pb.Stop()

	fmt.Println("Finished deploy")
}

func deploy(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	p := progress().PrependFunc(func(b *pb.Bar) string {
		return strutil.Resize(name+": "+steps[b.Current()-1], 30)
	}).PrependFunc(func(b *pb.Bar) string {
		if len(steps) == b.Current() {
			return emoji.Sprint(":rocket:")
		} else {
			return emoji.Sprint(":eyes:")
		}
	})

	for p.Incr() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	}
}

func progress() *pb.Bar {
	c := pb.AddBar(len(steps)).PrependElapsed().AppendCompleted()
	c.Width = 20

	return c
}
