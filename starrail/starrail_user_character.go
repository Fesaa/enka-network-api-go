package starrail

type UserCharacter struct {
	AvatarId  int
	Eidolon   int
	Level     int
	Ascension int
	LightCone *LightCone
	Relics    []Relic
}
