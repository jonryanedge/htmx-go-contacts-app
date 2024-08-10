package archiver

import (
	// "fmt"
	"math/rand"
	"sync"
	"time"
)

// Archiver represents the archiving process
type Archiver struct {
	mu              sync.Mutex
	archiveStatus   string
	archiveProgress int
	thread          *sync.WaitGroup
}

// NewArchiver creates a new Archiver instance
func NewArchiver() *Archiver {
	return &Archiver{
		archiveStatus:   "Waiting",
		archiveProgress: 0,
		thread:          &sync.WaitGroup{},
	}
}

// Status returns the current status of the archiving process
func (a *Archiver) Status() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.archiveStatus
}

// Progress returns the current progress of the archiving process
func (a *Archiver) Progress() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.archiveProgress
}

// Run starts the archiving process
func (a *Archiver) Run() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.archiveStatus == "Waiting" {
		a.archiveStatus = "Running"
		a.archiveProgress = 0
		a.thread.Add(1)
		go a.runImpl()
	}
}

// runImpl performs the archiving task
func (a *Archiver) runImpl() {
	defer a.thread.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
		a.mu.Lock()
		if a.archiveStatus != "Running" {
			a.mu.Unlock()
			return
		}
		a.archiveProgress = int(i+1) * 10
		a.mu.Unlock()
		// fmt.Printf("Working... %d%%\n", a.archiveProgress)
	}
	time.Sleep(1 * time.Second)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.archiveStatus == "Running" {
		a.archiveStatus = "Complete"
	}
}

// ArchiveFile returns the name of the file being archived
func (a *Archiver) ArchiveFile() string {
	return "internal/contacts/contacts.json"
}

// Reset resets the archiving process to its initial state
func (a *Archiver) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.archiveStatus = "Waiting"
	a.archiveProgress = 0
	if a.thread != nil {
		a.thread.Wait() // Wait for any running goroutines to finish
	}
	a.thread = &sync.WaitGroup{}
}