package combat

import (
	"math"
	"math/rand"
)

func GetPhysicalDamage(
	attacker Character,
	defender Character,
) int {
	if isCrit(attacker) {
		return int(math.Pow(float64(attacker.Strength), 2) / float64(attacker.Strength-defender.Defense) * 2)
	}
	if isDodged(attacker, defender) {
		return 0
	}
	variance := 0.9 + rand.Float64()*0.2
	return max(int(math.Pow(float64(attacker.Strength), 2)/float64(attacker.Strength-defender.Defense)*variance), 1)
}

func GetMagicalDamage(
	attacker Character,
	defender Character,
	element Element,
) int {
	variance := 0.9 + rand.Float64()*0.2
	return max(
		int(
			math.Pow(float64(attacker.Intelligence), 2)/
				float64(attacker.Intelligence-defender.Spirit)*
				variance/float64(defender.Resistences[element]),
		),
		1,
	)
}

func GetHealingAmount(caster Character) int {
	variance := 0.9 + rand.Float64()*0.2
	return int(float64(caster.Intelligence) * variance * 1.5)
}

func isDodged(
	attacker Character,
	defender Character,
) bool {
	return rand.Intn(100) > (60 + attacker.Accuracy - defender.Dexterity)
}

func isCrit(
	attacker Character,
) bool {
	return rand.Intn(StatCap) > (StatCap - attacker.Luck)
}
