package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeDependencies(t *testing.T) {

	d := make(Dependencies, 0)
	c := make(Category, 0)

	d["thirdpartyComponents"] = &c
	c = append(c, &Component{
		Name:     "mongo",
		Location: "percona/percona-server-mongodb",
		Version:  "5.0.14",
		Resources: &Resources{
			Limits: Limits{
				CPU:    "100m",
				Memory: "2G",
			},
			Requests: Requests{
				CPU:    "100m",
				Memory: "1G",
			},
		},
	})

	rd, err := NewEdgeDependency()
	if err != nil {
		t.Errorf("failed to read yaml dependencies file - %v", err)
	}
	assert.NotEmpty(t, rd)

	rc := rd["thirdpartyComponents"]

	name := rc.

}
