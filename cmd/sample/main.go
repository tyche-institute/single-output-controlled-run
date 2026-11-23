// Copyright 2026 Anton Sokolov / Tyche Institute, Tallinn, Estonia.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

// Command sample is the toy program that the controlled release runs build.
// It prints its name and the version that the release workflow passes in.
package main

import "fmt"

// gitVersion is set by the release workflow through -ldflags.
var gitVersion = "unknown"

func main() {
	fmt.Println("sample", gitVersion)
}
