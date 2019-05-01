package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	TEST_PREFIX    = "--- FAIL"
	PACKAGE_PREFIX = "FAIL"
)

func main() {
	tags := flag.String("tags", "", "tags")
	debug := flag.Bool("debug", false, "debug output to stderr")
	flag.Parse()
	this := ""
	if *tags != "" {
		this = fmt.Sprintf("-tags %q ", *tags)
	}
	scanner := bufio.NewScanner(os.Stdin)
	tests := map[string]map[string]struct{}{}
	current := map[string]struct{}{}
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(l, TEST_PREFIX) {
			parts := strings.Split(l, " ")
			if len(parts) < 3 {
				if *debug {
					fmt.Fprintln(os.Stderr, fmt.Sprintf("Thought it would be a failed test, but formatted strangely: %s", l))
				}
				continue
			}
			current[strings.Split(parts[2], "/")[0]] = struct{}{}
		} else if strings.HasPrefix(l, PACKAGE_PREFIX) {
			parts := strings.Split(l, "\t")
			if len(parts) < 2 {
				if *debug {
					fmt.Fprintln(os.Stderr, fmt.Sprintf("Thought it would be a failed package, but formatted strangely: %s", l))
				}
				continue
			}
			name := strings.Split(parts[1], " ")[0]
			tests[name] = current
			current = map[string]struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	packages := []string{}
	for k := range tests {
		packages = append(packages, k)
	}
	sort.Strings(packages)
	for _, p := range packages {
		sortedTests := []string{}
		for t := range tests[p] {
			sortedTests = append(sortedTests, t)
		}
		sort.Strings(sortedTests)
		beginning := fmt.Sprintf("go test %s%s", this, p)
		if len(sortedTests) == 0 {
			fmt.Printf("%s\n", beginning)
		}
		for _, t := range sortedTests {
			fmt.Printf("%s -run '^%s$'\n", beginning, t)
		}
	}
}
