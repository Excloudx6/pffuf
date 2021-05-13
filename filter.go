package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

func doSort() {
	if sorting == "Status" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Status < results[j].Status
		})
	}

	if sorting == "Length" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Length < results[j].Length
		})
	}

	if sorting == "Words" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Words < results[j].Words
		})
	}

	if sorting == "Lines" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Lines < results[j].Lines
		})
	}

	if sorting == "URL" {
		slice.Sort(results[:], func(i, j int) bool { return results[i].URL < results[j].URL })
	}

	if sorting == "Endpoint" {
		slice.Sort(results[:], func(i, j int) bool { return results[i].Endpoint < results[j].Endpoint })
	}

	if sorting == "None" {
		results = results[:0]
		results = origresults
		doFilter()

	}
}

func doFilter() {
	results = results[:0]
	results = origresults

	var resultstmp []NavResults

	for _, cur := range results {
		if !checkFilter(cur) {
			continue
		}
		resultstmp = append(resultstmp, cur)

	}
	results = resultstmp
	doSort()
}

func setFilter(cmdline string) {
	parts := strings.Split(cmdline, " ")
	cmd := parts[0]
	args := strings.Join(parts[1:], "")

	argsi := []int{}
	for _, arg := range strings.Split(args, ",") {
		if arg == " " || arg == "" {
			continue
		}
		argi, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid filter condition\n", arg)
			return
		}
		argsi = append(argsi, argi)
	}
	if cmd == "fc" {
		filters.StatusHide = argsi
	}
	if cmd == "fw" {
		filters.WordsHide = argsi
	}
	if cmd == "fl" {
		filters.LinesHide = argsi
	}
	if cmd == "fs" {
		filters.LenHide = argsi
	}
	if cmd == "mc" {
		filters.StatusMatch = argsi
	}
	if cmd == "mw" {
		filters.WordsMatch = argsi
	}
	if cmd == "ml" {
		filters.LenMatch = argsi
	}
	if cmd == "ms" {
		filters.LinesMatch = argsi
	}
	doFilter()
}

func showFilters() {
	fmt.Println("Current Filters:")
	fmt.Println("Hide STATUS: ", filters.StatusHide)
	fmt.Println("Hide WORDS: ", filters.WordsHide)
	fmt.Println("Hide LINES: ", filters.LinesHide)
	fmt.Println("Hide LENGTH: ", filters.LenHide)
	fmt.Println("Match STATUS: ", filters.StatusMatch)
	fmt.Println("Match WORDS: ", filters.WordsMatch)
	fmt.Println("Match LINES: ", filters.LinesMatch)
	fmt.Println("Match LENGTH: ", filters.LenMatch)
}
func clearFilters() {
	filters.StatusHide = filters.StatusHide[:0]
	filters.WordsHide = filters.WordsHide[:0]
	filters.LinesHide = filters.LinesHide[:0]
	filters.LenHide = filters.LenHide[:0]
	filters.StatusMatch = filters.StatusMatch[:0]
	filters.WordsMatch = filters.WordsMatch[:0]
	filters.LinesMatch = filters.LinesMatch[:0]
	filters.LenMatch = filters.LenMatch[:0]
	results = results[:0]
	results = origresults
	doSort()
	fmt.Println("All Filters Cleared !")
}

func ifMatch(line NavResults) bool {
	if contains(filters.StatusMatch, line.Status) {
		return true
	}
	if contains(filters.LenMatch, line.Length) {
		return true
	}
	if contains(filters.WordsMatch, line.Words) {
		return true
	}
	if contains(filters.LinesMatch, line.Lines) {
		return true
	}

	return false
}

// checks the current filters, true = show it, false = hide it
func checkFilter(line NavResults) bool {
	if ifMatch(line) {
		return true
	} else {

		if contains(filters.StatusHide, line.Status) {
			return false
		}
		if contains(filters.LenHide, line.Length) {
			return false
		}
		if contains(filters.WordsHide, line.Words) {
			return false
		}
		if contains(filters.LinesHide, line.Lines) {
			return false
		}
	}
	if len(filters.LinesMatch) == 0 && len(filters.StatusMatch) == 0 && len(filters.WordsMatch) == 0 && len(filters.LenMatch) == 0 {
		return true
	}

	return false
}
