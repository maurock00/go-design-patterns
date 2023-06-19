package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{name, attack, defense}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{creature: c}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Dubling attack", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDefenseModifier struct {
	CreatureModifier
	Defense int
}

func NewIncreaseDefenseModifier(c *Creature, defense int) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier: CreatureModifier{creature: c}, Defense: defense}
}

func (i *IncreaseDefenseModifier) Handle() {
	fmt.Println("Increasing defense", i.creature.Name, "\b's defense")
	i.creature.Defense += i.Defense
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	return
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	rootModifier := NewCreatureModifier(goblin)

	rootModifier.Add(NewNoBonusesModifier(goblin))
	rootModifier.Add(NewDoubleAttackModifier(goblin))
	rootModifier.Add(NewIncreaseDefenseModifier(goblin, 10))
	rootModifier.Add(NewDoubleAttackModifier(goblin))
	rootModifier.Add(NewIncreaseDefenseModifier(goblin, 5))
	rootModifier.Handle()

	fmt.Println(goblin.String())
}
