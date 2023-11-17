package scheduler

import (
	"github.com/robfig/cron/v3"
	persistence "homeTestAlfaBet/internal/db"
	"homeTestAlfaBet/internal/repository"
	"log"
	"sync"
	"time"
)

// todo maybe later add a column in the DB that indicates if the reminder was sent
// To keep track of events for which reminders have been sent
var sentReminders = make(map[uint]bool)
var mu sync.Mutex // Mutex to handle concurrent access to sentReminders

func SetupEventReminders() {
	db := persistence.SetupDatabase()
	eventRepo := repository.NewEventRepository(db)

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		sendReminders(eventRepo)
	})
	c.Start()
}

func sendReminders(repo repository.EventRepository) {
	now := time.Now()
	reminderTime := now.Add(30 * time.Minute)

	events, err := repo.GetEventsHappeningBetween(now, reminderTime)
	if err != nil {
		log.Println("Error retrieving events:", err)
		return
	}

	for _, event := range events {
		mu.Lock()
		if _, alreadySent := sentReminders[event.ID]; !alreadySent {
			// Send reminder logic
			log.Printf("Reminder: Event '%s' is starting at %s\n", event.Name, event.StartTime)
			// Mark as sent
			sentReminders[event.ID] = true
			// Here, integrate with an email service, notification system, or other reminder mechanism
		}
		mu.Unlock()
	}
}
