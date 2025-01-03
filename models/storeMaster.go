package models

import (
	"log"
	"encoding/csv"
	"os"
)

type Store struct {
	ID        string
	Name      string
	AreaCode  string
}

var StoreMaster = make(map[string]Store)  // store the information of the store in the map StoreMaster(in Memory)

func LoadStoreMaster(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records[1:] {

		if len(record) != 3 {
			log.Printf("Error on row %d: Invalid number of fields. Expected 3, got %d\n", i+2, len(record))
			continue // Skip this record and proceed to the next
		}

		store := Store{
			ID:       record[0],
			Name:     record[1],
			AreaCode: record[2],
		}
		StoreMaster[store.ID] = store
	}
	return nil
}
