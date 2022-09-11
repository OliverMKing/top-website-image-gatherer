package gather

import (
	"context"
	"log"
	"top-website-image-gatherer/pkg/screenshot"
	"top-website-image-gatherer/pkg/site"

	"golang.org/x/sync/errgroup"
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
	log.Print("starting to gather screenshots")

	eg, _ := errgroup.WithContext(context.TODO())
	sitesCh := make(chan site.Site)

	numGoroutines := 5
	for i := 0; i < numGoroutines; i++ {
		eg.Go(func(ch <-chan site.Site) func() error {
			return func() error {
				for site := range ch {
					if err := g.screenshoter.Screenshot(site, output); err != nil {
						return err
					}
				}

				return nil
			}
		}(sitesCh))
	}

	for _, site := range g.sites {
		sitesCh <- site
	}
	close(sitesCh)

	if err := eg.Wait(); err != nil {
		return err
	}

	log.Printf("finished gathering screenshots to %s", output)
	return nil
}
