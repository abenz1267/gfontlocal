package gfontlocal

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Fonts is a container struct
type Fonts struct {
	Fonts      []*Font
	CSSFolder  string
	FontFolder string
	URL        string
	SCSS       bool
	Dev        bool
}

// Font describes the according font you want to download
type Font struct {
	Name    string
	Weights string
}

// Download method for Fonts
func (f *Fonts) Download() {
	log.Println("Downloading fonts...")
	//var responses []string
	url := "https://fonts.googleapis.com/css?family="

	// replace spaces in fontname with "+" to make it url-friendly and build request url
	for i := range f.Fonts {
		f.Fonts[i].Name = strings.Replace(f.Fonts[i].Name, " ", "+", -1)
		url = url + f.Fonts[i].Name + ":" + f.Fonts[i].Weights

		if len(f.Fonts) > 1 && i < len(f.Fonts)-1 {
			url = url + "|"
		}
	}

	// download fonts
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")

	res, _ := client.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Fatal("can't get font. Status: " + strconv.Itoa(res.StatusCode))
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("(?s)\\/\\* latin \\*\\/.*?\\}")
	fontFaces := re.FindAll(data, -1)

	// extract all links
	var fontLinks []map[string][]byte

	for i, v := range fontFaces {
		fontMap := make(map[string][]byte)
		re = regexp.MustCompile("https?:\\/\\/?[\\da-z\\.-]+\\.[a-z\\.]{2,6}[\\/\\w \\.-]*\\/?")
		fontMap["link"] = re.Find(v)

		re = regexp.MustCompile("\\('.+?'\\)")

		// replace ' and space from name
		cleanName := bytes.Replace(re.Find(v), []byte(" "), []byte("_"), -1)
		cleanName = bytes.Replace(cleanName, []byte("('"), []byte(""), -1)
		cleanName = bytes.Replace(cleanName, []byte("')"), []byte(""), -1)
		fontMap["name"] = cleanName

		fontFaces[i] = bytes.Replace(fontFaces[i], fontMap["link"], []byte(f.URL+"/"+string(fontMap["name"])+".woff2"), 1)

		fontLinks = append(fontLinks, fontMap)
	}

	if f.Dev {
		// make css string
		var cssString []byte
		for _, v := range fontFaces {
			cssString = append(cssString, v...)
		}

		if f.SCSS {
			if err := ioutil.WriteFile(f.CSSFolder+"/_fonts.scss", cssString, 0644); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := ioutil.WriteFile(f.CSSFolder+"/fonts.css", cssString, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

	// download woff2
	for _, v := range fontLinks {
		req, _ := http.NewRequest("GET", string(v["link"]), nil)

		res, _ := client.Do(req)
		if res.StatusCode != http.StatusOK {
			log.Fatal("can't get font. Status: " + strconv.Itoa(res.StatusCode))
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(f.FontFolder+"/"+string(v["name"])+".woff2", data, 0644); err != nil {
			log.Fatal(err)
		}
		res.Body.Close()
	}
}
