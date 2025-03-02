package yamlparser

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/ghodss/yaml"
)

func ParseYAML(yamlPath string, out interface{}) error {
	if reflect.ValueOf(out).Kind() != reflect.Pointer {
		return fmt.Errorf("function ParseYAML requires to provide a pointer")
	}

	file, err := os.Open(yamlPath)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, out)
}
