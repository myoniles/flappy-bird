package game

type GameMode int

const (
	GameStart GameMode = iota
	GamePlaying
	GameEnd
)
