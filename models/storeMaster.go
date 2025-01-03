package models

import (
	"encoding/csv"
	"os"
)

type Store struct {
	ID        string
	Name      string
	AreaCode  string
}

var StoreMaster = make(map[string]Store)

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

	for _, record := range records[1:] {
		store := Store{
			ID:       record[0],
			Name:     record[1],
			AreaCode: record[2],
		}
		StoreMaster[store.ID] = store
	}
	return nil
}
