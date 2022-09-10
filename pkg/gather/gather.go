package gather

import (
	"fmt"
	"top-website-image-gatherer/pkg/screenshot"
	"top-website-image-gatherer/pkg/site"
)

type progress struct {
	complete int
	total    int
}

// Gatherer gathers screenshots of top websites and places them in the output directory
type Gatherer interface {
	Gather(output string) error
}

type gatherer struct {
	sites        []site.Site
	screenshoter screenshot.Screenshoter
}

var _ Gatherer = &gatherer{}

// New creates a new gatherer that will screenshot the provided sites
func New(sites []site.Site, screenshoter screenshot.Screenshoter) Gatherer {
	return &gatherer{sites: sites, screenshoter: screenshoter}
}

// Gather gathers the screenshots of top websites and places them in the output directory
func (g *gatherer) Gather(output string) error {
	for _, site := range g.sites {
		if err := g.screenshoter.Screenshot(site, output); err != nil {
			return fmt.Errorf("screenshotting: %w", err)
		}
	}

	return nil
}
