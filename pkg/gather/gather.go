package gather

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"top-website-image-gatherer/pkg/site"

	"github.com/chromedp/chromedp"
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
	for _, site := range g.sites {
		if err := screenshot(site); err != nil {
			return fmt.Errorf("screenshotting: %w", err)
		}
	}

	return nil
}

func screenshot(s site.Site) error {
	s.EnsureScheme()
	url := s.Url.String()

	log.Printf("................making request for screenshot using %s", url)
	var opts []chromedp.ExecAllocatorOption
	opts = append(opts, chromedp.WindowSize(1400, 900))
	opts = append(opts, chromedp.DefaultExecAllocatorOptions[:]...)

	actx, acancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	var buf []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(&buf, 102),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s-%d.%s", s.Url.Path, time.Now().UTC().Unix(), "png")
	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("..............saved screenshot to file %s", filename)

	return nil
}
