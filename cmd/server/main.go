/**
* * Server entry point
 */

package main

import (
	"log"
	"pulseDashboard/internal/bootstrap"
	"pulseDashboard/internal/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	app, err := bootstrap.New()

	if err != nil {
		log.Fatalf("bootstrap failed: %v", err)
	}

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
