package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type StroyLineUIData struct {
	ChronicleIconPath string    `json:"ChronicleIconPath"`
	Color             string    `json:"Color"`
	FigurePath        string    `json:"FigurePath"`
	Gender            string    `json:"Gender"`
	IconPath          string    `json:"IconPath"`
	LargeImgPath      string    `json:"LargeImgPath"`
	MediumImgPath     string    `json:"MediumImgPath"`
	Name              hash.Hash `json:"Name"`
	StoryLineID       float64   `json:"StoryLineID"`
}
type StroyLineUIDataAccessor struct {
	_data []StroyLineUIData
}

// LoadData retrieves the data. Must be called before StroyLineUIData.GroupData
func (a *StroyLineUIDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StroyLineUIData.json")
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
func (a *StroyLineUIDataAccessor) Raw() ([]StroyLineUIData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StroyLineUIData{}, err
		}
	}
	return a._data, nil
}
