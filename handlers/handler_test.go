package handlers

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

// return Indexd config
func makeImageConfigString() string {
	return `{
				"url": "http://indexd-service/",
				"username": "test",
				"password": "test"
		  	}`
}

// Test that Indexd config can be marshalled into IndexdInfo struct
func TestHandler(t *testing.T) {
	indexdInfo := new(IndexdInfo)
	if err := json.Unmarshal([]byte(makeImageConfigString()), indexdInfo); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, indexdInfo.URL, "http://indexd-service/")
	assert.Equal(t, indexdInfo.Username, "test")
	assert.Equal(t, indexdInfo.Password, "test")

}

// Test mustGetConfigInfo function inputting Indexd and Metadata Service configs
func TestMustGetConfigInfoUsingIndexdAndMDSCreds(t *testing.T) {
	jsonConfigInfo :=
		`
	{
		"url": "http://indexd-service/index",
		"username": "mr happy cat",
		"password": "whiskers",
		"metadataService": {
			"url": "http://revproxy-service/mds",
			"username": "mr friendly cat",
			"password": "paws"
		}
	}
	`
	err := os.Setenv("CONFIG_FILE", jsonConfigInfo)
	if err != nil {
		t.Fatal(err)
	}
	configInfo := mustGetConfigInfo()
	assert.Equal(t, configInfo.Indexd.URL, "http://indexd-service/index")
	assert.Equal(t, configInfo.Indexd.Username, "mr happy cat")
	assert.Equal(t, configInfo.Indexd.Password, "whiskers")

	assert.Equal(t, configInfo.MetadataService.URL, "http://revproxy-service/mds")
	assert.Equal(t, configInfo.MetadataService.Username, "mr friendly cat")
	assert.Equal(t, configInfo.MetadataService.Password, "paws")
}

// Test that mustGetConfigInfo function panics when no Indexd or Metadata Service
// config is present
func TestMustGetConfigInfoWithoutIndexdOrMDSCreds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expecting mustGetConfigInfo to panic since Indexd and Metadata Service creds were not present in CONFIG_FILE")
		}
	}()

	jsonConfigInfo :=
		`
	{}
	`
	err := os.Setenv("CONFIG_FILE", jsonConfigInfo)
	if err != nil {
		t.Fatal(err)
	}
	mustGetConfigInfo()
}

// Test that mustGetConfigInfo function panics when no Metadata Service config is
// present
func TestMustGetConfigInfoWithoutMDSCreds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expecting mustGetConfigInfo to panic since Metadata Service creds were not present in CONFIG_FILE")
		}
	}()

	jsonConfigInfo :=
		`
	{
		"url": "http://indexd-service/index",
		"username": "mr happy cat",
		"password": "whiskers"
	}
	`
	err := os.Setenv("CONFIG_FILE", jsonConfigInfo)
	if err != nil {
		t.Fatal(err)
	}
	mustGetConfigInfo()
}

// Test that getConfigInfo function panics when no Indexd config is present
func TestMustGetConfigInfoWithoutIndexdCreds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expecting mustGetConfigInfo to panic since Indexd creds were not present in CONFIG_FILE")
		}
	}()

	jsonConfigInfo :=
		`
	{
		"metadataService": {
			"url": "http://revproxy-service/mds",
			"username": "mr friendly cat",
			"password": "paws"
		}
	}
	`
	err := os.Setenv("CONFIG_FILE", jsonConfigInfo)
	if err != nil {
		t.Fatal(err)
	}
	mustGetConfigInfo()
}

// Test mustGetConfigInfo function inputting extra service config
func TestMustGetConfigInfoUsingExtraServiceInfo(t *testing.T) {
	jsonConfigInfo :=
		`
	{
		"url": "http://indexd-service/index",
		"username": "mr happy cat",
		"password": "whiskers",
		"metadataService": {
			"url": "http://revproxy-service/mds",
			"username": "mr friendly cat",
			"password": "paws",
			"extra": "Always happy to see a human!"
		},
		"futureService": {
			"url": "http://future-service",
			"username": "mr futuristic cat",
			"password": "hover craft",
			"extra": "2100 is going to be purrfectly awesome!"
		}
	}
	`
	err := os.Setenv("CONFIG_FILE", jsonConfigInfo)
	if err != nil {
		t.Fatal(err)
	}
	configInfo := mustGetConfigInfo()
	assert.Equal(t, configInfo.Indexd.URL, "http://indexd-service/index")
	assert.Equal(t, configInfo.Indexd.Username, "mr happy cat")
	assert.Equal(t, configInfo.Indexd.Password, "whiskers")

	assert.Equal(t, configInfo.MetadataService.URL, "http://revproxy-service/mds")
	assert.Equal(t, configInfo.MetadataService.Username, "mr friendly cat")
	assert.Equal(t, configInfo.MetadataService.Password, "paws")
}
