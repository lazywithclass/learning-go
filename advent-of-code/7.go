package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
)


func main() {
	file, err := os.Open("7.input")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	var allBagsContainingGoldenShinyBags = make(map[string]bool)
	var bagsThatCanContainGoldenShinyBags = make(map[string]bool)
	
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		bagsNames := extractBagsNames(line)
		found, name := maybeGetBagNameIfContainsGold(bagsNames)
		if found {
			bagsThatCanContainGoldenShinyBags[name] = true
		}
	}

	for _, line := range lines {
		bagsNames := extractBagsNames(line)

		entered := false
		for _, name := range bagsNames {
			if bagsThatCanContainGoldenShinyBags[name] {
				entered = true
				allBagsContainingGoldenShinyBags[bagsNames[0]] = true
			}		
		}

		fmt.Println(bagsThatCanContainGoldenShinyBags)
		fmt.Println(bagsNames, entered)
		fmt.Println("--------------")
	}

	fmt.Println(len(allBagsContainingGoldenShinyBags))
}

func maybeGetBagNameIfContainsGold(bagNames []string) (found bool, bagName string) {
	for i, name := range bagNames {
		if name == "shiny gold" && i != 0 {
			found = true
		}
	}
	return found, bagNames[0]
}

func extractBagsNames(line string) []string {
	split := strings.Split(line, " contain ")
	bag := extractBagName(split[0])
	var bags []string
	bags = append(bags, bag)
	for _, item := range strings.Split(split[1], ", ") {
		bags = append(bags, extractBagName(item))
	}
	return bags
}

func extractBagName(line string) string {
	re := regexp.MustCompile(`\d? ?([\w\s]+) bags?`)
	return re.FindAllStringSubmatch(line, -1)[0][1]
}
