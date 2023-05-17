package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Dependencies map[string]*Category

type Category map[string]*Component

type Component struct {
	Location              string     `yaml:"location"`
	Version               string     `yaml:"version"`
	ContextPath           string     `yaml:"contextPath,omitempty"`
	Optional              bool       `yaml:"optional,omitempty"`
	ComponentSearchString string     `yaml:"componentSearchString,omitempty"`
	RepositoryURL         string     `yaml:"repositoryUrl,omitempty"`
	Resources             *Resources `yaml:"resources,omitempty"`
}
type Limits struct {
	Cpu    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}
type Requests struct {
	Cpu    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}
type Resources struct {
	Limits   Limits   `yaml:"limits,omitempty"`
	Requests Requests `yaml:"requests,omitempty"`
}

func (d Dependencies) GetComponents(category Categories) *Category {
	return d[category.Name()]
}

func (d Dependencies) GetComponent(component Components) *Component {
	category := component.Category()
	r := *d[category.Name()]
	return r[component.Name()] //d.GetComponents(category)[component.Name()]
}

func NewDependency() (*Dependencies, error) {

	var (
		yamlDependenciesPath string = fmt.Sprintf("%s/dependencies.yaml", "/data/ach/go/src/github.com/nonus25/dependencies-yaml")
	)

	yml, err := os.Open(yamlDependenciesPath)
	if err != nil {
		return nil, err
	}
	defer yml.Close()

	ymlBytes, err := io.ReadAll(yml)
	if err != nil {
		return nil, err
	}

	d := &Dependencies{}

	err = yaml.Unmarshal(ymlBytes, d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

type Categories string

const (
	OPERATOR        Categories = "operator"
	CUMULOCITY_CORE Categories = "cumulocityCore"
)

func (c Categories) Name() string {
	return string(c)
}

func (c Categories) String() string {
	return string(c)
}

type Components uint

const (
	CUMULOCITY_IOT_EDGE_OPERATOR Components = iota
	CUMULOCITY_CORE_HELM_CHARTS
)

func (c Components) Category() Categories {
	switch c {
	case CUMULOCITY_IOT_EDGE_OPERATOR:
		return OPERATOR
	case CUMULOCITY_CORE_HELM_CHARTS:
		return CUMULOCITY_CORE
	default:
		return Categories("unknown")
	}
}

func (c Components) Name() string {
	switch c {
	case CUMULOCITY_IOT_EDGE_OPERATOR:
		return "cumulocity-iot-edge-operator"
	case CUMULOCITY_CORE_HELM_CHARTS:
		return "cumulocity-core-helm-charts"
	}
	return string(c)
}

func (c Components) String() string {
	return string(c)
}
