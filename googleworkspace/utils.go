package googleworkspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

// Returns the content of given file, or the inline JSON credential as it is
func pathOrContents(poc string) (string, error) {
	if len(poc) == 0 {
		return poc, nil
	}

	path, err := expandPath(poc)
	if err != nil {
		return path, err
	}

	// Check for valid file path
	if _, err := os.Stat(path); err == nil {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return string(contents), err
		}
		return string(contents), nil
	}

	// Return error if content is a file path and the file doesn't exist
	if len(path) > 1 && (path[0] == '/' || path[0] == '\\') {
		return "", fmt.Errorf("%s: no such file or dir", path)
	}

	// Return the inline content
	return poc, nil
}

// Expands the path to include the home directory if the path is prefixed with `~`
func expandPath(filePath string) (string, error) {
	// Check if the path has `~` to denote the home dir
	path := filePath
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, err
		}
	}
	return path, nil
}

type CredsConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func getConfigFromCreds(path string) (*CredsConfig, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config *CredsConfig
	err = json.Unmarshal(contents, &config)
	return config, nil
}
