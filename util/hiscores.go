package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/model"
	"strconv"
	"strings"
	"time"
)

const (
	HiscoresApiUrl      = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws"
	HiscoresFriendlyUrl = "https://secure.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws"
)

func ParseHiscore(username, result *string) (hs *model.Hiscore) {
	split := strings.Split(*result, "\n")
	return ParseHiscoreLines(username, &split)
}

func ParseHiscoreLines(username *string, lines *[]string) (hs *model.Hiscore) {
	hs = &model.Hiscore{Username: *username}

	for i, line := range *lines {
		if len(line) == 0 {
			continue
		}

		name := GetHiscoreName(i)

		if i < bhhunter {
			rank, level, xp := parseSkillLine(line)

			skill := model.HiscoreSkill{
				Name:       name,
				Rank:       rank,
				Level:      level,
				Experience: xp,
			}

			hs.Skills = append(hs.Skills, skill)
		} else {
			rank, score := parseMinigameLine(line)

			minigame := model.HiscoreMinigame{
				Name:  name,
				Rank:  rank,
				Score: score,
			}

			hs.Minigames = append(hs.Minigames, minigame)
		}

	}

	return hs
}

func GenerateHiscoresEmbed(hs *model.Hiscore) (embed *discordgo.MessageEmbed) {
	embed = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: hs.Username,
			URL:  *GetFriendlyHiscoresUrl(&hs.Username),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
	}

	var fields []*discordgo.MessageEmbedField

	for i, skill := range hs.Skills {
		field := &discordgo.MessageEmbedField{
			Name:   "\u200b",
			Value:  GetHiscoreEmoji(i) + " " + strconv.Itoa(skill.Level),
			Inline: true,
		}
		fields = append(fields, field)
	}

	embed.Fields = fields

	return embed
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

const (
	overall      = 0
	attack       = 1
	defence      = 2
	strength     = 3
	hitpoints    = 4
	ranged       = 5
	prayer       = 6
	magic        = 7
	cooking      = 8
	woodcutting  = 9
	fletching    = 10
	fishing      = 11
	firemaking   = 12
	crafting     = 13
	smithing     = 14
	mining       = 15
	herblore     = 16
	agility      = 17
	thieving     = 18
	slayer       = 19
	farming      = 20
	runecraft    = 21
	hunter       = 22
	construction = 23
	bhhunter     = 24
	bhrogue      = 25
	cluesall     = 26
	clueseasy    = 27
	cluesmedium  = 28
	clueshard    = 29
	clueselite   = 30
	cluesmaster  = 31
	lmsrank      = 32
)

func GetHiscoreName(index int) (name string) {

	switch index {
	case overall:
		return "Overall"
	case attack:
		return "Attack"
	case defence:
		return "Defence"
	case strength:
		return "Strength"
	case hitpoints:
		return "Hitpoints"
	case ranged:
		return "Ranged"
	case prayer:
		return "Prayer"
	case magic:
		return "Magic"
	case cooking:
		return "Cooking"
	case woodcutting:
		return "Woodcutting"
	case fletching:
		return "Fletching"
	case fishing:
		return "Fishing"
	case firemaking:
		return "Firemaking"
	case crafting:
		return "Crafting"
	case smithing:
		return "Smithing"
	case mining:
		return "Mining"
	case herblore:
		return "Herblore"
	case agility:
		return "Agility"
	case thieving:
		return "Thieving"
	case slayer:
		return "Slayer"
	case farming:
		return "Farming"
	case runecraft:
		return "Runecraft"
	case hunter:
		return "Hunter"
	case construction:
		return "Construction"
	case bhhunter:
		return "BH/Hunter"
	case bhrogue:
		return "BH/Rogue"
	case cluesall:
		return "Clues/All"
	case clueseasy:
		return "Clues/Easy"
	case cluesmedium:
		return "Clues/Medium"
	case clueshard:
		return "Clues/Hard"
	case clueselite:
		return "Clues/Elite"
	case cluesmaster:
		return "Clues/Master"
	case lmsrank:
		return "LMS/Rank"
	default:
		return "Error"
	}

}

func GetHiscoreEmoji(index int) (emoji string) {

	switch index {
	case overall:
		return "<:stats:529107862316908564>"
	case attack:
		return "<:attack:529105287664369674>"
	case defence:
		return "<:defence:529105287773421568>"
	case strength:
		return "<:strength:529105288020754452>"
	case hitpoints:
		return "<:hitpoints:529105289316663307>"
	case ranged:
		return "<:ranged:529105287534346246>"
	case prayer:
		return "<:prayer:529105287857307658>"
	case magic:
		return "<:magic:529105287861501989>"
	case cooking:
		return "<:cooking:529105287706312704>"
	case woodcutting:
		return "<:woodcutting:529105287681015831>"
	case fletching:
		return "<:fletching:529105287852982272>"
	case fishing:
		return "<:fishing:529105287878017051>"
	case firemaking:
		return "<:firemaking:529105287790067742>"
	case crafting:
		return "<:crafting:529105287727284224>"
	case smithing:
		return "<:smithing:529105287798325284>"
	case mining:
		return "<:mining:529105287819427850>"
	case herblore:
		return "<:herblore:529105287966097462>"
	case agility:
		return "<:agility:529105287718895616>"
	case thieving:
		return "<:thieving:529105287882342411>"
	case slayer:
		return "<:slayer:529105287488208898>"
	case farming:
		return "<:farming:529105287521501185>"
	case runecraft:
		return "<:runecrafting:529105287806976001>"
	case hunter:
		return "<:hunter:529105287601455119>"
	case construction:
		return "<:construction:529105287651786753>"
	case bhhunter:
		return "<:bh_hunter:529105998980448256>"
	case bhrogue:
		return "<:bh_rogue:529105998892236868>"
	case cluesall:
		return "<:cluescroll:529106218753720320>"
	case clueseasy:
		return "<:cluescroll:529106218753720320>"
	case cluesmedium:
		return "<:cluescroll:529106218753720320>"
	case clueshard:
		return "<:cluescroll:529106218753720320>"
	case clueselite:
		return "<:cluescroll:529106218753720320>"
	case cluesmaster:
		return "<:cluescroll:529106218753720320>"
	case lmsrank:
		return "<:lms:529108371811467264>"
	default:
		return ""
	}
}
