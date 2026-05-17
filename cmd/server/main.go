/**
* * Server entry point
 */

package main

import (
	"log"
	"pulseDashboard/internal/bootstrap"
	"pulseDashboard/internal/config"
)

// func init() {
// 	config.LoadEnvVariables()
// }

func main() {
	config.LoadEnvVariables()
	app, err := bootstrap.New()

	if err != nil {
		log.Fatalf("bootstrap failed: %v", err)
	}

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
