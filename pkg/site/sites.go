package site

import (
	"bytes"
	_ "embed"
	"encoding/gob"
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

type Site struct {
	Url *url.URL
}

// Top returns the top websites with an offset and number of sites
func Top(top, offset int) []Site {
	return sites[offset : offset+top]
}
