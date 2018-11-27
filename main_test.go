package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// this is where we make a fake struct to replace a real bigquery config
type fakebqclient struct {
	context string
	client string
}

// this is a fake bigquery method that will always return what we want
func (bq fakebqclient) getDatasets() string {
	return "hello i am a fake list of datasets"
}

func TestEnsureDataset(t *testing.T) {
	bq := BQClient{
		Bq: fakebqclient{},
	}
	datasetlist := bq.ensureDataset()
	assert.Equal(t, "hello i am a fake list of datasets", datasetlist, "not using the mocked method")
}