package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron/v3"
)

// cleanPaymentData simulates a data cleaning task.
// It takes a job ID, simulates work, and logs its duration.
// This function represents the actual work that a cron scheduler would trigger.
func cleanPaymentData(jobID int) {
	startTime := time.Now()
	fmt.Printf("[%s] Job %d: Starting payment data cleaning...\n", startTime.Format("15:04:05"), jobID)

	// Simulate work (e.g., database operations, data processing).
	// The variable sleep duration here simulates the "cost" and "latency"
	// of a real-world data cleaning task, which is a key focus of the article.
	processingTime := time.Duration(rand.Intn(3)+1) * time.Second // Simulate 1-3 seconds of work
	time.Sleep(processingTime)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("[%s] Job %d: Finished payment data cleaning. Duration: %s\n", endTime.Format("15:04:05"), jobID, duration)
}

func main() {
	fmt.Println("Starting Go Cron Scheduler for Payment Data Cleaning Demo...")

	// Create a new cron scheduler instance.
	// We use `cron.WithSeconds()` to enable second-level scheduling,
	// which is useful for demonstrating frequent tasks and observing behavior.
	c := cron.New(cron.WithSeconds())

	// Initialize random seed for simulating variable work duration.
	rand.Seed(time.Now().UnixNano())

	// Keep track of job IDs for clearer logging.
	jobCounter := 0

	// Schedule the data cleaning task to run every 5 seconds.
	// The cron expression "*/5 * * * * *" means:
	//   - Every 5th second (0, 5, 10, ..., 55)
	//   - Of every minute
	//   - Of every hour
	//   - Of every day of the month
	//   - Of every month
	//   - Of every day of the week
	// This demonstrates a recurring scheduled task, central to the article's topic.
	_, err := c.AddFunc("*/5 * * * * *", func() {
		jobCounter++
		cleanPaymentData(jobCounter) // Execute the simulated cleaning task
	})
	if err != nil {
		fmt.Printf("Error scheduling job: %v\n", err)
		return
	}

	// Start the cron scheduler.
	// It runs in its own goroutine, allowing the `main` function to continue.
	c.Start()

	fmt.Println("Scheduler started. Scheduled to run every 5 seconds. Press Ctrl+C to exit.")

	// Keep the main goroutine alive indefinitely to allow the scheduler to run continuously.
	// In a real application, this might be replaced by a graceful shutdown mechanism.
	select {}
}
