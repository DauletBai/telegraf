package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ApplyMigrations(db *sql.DB, migrationsDir string) error {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("Failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationFilePath := filepath.Join(migrationsDir, file.Name())
			migrationContent, err := ioutil.ReadFile(migrationFilePath)
			if err != nil {
				return fmt.Errorf("Failed to read migration files %s: %v", file.Name(), err)
			}

			log.Printf("Application of migration: %s", file.Name())
			_, err = db.Exec(string(migrationContent))
			if err != nil {
				return fmt.Errorf("Error applying migration %s: %v", file.Name(), err)
			}
		}
	}

	log.Println("Powerful graces successfully applied")
	return nil
}