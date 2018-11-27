package main

// this is the interface type that will need to be mocked in tests
type BQClientinterface interface {
	getDatasets() (string)
}

// this is the actual struct where I would construct a real bq client object, prolly using a build method
type BQClientStruct struct {
	context string
	client string
}

// this is the structure abstraction that we actually use in our code
type BQClient struct {
	BQClientinterface
}

// this is an example of a method that in real life would reach out to an external api so it would need to be mocked
func (bq BQClientStruct) getDatasets() string {
	return "real list of datasets"
}

func main() {
	bq := BQClient{
		BQClientStruct{
			client: "mooclient", 
			context: "moocontext",
		},
	}
	bq.ensureDataset()
}

// this is an example of a method that we want to test that calls a method that calls an external api (CALLS!)
func (u *BQClient) ensureDataset() string {
	return u.getDatasets()
}

