# gfontlocal

Gfontlocal allows you to serve up-to-date Google Fonts locally.

Only works for latin and woff2/ttf.

# Usage
```
go get -u github.com/abenz1267/gfontlocal
```


Example:

```
fontfolder := "public/css/"

font := gfontlocal.Font{"Open Sans", []int{400, 600}, fontfolder}
font2 := gfontlocal.Font{"Montserrat", []int{200}, "woff2",fontfolder} // woff2 is the default type, you can also just "" instead of "woff2"

fonts := gfontlocal.Fonts{[]gfontlocal.Font{font, font2}, "public/css/font.css"}

err := gfontlocal.GetFont(fonts)
if err != nil {
  fmt.Println(err)
}
```


# Disclaimer

This is an initial release.

I've created this project for personal usage and learning. Suggestions, feature requests or general feedback is appreciated.
