package main

import (
	"fmt"
	"math/rand"
)

var packages = []string{
	"Fetching UND",
	"Fetching Cosmovisor",
	"Committing Node Settings",
    "Done!",
}

func getPackages() []string {
	pkgs := packages
	copy(pkgs, packages)

	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

	for k := range pkgs {
		pkgs[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)) //nolint:gosec
	}
	return pkgs
}
