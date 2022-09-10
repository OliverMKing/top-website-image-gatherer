//go:build exclude

// Generates a slice of all sites
package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"encoding/gob"
	"io"
	"log"
	"net/url"
	"os"
	"top-website-image-gatherer/pkg/site"
)

const (
	filename = "sites.gob"
)

var (
	//go:embed sites.csv
	sitesCsv []byte
)

func main() {
	r := csv.NewReader(bytes.NewReader(sitesCsv))

	// skip headings (first line)
	if _, err := r.Read(); err != nil {
		log.Fatalf("reading first line: %s", err.Error())
	}

	var s []site.Site
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("reading line: %s", err.Error())
		}

		u, err := url.Parse(record[2])
		if err != nil {
			log.Fatalf("parsing url: %s", err.Error())
		}
		s = append(s, site.Site{Url: u})
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&s); err != nil {
		log.Fatalf("encoding sites: %s", err.Error())
	}

	if err := os.WriteFile(filename, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("writing %s: %s", filename, err.Error())
	}
}
