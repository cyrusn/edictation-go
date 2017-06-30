package voice

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func GetVoice(title string) string {
	resp, err := http.Get("https://en.oxforddictionaries.com/definition/" + title)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	var srcs []string
	scrap(doc, &srcs)
	if len(srcs) != 0 {
		return srcs[0]
	}

	return ""
}

// TODO: break the loop once src found
func scrap(node *html.Node, srcs *[]string) {
	if node.Type == html.ElementNode && node.Data == "audio" {
		for _, a := range node.Attr {
			if a.Val != "" {
				*srcs = append(*srcs, a.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		scrap(c, srcs)
	}
}
