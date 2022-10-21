package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// DB Params (Specify credentials)
const (
	host     = "localhost"
	port     = "5432"
	user     = ""
	password = ""
	dbname   = "rusty-blog"
)

func initDbConnection() *sql.DB {
	psqlParams := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlParams)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Notes Table (https://stackoverflow.com/questions/16330490/in-go-how-can-i-convert-a-struct-to-a-byte-array)
// First letter needs to be capitalized as a the first letter cap means the name is exported and can be used anywhere
type Note struct {
	Note_id      string    `json:"note_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Notebook_id  string    `json:"notebook_id"`
}

func getAllNotes(w http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		db := initDbConnection()

		rows, err := db.Query("select * from notes")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var notes []Note

		for rows.Next() {
			var note Note

			if err := rows.Scan(&note.Note_id, &note.Content, &note.Title, &note.Notebook_id); err != nil {
				log.Println(err.Error())
			}

			notes = append(notes, note)
		}

		notesBytes, _ := json.MarshalIndent(notes, "", "\t")

		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(notesBytes)
		defer db.Close()

	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

func updateNoteContent(w http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		db := initDbConnection()

		w.Header().Set("Content-Type", "application/json")
		// Get JSON data from request
		decoder := json.NewDecoder(request.Body)

		var note Note

		err := decoder.Decode(&note)
		if err != nil {
			log.Fatal(err)
		}
		query := "update notes set content=$1;"

		_, err = db.Exec(query, note.Content)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

// type Notebook struct {
// 	Notebook_id string `json:"notebook_id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

// func getAllNotebooks(w http.ResponseWriter, request *http.Request) {

// 	if request.Method == "GET" {
// 		db := initDbConnection()

// 		rows, err := db.Query("select * from notebooks")
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		defer rows.Close()

// 		var notebooks []Notebook

// 		for rows.Next() {
// 			var notebook Notebook

// 			if err := rows.Scan(&notebook.Notebook_id, &notebook.Title, &notebook.Description); err != nil {
// 				log.Println(err.Error())
// 			}

// 			notebooks = append(notebooks, notebook)
// 		}

// 		notesBytes, _ := json.MarshalIndent(notebooks, "", "\t")

// 		if err != nil {
// 			log.Println(err.Error())
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(notesBytes)
// 		defer db.Close()

// 	} else {
// 		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
// 	}
// }

func main() {
	http.HandleFunc("/get-all-notes", getAllNotes)
	// http.HandleFunc("/get-all-notebooks", getAllNotebooks)
	http.HandleFunc("/update-note-content", updateNoteContent)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
