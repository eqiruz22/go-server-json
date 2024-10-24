package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
}


type Database struct {
	Posts []Post `json:"posts"`
}

var Db Database

func LoadDB() error {
	file, err := os.Open("db.json")
	if err != nil {
		return fmt.Errorf("failed to open database: %w",err)
	}
	defer file.Close()
	db, _ := io.ReadAll(file)
	err = json.Unmarshal(db, &Db)
	if err != nil {
		return fmt.Errorf("failed to parse json database: %w",err)
	}
	return nil
}

// save for each data
func SaveDB() error {
	file,err := json.MarshalIndent(Db,"","  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("db.json",file,0644)
	if err != nil {
		return err
	}
	return nil;
}

