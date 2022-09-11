package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"
	"twig/pkg/site"

	"github.com/chromedp/chromedp"
)

// Screenshoter screenshots a website and places the screenshot in the output directory
type Screenshoter interface {
	Screenshot(s site.Site, output string) error
}

type screenshoter struct {
	wait time.Duration
}

var _ Screenshoter = &screenshoter{}

// New creates a new screenshoter. Wait is the amount of time to wait for a page to load before screenshotting
func New(wait time.Duration) Screenshoter {
	return &screenshoter{wait: wait}
}

func (ss *screenshoter) Screenshot(s site.Site, output string) error {
	s.EnsureScheme()
	url := s.Url.String()

	log.Printf("making request for screenshot using %s", url)
	var opts []chromedp.ExecAllocatorOption
	opts = append(opts, chromedp.WindowSize(1400, 900))
	opts = append(opts, chromedp.DefaultExecAllocatorOptions[:]...)

	actx, acancel := chromedp.NewExecAllocator(context.TODO(), opts...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	var buf []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(ss.wait),
		chromedp.FullScreenshot(&buf, 102),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return fmt.Errorf("running screenshot tasks for %s: %w", url, err)
	}

	filename := fmt.Sprintf("%s-%d.%s", s.Url.Path, time.Now().UTC().Unix(), "png")
	filepath := path.Join(output, filename)
	if err := ioutil.WriteFile(filepath, buf, 0644); err != nil {
		return fmt.Errorf("writing image %s for %s: %w", filepath, url, err)
	}
	log.Printf("saved screenshot to file %s", filepath)

	return nil
}
