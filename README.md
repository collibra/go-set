# GO set
[![Quality Gate Status](https://sonarqube.collibra.dev/api/project_badges/measure?project=go-set&metric=alert_status&token=sqb_4c7f68da4a4cf86d07d42f42f9c61e5bf1f971a4)](https://sonarqube.collibra.dev/dashboard?id=go-set)

## Introduction
This library introduce sets in golang.

Set is a container that store unique elements in no particular order.
Sets are an alias for `map[T]struct{}` where T is a comparable type.

## Getting Started
Add this library as a dependency via `go get github.com/collibra/go-set`

## Examples
```go
import github.com/collibra/go-set/set

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
