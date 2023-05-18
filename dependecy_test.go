package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependenciesEnums(t *testing.T) {

	// testing Categories
	want := "thirdparty"
	got := THIRDPARTY.Name()
	assert.Equal(t, want, got)
	want = "operator"
	got = OPERATOR.Name()
	assert.Equal(t, want, got)
	want = "cumulocityCore"
	got = CUMULOCITY_CORE.Name()
	assert.Equal(t, want, got)
	want = "cumulocityMicroservices"
	got = CUMULOCITY_MICROSERVICES.Name()
	assert.Equal(t, want, got)
	want = "cumulocityWebapps"
	got = CUMULOCITY_WEBAPPS.Name()
	assert.Equal(t, want, got)

	// testing components
	want = "cumulocity-iot-edge-operator"
	got = CUMULOCITY_IOT_EDGE_OPERATOR.Name()
	assert.Equal(t, want, got)
	want = "cacerts"
	got = CACERTS.Name()
	assert.Equal(t, want, got)

	want = "cumulocity-core-helm-charts"
	got = CUMULOCITY_CORE_HELM_CHARTS.Name()
	assert.Equal(t, want, got)
	want = "cumulocity-core-image"
	got = CUMULOCITY_CORE_IMAGE.Name()
	assert.Equal(t, want, got)
	want = "openresty"
	got = OPENRESTY.Name()
	assert.Equal(t, want, got)
	want = "registry"
	got = REGISTRY.Name()
	assert.Equal(t, want, got)
	want = "alpine_acl"
	got = ALPINE_ACL.Name()
	assert.Equal(t, want, got)
	want = "busybox"
	got = BUSYBOX.Name()
	assert.Equal(t, want, got)

	want = "apama-ctrl"
	got = APAMA_CTRL.Name()
	assert.Equal(t, want, got)
	want = "smartrule"
	got = SMARTRULE.Name()
	assert.Equal(t, want, got)
	want = "opcua-mgmt-service"
	got = OPCUA_MGMT_SERVICE.Name()
	assert.Equal(t, want, got)
	want = "device-simulator"
	got = DEVICE_SIMULATOR.Name()
	assert.Equal(t, want, got)

	want = "administration"
	got = ADMINISTRATION.Name()
	assert.Equal(t, want, got)
	want = "cockpit"
	got = COCKPIT.Name()
	assert.Equal(t, want, got)
	want = "devicemanagement"
	got = DEVICE_MANAGEMENT.Name()
	assert.Equal(t, want, got)
	want = "streaming-analytics"
	got = STREAMING_ANALYTICS.Name()
	assert.Equal(t, want, got)

	want = "mongo-operator-helm-charts"
	got = MONGO_OPERATOR_HELM_CHARTS.Name()
	assert.Equal(t, want, got)
	want = "mongo-server-helm-charts"
	got = MONGO_SERVER_HELM_CHARTS.Name()
	assert.Equal(t, want, got)
	want = "mongo"
	got = MONGO.Name()
	assert.Equal(t, want, got)
	want = "logging-operator-helm-charts"
	got = LOGGING_OPERATOR_HELM_CHARTS.Name()
	assert.Equal(t, want, got)

	// testing Category grouping
	want = "operator"
	got = CUMULOCITY_IOT_EDGE_OPERATOR.Category().Name()
	assert.Equal(t, want, got)
	got = CACERTS.Category().Name()
	assert.Equal(t, want, got)

	want = "cumulocityCore"
	got = CUMULOCITY_CORE_HELM_CHARTS.Category().Name()
	assert.Equal(t, want, got)
	got = CUMULOCITY_CORE_IMAGE.Category().Name()
	assert.Equal(t, want, got)
	got = OPENRESTY.Category().Name()
	assert.Equal(t, want, got)
	got = REGISTRY.Category().Name()
	assert.Equal(t, want, got)
	got = ALPINE_ACL.Category().Name()
	assert.Equal(t, want, got)
	got = BUSYBOX.Category().Name()
	assert.Equal(t, want, got)

	want = "cumulocityMicroservices"
	got = APAMA_CTRL.Category().Name()
	assert.Equal(t, want, got)
	got = SMARTRULE.Category().Name()
	assert.Equal(t, want, got)
	got = OPCUA_MGMT_SERVICE.Category().Name()
	assert.Equal(t, want, got)
	got = DEVICE_SIMULATOR.Category().Name()
	assert.Equal(t, want, got)

	want = "cumulocityWebapps"
	got = ADMINISTRATION.Category().Name()
	assert.Equal(t, want, got)
	got = COCKPIT.Category().Name()
	assert.Equal(t, want, got)
	got = DEVICE_MANAGEMENT.Category().Name()
	assert.Equal(t, want, got)
	got = STREAMING_ANALYTICS.Category().Name()
	assert.Equal(t, want, got)

	want = "thirdparty"
	got = MONGO_OPERATOR_HELM_CHARTS.Category().Name()
	assert.Equal(t, want, got)
	got = MONGO_SERVER_HELM_CHARTS.Category().Name()
	assert.Equal(t, want, got)
	got = MONGO.Category().Name()
	assert.Equal(t, want, got)
	got = LOGGING_OPERATOR_HELM_CHARTS.Category().Name()
	assert.Equal(t, want, got)

}

