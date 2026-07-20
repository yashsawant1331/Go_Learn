package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	masking := false

	for i := 0; i < size; i++ {

		// Detect the beginning of a link
		if !masking && len(text[i:]) >= nlink && text[i:i+nlink] == link {
			masking = true

			// Keep the http:// prefix
			buf = append(buf, link...)

			// Skip the characters of "http://"
			i += nlink - 1
			continue
		}

		if masking {
			// Stop masking when whitespace is found
			if text[i] == ' ' || text[i] == '\n' || text[i] == '\t' {
				masking = false
				buf = append(buf, text[i])
			} else {
				buf = append(buf, '*')
			}
		} else {
			buf = append(buf, text[i])
		}
	}

	fmt.Println(string(buf))
}
