package gfontlocal

import "testing"

func TestGetFont(*testing.T) {
	fonts := Fonts{}
	fonts.CssFile = "onefont.css"

	fonts2 := Fonts{}
	fonts2.CssFile = "twofonts.css"

	fonts3 := Fonts{}
	fonts3.CssFile = "twofonts2.css"

	font := Font{}
	font.Name = "Open Sans"
	font.FontPath = ""
	font.Size = append(font.Size, 300, 600)

	font2 := Font{}
	font2.Name = "Montserrat"
	font2.FontPath = ""
	font2.Size = append(font2.Size, 200)

	fonts.Fonts = append(fonts.Fonts, font)
	fonts2.Fonts = append(fonts2.Fonts, font, font2)
	fonts3.Fonts = append(fonts3.Fonts, font2, font)

	err := GetFont(fonts)
	if err != nil {
		panic(err)
	}
	err = GetFont(fonts2)
	if err != nil {
		panic(err)
	}
	err = GetFont(fonts3)
	if err != nil {
		panic(err)
	}
}
