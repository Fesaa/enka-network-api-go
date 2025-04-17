package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ClockParkChapterConfig struct {
	ChapterAutoUnlock              float64   `json:"ChapterAutoUnlock"`
	ChapterGamePlayRoundRandomList []float64 `json:"ChapterGamePlayRoundRandomList"`
	ChapterID                      float64   `json:"ChapterID"`
	ChapterRoundIDList             []float64 `json:"ChapterRoundIDList"`
	ChapterStoryIDList             []float64 `json:"ChapterStoryIDList"`
	ChapterTitle                   hash.Hash `json:"ChapterTitle"`
	ChapterType                    string    `json:"ChapterType"`
	CheckPointList                 []float64 `json:"CheckPointList"`
	NextChapterID                  []float64 `json:"NextChapterID"`
	RewardID                       float64   `json:"RewardID"`
	RewardProgress                 float64   `json:"RewardProgress"`
	SuccessToRoundID               float64   `json:"SuccessToRoundID"`
}
type ClockParkChapterConfigAccessor struct {
	_data          []ClockParkChapterConfig
	_dataChapterID map[float64]ClockParkChapterConfig
}

// LoadData retrieves the data. Must be called before ClockParkChapterConfig.GroupData
func (a *ClockParkChapterConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkChapterConfig.json")
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
func (a *ClockParkChapterConfigAccessor) Raw() ([]ClockParkChapterConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkChapterConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkChapterConfigAccessor.LoadData to preload everything
func (a *ClockParkChapterConfigAccessor) GroupData() {
	a._dataChapterID = map[float64]ClockParkChapterConfig{}
	for _, d := range a._data {
		a._dataChapterID[d.ChapterID] = d
	}
}

// ByChapterID returns the ClockParkChapterConfig uniquely identified by ChapterID
//
// Error is only non-nil if the source errors out
func (a *ClockParkChapterConfigAccessor) ByChapterID(identifier float64) (ClockParkChapterConfig, error) {
	if a._dataChapterID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkChapterConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChapterID[identifier], nil
}
