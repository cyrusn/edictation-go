package voice

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"runtime"
)

// getVoiceSource will use a simple node app in ./voice.js
// ./voice.js will return a mp3 source from translate.google.come
func GetVoiceSource(title string, speed float64) (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("No caller information")
	}

	jsFilePath := path.Join(path.Dir(filename), "./voice.js")

	args := []string{
		jsFilePath,
		fmt.Sprintf(`-t="%s"`, title),
		fmt.Sprintf(`-s=%v`, speed),
	}
	cmd := exec.Command("node", args...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
