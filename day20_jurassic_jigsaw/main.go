package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var monster = [][2]int{
	{0, 18},
	{1, 0},
	{1, 5},
	{1, 6},
	{1, 11},
	{1, 12},
	{1, 17},
	{1, 18},
	{1, 19},
	{2, 1},
	{2, 4},
	{2, 7},
	{2, 10},
	{2, 13},
	{2, 16},
}

type Row [10]bool

type Grid struct {
	id    int
	grid  [10]Row
	edgeN [2]int
	edgeE [2]int
	edgeS [2]int
	edgeW [2]int
}

type GridN struct {
	grid [][]bool
}

func main() {
	input := readInputs("input.txt")

	// Part 1 - find corners by finding and coutning pairs of matching tiles.
	// Grouped into tiles which have 2 (corner), 3 (edge) or 4 (centre).
	tiles, count, links := parseInputs(input)
	groupedPieces := definePieces(tiles, count)

	var corners string
	product := 1
	for _, cornerID := range groupedPieces[2] {
		corners += fmt.Sprintf(" %d", cornerID.id)
		product *= cornerID.id
	}
	fmt.Printf("Part 1 - corner tile IDs identified as:%s, with product of %d\n", corners, product)

	// Part 2 - assemble image, identify sea monsters and calculate water roughness.
	// Seed image with random corner piece in any orientation
	firstCorner := groupedPieces[2][0]
	// Construct image from this corner
	grid := buildGrid(tiles, firstCorner.id, links)

	// Find sea monsters
	total := countMonsters(grid)
	imageCorrection := total * 15
	waterRoughness := grid.count() - imageCorrection
	fmt.Printf("Part 2 - found %d monsters and water roughness of %d\n", total, waterRoughness)
}

func countMonsters(grid *GridN) int {
	// Initial check
	count := grid.findSeaMonsters()
	if count > 0 {
		return count
	}
	// Three rotations
	for x := 0; x < 3; x++ {
		grid.rotate()
		count = grid.findSeaMonsters()
		if count > 0 {
			return count
		}
	}
	// Flip
	grid.flip()
	count = grid.findSeaMonsters()
	if count > 0 {
		return count
	}
	// Three rotations
	for x := 0; x < 3; x++ {
		grid.rotate()
		count = grid.findSeaMonsters()
		if count > 0 {
			return count
		}
	}
	return count
}

func newGrid(size int) *GridN {
	var grid = make([][]bool, size, size)
	for x := 0; x < size; x++ {
		grid[x] = make([]bool, size, size)
	}
	return &GridN{grid: grid}
}

func (n *GridN) rotate() {
	size := len(n.grid)
	g := newGrid(size)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if n.grid[col][size-1-row] {
				g.grid[row][col] = true
			}
		}
	}
	n.grid = g.grid
}

func (n *GridN) flip() {
	size := len(n.grid)
	g := newGrid(size)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			g.grid[row][col], g.grid[row][size-1-col] = n.grid[row][size-1-col], n.grid[row][col]
		}
	}
	n.grid = g.grid
}

func (n *GridN) count() int {
	size := len(n.grid)
	var total int
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			// Sum all #'s (true)
			if n.grid[row][col] {
				total++
			}
		}
	}
	return total
}

func (n *GridN) addInner(rOffset, cOffset int, subGrid *Grid) {
	for r := 1; r < 9; r++ {
		for c := 1; c < 9; c++ {
			n.grid[rOffset+r][cOffset+c] = subGrid.grid[r][c]
		}
	}
}

func (n *GridN) findSeaMonsters() int {
	size := len(n.grid)
	var monsters int
	for row := 0; row < size-3; row++ {
		for col := 0; col < size-20; col++ {
			// Search for monster at position
			hits := 0
			for _, pos := range monster {
				r := row + pos[0]
				c := col + pos[1]
				if n.grid[r][c] {
					hits++
				}
			}
			// If monster found then remove from image
			if hits == 15 {
				monsters++
			}
		}
	}
	return monsters
}

