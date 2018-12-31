package model

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