func TestNewDependencies(t *testing.T) {

	// static representation of dependencies Yaml file
	// if u modify this string make sure it does not contain TABs
	// otherwise the test will fail
	input := []byte(`---

operator:
  cumulocity-iot-edge-operator:
    name: cumulocity-iot-edge-operator
    location: chartrepo/edge/cumulocity-iot-edge-operator
    version: 1017.0.0-2224

  cacerts:
    name: cacerts
    location: sag/kubernetes-edge/1017.0.0/cacerts                       
    version: latest                                           

cumulocityCore:
  cumulocity-core-helm-charts:
    name: cumulocity-core-helm-charts
    componentSearchString: cumulocity-core helm chart
    location: chartrepo/stack/cumulocity-core
    version: 1017.0.13

  cumulocity-core-image:
    name:  cumulocity-core-image
    resources:
      limits:
        cpu: 2000m
        memory: 6G
      requests:
        cpu: 2000m
        memory: 2G
    componentSearchString: cumulocity-core container image
    location: stack/core
    version: 1017.0.13-1

  openresty:
    name: openresty
    location: stack/openresty
    version: 1.19.3.2-20.el7.c8y.1011.0.0

  registry:
    name: registry
    resources:
      limits:
        cpu: 500m
        memory: 1G
      requests:
        cpu: 250m
        memory: 500Mi
    location: cumulocity/registry
    version: 2.8.2

  alpine_acl:
    name: alpine_acl
    location: stack/alpine_acl
    version: latest                                         

  busybox:
    name: busybox 
    location: stack/busybox
    version: latest                                         

cumulocityMicroservices:
  apama-ctrl:
    name: apama-ctrl
    contextPath: cep
    optional: false
    resources:
      limits:
        cpu: 1000m
        memory: 1G
    componentSearchString: apama-ctrl-1c-4g
    location: sag/kubernetes-edge/1017.0.0/apama-ctrl-1c-4g-10.17.0.0.20230131-1844-3afe51bcf.zip
    version: 1017.0.0

  smartrule:
    name: smartrule
    contextPath: smartrule
    optional: false
    resources:
        limits:
        cpu: 1000m
        memory: 1G
    componentSearchString: smartrule
    location: sag/kubernetes-edge/1017.0.0/smartrule-1017.0.13.zip
    version: 1017.0.0

  opcua-mgmt-service:
    name: opcua-mgmt-service
    contextPath: opcua-mgmt-service
    optional: false
    resources:
      limits:
        cpu: 1000m
        memory: 1G
    componentSearchString: opcua-mgmt-service
    location: sag/kubernetes-edge/1017.0.0/opcua-mgmt-service-1017.0.13.zip
    version: 1017.0.0

  device-simulator:
    name: device-simulator
    contextPath: device-simulator
    optional: false
    resources:
      limits:
        cpu: 250m
        memory: 500Mi
    componentSearchString: device-simulator
    location: sag/kubernetes-edge/1017.0.0/device-simulator-1017.0.13.zip
    version: 1017.0.0

cumulocityWebapps:
  administration:
    name: administration
    optional: false
    componentSearchString: administration
    location: sag/kubernetes-edge/1017.0.0/administration-1017.0.13.zip
    version: 1017.0.0

  cockpit:
    name: cockpit
    optional: false
    componentSearchString: cockpit
    location: sag/kubernetes-edge/1017.0.0/cockpit-1017.0.13.zip
    version: 1017.0.0

  devicemanagement:
    name: "devicemanagement"
    optional: false
    componentSearchString: device-management
    location: sag/kubernetes-edge/1017.0.0/devicemanagement-1017.0.13.zip
    version: 1017.0.0

  streaming-analytics:
    name: "Streaming Analytics"
    optional: false
    componentSearchString: streaming-analytics-app
    location: sag/kubernetes-edge/1017.0.0/streaming-analytics-app-10.17.0.0.20230131-1844-3afe51bcf.zip
    version: 1017.0.0

thirdparty:                                                 
  mongo-operator-helm-charts:
    name: mongo-operator-helm-charts
    repositoryUrl: https://percona.github.io/percona-helm-charts/
    location: percona/psmdb-operator 
    version: 1.14.0

  mongo-server-helm-charts:
    name: mongo-server-helm-charts
    repositoryUrl: https://percona.github.io/percona-helm-charts/
    location: percona/psmdb-db 
    version: 1.14.0

  mongo:
    name: mongo
    resources:
      limits:
        cpu: 1000m
        memory: 2G
      requests:
        cpu: 1000m
        memory: 1G
    location: percona/percona-server-mongodb                          
    version: 5.0.14

  logging-operator-helm-charts:
    name: logging-operator-helm-charts
    repositoryUrl: https://kubernetes-charts.banzaicloud.com         
    location: banzaicloud-stable/logging-operator 
    version: 3.17.10
`)

	// stub object creation to compare with parsing results
	d := make(Dependencies, 0)
	c := make(Category, 0)

	c[MONGO_OPERATOR_HELM_CHARTS.Name()] = &Component{
		Name:          "mongo-operator-helm-charts",
		RepositoryURL: "https://percona.github.io/percona-helm-charts/",
		Location:      "percona/psmdb-operator",
		Version:       "1.14.0",
	}
	c[MONGO_SERVER_HELM_CHARTS.Name()] = &Component{
		Name:          "mongo-server-helm-charts",
		RepositoryURL: "https://percona.github.io/percona-helm-charts/",
		Location:      "percona/psmdb-db",
		Version:       "1.14.0",
	}
	c[MONGO.Name()] = &Component{
		Name:     "mongo",
		Location: "percona/percona-server-mongodb",
		Version:  "5.0.14",
		Resources: &Resources{
			Limits: Limits{
				Cpu:    "1000m",
				Memory: "2G",
			},
			Requests: Requests{
				Cpu:    "1000m",
				Memory: "1G",
			},
		},
	}
	c[LOGGING_OPERATOR_HELM_CHARTS.Name()] = &Component{
		Name:          "logging-operator-helm-charts",
		RepositoryURL: "https://kubernetes-charts.banzaicloud.com",
		Location:      "banzaicloud-stable/logging-operator",
		Version:       "3.17.10",
	}

	d[THIRDPARTY.Name()] = &c

	assert.NotEmpty(t, d)

	// testing if file exists and can be parsed by NewDependencies
	// nd as NewDependency (ND)
	nd, err := NewDependency()
	if err != nil {
		t.Errorf("failed to read yaml dependencies file - %v", err)
	}
	assert.NotEmpty(t, nd)

	// loading static structure as string
	// pdy as parseDependenciesYaml
	pdy, err := parseDependenciesYaml(input)
	if err != nil {
		t.Errorf("failed to parse yaml dependencies file - %v", err)
	}
	assert.NotEmpty(t, pdy)

	wantComponents := d[THIRDPARTY.Name()]
	gotComponents := pdy.GetComponents(THIRDPARTY)
	assert.Equal(t, wantComponents, gotComponents)

	wantComponent := c[MONGO.Name()]
	gotComponent := pdy.GetComponent(MONGO)
	assert.Equal(t, wantComponent, gotComponent)
}
