mport (
	"edge-k8s-operator/edge/constants"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
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

func (d *Dependencies) GetComponents(category Categories) map[string]*Component {
	c := *d
	return c[category.String()]
}

func (d *Dependencies) GetComponent(category Categories, component Components) *Component {

}

func NewDependency() (*Dependencies, error) {

	var (
		yamlDependenciesPath string = fmt.Sprintf("%s/dependencies.yaml", constants.OPERATOR_ASSETS_DIR)
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

type Categories string

const (
	OPERATOR string = "operator"
)

func (c Categories) Name() {
	return 
}