func (n *GridN) print() {
	for _, row := range n.grid {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func buildGrid(tiles map[int]*Grid, nextTileID int, links map[int][]int) *GridN {
	var imageGrid = newGrid(96) // Grid of 12x12 8x8
	var nextTile *Grid
	var nextWest int
	var nextNorth []int
	var lastID []int
	for row := 0; row < 12; row++ {
		for col := 0; col < 12; col++ {
			switch {
			case row == 0 && col == 0:
				// first corner - special case
				nextTile = tiles[nextTileID]
				nextTile.adjust(0, -1, -1, 0)
			case row == 0:
				// first row - special case
				nextTileID = getNextPiece(lastID[col-1], nextWest, links)
				nextTile = tiles[nextTileID]
				nextTile.adjust(0, -1, -1, nextWest)
			case col == 0:
				// first col (after first row) - special case
				nextTileID = getNextPiece(lastID[(row-1)*12+col], nextNorth[(row-1)*12+col], links)
				nextTile = tiles[nextTileID]
				nextTile.adjust(nextNorth[(row-1)*12+col], -1, -1, 0)
			default:
				// all other tile positions
				nextTileID = getNextPiece(lastID[(row-1)*12+col], nextNorth[(row-1)*12+col], links)
				nextTile = tiles[nextTileID]
				nextTile.adjust(nextNorth[(row-1)*12+col], -1, -1, nextWest)
			}
			// Add subtile into larger imageGrid
			imageGrid.addInner(row*8-1, col*8-1, nextTile)
			nextWest = nextTile.edgeE[0]
			nextNorth = append(nextNorth, nextTile.edgeS[0])
			lastID = append(lastID, nextTileID)
		}
	}
	return imageGrid
}

func getNextPiece(id, side int, links map[int][]int) int {
	pairPieces := links[side]
	var nextPiece int
	if pairPieces[0] == id {
		nextPiece = pairPieces[1]
	} else {
		nextPiece = pairPieces[0]
	}
	return nextPiece
}

func definePieces(tiles map[int]*Grid, count map[int]int) map[int][]*Grid {
	sideGroups := make(map[int][]*Grid)

	for _, piece := range tiles {
		sides := piece.combo(count)
		sideGroups[sides] = append(sideGroups[sides], piece)
	}

	return sideGroups
}

func checkEdge(cases [2]int, target int) bool {
	// Return true for targets we do not want to match
	if target == -1 {
		return true
	}
	// Special case for target == 0 (here we only check the inner 8 bits which have been zeroed out)
	if ((cases[0] & 0b0111111110) == target) || ((cases[1] & 0b0111111110) == target) {
		return true
	}
	// Compare forward calculated edge pattern to target
	if (cases[0] == target) || (cases[1] == target) {
		return true
	}
	return false
}

func (g *Grid) rotate() {
	// Single clockwise rotation
	rows := [10]Row{}
	for r := 0; r < 10; r++ {
		row := [10]bool{}
		for c := 0; c < 10; c++ {
			if g.grid[c][9-r] {
				row[c] = true
			}
		}
		rows[r] = row
	}
	g.grid = rows
	g.calculateEdges()
}

func (g *Grid) flip() {
	for row := 0; row < 10; row++ {
		for col := 0; col < 5; col++ {
			g.grid[row][col], g.grid[row][9-col] = g.grid[row][9-col], g.grid[row][col]
		}
	}
	g.calculateEdges()
}

// Algorithm to flip and/or rotate tile to find target edges
// Completes 3 rotations, then flip, then 3 rotations
func (g *Grid) adjust(n, e, s, w int) bool {
	// Check initial state
	if checkEdge(g.edgeN, n) && checkEdge(g.edgeE, e) && checkEdge(g.edgeS, s) && checkEdge(g.edgeW, w) {
		return true
	}
	// Rotate and check 3 times
	for r := 0; r < 3; r++ {
		g.rotate()
		if checkEdge(g.edgeN, n) && checkEdge(g.edgeE, e) && checkEdge(g.edgeS, s) && checkEdge(g.edgeW, w) {
			return true
		}
	}
	// Flip and check
	g.flip()
	if checkEdge(g.edgeN, n) && checkEdge(g.edgeE, e) && checkEdge(g.edgeS, s) && checkEdge(g.edgeW, w) {
		return true
	}
	// Finally rotate and check 3 times on reverse side
	for r := 0; r < 3; r++ {
		g.rotate()
		if checkEdge(g.edgeN, n) && checkEdge(g.edgeE, e) && checkEdge(g.edgeS, s) && checkEdge(g.edgeW, w) {
			return true
		}
	}
	return false
}

func (g *Grid) rowEdge(row int) [2]int {
	var edge int
	var edgeR int
	for col := 0; col < 10; col++ {
		bit := g.grid[row][col]
		if bit {
			edge = (edge << 1) | 1
		} else {
			edge = edge << 1
		}
	}
	for col := 9; col >= 0; col-- {
		bit := g.grid[row][col]
		if bit {
			edgeR = (edgeR << 1) | 1
		} else {
			edgeR = edgeR << 1
		}
	}
	return [2]int{edge, edgeR}
}

func (g *Grid) colEdge(col int) [2]int {
	var edge int
	var edgeR int
	for row := 0; row < 10; row++ {
		bit := g.grid[row][col]
		if bit {
			edge = (edge << 1) | 1
		} else {
			edge = edge << 1
		}
	}
	for row := 9; row >= 0; row-- {
		bit := g.grid[row][col]
		if bit {
			edgeR = (edgeR << 1) | 1
		} else {
			edgeR = edgeR << 1
		}
	}
	return [2]int{edge, edgeR}
}

func (g *Grid) calculateEdges() {
	g.edgeN = g.rowEdge(0)
	g.edgeE = g.colEdge(9)
	g.edgeS = g.rowEdge(9)
	g.edgeW = g.colEdge(0)
}

func (g *Grid) print() {
	fmt.Printf("[ID:%d] N:%d E:%d S:%d W:%d\n", g.id, g.edgeN, g.edgeE, g.edgeS, g.edgeW)
	//fmt.Printf("%d %d %d %d \n", g.edgeN, g.edgeE, g.edgeS, g.edgeW)
}

func (g *Grid) printGrid() {
	for _, row := range g.grid {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func hasMatchingPair(term [2]int, unique map[int]int) bool {
	count0, _ := unique[term[0]]
	count1, _ := unique[term[1]]

	if count0 != count1 {
		log.Fatalf("non match")
	}

	// Check if this side of piece has a matching pair (edge is non-outer)
	if count0 == 2 {
		return true
	}
	return false
}

func (g *Grid) clearRow(row int) {
	for col := 1; col < 9; col++ {
		g.grid[row][col] = false
	}
}

func (g *Grid) clearCol(col int) {
	for row := 1; row < 9; row++ {
		g.grid[row][col] = false
	}
}

func (g *Grid) combo(unique map[int]int) int {
	var count int
	if hasMatchingPair(g.edgeN, unique) {
		count++
	} else {
		g.clearRow(0)
		g.calculateEdges()
	}
	if hasMatchingPair(g.edgeE, unique) {
		count++
	} else {
		g.clearCol(9)
		g.calculateEdges()
	}
	if hasMatchingPair(g.edgeS, unique) {
		count++
	} else {
		g.clearRow(9)
		g.calculateEdges()
	}
	if hasMatchingPair(g.edgeW, unique) {
		count++
	} else {
		g.clearCol(0)
		g.calculateEdges()
	}
	return count
}

func addMap(key, id int, count map[int]int, links map[int][]int) {
	if _, ok := count[key]; ok {
		count[key]++
	} else {
		count[key] = 1
	}
	if _, ok := links[key]; ok {
		links[key] = append(links[key], id)
	} else {
		links[key] = []int{id}
	}
}

func (g *Grid) dump(count map[int]int, links map[int][]int) {
	addMap(g.edgeN[0], g.id, count, links)
	addMap(g.edgeE[0], g.id, count, links)
	addMap(g.edgeS[0], g.id, count, links)
	addMap(g.edgeW[0], g.id, count, links)
	addMap(g.edgeN[1], g.id, count, links)
	addMap(g.edgeE[1], g.id, count, links)
	addMap(g.edgeS[1], g.id, count, links)
	addMap(g.edgeW[1], g.id, count, links)
}

func parseInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln("error parsing int")
	}
	return val
}

func parseInputs(inputs []string) (map[int]*Grid, map[int]int, map[int][]int) {
	var tiles = make(map[int]*Grid)
	var count = make(map[int]int)
	var links = make(map[int][]int)

	for line := 0; line < len(inputs); line++ {
		if inputs[line] == "\r\n" {
			continue
		}
		// If tile found then process
		if strings.Contains(inputs[line], "Tile") {
			name := strings.Replace(inputs[line], "Tile ", "", -1)
			name = strings.Replace(name, ":", "", -1)
			id := parseInt(name)

			rows := [10]Row{}
			for r := 0; r < 10; r++ {
				line++
				row := [10]bool{}
				for c := 0; c < 10; c++ {
					if inputs[line][c] == '#' {
						row[c] = true
					}
				}
				rows[r] = row
			}
			var data = &Grid{id: id, grid: rows}
			data.calculateEdges()
			data.dump(count, links)
			tiles[id] = data
		}
	}
	return tiles, count, links
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
