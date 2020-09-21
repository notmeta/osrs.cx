package model

type HiscoreSkill struct {
	Name       string
	Rank       int
	Level      int
	Experience int
}

type HiscoreActivity struct {
	Name  string
	Rank  int
	Score int
}

type Hiscore struct {
	Username   string
	Skills     []HiscoreSkill
	Activities []HiscoreActivity
}

const (
	HiscoresApiUrl      = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws"
	HiscoresFriendlyUrl = "https://secure.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws"
)

const (
	MinigamesEmoji = "<:minigames:530206797215301632>"
)

const (
	Overall = iota
	Attack
	Defence
	Strength
	Hitpoints
	Ranged
	Prayer
	Magic
	Cooking
	Woodcutting
	Fletching
	Fishing
	Firemaking
	Crafting
	Smithing
	Mining
	Herblore
	Agility
	Thieving
	Slayer
	Farming
	Runecraft
	Hunter
	Construction
	League
	Bhhunter
	Bhrogue
	Cluesall
	Cluesbeginner
	Clueseasy
	Cluesmedium
	Clueshard
	Clueselite
	Cluesmaster
	Lmsrank
	Sire
	Hydra
	Barrows
	Bryophyta
	Callisto
	Cerberus
	Xeric
	Xericchallenge
	Chaoselemental
	Chaosfanatic
	Zilyana
	Corporealbeast
	Crazyarchaeologist
	Dkprime
	Dkrex
	Dksupreme
	Derangedarchaeologist
	Graardor
	Mole
	Grotesqueguardians
	Hespori
	Kalphitequeen
	Kingblackdragon
	Kraken
	Kreearra
	Kril
	Mimic
	Nightmare
	Obor
	Sarachnis
	Scorpia
	Skotizo
	Gauntlet
	Corruptedgauntlet
	Theatreofblood
	Thermonucleardevil
	Tzkalzuk
	Tztokjad
	Venenatis
	Vetion
	Vorkath
	Wintertodt
	Zulcano
	Zulrah
	ActivityOffset = League
	BossOffset     = Sire
)

func GetHiscoreName(index int) (name string) {
	switch index {
	case Overall:
		return "Overall"
	case Attack:
		return "Attack"
	case Defence:
		return "Defence"
	case Strength:
		return "Strength"
	case Hitpoints:
		return "Hitpoints"
	case Ranged:
		return "Ranged"
	case Prayer:
		return "Prayer"
	case Magic:
		return "Magic"
	case Cooking:
		return "Cooking"
	case Woodcutting:
		return "Woodcutting"
	case Fletching:
		return "Fletching"
	case Fishing:
		return "Fishing"
	case Firemaking:
		return "Firemaking"
	case Crafting:
		return "Crafting"
	case Smithing:
		return "Smithing"
	case Mining:
		return "Mining"
	case Herblore:
		return "Herblore"
	case Agility:
		return "Agility"
	case Thieving:
		return "Thieving"
	case Slayer:
		return "Slayer"
	case Farming:
		return "Farming"
	case Runecraft:
		return "Runecraft"
	case Hunter:
		return "Hunter"
	case Construction:
		return "Construction"
	case Bhhunter:
		return "BH/Hunter"
	case Bhrogue:
		return "BH/Rogue"
	case Cluesall:
		return "Clues/All"
	case Cluesbeginner:
		return "Clues/Beginner"
	case Clueseasy:
		return "Clues/Easy"
	case Cluesmedium:
		return "Clues/Medium"
	case Clueshard:
		return "Clues/Hard"
	case Clueselite:
		return "Clues/Elite"
	case Cluesmaster:
		return "Clues/Master"
	case Lmsrank:
		return "LMS/Rank"
	case Sire:
		return "Abyssal Sire"
	case Hydra:
		return "Alchemical Hydra"
	case Barrows:
		return "Barrows Chests"
	case Bryophyta:
		return "Bryophyta"
	case Callisto:
		return "Callisto"
	case Cerberus:
		return "Cerberus"
	case Xeric:
		return "Chambers of Xeric"
	case Xericchallenge:
		return "Chambers of Xeric: Challenge Mode"
	case Chaoselemental:
		return "Chaos Elemental"
	case Chaosfanatic:
		return "Chaos Fanatic"
	case Zilyana:
		return "Commander Zilyana"
	case Corporealbeast:
		return "Corporeal Beast"
	case Crazyarchaeologist:
		return "Crazy Archaeologist"
	case Dkprime:
		return "Dagannoth Prime"
	case Dkrex:
		return "Dagannoth Rex"
	case Dksupreme:
		return "Dagannoth Supreme"
	case Derangedarchaeologist:
		return "Deranged Archaeologist"
	case Graardor:
		return "General Graardor"
	case Mole:
		return "Giant Mole"
	case Grotesqueguardians:
		return "Grotesque Guardians"
	case Hespori:
		return "Hespori"
	case Kalphitequeen:
		return "Kalphite Queen"
	case Kingblackdragon:
		return "King Black Dragon"
	case Kraken:
		return "Kraken"
	case Kreearra:
		return "Kree'arra"
	case Kril:
		return "K'ril Tsutsaroth"
	case Mimic:
		return "Mimic"
	case Nightmare:
		return "Nightmare"
	case Obor:
		return "Obor"
	case Sarachnis:
		return "Sarachnis"
	case Scorpia:
		return "Scorpia"
	case Skotizo:
		return "Skotizo"
	case Gauntlet:
		return "The Gauntlet"
	case Corruptedgauntlet:
		return "The Corrupted Gauntlet"
	case Theatreofblood:
		return "Theatre of Blood"
	case Thermonucleardevil:
		return "Thermonuclear Smoke Devil"
	case Tzkalzuk:
		return "TzKal-Zuk"
	case Tztokjad:
		return "TzTok-Jad"
	case Venenatis:
		return "Venenatis"
	case Vetion:
		return "Vet'ion"
	case Vorkath:
		return "Vorkath"
	case Wintertodt:
		return "Wintertodt"
	case Zulcano:
		return "Zulcano"
	case Zulrah:
		return "Zulrah"
	default:
		return "Error"
	}
}

