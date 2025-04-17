package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlayerReturnConfig struct {
	ActivityModuleID      float64                                   `json:"ActivityModuleID"`
	DailyDoubleTime       float64                                   `json:"DailyDoubleTime"`
	DispatchLink          string                                    `json:"DispatchLink"`
	ExtraMultipleDropList []PlayerReturnConfigExtraMultipleDropList `json:"ExtraMultipleDropList"`
	FarmMultipleDropID    float64                                   `json:"FarmMultipleDropID"`
	KeyPointID            []float64                                 `json:"KeyPointID"`
	LimitTime             float64                                   `json:"LimitTime"`
	LoginReward           []float64                                 `json:"LoginReward"`
	PlayerReturnID        float64                                   `json:"PlayerReturnID"`
	QuestGroupID          []float64                                 `json:"QuestGroupID"`
	RecommendActivity     []interface{}                             `json:"RecommendActivity"`
	RecommendAvatar       []interface{}                             `json:"RecommendAvatar"`
	RecommendMission      []interface{}                             `json:"RecommendMission"`
	ReturnRewardID        float64                                   `json:"ReturnRewardID"`
	TotalDoubleTime       float64                                   `json:"TotalDoubleTime"`
	ValidityPeriod        float64                                   `json:"ValidityPeriod"`
}
type PlayerReturnConfigExtraMultipleDropList struct {
	DMNHECGLFNM float64 `json:"DMNHECGLFNM"`
	HIDNKBOMEPO float64 `json:"HIDNKBOMEPO"`
	PDLDCNMHPGN float64 `json:"PDLDCNMHPGN"`
}
type PlayerReturnConfigAccessor struct {
	_data               []PlayerReturnConfig
	_dataPlayerReturnID map[float64]PlayerReturnConfig
}

// LoadData retrieves the data. Must be called before PlayerReturnConfig.GroupData
func (a *PlayerReturnConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnConfig.json")
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
func (a *PlayerReturnConfigAccessor) Raw() ([]PlayerReturnConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnConfigAccessor.LoadData to preload everything
func (a *PlayerReturnConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPlayerReturnID[d.PlayerReturnID] = d
	}
}

// ByPlayerReturnID returns the PlayerReturnConfig uniquely identified by PlayerReturnID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnConfigAccessor) ByPlayerReturnID(identifier float64) (PlayerReturnConfig, error) {
	if a._dataPlayerReturnID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPlayerReturnID[identifier], nil
}
