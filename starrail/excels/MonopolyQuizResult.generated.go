package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyQuizResult struct {
	Desc         hash.Hash `json:"Desc"`
	ID           float64   `json:"ID"`
	PlayerIDList []float64 `json:"PlayerIDList"`
	QuizID       float64   `json:"QuizID"`
}
type MonopolyQuizResultAccessor struct {
	_data   []MonopolyQuizResult
	_dataID map[float64]MonopolyQuizResult
}

// LoadData retrieves the data. Must be called before MonopolyQuizResult.GroupData
func (a *MonopolyQuizResultAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyQuizResult.json")
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
func (a *MonopolyQuizResultAccessor) Raw() ([]MonopolyQuizResult, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyQuizResult{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyQuizResultAccessor.LoadData to preload everything
func (a *MonopolyQuizResultAccessor) GroupData() {
	a._dataID = map[float64]MonopolyQuizResult{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MonopolyQuizResult uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MonopolyQuizResultAccessor) ByID(identifier float64) (MonopolyQuizResult, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonopolyQuizResult{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
