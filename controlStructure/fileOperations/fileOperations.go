package fileOperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	balanceText := string(data)
	value, err := strconv.ParseFloat(balanceText, 64) // string conversion to a float type

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil // return nil if no error
}

// storing data in files
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
