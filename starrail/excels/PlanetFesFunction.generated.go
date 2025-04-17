package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PlanetFesFunction struct {
	Description  map[string]json.Number `json:"Description"`
	FunctionType string                 `json:"FunctionType"`
	ParamList    []json.Number          `json:"ParamList"`
	SkillID      json.Number            `json:"SkillID"`
}
type PlanetFesFunctionAccessor struct {
	_data []PlanetFesFunction
}

// LoadData retrieves the data. Must be called before PlanetFesFunction.GroupData
func (a *PlanetFesFunctionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesFunction.json")
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
func (a *PlanetFesFunctionAccessor) Raw() ([]PlanetFesFunction, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesFunction{}, err
		}
	}
	return a._data, nil
}
