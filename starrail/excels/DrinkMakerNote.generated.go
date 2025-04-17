package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type DrinkMakerNote struct {
	DrinkMakerNoteList []map[string]json.Number `json:"DrinkMakerNoteList"`
	GuestID            json.Number              `json:"GuestID"`
	UnlockDay          json.Number              `json:"UnlockDay"`
}
type DrinkMakerNoteAccessor struct {
	_data []DrinkMakerNote
}

// LoadData retrieves the data. Must be called before DrinkMakerNote.GroupData
func (a *DrinkMakerNoteAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DrinkMakerNote.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *DrinkMakerNoteAccessor) Raw() ([]DrinkMakerNote, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DrinkMakerNote{}, err
		}
	}
	return a._data, nil
}
