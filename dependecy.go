package main

import (
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Dependencies map[string]Category

type Category map[string]Component

type Component struct {
	Name                  string     `yaml:"name"`
	Location              string     `yaml:"location"`
	Version               string     `yaml:"version"`
	ContextPath           string     `yaml:"contextPath,omitempty"`
	Optional              bool       `yaml:"optional,omitempty"`
	ComponentSearchString string     `yaml:"componentSearchString,omitempty"`
	RepositoryURL         string     `yaml:"repositoryUrl,omitempty"`
	Resources             *Resources `yaml:"resources,omitempty"`
}
type Limits struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}
type Requests struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}
type Resources struct {
	Limits   Limits   `yaml:"limits,omitempty"`
	Requests Requests `yaml:"requests,omitempty"`
}

func (resources *Resources) GetLimits() *Limits {
	return &resources.Limits
}

func (resources *Resources) GetRequests() *Requests {
	return &resources.Requests
}

func (resources *Resources) SetLimits(limits Limits) *Resources {
	resources.Limits = limits
	return resources
}

func (resources *Resources) SetRequests(requests Requests) *Resources {
	resources.Requests = requests
	return resources
}
func (requests *Requests) GetCPU() string {
	return requests.CPU
}

func (requests *Requests) GetMemory() string {
	return requests.Memory
}

func (requests *Requests) SetCPU(cpu string) *Requests {
	requests.CPU = cpu
	return requests
}

func (requests *Requests) SetMemory(memory string) *Requests {
	requests.Memory = memory
	return requests
}
func (limits *Limits) GetCPU() string {
	return limits.CPU
}

func (limits *Limits) GetMemory() string {
	return limits.Memory
}

func (limits *Limits) SetCPU(cpu string) *Limits {
	limits.CPU = cpu
	return limits
}

func (limits *Limits) SetMemory(memory string) *Limits {
	limits.Memory = memory
	return limits
}
func (component *Component) GetName() string {
	return component.Name
}

func (component *Component) GetLocation() string {
	return component.Location
}

func (component *Component) GetVersion() string {
	return component.Version
}

func (component *Component) GetContextPath() string {
	return component.ContextPath
}

func (component *Component) GetOptional() bool {
	return component.Optional
}

func (component *Component) GetComponentSearchString() string {
	return component.ComponentSearchString
}

func (component *Component) GetRepositoryURL() string {
	return component.RepositoryURL
}

func (component *Component) GetResources() *Resources {
	return component.Resources
}

func (component *Component) SetName(name string) *Component {
	component.Name = name
	return component
}

func (component *Component) SetLocation(location string) *Component {
	component.Location = location
	return component
}

func (component *Component) SetVersion(version string) *Component {
	component.Version = version
	return component
}

func (component *Component) SetContextPath(contextPath string) *Component {
	component.ContextPath = contextPath
	return component
}

func (component *Component) SetOptional(optional bool) *Component {
	component.Optional = optional
	return component
}

func (component *Component) SetComponentSearchString(componentSearchString string) *Component {
	component.ComponentSearchString = componentSearchString
	return component
}

func (component *Component) SetRepositoryURL(repositoryURL string) *Component {
	component.RepositoryURL = repositoryURL
	return component
}

func (component *Component) SetResources(resources *Resources) *Component {
	component.Resources = resources
	return component
}

func NewEdgeDependency() (*Dependencies, error) {

	var (
		yamlDependenciesPath string = "/home/auc/go/src/github.com/nXnus25/yaml-conversion/dependencies.yaml"
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

	//fmt.Println(string(ymlBytes))

	d := &Dependencies{}

	err = yaml.Unmarshal(ymlBytes, d)
	if err != nil {
		return nil, err
	}

	fmt.Println("---\n", d)

	return d, nil
}

func (d Dependencies) GetCategory(name string) Category {
	if name != "" {

		fmt.Println(name, "  ")

		return d[name]
	}
	return nil
}

func (d Dependencies) GetComponentsForCategory(categoryName string) Component {
	c := d.GetCategory(categoryName)

	return c["operator"]
}