func GetHiscoreEmoji(index int) (emoji string) {
	switch index {
	case Overall:
		return "<:stats:529107862316908564>"
	case Attack:
		return "<:Attack:529105287664369674>"
	case Defence:
		return "<:Defence:529105287773421568>"
	case Strength:
		return "<:Strength:529105288020754452>"
	case Hitpoints:
		return "<:Hitpoints:529105289316663307>"
	case Ranged:
		return "<:Ranged:529105287534346246>"
	case Prayer:
		return "<:Prayer:529105287857307658>"
	case Magic:
		return "<:Magic:529105287861501989>"
	case Cooking:
		return "<:Cooking:529105287706312704>"
	case Woodcutting:
		return "<:Woodcutting:529105287681015831>"
	case Fletching:
		return "<:Fletching:529105287852982272>"
	case Fishing:
		return "<:Fishing:529105287878017051>"
	case Firemaking:
		return "<:Firemaking:529105287790067742>"
	case Crafting:
		return "<:Crafting:529105287727284224>"
	case Smithing:
		return "<:Smithing:529105287798325284>"
	case Mining:
		return "<:Mining:529105287819427850>"
	case Herblore:
		return "<:Herblore:529105287966097462>"
	case Agility:
		return "<:Agility:529105287718895616>"
	case Thieving:
		return "<:Thieving:529105287882342411>"
	case Slayer:
		return "<:Slayer:529105287488208898>"
	case Farming:
		return "<:Farming:529105287521501185>"
	case Runecraft:
		return "<:runecrafting:529105287806976001>"
	case Hunter:
		return "<:Hunter:529105287601455119>"
	case Construction:
		return "<:Construction:529105287651786753>"
	case Bhhunter:
		return "<:bh_hunter:529105998980448256>"
	case Bhrogue:
		return "<:bh_rogue:529105998892236868>"
	case Cluesall:
		return "<:cluescroll:529106218753720320>"
	case Cluesbeginner:
		return "<:cluescroll:529106218753720320>"
	case Clueseasy:
		return "<:cluescroll:529106218753720320>"
	case Cluesmedium:
		return "<:cluescroll:529106218753720320>"
	case Clueshard:
		return "<:cluescroll:529106218753720320>"
	case Clueselite:
		return "<:cluescroll:529106218753720320>"
	case Cluesmaster:
		return "<:cluescroll:529106218753720320>"
	case Lmsrank:
		return "<:lms:529108371811467264>"
	default:
		return ""
	}
}
