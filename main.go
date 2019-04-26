package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TEST_PREFIX    = "--- FAIL"
	PACKAGE_PREFIX = "FAIL"
)

func main() {
	ent := flag.Bool("ent", false, "ent build")
	flag.Parse()
	tags := ""
	if *ent {
		tags = "-tags 'ent prem' "
	}
	scanner := bufio.NewScanner(os.Stdin)
	tests := map[string][]string{}
	current := []string{}
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(l, TEST_PREFIX) {
			parts := strings.Split(l, " ")
			if len(parts) < 3 {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("Thought it would be a failed test, but formatted strangely: %s", l))
				continue
			}
			current = append(current, strings.Split(parts[2], "/")[0])
		} else if strings.HasPrefix(l, PACKAGE_PREFIX) {
			parts := strings.Split(l, "\t")
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("Thought it would be a failed package, but formatted strangely: %s", l))
				continue
			}
			name := strings.Split(parts[1], " ")[0]
			tests[name] = current
			current = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	for p, ts := range tests {
		beginning := fmt.Sprintf("go test %s%s", tags, p)
		if len(ts) == 0 {
			fmt.Printf("%s\n", beginning)
		}
		for _, t := range ts {
			fmt.Printf("%s -run '^%s$'\n", beginning, t)
		}
	}
}
