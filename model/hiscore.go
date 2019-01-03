package model

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"math"
	"strconv"
	"strings"
	"time"
)

type HiscoreSkill struct {
	Name       string
	Rank       int
	Level      int
	Experience int
}

type HiscoreMinigame struct {
	Name  string
	Rank  int
	Score int
}

type Hiscore struct {
	Username  string
	Skills    []HiscoreSkill
	Minigames []HiscoreMinigame
}

const (
	HiscoresApiUrl      = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws"
	HiscoresFriendlyUrl = "https://secure.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws"
)

const (
	MinigamesEmoji = "<:minigames:530206797215301632>"
)

const (
	Overall      = 0
	Attack       = 1
	Defence      = 2
	Strength     = 3
	Hitpoints    = 4
	Ranged       = 5
	Prayer       = 6
	Magic        = 7
	Cooking      = 8
	Woodcutting  = 9
	Fletching    = 10
	Fishing      = 11
	Firemaking   = 12
	Crafting     = 13
	Smithing     = 14
	Mining       = 15
	Herblore     = 16
	Agility      = 17
	Thieving     = 18
	Slayer       = 19
	Farming      = 20
	Runecraft    = 21
	Hunter       = 22
	Construction = 23
	Bhhunter     = 24
	Bhrogue      = 25
	Cluesall     = 26
	Clueseasy    = 27
	Cluesmedium  = 28
	Clueshard    = 29
	Clueselite   = 30
	Cluesmaster  = 31
	Lmsrank      = 32
	Offset       = Bhhunter
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

func (hs *Hiscore) GenerateHiscoresEmbed() (embed *discordgo.MessageEmbed) {
	embed = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: hs.Username,
			URL:  *GetFriendlyHiscoresUrl(&hs.Username),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Color: 0x00FF00,
	}

	embed.Fields = append(embed.Fields,
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d",
				GetHiscoreEmoji(Attack), hs.Skills[Attack].Level,
				GetHiscoreEmoji(Strength), hs.Skills[Strength].Level,
				GetHiscoreEmoji(Defence), hs.Skills[Defence].Level,
				GetHiscoreEmoji(Ranged), hs.Skills[Ranged].Level,
				GetHiscoreEmoji(Prayer), hs.Skills[Prayer].Level,
				GetHiscoreEmoji(Magic), hs.Skills[Magic].Level,
				GetHiscoreEmoji(Runecraft), hs.Skills[Runecraft].Level,
				GetHiscoreEmoji(Construction), hs.Skills[Construction].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d",
				GetHiscoreEmoji(Hitpoints), hs.Skills[Hitpoints].Level,
				GetHiscoreEmoji(Agility), hs.Skills[Agility].Level,
				GetHiscoreEmoji(Herblore), hs.Skills[Herblore].Level,
				GetHiscoreEmoji(Thieving), hs.Skills[Thieving].Level,
				GetHiscoreEmoji(Crafting), hs.Skills[Crafting].Level,
				GetHiscoreEmoji(Fletching), hs.Skills[Fletching].Level,
				GetHiscoreEmoji(Slayer), hs.Skills[Slayer].Level,
				GetHiscoreEmoji(Hunter), hs.Skills[Hunter].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d",
				GetHiscoreEmoji(Mining), hs.Skills[Mining].Level,
				GetHiscoreEmoji(Smithing), hs.Skills[Smithing].Level,
				GetHiscoreEmoji(Fishing), hs.Skills[Fishing].Level,
				GetHiscoreEmoji(Cooking), hs.Skills[Cooking].Level,
				GetHiscoreEmoji(Firemaking), hs.Skills[Firemaking].Level,
				GetHiscoreEmoji(Woodcutting), hs.Skills[Woodcutting].Level,
				GetHiscoreEmoji(Farming), hs.Skills[Farming].Level,
				GetHiscoreEmoji(Overall), hs.Skills[Overall].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Overall", GetHiscoreEmoji(Overall)),
			Value: fmt.Sprintf(
				"**Rank:** %s\n**Level:** %s\n**Exp:** %s\n\n**Combat Level:** %s",
				util.RenderInteger("#,###.", hs.Skills[Overall].Rank),
				util.RenderInteger("#,###.", hs.Skills[Overall].Level),
				util.RenderInteger("#,###.", hs.Skills[Overall].Experience),
				util.RenderFloat("", CalculateCombatLevel(hs)),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Minigames", MinigamesEmoji),
			Value: fmt.Sprintf(
				"%s **BH/Hunter:** %s\n%s **BH/Rogue:** %s\n%s **LMS:** %s",
				GetHiscoreEmoji(Bhhunter), util.RenderInteger("#,###.", hs.Minigames[Bhhunter-Offset].Score),
				GetHiscoreEmoji(Bhrogue), util.RenderInteger("#,###.", hs.Minigames[Bhrogue-Offset].Score),
				GetHiscoreEmoji(Lmsrank), util.RenderInteger("#,###.", hs.Minigames[Lmsrank-Offset].Score),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Clues", GetHiscoreEmoji(Cluesall)),
			Value: fmt.Sprintf(
				"**All:** %s\n**Easy:** %s\n**Medium:** %s\n**Hard:** %s\n**Elite:** %s\n**Master:** %s",
				util.RenderInteger("#,###.", hs.Minigames[Cluesall-Offset].Score),
				util.RenderInteger("#,###.", hs.Minigames[Clueseasy-Offset].Score),
				util.RenderInteger("#,###.", hs.Minigames[Cluesmedium-Offset].Score),
				util.RenderInteger("#,###.", hs.Minigames[Clueshard-Offset].Score),
				util.RenderInteger("#,###.", hs.Minigames[Clueselite-Offset].Score),
				util.RenderInteger("#,###.", hs.Minigames[Cluesmaster-Offset].Score),
			),
			Inline: true,
		},
	)

	return embed
}

func ParseHiscore(username, result *string) (hs *Hiscore) {
	split := strings.Split(*result, "\n")
	return ParseHiscoreLines(username, &split)
}

func ParseHiscoreLines(username *string, lines *[]string) (hs *Hiscore) {
	hs = &Hiscore{Username: *username}

	for i, line := range *lines {
		if len(line) == 0 {
			continue
		}

		name := GetHiscoreName(i)

		if i < Bhhunter {
			rank, level, xp := parseSkillLine(line)

			skill := HiscoreSkill{
				Name:       name,
				Rank:       rank,
				Level:      level,
				Experience: xp,
			}

			hs.Skills = append(hs.Skills, skill)
		} else {
			rank, score := parseMinigameLine(line)

			minigame := HiscoreMinigame{
				Name:  name,
				Rank:  rank,
				Score: score,
			}

			hs.Minigames = append(hs.Minigames, minigame)
		}

	}

	return hs
}

func GetHiscoresApiUrl(username *string) *string {
	url := fmt.Sprintf("%s?player=%s", HiscoresApiUrl, *username)
	return &url
}

func GetFriendlyHiscoresUrl(username *string) *string {
	url := fmt.Sprintf("%s?user1=%s", HiscoresFriendlyUrl, *username)
	return &url
}

func parseSkillLine(line string) (rank, level, xp int) {
	split := strings.Split(line, ",")

	rank, _ = strconv.Atoi(split[0])
	level, _ = strconv.Atoi(split[1])
	xp, _ = strconv.Atoi(split[2])

	return rank, level, xp
}

func parseMinigameLine(line string) (rank, score int) {
	split := strings.Split(line, ",")

	rank, _ = strconv.Atoi(split[0])
	score, _ = strconv.Atoi(split[1])

	return rank, score
}

func CalculateCombatLevel(hs *Hiscore) float64 {
	base := 0.25 * float64(hs.Skills[Defence].Level+hs.Skills[Hitpoints].Level+(hs.Skills[Prayer].Level/2))
	melee := 0.325 * float64(hs.Skills[Attack].Level+hs.Skills[Strength].Level)
	ranged := 0.325 * float64((hs.Skills[Ranged].Level/2)+hs.Skills[Ranged].Level)
	magic := 0.325 * float64((hs.Skills[Magic].Level/2)+hs.Skills[Magic].Level)

	return base + (math.Max(melee, math.Max(magic, ranged)))
}
