package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)

type Person struct {
	name string
	// other attributes
}

type RelationInfo struct {
	from     *Person
	relation Relationship
	to       *Person
}

// Relationships low level module
type Relationships struct {
	relations []RelationInfo
}

type RelationshipsBrowser interface {
	findAllParentChild(parentName string) []*Person
}

func (r *Relationships) addParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, RelationInfo{parent, Parent, child})
	r.relations = append(r.relations, RelationInfo{child, Child, parent})
}

func (r *Relationships) findAllParentChild(parentName string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.from.name == parentName && v.relation == Parent {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

// RelationshipsVersion2 low level module
type RelationshipsVersion2 struct {
}

func (r *RelationshipsVersion2) findAllParentChild(parentName string) []*Person {
	result := []*Person{
		{
			name: "Juanito",
		},
		{
			name: "Sebastian",
		},
	}

	return result
}

// Research high level module
type Research struct {
	// break DIP
	relationships Relationships
	// dont Break DIP
	browser RelationshipsBrowser
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "Jhon" && rel.relation == Parent {
			fmt.Println("John has a child called ", rel.to.name)
		}
	}
}

func (r *Research) InvestigateGoodWay() {
	childs := r.browser.findAllParentChild("Jhon")

	for _, child := range childs {
		fmt.Println("Jhon has a child called ", child.name)
	}
}

func main() {
	parent := Person{name: "Jhon"}
	child1 := Person{name: "Juan"}
	child2 := Person{name: "Carlitos"}

	rels := Relationships{}
	rels2 := RelationshipsVersion2{}

	rels.addParentAndChild(&parent, &child1)
	rels.addParentAndChild(&parent, &child2)

	r := Research{relationships: rels, browser: &rels2}
	r.Investigate()

	fmt.Println("")

	r.InvestigateGoodWay()
}
