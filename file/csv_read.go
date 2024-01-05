package csv_read

import (
	"encoding/csv"
	"os"
)

func LoadCsv(path string, delim string) error {

	//open csv file
	file, err := os.Open(path)
	if err != nil {
		// TODO: Error return
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return err
	}

	headerMap := HeaderMap(headers)
	println(headerMap)

	return nil
}

// produce map of headers from headers array
func HeaderMap(headers []string) map[string]int {
	headerMap := make(map[string]int)

	for col, header := range headers {
		headerMap[header] = col
	}
	return headerMap
}
