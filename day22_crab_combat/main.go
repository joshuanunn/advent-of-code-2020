package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Player struct {
	cards []int
	cache map[string]bool
	loser bool
}

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")

	// Part 1 - count number of ingredients that cannot be an allergen.
	players1 := parseInputs(input)
	winner, deck := playPart1(players1)
	fmt.Printf("Part 1 - player %d is the winner, with a score of %d\n", winner+1, calcScore(deck))

	// Part 2 - match ingredients to allergens.
	players2 := parseInputs(input)
	winner, deck = playPart2(players2)
	fmt.Printf("Part 2 - player %d is the winner, with a score of %d\n", winner+1, calcScore(deck))
}

func playPart1(players []*Player) (int, []int) {
	// Play out game until winner found
	for {
		// Check if we have a winner and return
		if players[0].Lost() {
			return 1, players[1].cards
		}
		if players[1].Lost() {
			return 0, players[0].cards
		}

		// Draw cards and play
		round := []int{}
		for _, player := range players {
			// Draw cards
			card, _ := player.Get()
			round = append(round, card)
		}
		winner, cards := findWinner(round)
		for _, card := range cards {
			players[winner].Put(card)
		}
	}
	return -1, nil
}

func playPart2(players []*Player) (int, []int) {
	// Play out game until winner found
	for {
		// Check if either card deck has been played before and abort if so
		if players[0].Cache() || players[1].Cache() {
			return 0, nil
		}

		// Check if we have a winner and return
		if players[0].Lost() {
			return 1, players[1].cards
		}
		if players[1].Lost() {
			return 0, players[0].cards
		}

		// Draw cards and play
		var subDeck []*Player
		round := []int{}
		recurse := true
		for _, player := range players {
			// Draw cards
			card, _ := player.Get()
			round = append(round, card)
			// Check can recurse
			if len(player.cards) >= card {
				subDeck = append(subDeck, copy(player, card))
			} else {
				recurse = false
			}
		}

		// If able to recurse, then recurse
		if recurse {
			winner, _ := playPart2(subDeck)
			if winner == 0 {
				players[winner].Put(round[0])
				players[winner].Put(round[1])
			} else {
				players[winner].Put(round[1])
				players[winner].Put(round[0])
			}
		} else {
			winner, cards := findWinner(round)
			for _, card := range cards {
				players[winner].Put(card)
			}
		}
	}
	return -1, nil
}

func newPlayer() *Player {
	player := &Player{}
	player.cache = make(map[string]bool)
	return player
}

func (p *Player) Cache() bool {
	// Calculate hash for internal card state
	var hash string
	for _, card := range p.cards {
		hash += fmt.Sprintf("%d,", card)
	}

	// Check if state seen before and cache if not
	if _, ok := p.cache[hash]; ok {
		return true
	}
	p.cache[hash] = true
	return false
}

func (p *Player) Lost() bool {
	if p.isEmpty() {
		p.loser = true
	}
	return p.loser
}

func copy(player *Player, limit int) *Player {
	var play = newPlayer()
	for i, card := range player.cards {
		if i < limit {
			play.cards = append(play.cards, card)
		}
	}
	return play
}

func calcScore(cards []int) int {
	var score int
	for i, card := range cards {
		score += (len(cards) - i) * card
	}
	return score
}

func findWinner(cards []int) (int, []int) {
	max := -1
	winner := -1
	for player, card := range cards {
		if card > max {
			max = card
			winner = player
		}
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	return winner, cards
}

func (p *Player) isEmpty() bool {
	return len(p.cards) == 0
}

func (p *Player) Put(card int) {
	p.cards = append(p.cards, card)
}

func (p *Player) Get() (int, bool) {
	if p.isEmpty() {
		return 0, false
	}
	card := (p.cards)[0]
	p.cards = p.cards[1:]
	return card, true
}

func parseInt(line string) int {
	value, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("could not parse int on line")
	}
	return value
}

func parseInputs(inputs []string) []*Player {
	var players []*Player

	player := -1
	for _, line := range inputs {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "Player") {
			player++
			players = append(players, newPlayer())
			continue
		}
		players[player].Put(parseInt(line))
	}
	return players
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
