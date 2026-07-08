package main

import (
	"fmt"
	"strings"
	"time"
)

var digits = map[rune][5]string{
	'0': {
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	},
	'1': {
		" ██",
		"  █",
		"  █",
		"  █",
		"███",
	},
	'2': {
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	},
	'3': {
		"███",
		"  █",
		"███",
		"  █",
		"███",
	},
	'4': {
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	},
	'5': {
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	},
	'6': {
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	},
	'7': {
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	},
	'8': {
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	},
	'9': {
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	},
	':': {
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	},
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func renderClock(currentTime string) {
	for row := 0; row < 5; row++ {
		var line strings.Builder

		for _, ch := range currentTime {
			pattern := digits[ch]
			line.WriteString(pattern[row])
			line.WriteString("  ")
		}

		fmt.Println(line.String())
	}
}

func main() {
	for {
		clearScreen()

		now := time.Now().Format("15:04:05")

		renderClock(now)

		time.Sleep(time.Second)
	}
}
