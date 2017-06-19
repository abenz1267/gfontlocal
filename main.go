package gfontlocal

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Fonts type
type Fonts struct {
	Fonts   []Font
	CssFile string
}

//Font type
type Font struct {
	Name     string
	Size     []int
	FontPath string
}

type fontLink struct {
	link     string
	filename string
}

// GetFont from google as woff2
func GetFont(fonts Fonts) error {
	var link string
	var fontStrings []string
	var fontLinks []fontLink
	var cssFile string
	var err error

	for _, font := range fonts.Fonts {
		if font.Size != nil {
			for _, v := range font.Size {
				size := strconv.Itoa(v)
				font.Name = strings.Replace(font.Name, " ", "+", -1)
				filename := font.FontPath + font.Name + "_" + size + ".woff2"
				link = "https://fonts.googleapis.com/css?family=" + font.Name + ":" + size
				fontStrings, fontLinks, err = fontData(fontStrings, fontLinks, filename, link)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("no font-size set")
		}
	}

	for i, v := range fontLinks {
		fontStrings[i] = strings.Replace(fontStrings[i], v.link, "/"+v.filename, -1)
		err := downloadFile(v.filename, v.link)
		if err != nil {
			return err
		}
	}

	for _, v := range fontStrings {
		cssFile = cssFile + strings.TrimSpace(v)
	}

	if _, err := os.Stat(fonts.CssFile); os.IsNotExist(err) {
		err = ioutil.WriteFile(fonts.CssFile, []byte(cssFile), 0644)
		if err != nil {
			return err
		}
	} else {
		err := os.Remove(fonts.CssFile)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fonts.CssFile, []byte(cssFile), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func fontData(fontStrings []string, fontLinks []fontLink, filename, link string) ([]string, []fontLink, error) {
	re := regexp.MustCompile("https?:\\/\\/?[\\da-z\\.-]+\\.[a-z\\.]{2,6}[\\/\\w \\.-]*\\/?")

	fontString, err := getFontCSS(link)
	if err != nil {
		return nil, nil, err
	}

	fontLink := fontLink{re.FindString(fontString), filename}
	fontStrings = append(fontStrings, fontString)
	fontLinks = append(fontLinks, fontLink)

	return fontStrings, fontLinks, nil
}

func getFontCSS(link string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")

	res, _ := client.Do(req)
	if res.StatusCode != http.StatusOK {
		return "", errors.New("can't get font")
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	firstSplit := strings.SplitAfterN(responseString, "/* latin */", -1)
	firstFont := strings.SplitAfterN(firstSplit[1], "}", -1)[0]

	return firstFont, err
}

// credit to https://stackoverflow.com/users/1511332/pablo-jomer
func downloadFile(filepath string, link string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
