package data

import (
	"fmt"
	"sync"
	"time"
)

type Archive struct {
	Status   string
	Progress int
}

type Archiver interface {
	GetArchive() Archive
	GetArchiveStatus() string
}

var (
	archive Archive
	mu      sync.Mutex
)

var archiver = Archive{
	Status:   "Waiting",
	Progress: 0,
}

func GetArchiver() Archive {
	mu.Lock()
	defer mu.Unlock()
	// a := Archive{
	// 	Status:   "Waiting",
	// 	Progress: 0,
	// }
	return archiver
}

func (a *Archive) GetArchive() Archive {
	fmt.Printf("archive: %s", archive.Status)
	mu.Lock()
	defer mu.Unlock()
	return archive
}

func (a *Archive) GetArchiveStatus() string {
	mu.Lock()
	defer mu.Unlock()
	return archive.Status
}

func (a *Archive) GetArchiveProgress() int {
	mu.Lock()
	defer mu.Unlock()
	return archive.Progress
}

func (a *Archive) RunArchive() {
	mu.Lock()
	defer mu.Unlock()
	for i := range [10]int{} {
		time.Sleep(time.Second)
		if a.Status != "Running" {
			return
		}
		a.Progress = (i + 1) * 10
		fmt.Printf("Here... %d\n", a.Progress)
	}
	time.Sleep(time.Second)
	if a.Status != "Running" {
		return
	}
	a.Status = "Complete"
}

func (a *Archive) Run() {
	mu.Lock()
	defer mu.Unlock()
	if a.Status == "Waiting" {
		a.Status = "Running"
		a.Progress = 0
	}
	go a.RunArchive()
}

func (a *Archive) File() string {
	return contactsFile
}

func (a *Archive) Reset() {
	mu.Lock()
	defer mu.Unlock()
	a.Status = "Waiting"
	a.Progress = 0
}
