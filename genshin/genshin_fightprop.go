package genshin

type FightProp struct {
	Id   string
	Name string
}

var fightProps = []FightProp{
	{"1", "Base HP"},
	{"4", "Base ATK"},
	{"7", "Base DEF"},
	{"10", "Base SPD"},
	{"3", "HP%"},
	{"6", "ATK%"},
	{"9", "DEF%"},
	{"11", "SPD%"},
	{"2", "HP"},
	{"5", "ATK"},
	{"8", "DEF"},
	{"20", "CRIT Rate"},
	{"22", "CRIT DMG"},
	{"23", "Energy Recharge"},
	{"26", "Healing Bonus"},
	{"27", "Incoming Healing Bonus"},
	{"28", "Elemental Mastery"},
	{"80", "Cooldown reduction"},
	{"81", "Shield Strength"},
	{"1010", "Current HP"},
	{"2000", "Max HP"},
	{"30", "Physical DMG Bonus"},
	{"40", "Pyro DMG Bonus"},
	{"41", "Electro DMG Bonus"},
	{"42", "Hydro DMG Bonus"},
	{"43", "Dendro DMG Bonus"},
	{"44", "Anemo DMG Bonus"},
	{"45", "Geo DMG Bonus"},
	{"46", "Cryo DMG Bonus"},
	{"29", "Physical RES"},
	{"50", "Pyro RES"},
	{"51", "Electro RES"},
	{"52", "Hydro RES"},
	{"53", "Dendro RES"},
	{"54", "Anemo RES"},
	{"55", "Geo RES"},
	{"56", "Cryo RES"},
	{"70", "Pyro Energy Cost"},
	{"71", "Electro Energy Cost"},
	{"72", "Hydro Energy Cost"},
	{"73", "Dendro Energy Cost"},
	{"74", "Anemo Energy Cost"},
	{"75", "Cryo Energy Cost"},
	{"76", "Geo Energy Cost"},
	{"1000", "Current Pyro Energy"},
	{"1001", "Current Electro Energy"},
	{"1002", "Current Hydro Energy"},
	{"1003", "Current Dendro Energy"},
	{"1004", "Current Anemo Energy"},
	{"1005", "Current Cryo Energy"},
	{"1006", "Current Geo Energy"},
	{"2001", "ATK"},
	{"2002", "DEF"},
	{"2003", "SPD"},
	{"3025", "Elemental reaction CRIT Rate"},
	{"3026", "Elemental reaction CRIT DMG"},
	{"3027", "(Overloaded) CRIT Rate"},
	{"3028", "(Overloaded) CRIT DMG"},
	{"3029", "(Swirl) CRIT Rate"},
	{"3030", "(Swirl) CRIT DMG"},
	{"3031", "(Electro-Charged) CRIT Rate"},
	{"3032", "(Electro-Charged) CRIT DMG"},
	{"3033", "(Superconduct) CRIT Rate"},
	{"3034", "(Superconduct) CRIT DMG"},
	{"3035", "(Burning) CRIT Rate"},
	{"3036", "(Burning) CRIT DMG"},
	{"3037", "(Frozen/Shattered) CRIT Rate"},
	{"3038", "(Frozen/Shattered) CRIT DMG"},
	{"3039", "(Bloom) CRIT Rate"},
	{"3040", "(Bloom) CRIT DMG"},
	{"3041", "(Burgeon) CRIT Rate"},
	{"3042", "(Burgeon) CRIT DMG"},
	{"3043", "(Hyperbloom) CRIT Rate"},
	{"3044", "(Hyperbloom) CRIT DMG"},
	{"3045", "Base Elemental reaction CRIT Rate"},
	{"3046", "Base Elemental reaction CRIT DMG"},
}

func fromId(id string) *FightProp {
	for _, prop := range fightProps {
		if prop.Id == id {
			return &prop
		}
	}
	return nil
}
