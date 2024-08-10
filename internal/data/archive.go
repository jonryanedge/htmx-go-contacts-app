package data

import (
	"fmt"
	"strconv"
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

func GetArchiver() Archive {
	a := Archive{
		Status:   "Waiting",
		Progress: 0,
	}
	return a
}

func (a *Archive) GetArchive() Archive {
	return *a
}

func (a *Archive) GetArchiveStatus() string {
	status := &a.Status
	return *status
}

func (a *Archive) GetArchiveProgress() int {
	progress := &a.Progress
	return *progress
}

func (a *Archive) RunArchive() {
	for i := range [10]int{} {
		time.Sleep(time.Second)
		if a.Status != "Running" {
			return
		}
		a.Progress = (i + 1) / 10
		fmt.Printf("Here... %s", strconv.Itoa(a.Progress))
	}
	time.Sleep(time.Second)
	if a.Status != "Running" {
		return
	}
	a.Status = "Complete"
}

func (a *Archive) Run() {
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
	a.Status = "Waiting"
}
