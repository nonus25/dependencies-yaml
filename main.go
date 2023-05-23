package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s, err := NewDependency()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s.GetComponents(CUMULOCITY_WEBAPPS))
	c := s.GetComponent(STREAMING_ANALYTICS)
	fmt.Printf("%#v\n", c)
	//o := s.GetCategory("operator")
	fmt.Println(c.Location)
	fmt.Println(c.Name)

	g := filepath.Join("/var/assets", "/dependency.yaml")
	fmt.Println(g)

	/*d := make(Dependencies, 0)
	c := make(Category, 0)

	d["thirdpartyComponents"] = c
	c["mongo"] = Component{
		Name:        "mongo",
		Location:    "location",
		Version:     "version",
		ContextPath: "ctxPath",
	}

	yamlData, err := yaml.Marshal(&d)
	if err != nil {
		log.Println("", err)
	}
	fmt.Println(" --- YAML ---")
	fmt.Println(string(yamlData))
	*/

}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func (s singleton) String() string {
	return fmt.Sprintf("called: %d", s.count)
}
