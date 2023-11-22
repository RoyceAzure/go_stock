package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("read json, path is empty")
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read json, get err : %w", err)
	}
	return file, nil
}

func WriteJsonFile(path string, data any) error {
	if path == "" {
		return fmt.Errorf("write json, path is empty")
	}
	var byteData []byte
	var err error
	var ok bool
	if byteData, ok = data.([]byte); !ok {
		byteData, err = json.Marshal(data)
		if err != nil {
			return fmt.Errorf("write json file, marshal json get err : %w", err)
		}
	}

	err = os.WriteFile(path, byteData, 0644)
	if err != nil {
		return fmt.Errorf("write json file, get err : %w", err)
	}

	return nil
}
