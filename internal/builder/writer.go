package builder

import (
	"encoding/json"
	"os"
)

// WriteJSON marshals data to indented JSON and writes it to path.
func WriteJSON(data any, path string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}
