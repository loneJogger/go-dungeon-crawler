package combat

type TargetType int

const (
	TargetSelf TargetType = iota
	TargetSingleEnemy
	TargetSingleAlly
	TargetAllEnemies
	TargetAllAllies
	TargetAllAll
)
