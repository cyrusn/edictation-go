package voice

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetVoiceResponse ...
func GetVoiceResponse(title string) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:4000/voice/%s", title)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return http.Get(string(src))
}

// func GetVoiceSrc(title string) (*http.Response, error) {
//
// 	prefixs := []string{"1", "1_8", "2"}
//
// 	for _, prefix := range prefixs {
// 		src := "http://audio.oxforddictionaries.com/en/mp3/" + title + "_gb_" + prefix + ".mp3"
// 		resp, err := http.Get(src)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if resp.StatusCode != 403 {
// 			return resp, nil
// 		}
// 	}
//
// 	resp, err := http.Get("https://en.oxforddictionaries.com/definition/" + title)
// 	defer resp.Body.Close()
//
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	doc, err := html.Parse(bytes.NewReader(body))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var src string
// 	scrap(doc, &src)
// 	return http.Get(src)
// }
//
// func scrap(node *html.Node, src *string) {
// 	if node.Type == html.ElementNode && node.Data == "audio" {
// 		for _, a := range node.Attr {
// 			if a.Val != "" {
// 				*src = a.Val
// 				break
// 			}
// 		}
// 	}
//
// 	for c := node.FirstChild; c != nil && *src == ""; c = c.NextSibling {
// 		scrap(c, src)
// 	}
// }
