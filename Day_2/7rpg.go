package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Character type
type Character struct {
	Name      string
	Class     string
	Health    int
	Strength  int
	Dexterity int
}

// Attack method
func (c Character) Attack() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rollBonus := r.Intn(6) + 1 // 1-6 random bonus
	return c.Strength + rollBonus
}

// DefendAgainst method
func (c *Character) DefendAgainst(attackValue int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	defense := c.Dexterity + r.Intn(6) + 1 // Dexterity + 1-6 random bonus

	damage := attackValue - defense
	if damage < 0 {
		damage = 0
	}

	c.Health -= damage

	fmt.Printf("%s defends! (Attack: %d, Defense: %d)\n", c.Name, attackValue, defense)
	fmt.Printf("%s takes %d damage.\n", c.Name, damage)
}

// IsAlive method
func (c Character) IsAlive() bool {
	return c.Health > 0
}

// DisplayStats method
func (c Character) DisplayStats() {
	fmt.Printf("--- %s the %s ---\n", c.Name, c.Class)
	fmt.Printf("Health: %d\n", c.Health)
	fmt.Printf("Strength: %d\n", c.Strength)
	fmt.Printf("Dexterity: %d\n", c.Dexterity)
	fmt.Println("------------------")
}

func main() {
	fmt.Println("⚔️ CHARACTER BATTLE SIMULATOR ⚔️")

	// Create our hero
	var name string
	fmt.Print("Enter your character's name: ")
	fmt.Scanln(&name)

	fmt.Println("\nChoose a class:")
	fmt.Println("1: Warrior (High strength, average dexterity)")
	fmt.Println("2: Rogue (High dexterity, average strength)")
	fmt.Print("Your choice (1-2): ")

	var classChoice int
	fmt.Scanln(&classChoice)

	hero := Character{Name: name}

	switch classChoice {
	case 1:
		hero.Class = "Warrior"
		hero.Health = 100
		hero.Strength = 15
		hero.Dexterity = 10
	case 2:
		hero.Class = "Rogue"
		hero.Health = 80
		hero.Strength = 10
		hero.Dexterity = 15
	default:
		hero.Class = "Adventurer"
		hero.Health = 90
		hero.Strength = 12
		hero.Dexterity = 12
	}

	// Create enemy
	enemy := Character{
		Name:      "Goblin",
		Class:     "Monster",
		Health:    70,
		Strength:  12,
		Dexterity: 8,
	}

	fmt.Println("\n--- Battle Begins! ---")
	hero.DisplayStats()
	enemy.DisplayStats()

	// Battle loop
	round := 1
	for hero.IsAlive() && enemy.IsAlive() {
		fmt.Printf("\n--- Round %d ---\n", round)

		// Hero attacks enemy
		fmt.Printf("%s attacks %s!\n", hero.Name, enemy.Name)
		attackValue := hero.Attack()
		enemy.DefendAgainst(attackValue)

		// Check if enemy is defeated
		if !enemy.IsAlive() {
			fmt.Printf("\n%s has defeated %s! Victory!\n", hero.Name, enemy.Name)
			break
		}

		// Enemy attacks hero
		fmt.Printf("\n%s attacks %s!\n", enemy.Name, hero.Name)
		attackValue = enemy.Attack()
		hero.DefendAgainst(attackValue)

		// Display current stats
		fmt.Printf("\nAfter round %d:\n", round)
		fmt.Printf("%s: %d health\n", hero.Name, hero.Health)
		fmt.Printf("%s: %d health\n", enemy.Name, enemy.Health)

		round++

		// Pause between rounds
		fmt.Println("\nPress Enter for next round...")
		fmt.Scanln()
	}

	// Check final battle result
	if !hero.IsAlive() {
		fmt.Printf("\n%s has been defeated by %s! Game over.\n", hero.Name, enemy.Name)
	}
}
