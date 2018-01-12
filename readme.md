# gfontlocal

Gfontlocal allows you to serve up-to-date Google Fonts locally.

Only works for latin woff2 fonts.

# Usage

```
go get -u github.com/abenz1267/gfontlocal
```

Example:

```
firstFont := Font{Name: "Open Sans", Weights: "300,400"}
secondFont := Font{Name: "Montserrat", Weights: "200,300"}

fonts := Fonts{Fonts: []*Font{&firstFont, &secondFont}, CSSFolder: "./test/css", FontFolder: "./test/fonts", URL: "/public/fonts", SCSS: true}

fonts.Download()
```

# Disclaimer

This is an initial release.

I've created this project for personal usage and learning. Suggestions, feature requests or general feedback is appreciated.
