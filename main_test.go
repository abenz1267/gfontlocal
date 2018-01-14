package gfontlocal

import (
	"testing"
)

func TestDownload(*testing.T) {
	firstFont := Font{Name: "Open Sans", Weights: "300,400"}
	secondFont := Font{Name: "Montserrat", Weights: "200,300"}

	fonts := Fonts{Fonts: []*Font{&firstFont, &secondFont}, CSSFolder: "./test/css", FontFolder: "./test/fonts", URL: "/public/fonts", SCSS: true, Dev: true}

	fonts.Download()
}
