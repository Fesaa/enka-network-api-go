package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AchievementData struct {
	AchievementDesc     hash.Hash    `json:"AchievementDesc"`
	AchievementDescPS   hash.Hash    `json:"AchievementDescPS"`
	AchievementID       float64      `json:"AchievementID"`
	AchievementTitle    hash.Hash    `json:"AchievementTitle"`
	AchievementTitlePS  hash.Hash    `json:"AchievementTitlePS"`
	HideAchievementDesc hash.Hash    `json:"HideAchievementDesc"`
	LinearQuestID       float64      `json:"LinearQuestID"`
	PSTrophyID          string       `json:"PSTrophyID"`
	ParamList           []hash.Value `json:"ParamList"`
	Priority            float64      `json:"Priority"`
	QuestID             float64      `json:"QuestID"`
	Rarity              string       `json:"Rarity"`
	RecordText          hash.Hash    `json:"RecordText"`
	RecordType          string       `json:"RecordType"`
	SeriesID            float64      `json:"SeriesID"`
	ShowType            string       `json:"ShowType"`
}
type AchievementDataAccessor struct {
	_data              []AchievementData
	_dataAchievementID map[float64]AchievementData
	_dataLinearQuestID map[float64]AchievementData
	_dataQuestID       map[float64]AchievementData
}

// LoadData retrieves the data. Must be called before AchievementData.GroupData
func (a *AchievementDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AchievementData.json")
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
func (a *AchievementDataAccessor) Raw() ([]AchievementData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AchievementData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AchievementDataAccessor.LoadData to preload everything
func (a *AchievementDataAccessor) GroupData() {
	a._dataAchievementID = map[float64]AchievementData{}
	a._dataLinearQuestID = map[float64]AchievementData{}
	a._dataQuestID = map[float64]AchievementData{}
	for _, d := range a._data {
		a._dataAchievementID[d.AchievementID] = d
		a._dataLinearQuestID[d.LinearQuestID] = d
		a._dataQuestID[d.QuestID] = d
	}
}

// ByAchievementID returns the AchievementData uniquely identified by AchievementID
//
// Error is only non-nil if the source errors out
func (a *AchievementDataAccessor) ByAchievementID(identifier float64) (AchievementData, error) {
	if a._dataAchievementID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AchievementData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAchievementID[identifier], nil
}

// ByLinearQuestID returns the AchievementData uniquely identified by LinearQuestID
//
// Error is only non-nil if the source errors out
func (a *AchievementDataAccessor) ByLinearQuestID(identifier float64) (AchievementData, error) {
	if a._dataLinearQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AchievementData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLinearQuestID[identifier], nil
}

// ByQuestID returns the AchievementData uniquely identified by QuestID
//
// Error is only non-nil if the source errors out
func (a *AchievementDataAccessor) ByQuestID(identifier float64) (AchievementData, error) {
	if a._dataQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AchievementData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestID[identifier], nil
}
