package couchbase

import (
	"log"
	"os"
	"testing"
	"time"
)

var store *CouchbaseStore

type TestData struct {
	Data int
}

// Test setup and tear down to avoid repetition.
// It also deletes the document that was shared across functions for reading.
func TestMain(m *testing.M) {
	store = NewCouchbaseStore("couchbase://localhost", "test", "", "")
	if err := store.ConnectBucket(); err != nil {
		log.Fatalf("could not connect to host or bucket: %s\n", err)
	}

	// This data is created to test the Read query later.
	data := &TestData{Data: 1}
	readObj := NewCouchbaseObject("nickel")
	readObj.SetData(data)
	err := store.Create(readObj)
	if err != nil {
		log.Fatalf("could not make a TestData object in couchbase: %s\n", err)
	}
	// Ensure the data is written to couchbase before tests are run.
	time.Sleep(1 * time.Second)

	exitCode := m.Run()

	err = store.DestroyOne(readObj.key)
	if err != nil {
		log.Fatalf("could not destroy the TestData: %s\n", err)
	}
	store.ShutdownBucket()

	os.Exit(exitCode)
}

func checkFatal(t *testing.T, msg string, err error) {
	if err != nil {
		t.Fatalf("%s: %s\n", msg, err)
	}
}

func TestCRUD(t *testing.T) {
	testObj := NewCouchbaseObject("delete-me")
	testObj.SetData("first")

	err := store.Create(testObj)
	checkFatal(t, "could not create an object", err)

	err, updatedObj := store.ReadOne(testObj.key)
	checkFatal(t, "could not read an object", err)

	updatedObj.SetData("second")
	err = store.UpdateOne(updatedObj)
	checkFatal(t, "could not update object", err)

	err = store.DestroyOne(testObj.key)
	checkFatal(t, "could not destroy object", err)
}

func TestRead(t *testing.T) {
	query := NewNickelQuery("SELECT * FROM test WHERE Data = 1", store.cluster)
	err, results := store.Read(query)
	if err != nil {
		t.Fatalf("could not perform read query: %s\n", err)
	}

	if len(results) != 1 {
		t.Fatalf("expcted result length of one but got: %d\n", len(results))
	}
}
