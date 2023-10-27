package enkanetworkapigo

type StarRailUserCharacter struct {
	AvatarId  int
	Eidolon   int
	Level     int
	Ascension int
	LightCone *StarRailLightCone
	Relics    []StarRailRelic
}

func (c *StarRailUserCharacter) GetGameData() {

}
