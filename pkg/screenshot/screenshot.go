package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"top-website-image-gatherer/pkg/site"

	"github.com/chromedp/chromedp"
)

// Screenshoter screenshots a website and places the screenshot in the output directory
type Screenshoter interface {
	Screenshot(s site.Site, output string) error
}

type screenshoter struct{}

var _ Screenshoter = &screenshoter{}

func New() Screenshoter {
	return &screenshoter{}
}

func (ss *screenshoter) Screenshot(s site.Site, output string) error {
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
