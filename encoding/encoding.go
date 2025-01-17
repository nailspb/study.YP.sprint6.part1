package encoding

import (
	"encoding/json"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	file, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &j.DockerCompose)
	if err != nil {
		return err
	}
	output, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(j.FileOutput, output, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	file, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &y.DockerCompose)
	if err != nil {
		return err
	}
	output, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(y.FileOutput, output, 0644)
	if err != nil {
		return err
	}
	return nil
}
