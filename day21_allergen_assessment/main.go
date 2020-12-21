package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

type Allergen struct {
	name    string
	matches map[string]bool
}

type Allergens map[string]*Allergen

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")
	foodSlice := parseInputs(input)

	// Part 1 - count number of ingredients that cannot be an allergen.
	allergens := getAllergens(foodSlice)
	count := countNonAllergens(allergens, foodSlice)
	fmt.Printf("Part 1 - number of ingredients that cannot be an allergen = %d\n", count)

	// Part 2 - match ingredients to allergens.
	assignments := assignAllergens(allergens)
	dangerousIngredients := summarise(assignments)
	fmt.Printf("Part 2 - canonical dangerous ingredient list = %s\n", dangerousIngredients)
}

func summarise(assignments map[string]string) string {
	// Sort assignments by allergen name
	var keys []string
	for key := range assignments {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var ingredients []string
	for _, key := range keys {
		ingredients = append(ingredients, assignments[key])
	}
	return strings.Join(ingredients, ",")
}

func assignAllergens(allergens Allergens) map[string]string {
	// Loop over all allergens and iteratively assign to ingredients
	var assignments = make(map[string]string)
	for len(allergens) > 0 {
		for name, allergen := range allergens {
			// Attempt assignment (where only one ingredient for an allergen)
			if len(allergen.matches) == 1 {
				match := allergen.Ingredient()
				assignments[name] = match
				// If success, then remove ingredient from other allergens and remove allergen
				allergens.removeIngredient(match)
				delete(allergens, name)
				break
			}
		}
	}
	return assignments
}

func (a *Allergens) removeIngredient(remove string) {
	for _, allergen := range *a {
		if _, ok := allergen.matches[remove]; ok {
			delete(allergen.matches, remove)
		}
	}
}

func (a *Allergen) Check(ingredients []string) {

	if a.matches == nil {
		a.matches = map[string]bool{}
	}

	// If have ingredients, compute union with new ingredients, else add all
	var union = make(map[string]bool)
	if len(a.matches) > 0 {
		for _, ingredient := range ingredients {
			if _, ok := a.matches[ingredient]; ok {
				union[ingredient] = true
			}
		}
	} else {
		for _, ingredient := range ingredients {
			union[ingredient] = true
		}
	}
	a.matches = union
}

func (a *Allergen) Ingredient() string {
	for ingredient := range a.matches {
		return ingredient
	}
	return ""
}

func countNonAllergens(allergens Allergens, food []*Food) int {
	// Get all ingredients that might be an allergen
	possible := map[string]bool{}
	for _, allergen := range allergens {
		for ingredient := range allergen.matches {
			possible[ingredient] = true
		}
	}
	// Get all ingredients that cannot be an allergen
	var notFound int
	for _, food := range food {
		for _, ingredient := range food.ingredients {
			if _, ok := possible[ingredient]; !ok {
				notFound++
			}
		}
	}
	return notFound
}

func getAllergens(inputs []*Food) Allergens {
	// Get all ingredients that do have a match
	var allergens = make(map[string]*Allergen)
	for _, food := range inputs {
		for _, allergen := range food.allergens {
			if _, ok := allergens[allergen]; !ok {
				allergens[allergen] = &Allergen{name: allergen}
			}
			allergens[allergen].Check(food.ingredients)
		}
	}
	return allergens
}

func parseInputs(inputs []string) []*Food {
	var foodSlice = []*Food{}

	for _, line := range inputs {
		if len(line) > 0 {
			line = strings.Replace(line, "(", "", -1)
			line = strings.Replace(line, ")", "", -1)
			line = strings.Replace(line, ",", "", -1)
			split := strings.Split(line, " contains ")

			ingredients := strings.Split(split[0], " ")
			allergens := strings.Split(split[1], " ")

			food := &Food{ingredients, allergens}
			foodSlice = append(foodSlice, food)
		}
	}
	return foodSlice
}

func readInputs(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	lines := string(b)

	var inputs []string
	for _, line := range strings.Split(lines, "\n") {
		inputs = append(inputs, line)
	}
	return inputs
}
