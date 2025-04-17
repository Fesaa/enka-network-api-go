package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FightFestPaperInterview struct {
	Comment        hash.Hash `json:"Comment"`
	Detail         hash.Hash `json:"Detail"`
	IconPath       string    `json:"IconPath"`
	Info           hash.Hash `json:"Info"`
	Name           hash.Hash `json:"Name"`
	PaperID        float64   `json:"PaperID"`
	SortWeight     float64   `json:"SortWeight"`
	TextJoinItemID float64   `json:"TextJoinItemID"`
}
type FightFestPaperInterviewAccessor struct {
	_data               []FightFestPaperInterview
	_dataTextJoinItemID map[float64]FightFestPaperInterview
}

// LoadData retrieves the data. Must be called before FightFestPaperInterview.GroupData
func (a *FightFestPaperInterviewAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestPaperInterview.json")
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
func (a *FightFestPaperInterviewAccessor) Raw() ([]FightFestPaperInterview, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestPaperInterview{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestPaperInterviewAccessor.LoadData to preload everything
func (a *FightFestPaperInterviewAccessor) GroupData() {
	for _, d := range a._data {
		a._dataTextJoinItemID[d.TextJoinItemID] = d
	}
}

// ByTextJoinItemID returns the FightFestPaperInterview uniquely identified by TextJoinItemID
//
// Error is only non-nil if the source errors out
func (a *FightFestPaperInterviewAccessor) ByTextJoinItemID(identifier float64) (FightFestPaperInterview, error) {
	if a._dataTextJoinItemID == nil {
		err := a.LoadData()
		if err != nil {
			return FightFestPaperInterview{}, err
		}
		a.GroupData()
	}
	return a._dataTextJoinItemID[identifier], nil
}
