package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	hgtCheck  = regexp.MustCompile(`^([0-9]+)(cm|in)$`).FindStringSubmatch
	hclCheck  = regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString
	eclCheck  = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString
	pidCheck  = regexp.MustCompile(`^[0-9]{9}$`).MatchString
	lineCheck = regexp.MustCompile(`(\w{3}):(\S+)\s?`).FindAllStringSubmatch
)

type Passport struct {
	byr int    // Birth Year
	iyr int    // Issue Year
	eyr int    // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

func main() {
	// Read inputs into slice of string, then parse into slice of *Passport
	blocks := readInputs("input.txt")
	passports := parseInputs(blocks)

	// Part 1 - count number of valid passports using ruleset 1
	count := countValidPart1(passports)
	fmt.Printf("Part 1 - found a total of %d/%d valid passports.\n", count, len(passports))

	// Part 2 - count number of valid passwords using ruleset 2
	count = countValidPart2(passports)
	fmt.Printf("Part 2 - found a total of %d/%d valid passports.\n", count, len(passports))
}

func (p *Passport) Set(field, value string) {
	if field == "byr" {
		year, err := strconv.Atoi(value)
		if err != nil {
			p.byr = 0
		}
		p.byr = year
	}
	if field == "iyr" {
		year, err := strconv.Atoi(value)
		if err != nil {
			p.iyr = 0
		}
		p.iyr = year
	}
	if field == "eyr" {
		year, err := strconv.Atoi(value)
		if err != nil {
			p.eyr = 0
		}
		p.eyr = year
	}
	if field == "hgt" {
		p.hgt = value
	}
	if field == "hcl" {
		p.hcl = value
	}
	if field == "ecl" {
		p.ecl = value
	}
	if field == "pid" {
		p.pid = value
	}
	if field == "cid" {
		p.cid = value
	}
}

func (p *Passport) Validate() bool {
	// Birth Year: four digits; at least 1920 and at most 2002.
	if p.byr < 1920 || p.byr > 2002 {
		return false
	}
	// Issue Year: four digits; at least 2010 and at most 2020.
	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}
	// Expiration Year: four digits; at least 2020 and at most 2030.
	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}
	// Height: a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	hgt := hgtCheck(p.hgt)
	if hgt == nil {
		return false
	}
	height, err := strconv.Atoi(hgt[1])
	if err != nil {
		return false
	}
	unit := hgt[2]
	if unit == "cm" {
		if height < 150 || height > 193 {
			return false
		}
	} else if unit == "in" {
		if height < 59 || height > 76 {
			return false
		}
	}
	// Hair Color: a # followed by exactly six characters 0-9 or a-f.
	if !hclCheck(p.hcl) {
		return false
	}
	// Eye Color: exactly one of: amb blu brn gry grn hzl oth.
	if !eclCheck(p.ecl) {
		return false
	}
	// Passport ID: a nine-digit number, including leading zeroes.
	if !pidCheck(p.pid) {
		return false
	}
	return true
}

func (p *Passport) Glance() bool {
	// Check all fields are defined, except cid
	if p.byr == 0 {
		return false
	}
	if p.iyr == 0 {
		return false
	}
	if p.eyr == 0 {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}
	return true
}

func countValidPart1(passports []*Passport) int {
	var total int
	for _, passport := range passports {
		if passport.Glance() {
			total++
		}
	}
	return total
}

func countValidPart2(passports []*Passport) int {
	var total int
	for _, passport := range passports {
		if passport.Validate() {
			total++
		}
	}
	return total
}

// Parse passport string blocks into slice of *Passport
func parseInputs(blocks []string) []*Passport {
	var passports []*Passport
	for _, block := range blocks {
		fields := lineCheck(block, -1)
		if len(fields) > 0 {
			passport := &Passport{}
			for _, field := range fields {
				passport.Set(field[1], field[2])
			}
			passports = append(passports, passport)
		} else {
			log.Fatalf("failed to parse regexp for block")
		}
	}
	return passports
}

// Read in file and split into passport string blocks
func readInputs(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var blocks []string
	var block string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			block += " " + line
			continue
		}
		blocks = append(blocks, block)
		block = ""
	}
	if block != "" {
		blocks = append(blocks, block)
	}
	return blocks
}
