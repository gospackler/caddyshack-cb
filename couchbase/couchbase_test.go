package couchbase

import (
	"log"
	"os"
	"testing"
)

var (
	store   *CouchbaseStore
	testObj *CouchbaseObject
)

// Test setup and tear down to avoid repetition.
func TestMain(m *testing.M) {
	store = NewCouchbaseStore("couchbase://localhost", "default")
	if err := store.ConnectBucket(); err != nil {
		log.Fatalf("could not connect to host or bucket: %s\n", err)
	}
	testObj = NewCouchbaseObject("test")
	testObj.SetData("data")
	exitCode := m.Run()
	store.ShutdownBucket()
	os.Exit(exitCode)
}

func TestCreateDocument(t *testing.T) {
	err := store.Create(testObj)
	if err != nil {
		t.Fatalf("could not create an object: %s\n", err)
	}
}

func TestDeleteDocument(t *testing.T) {
	err := store.DestroyOne(testObj.key)
	if err != nil {
		t.Fatalf("could not destroy object: %s\n", err)
	}
}
