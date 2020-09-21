package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/model"
	"math"
	"strconv"
	"strings"
	"time"
)

type Hiscore struct {
	*model.Hiscore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumBossKills(hs *Hiscore) int {
	result := 0
	for i := model.BossOffset - model.ActivityOffset; i < len(hs.Activities); i++ {
		result += hs.Activities[i].Score
	}
	return result
}

func (hs *Hiscore) GenerateBossEmbed() (embed *discordgo.MessageEmbed) {
	embed = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: *hs.GetFriendlyUsername(),
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
				"**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s",
				model.GetHiscoreName(model.Sire), RenderInteger("#,###.", hs.Activities[model.Sire-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Hydra), RenderInteger("#,###.", hs.Activities[model.Hydra-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Barrows), RenderInteger("#,###.", hs.Activities[model.Barrows-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Bryophyta), RenderInteger("#,###.", hs.Activities[model.Bryophyta-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Callisto), RenderInteger("#,###.", hs.Activities[model.Callisto-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Xeric), RenderInteger("#,###.", hs.Activities[model.Xeric-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Xericchallenge), RenderInteger("#,###.", hs.Activities[model.Xericchallenge-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Chaoselemental), RenderInteger("#,###.", hs.Activities[model.Chaoselemental-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Chaosfanatic), RenderInteger("#,###.", hs.Activities[model.Chaosfanatic-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Zilyana), RenderInteger("#,###.", hs.Activities[model.Zilyana-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Corporealbeast), RenderInteger("#,###.", hs.Activities[model.Corporealbeast-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Crazyarchaeologist), RenderInteger("#,###.", hs.Activities[model.Crazyarchaeologist-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Dkprime), RenderInteger("#,###.", hs.Activities[model.Dkprime-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Dkrex), RenderInteger("#,###.", hs.Activities[model.Dkrex-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Dksupreme), RenderInteger("#,###.", hs.Activities[model.Dksupreme-model.ActivityOffset].Score),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s",
				model.GetHiscoreName(model.Derangedarchaeologist), RenderInteger("#,###.", hs.Activities[model.Derangedarchaeologist-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Graardor), RenderInteger("#,###.", hs.Activities[model.Graardor-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Mole), RenderInteger("#,###.", hs.Activities[model.Mole-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Grotesqueguardians), RenderInteger("#,###.", hs.Activities[model.Grotesqueguardians-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Hespori), RenderInteger("#,###.", hs.Activities[model.Hespori-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Kalphitequeen), RenderInteger("#,###.", hs.Activities[model.Kalphitequeen-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Kingblackdragon), RenderInteger("#,###.", hs.Activities[model.Kingblackdragon-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Kraken), RenderInteger("#,###.", hs.Activities[model.Kraken-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Kreearra), RenderInteger("#,###.", hs.Activities[model.Kreearra-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Kril), RenderInteger("#,###.", hs.Activities[model.Kril-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Mimic), RenderInteger("#,###.", hs.Activities[model.Mimic-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Nightmare), RenderInteger("#,###.", hs.Activities[model.Nightmare-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Obor), RenderInteger("#,###.", hs.Activities[model.Obor-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Sarachnis), RenderInteger("#,###.", hs.Activities[model.Sarachnis-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Scorpia), RenderInteger("#,###.", hs.Activities[model.Scorpia-model.ActivityOffset].Score),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n**%s:** %s\n",
				model.GetHiscoreName(model.Skotizo), RenderInteger("#,###.", hs.Activities[model.Skotizo-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Gauntlet), RenderInteger("#,###.", hs.Activities[model.Gauntlet-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Corruptedgauntlet), RenderInteger("#,###.", hs.Activities[model.Corruptedgauntlet-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Theatreofblood), RenderInteger("#,###.", hs.Activities[model.Theatreofblood-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Thermonucleardevil), RenderInteger("#,###.", hs.Activities[model.Thermonucleardevil-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Tzkalzuk), RenderInteger("#,###.", hs.Activities[model.Tzkalzuk-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Tztokjad), RenderInteger("#,###.", hs.Activities[model.Tztokjad-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Venenatis), RenderInteger("#,###.", hs.Activities[model.Venenatis-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Vetion), RenderInteger("#,###.", hs.Activities[model.Vetion-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Vorkath), RenderInteger("#,###.", hs.Activities[model.Vorkath-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Wintertodt), RenderInteger("#,###.", hs.Activities[model.Wintertodt-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Zulcano), RenderInteger("#,###.", hs.Activities[model.Zulcano-model.ActivityOffset].Score),
				model.GetHiscoreName(model.Zulrah), RenderInteger("#,###.", hs.Activities[model.Zulrah-model.ActivityOffset].Score),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"**%s:** %s",
				"Total kills", RenderInteger("#,###.", SumBossKills(hs)),
			),
			Inline: true,
		},
	)
	return embed
}

func (hs *Hiscore) GenerateHiscoresEmbed() (embed *discordgo.MessageEmbed) {
	embed = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: *hs.GetFriendlyUsername(),
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
				model.GetHiscoreEmoji(model.Attack), hs.Skills[model.Attack].Level,
				model.GetHiscoreEmoji(model.Strength), hs.Skills[model.Strength].Level,
				model.GetHiscoreEmoji(model.Defence), hs.Skills[model.Defence].Level,
				model.GetHiscoreEmoji(model.Ranged), hs.Skills[model.Ranged].Level,
				model.GetHiscoreEmoji(model.Prayer), hs.Skills[model.Prayer].Level,
				model.GetHiscoreEmoji(model.Magic), hs.Skills[model.Magic].Level,
				model.GetHiscoreEmoji(model.Runecraft), hs.Skills[model.Runecraft].Level,
				model.GetHiscoreEmoji(model.Construction), hs.Skills[model.Construction].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d",
				model.GetHiscoreEmoji(model.Hitpoints), hs.Skills[model.Hitpoints].Level,
				model.GetHiscoreEmoji(model.Agility), hs.Skills[model.Agility].Level,
				model.GetHiscoreEmoji(model.Herblore), hs.Skills[model.Herblore].Level,
				model.GetHiscoreEmoji(model.Thieving), hs.Skills[model.Thieving].Level,
				model.GetHiscoreEmoji(model.Crafting), hs.Skills[model.Crafting].Level,
				model.GetHiscoreEmoji(model.Fletching), hs.Skills[model.Fletching].Level,
				model.GetHiscoreEmoji(model.Slayer), hs.Skills[model.Slayer].Level,
				model.GetHiscoreEmoji(model.Hunter), hs.Skills[model.Hunter].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: "\u200b",
			Value: fmt.Sprintf(
				"%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d\n%s %d",
				model.GetHiscoreEmoji(model.Mining), hs.Skills[model.Mining].Level,
				model.GetHiscoreEmoji(model.Smithing), hs.Skills[model.Smithing].Level,
				model.GetHiscoreEmoji(model.Fishing), hs.Skills[model.Fishing].Level,
				model.GetHiscoreEmoji(model.Cooking), hs.Skills[model.Cooking].Level,
				model.GetHiscoreEmoji(model.Firemaking), hs.Skills[model.Firemaking].Level,
				model.GetHiscoreEmoji(model.Woodcutting), hs.Skills[model.Woodcutting].Level,
				model.GetHiscoreEmoji(model.Farming), hs.Skills[model.Farming].Level,
				model.GetHiscoreEmoji(model.Overall), hs.Skills[model.Overall].Level,
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Overall", model.GetHiscoreEmoji(model.Overall)),
			Value: fmt.Sprintf(
				"**Rank:** %s\n**Level:** %s\n**Exp:** %s\n\n**Combat Level:** %s",
				RenderInteger("#,###.", hs.Skills[model.Overall].Rank),
				RenderInteger("#,###.", hs.Skills[model.Overall].Level),
				RenderInteger("#,###.", hs.Skills[model.Overall].Experience),
				RenderFloat("", CalculateCombatLevel(hs)),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Activities", model.MinigamesEmoji),
			Value: fmt.Sprintf(
				"%s **BH/Hunter:** %s\n%s **BH/Rogue:** %s\n%s **LMS:** %s",
				model.GetHiscoreEmoji(model.Bhhunter), RenderInteger("#,###.", hs.Activities[model.Bhhunter-model.ActivityOffset].Score),
				model.GetHiscoreEmoji(model.Bhrogue), RenderInteger("#,###.", hs.Activities[model.Bhrogue-model.ActivityOffset].Score),
				model.GetHiscoreEmoji(model.Lmsrank), RenderInteger("#,###.", hs.Activities[model.Lmsrank-model.ActivityOffset].Score),
			),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s Clues", model.GetHiscoreEmoji(model.Cluesall)),
			Value: fmt.Sprintf(
				"**All:** %s\n**Beginner:** %s\n**Easy:** %s\n**Medium:** %s\n**Hard:** %s\n**Elite:** %s\n**Master:** %s",
				RenderInteger("#,###.", hs.Activities[model.Cluesall-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Cluesbeginner-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Clueseasy-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Cluesmedium-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Clueshard-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Clueselite-model.ActivityOffset].Score),
				RenderInteger("#,###.", hs.Activities[model.Cluesmaster-model.ActivityOffset].Score),
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
	hs = &Hiscore{&model.Hiscore{Username: *username}}
	for i, line := range *lines {
		if len(line) == 0 {
			continue
		}
		name := model.GetHiscoreName(i)
		if i < model.League {
			rank, level, xp := parseSkillLine(line)
			skill := model.HiscoreSkill{
				Name:       name,
				Rank:       max(0, rank),
				Level:      max(0, level),
				Experience: max(0, xp),
			}
			hs.Skills = append(hs.Skills, skill)
		} else {
			rank, score := parseMinigameLine(line)
			minigame := model.HiscoreActivity{
				Name:  name,
				Rank:  max(0, rank),
				Score: max(0, score),
			}
			hs.Activities = append(hs.Activities, minigame)
		}
	}
	return hs
}

func (hs *Hiscore) GetFriendlyUsername() *string {
	username := strings.Replace(hs.Username, "+", " ", -1)
	return &username
}

func GetHiscoresApiUrl(username *string) *string {
	url := fmt.Sprintf("%s?player=%s", model.HiscoresApiUrl, *username)
	return &url
}

func GetFriendlyHiscoresUrl(username *string) *string {
	url := fmt.Sprintf("%s?user1=%s", model.HiscoresFriendlyUrl, *username)
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
	base := 0.25 * float64(hs.Skills[model.Defence].Level+hs.Skills[model.Hitpoints].Level+(hs.Skills[model.Prayer].Level/2))
	melee := 0.325 * float64(hs.Skills[model.Attack].Level+hs.Skills[model.Strength].Level)
	ranged := 0.325 * float64((hs.Skills[model.Ranged].Level/2)+hs.Skills[model.Ranged].Level)
	magic := 0.325 * float64((hs.Skills[model.Magic].Level/2)+hs.Skills[model.Magic].Level)
	return base + (math.Max(melee, math.Max(magic, ranged)))
}
