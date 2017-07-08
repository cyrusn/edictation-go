package voice

import (
	"fmt"
	"testing"
)

func TestGetVoiceSource(t *testing.T) {
	title := "hello"
	speed := 0.24

	src, err := GetVoiceSource(title, speed)
	fmt.Println("result: ", src)
	if err != nil {
		t.Error(src)
	}
}
