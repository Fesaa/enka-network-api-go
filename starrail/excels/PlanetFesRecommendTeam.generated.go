package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PlanetFesRecommendTeam struct {
	Business   []map[string]json.Number `json:"Business"`
	Exhibition []map[string]json.Number `json:"Exhibition"`
	FesLevel   json.Number              `json:"FesLevel"`
	Game       []map[string]json.Number `json:"Game"`
	ID         json.Number              `json:"ID"`
}
type PlanetFesRecommendTeamAccessor struct {
	_data []PlanetFesRecommendTeam
}

// LoadData retrieves the data. Must be called before PlanetFesRecommendTeam.GroupData
func (a *PlanetFesRecommendTeamAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesRecommendTeam.json")
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
func (a *PlanetFesRecommendTeamAccessor) Raw() ([]PlanetFesRecommendTeam, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesRecommendTeam{}, err
		}
	}
	return a._data, nil
}
