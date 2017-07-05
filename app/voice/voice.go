package voice

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetVoiceResponse will return the mp3 source url with *http.Response format
func GetVoiceResponse(title string, port string) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost"+port+"/voice/%s", title)
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
