package routes

import (
	"axo/axo"
	"axo/database"
	"axo/models"
	"encoding/json"
	"net/http"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	//Find all notes
	var notes []models.Note
	database.DB.Find(&notes)
	//Return note
	json.NewEncoder(w).Encode(notes)
}

func PostNote(w http.ResponseWriter, r *http.Request) {
	var note = r.URL.Query().Get("note")
	//Check if note is empty
	if note == "" {
		axo.Error(w, "note is empty", http.StatusBadRequest)
		return
	}
	//Create note
	var newNote models.Note
	newNote.Title = note
	//Save note
	database.DB.Create(&newNote)
	//Return note
	json.NewEncoder(w).Encode(newNote)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	//Get note id
	var noteId = r.URL.Query().Get("id")
	//Check if note id is empty
	if noteId == "" {
		axo.Error(w, "note id is empty", http.StatusBadRequest)
		return
	}
	//Delete note
	database.DB.Delete(&models.Note{}, noteId)
}
