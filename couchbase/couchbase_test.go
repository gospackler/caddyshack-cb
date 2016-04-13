package couchbase

import (
	"log"
	"os"
	"testing"
)

var store *CouchbaseStore

type TestData struct {
	Data int
}

// Test setup and tear down to avoid repetition.
// It also deletes the document that was shared across functions for reading.
func TestMain(m *testing.M) {
	store = NewCouchbaseStore("couchbase://localhost", "default")
	if err := store.ConnectBucket(); err != nil {
		log.Fatalf("could not connect to host or bucket: %s\n", err)
	}
	exitCode := m.Run()
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
	t.Skip("TBD")
	data := &TestData{Data: 1}
	testObj := NewCouchbaseObject("nickel")
	testObj.SetData(data)

	err := store.Create(testObj)
	checkFatal(t, "could not insert a TestData object", err)

	query := NewNickelQuery("SELECT * FROM default WHERE Data = 1", store.bucket)
	err, results := store.Read(query)
	t.Log(results)

	err = store.DestroyOne(testObj.key)
	checkFatal(t, "could not destroy TestData object", err)
}
