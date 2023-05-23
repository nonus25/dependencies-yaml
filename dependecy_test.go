package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeDependencies(t *testing.T) {

	d := make(Dependencies, 0)
	c := make(Category, 0)

	d["thirdpartyComponents"] = &c
	/*c = append(c, &Component{
		Name:     "mongo",
		Location: "percona/percona-server-mongodb",
		Version:  "5.0.14",
		Resources: &Resources{
			Limits: Limits{
				Cpu:    "100m",
				Memory: "2G",
			},
			Requests: Requests{
				Cpu:    "100m",
				Memory: "1G",
			},
		},
	})*/

	rd, err := NewDependency()
	if err != nil {
		t.Errorf("failed to read yaml dependencies file - %v", err)
	}
	assert.NotEmpty(t, rd)

	//rc := rd["thirdpartyComponents"

}

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		//Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(),not nil")
	}
	fmt.Println("1", counter1)

	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
	fmt.Println("2", counter2)

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
