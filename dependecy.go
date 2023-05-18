package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Dependencies stores categories
// the map represents the file called dependencies.yaml
type Dependencies map[string]*Category

// Category represents map of root level tags in dependencies.yaml file
// Operator version
// Core components versions and their respective locations
// External 3rd party components (e.g. registry, Percona operator, MongoDB, etc.) versions and their respective locations
type Category map[string]*Component

// Component represents the structure and details
// Name, Location and Version are mandatory fields
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

type Resources struct {
	Limits   Limits   `yaml:"limits,omitempty"`
	Requests Requests `yaml:"requests,omitempty"`
}
type Limits struct {
	Cpu    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}
type Requests struct {
	Cpu    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}

// getter for all components in a category
// by specifying category we are getting map of components
func (d Dependencies) GetComponents(category Categories) *Category {
	return d[category.Name()]
}

// getter for specific component
// by specifying component name
func (d Dependencies) GetComponent(component Components) *Component {
	category := component.Category()
	r := *d[category.Name()]
	return r[component.Name()]
}

// loads and parse YAML dependencies file
func NewDependency() (*Dependencies, error) {

	var (
		yamlDependenciesPath string = fmt.Sprintf("%s/dependencies.yaml", "/data/ach/go/src/github.com/nonus25/dependencies-yaml")
	)

	yml, err := os.Open(yamlDependenciesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open [%v] - %v", yamlDependenciesPath, err)
	}
	defer yml.Close()

	ymlBytes, err := io.ReadAll(yml)
	if err != nil {
		return nil, fmt.Errorf("failed to read [%v] - %v", yamlDependenciesPath, err)
	}

	if d, err := parseDependenciesYaml(ymlBytes); err != nil {
		return nil, fmt.Errorf("failed to parse yaml [%v] - %v", yamlDependenciesPath, err)
	} else {
		return d, nil
	}
}

// function is parsing dependencies.yaml file
// split it into separate function for unit tests
func parseDependenciesYaml(input []byte) (*Dependencies, error) {
	d := &Dependencies{}
	fmt.Println(string(input))

	err := yaml.Unmarshal(input, d)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml - %v", err)
	}

	return d, nil
}

// all categories represented via enum
// listed in dependencies.yaml file
// new added category in Yaml needs to be also added here and vice versa
type Categories string

const (
	OPERATOR                 Categories = "operator"
	CUMULOCITY_CORE          Categories = "cumulocityCore"
	CUMULOCITY_MICROSERVICES Categories = "cumulocityMicroservices"
	CUMULOCITY_WEBAPPS       Categories = "cumulocityWebapps"
	THIRDPARTY               Categories = "thirdparty"
)

func (c Categories) Name() string {
	return string(c)
}

func (c Categories) String() string {
	return string(c)
}

// all components listed in dependencies.yaml file
// new dependency added to the file need to be also added here and vice versa
type Components uint

const (
	CUMULOCITY_IOT_EDGE_OPERATOR Components = iota + 1
	CACERTS
	CUMULOCITY_CORE_HELM_CHARTS
	CUMULOCITY_CORE_IMAGE
	OPENRESTY
	REGISTRY
	ALPINE_ACL
	BUSYBOX
	APAMA_CTRL
	SMARTRULE
	OPCUA_MGMT_SERVICE
	DEVICE_SIMULATOR
	ADMINISTRATION
	COCKPIT
	DEVICE_MANAGEMENT
	STREAMING_ANALYTICS
	MONGO_OPERATOR_HELM_CHARTS
	MONGO_SERVER_HELM_CHARTS
	MONGO
	LOGGING_OPERATOR_HELM_CHARTS
)

// returns the category to specific component
func (c Components) Category() Categories {
	switch c {
	case CUMULOCITY_IOT_EDGE_OPERATOR:
		return OPERATOR
	case CACERTS:
		return OPERATOR
	case CUMULOCITY_CORE_HELM_CHARTS:
		return CUMULOCITY_CORE
	case CUMULOCITY_CORE_IMAGE:
		return CUMULOCITY_CORE
	case OPENRESTY:
		return CUMULOCITY_CORE
	case REGISTRY:
		return CUMULOCITY_CORE
	case ALPINE_ACL:
		return CUMULOCITY_CORE
	case BUSYBOX:
		return CUMULOCITY_CORE
	case APAMA_CTRL:
		return CUMULOCITY_MICROSERVICES
	case SMARTRULE:
		return CUMULOCITY_MICROSERVICES
	case OPCUA_MGMT_SERVICE:
		return CUMULOCITY_MICROSERVICES
	case DEVICE_SIMULATOR:
		return CUMULOCITY_MICROSERVICES
	case ADMINISTRATION:
		return CUMULOCITY_WEBAPPS
	case COCKPIT:
		return CUMULOCITY_WEBAPPS
	case DEVICE_MANAGEMENT:
		return CUMULOCITY_WEBAPPS
	case STREAMING_ANALYTICS:
		return CUMULOCITY_WEBAPPS
	case MONGO_OPERATOR_HELM_CHARTS:
		return THIRDPARTY
	case MONGO_SERVER_HELM_CHARTS:
		return THIRDPARTY
	case MONGO:
		return THIRDPARTY
	case LOGGING_OPERATOR_HELM_CHARTS:
		return THIRDPARTY
	default:
		return Categories("unknown")
	}
}

func (c Components) Name() string {
	switch c {
	case CUMULOCITY_IOT_EDGE_OPERATOR:
		return "cumulocity-iot-edge-operator"
	case CACERTS:
		return "cacerts"
	case CUMULOCITY_CORE_HELM_CHARTS:
		return "cumulocity-core-helm-charts"
	case CUMULOCITY_CORE_IMAGE:
		return "cumulocity-core-image"
	case OPENRESTY:
		return "openresty"
	case REGISTRY:
		return "registry"
	case ALPINE_ACL:
		return "alpine_acl"
	case BUSYBOX:
		return "busybox"
	case APAMA_CTRL:
		return "apama-ctrl"
	case SMARTRULE:
		return "smartrule"
	case OPCUA_MGMT_SERVICE:
		return "opcua-mgmt-service"
	case DEVICE_SIMULATOR:
		return "device-simulator"
	case ADMINISTRATION:
		return "administration"
	case COCKPIT:
		return "cockpit"
	case DEVICE_MANAGEMENT:
		return "devicemanagement"
	case STREAMING_ANALYTICS:
		return "streaming-analytics"
	case MONGO_OPERATOR_HELM_CHARTS:
		return "mongo-operator-helm-charts"
	case MONGO_SERVER_HELM_CHARTS:
		return "mongo-server-helm-charts"
	case MONGO:
		return "mongo"
	case LOGGING_OPERATOR_HELM_CHARTS:
		return "logging-operator-helm-charts"
	default:
		return "unknown"
	}
}
