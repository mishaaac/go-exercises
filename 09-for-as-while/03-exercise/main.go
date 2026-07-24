/*
Exercise 3: Turn-Based Console RPG Combat Engine

Design the core of a turn-based console RPG game using Go 1.22+.

Requirements:
- Initialize combatants: Hero with 100 HP, Monster with 80 HP.
- In each round:
  - Hero attacks dealing random damage between 5 and 15 HP (range of 11 possible values: [5..15]).
  - Monster counterattacks dealing random damage between 5 and 12 HP (range of 8 possible values: [5..12]).
  - Standard RPG logic check: If the monster dies during the hero's attack, the combat loop terminates
    immediately so the dead monster cannot counterattack!
  - Use Go 1.22+ `math/rand/v2` (`rand.N(N)` / `rand.IntN(N)`) for random damage calculations.
  - Display status after each round, ensuring HP never drops below 0.
  - Declare victory for the surviving fighter.
*/

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// Combatant initial stats.
	heroHP := 100
	monsterHP := 80
	roundCount := 1

	fmt.Println("===========================================")
	fmt.Println("  ⚔️  BATTLE INITIATED: HERO vs MONSTER  ⚔️ ")
	fmt.Println("===========================================")
	fmt.Printf("Hero HP: %d | Monster HP: %d\n\n", heroHP, monsterHP)

	// Infinite combat loop terminated dynamically when HP <= 0.
	for {
		fmt.Printf("--- 🗡️ Round %d --- \n", roundCount)

		// 1. Hero Attacks First!
		// Damage range 5 to 15: rand.IntN(11) yields [0..10] + 5 => [5..15].
		heroDamage := rand.IntN(11) + 5
		monsterHP -= heroDamage

		// Clamp HP floor to 0 so HP never reads negative.
		if monsterHP < 0 {
			monsterHP = 0
		}

		fmt.Printf(" Hero strikes the monster for %d damage! (Monster HP: %d)\n", heroDamage, monsterHP)

		// Check if monster was slain before it can counterattack!
		if monsterHP == 0 {
			fmt.Println("\n💥 The monster collapses from the strike!")
			break
		}

		// 2. Monster Counterattacks (if still alive)!
		// Damage range 5 to 12: rand.IntN(8) yields [0..7] + 5 => [5..12].
		monsterDamage := rand.IntN(8) + 5
		heroHP -= monsterDamage

		if heroHP < 0 {
			heroHP = 0
		}

		fmt.Printf(" Monster retaliates for %d damage! (Hero HP: %d)\n", monsterDamage, heroHP)

		// Check if hero was slain in battle.
		if heroHP == 0 {
			fmt.Println("\n💥 The hero falls in combat!")
			break
		}

		// End of turn summary status.
		fmt.Printf("\n📊 End of Round %d Status -> Hero HP: %d | Monster HP: %d\n\n",
			roundCount, heroHP, monsterHP)

		// Small delay to simulate turn pacing in CLI.
		time.Sleep(500 * time.Millisecond)
		roundCount++
	}

	// Post-combat resolution block.
	fmt.Println("\n===========================================")
	if heroHP > 0 {
		fmt.Printf("🏆 VICTORY! The Hero vanquished the monster in %d rounds!\n", roundCount)
	} else {
		fmt.Printf("💀 GAME OVER! The Hero was defeated in round %d.\n", roundCount)
	}
	fmt.Println("===========================================")
}
