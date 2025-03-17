package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"waiter-app/db"
	"waiter-app/models"

	"github.com/jinzhu/gorm"
)

func GetAllTables(w http.ResponseWriter, r *http.Request) {
	var tables []models.Table
	err := db.DB.Find(&tables).Error
	if err != nil {
		http.Error(w, "Failed to retrieve tables", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}

func GetTable(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tables/"):]
	tableID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid table ID", http.StatusBadRequest)
		return
	}

	var table models.Table
	err = db.DB.First(&table, tableID).Error
	if err != nil {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(table)
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON to a Table struct
	var newTable models.Table
	err := json.NewDecoder(r.Body).Decode(&newTable)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Insert the new table into the database
	if err := db.DB.Create(&newTable).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the created table data as the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTable)
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tables/"):]
	tableID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid table ID", http.StatusBadRequest)
		return
	}

	// Decode the incoming JSON request body into a Table struct
	var updatedTable models.Table
	err = json.NewDecoder(r.Body).Decode(&updatedTable)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Find the existing table by ID
	var table models.Table
	err = db.DB.First(&table, tableID).Error
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving table: %v", err), http.StatusInternalServerError)
		return
	}

	// Update the table fields
	table.Name = updatedTable.Name
	table.Capacity = updatedTable.Capacity

	// Save the updated table
	err = db.DB.Save(&table).Error
	if err != nil {
		http.Error(w, "Failed to update table", http.StatusInternalServerError)
		return
	}

	// Return the updated table as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(table)
}

func DeleteTable(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tables/"):]
	tableID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid table ID", http.StatusBadRequest)
		return
	}

	err = db.DB.Delete(&models.Table{}, tableID).Error
	if err != nil {
		http.Error(w, "Failed to delete table", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
