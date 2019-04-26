// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"sort"
	"strings"
)

// summarize prints the report and errors if any
func summarize(r *report, errs ...error) {
	sort.Strings(r.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range r.domains {
		parsed := r.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", r.total.visits)

	// only handle the errors once
	dumpErrs(errs...)
}

// this variadic func simplifies the multiple error handling
func dumpErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
		}
	}
}