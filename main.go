package main

import (
	"fmt"
)

func main() {
	s, err := NewEdgeDependency()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s["cumulocityComponents"])
	//o := s.GetCategory("operator")
	//fmt.Println(s.GetComponentsForCategory("operator"))

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
