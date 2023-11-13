package mypackage

import "fmt"

var allPackages = []Package{
	{Title: "zero"},
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Package struct {
	Title string
}

func (p *Package) String() string {
	return fmt.Sprintf("{%s}", p.Title)
}
