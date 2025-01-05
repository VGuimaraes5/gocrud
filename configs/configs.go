package configs

import "log"

var (
	Database database
)

func init() {
	err := newDatabase()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
