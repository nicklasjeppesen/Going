package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run add_migration.go <command>")
		return
	}

	command := os.Args[1]
	fmt.Println(command)

	switch command {
	case "add":
		Add_migration(os.Args[2])
	case "run":
		Run_migrations()
	default:
		fmt.Println("Unknown command idiot")
	}

}

func Add_migration(name string) {
	if len(os.Args) < 2 {
		fmt.Println("Brug: go run add_migration.go <navn>")
		return
	}

	// Timestamp i format YYYYMMDDHHMMSS
	timestamp := time.Now().Format("20060102150405")

	dir := "scripts"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Printf("Kunne ikke oprette mappe %s: %v\n", dir, err)
			return
		}
	}

	// Nyt filnavn
	filename := filepath.Join(dir, fmt.Sprintf("%s_%s.sql", timestamp, name))

	// Opret tom fil
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Kunne ikke oprette fil: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("Ny migration oprettet: %s\n", filename)
}

func Run_migrations() {

	db, err := sql.Open("sqlite3", "../../../data/connect.db")
	if err != nil {
		log.Fatalf("Error trying opening the database: %v", err)
	}
	defer db.Close()

	// Sikr at migrations tabellen findes
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY key,
			filename TEXT UNIQUE NOT NULL,
			applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("Error connection to migration tabel: %v", err)
	}

	// Reading all files in ./scripts folder
	files, err := os.ReadDir("./scripts")
	if err != nil {
		log.Fatalf("Could not read the folder: %v", err)
	}

	// Filtering only *.SQL
	var migrations []string
	for _, f := range files {
		if !f.IsDir() && filepathExt(f.Name()) == ".sql" {
			migrations = append(migrations, f.Name())
		}
	}

	// Sorting in ASC order
	sort.Strings(migrations)

	// Running each migration
	for _, m := range migrations {

		// Tjek om filen allerede findes i migrations tabellen
		var migrationAlreadyRun bool
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM migrations WHERE filename = $1)", m).Scan(&migrationAlreadyRun)
		if err != nil {
			log.Fatalf("Error by checking migration tabel: %v", err)
		}

		if migrationAlreadyRun {
			continue
		}

		fmt.Printf("Running migration: %s\n", m)

		sqlBytes, err := os.ReadFile("./scripts/" + m)
		if err != nil {
			log.Fatalf("Could not read %s: %v", m, err)
		}

		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			//log.Fatalf("Error by running %s: %v", m, err)
			fmt.Printf("Error by running %s: %v", m, err)
		}

		// Indsæt i migrations tabellen
		_, err = db.Exec("INSERT INTO migrations (filename) VALUES ($1)", m)
		if err != nil {
			log.Fatalf("Error trying insert migration into migration tabel: %v", err)
		}

	}

	fmt.Println("All Migration executed successfully")
}

func filepathExt(path string) string {
	if len(path) < 4 {
		return ""
	}
	return path[len(path)-4:]
}
