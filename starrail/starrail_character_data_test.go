package starrail

import (
	"encoding/json"
	"os"
	"testing"
)

func TestResourceData(t *testing.T) {
	file, err := os.ReadFile("../resources/honkai_characters.json")
	if err != nil {
		t.FailNow()
	}

	var starRailCharacterData map[string]*CharacterData
	err = json.Unmarshal(file, &starRailCharacterData)
	if err != nil {
		t.FailNow()
	}

}
