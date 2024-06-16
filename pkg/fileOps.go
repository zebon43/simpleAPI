package fileOps

import (
	"encoding/csv"
	"log"
	"os"
)

func GetData(file_path string) (records [][]string) {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatal("Error:", err)
	}

	return
}

func WriteFile(request_data []string) {
	file, err := os.OpenFile("request_log.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(request_data)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
