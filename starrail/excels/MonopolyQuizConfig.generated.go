package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyQuizConfig struct {
	Duration       float64   `json:"Duration"`
	QuizDesc       hash.Hash `json:"QuizDesc"`
	QuizID         float64   `json:"QuizID"`
	QuizName       hash.Hash `json:"QuizName"`
	QuizTaskIDList []float64 `json:"QuizTaskIDList"`
}
type MonopolyQuizConfigAccessor struct {
	_data       []MonopolyQuizConfig
	_dataQuizID map[float64]MonopolyQuizConfig
}

// LoadData retrieves the data. Must be called before MonopolyQuizConfig.GroupData
func (a *MonopolyQuizConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyQuizConfig.json")
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
func (a *MonopolyQuizConfigAccessor) Raw() ([]MonopolyQuizConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyQuizConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyQuizConfigAccessor.LoadData to preload everything
func (a *MonopolyQuizConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataQuizID[d.QuizID] = d
	}
}

// ByQuizID returns the MonopolyQuizConfig uniquely identified by QuizID
//
// Error is only non-nil if the source errors out
func (a *MonopolyQuizConfigAccessor) ByQuizID(identifier float64) (MonopolyQuizConfig, error) {
	if a._dataQuizID == nil {
		err := a.LoadData()
		if err != nil {
			return MonopolyQuizConfig{}, err
		}
		a.GroupData()
	}
	return a._dataQuizID[identifier], nil
}
