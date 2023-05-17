package main

import (
	"fmt"
)

func main() {
	s, err := NewDependency()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s.GetComponents(CUMULOCITY_CORE))
	c := s.GetComponent(CUMULOCITY_IOT_EDGE_OPERATOR)
	fmt.Printf("%#v\n", c)
	//o := s.GetCategory("operator")
	fmt.Println(c.Location)

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
