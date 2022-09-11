package site

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
	"log"
	"net/url"
)

//go:generate go run ./gen/gen.go
var (
	//go:embed gen/sites.gob
	sitesGob []byte
	sites    = func() (s []Site) {
		dec := gob.NewDecoder(bytes.NewReader(sitesGob))
		if err := dec.Decode(&s); err != nil {
			log.Fatalf("decoding sites gob: %s", err.Error())
		}
		return
	}()
)

// Site is a struct representing a website
type Site struct {
	Url *url.URL
}

// EnsureScheme checks if a site has a url scheme set then sets the scheme to https if it doesn't
func (s Site) EnsureScheme() {
	if s.Url.Scheme == "" {
		s.Url.Scheme = "https"
	}
}

// Top returns the top websites with an offset and number of sites
func Top(num, offset int) ([]Site, error) {
	if l := len(sites); offset+num >= l {
		return nil, fmt.Errorf("num %d + offset %d out of sites range 0 to %d", num, offset, l)
	}

	return sites[offset : offset+num], nil
}
