package gather

import (
	"fmt"
	"top-website-image-gatherer/pkg/site"
)

type progress struct {
	complete int
	total    int
}

type Gatherer interface {
	Gather() error
}

type gatherer struct {
	sites  []site.Site
	output string
}

var _ Gatherer = &gatherer{}

// New creates a new gatherer from output directory, top (number of sites), and offset from the top of the list
func New(output string, top, offset int) Gatherer {
	sites := site.Top(top, offset)
	return &gatherer{sites: sites, output: output}
}

func (g *gatherer) Gather() error {
	fmt.Print(g.sites)

	return nil
}
