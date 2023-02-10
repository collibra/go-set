# RAITO - GO set

[![Go Report Card](https://goreportcard.com/badge/github.com/raito-io/golang-set)](https://goreportcard.com/report/github.com/raito-io/golang-set)![Version](https://img.shields.io/github/v/tag/raito-io/golang-set?sort=semver&label=version&color=651FFF)
[![Build](https://img.shields.io/github/actions/workflow/status/raito-io/golang-set/build.yml?branch=main)](https://github.com/raito-io/golang-set/actions/workflows/build.yml)
[![Coverage](https://img.shields.io/codecov/c/github/raito-io/golang-set?label=coverage)](https://app.codecov.io/gh/raito-io/golang-set)
[![Contribute](https://img.shields.io/badge/Contribute-🙌-green.svg)](/CONTRIBUTING.md)
[![Go version](https://img.shields.io/github/go-mod/go-version/raito-io/golang-set?color=7fd5ea)](https://golang.org/)
[![Software License](https://img.shields.io/badge/license-Apache%202-brightgreen.svg?label=license)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/raito-io/golang-set.svg)](https://pkg.go.dev/github.com/raito-io/golang-set)

## Introduction
This library introduce sets in golang.

Set is a container that store unique elements in no particular order.
Sets are an alias for `map[T]struct{}` where T is a comparable type.

## Getting Started
Add this library as a dependency via `go get github.com/raito-io/golang-set`

## Examples
```go
import github.com/raito-io/golang-set/set

func Foo() {
	var a set.Set[int]
	
	a = set.NewSet[int](2, 5, 9, 7)
	
	l := len(a) //l = 4
	
	contains := a.Contains(4) //false
	
	for i := range(a) {
	    //loop over all elements in the set	
    }
}
```
