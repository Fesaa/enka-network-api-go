package genshin

import "github.com/Fesaa/enka-network-api-go/localization"

type ArtifactType int

const (
	FLOWER ArtifactType = iota
	FEATHER
	SANDS
	GOBLET
	CIRCLET
)

type Artifact struct {
	ArtifactType ArtifactType
	Level        int
	MainStats    ArtifactStat
	SubStats     []ArtifactStat
	SetNameHash  string
	IconKey      string
}

func (a *Artifact) GetHash() string {
	return a.SetNameHash
}

// SetName returns the name of the artifact set
func (a *Artifact) SetName() string {
	return localization.GetGenshinLocaleOrHash(a)
}

type ArtifactStat struct {
	Stat  string
	Value float64
}

func TypeFromId(identifier string) ArtifactType {
	switch identifier {
	case "EQUIP_BRACER":
		return FLOWER
	case "EQUIP_NECKLACE":
		return FEATHER
	case "EQUIP_RING":
		return GOBLET
	case "EQUIP_SHOES":
		return SANDS
	case "EQUIP_DRESS":
		return CIRCLET
	}
	return -1
}
