package couchbase

import "github.com/couchbase/gocb"

// The data stored in this object must be serializable to work with the client.
// The key is a unique identifier for the document you wish to store.
type CouchbaseObject struct {
	key      string
	data     interface{}
	objectId gocb.Cas
	expiry   uint32
}

func NewCouchbaseObject(key string) *CouchbaseObject {
	return &CouchbaseObject{
		key: key,
	}
}

func (c *CouchbaseObject) GetKey() string {
	return c.key
}

func (c *CouchbaseObject) SetKey(key string) {
	c.key = key
}

func (c *CouchbaseObject) SetData(data interface{}) {
	c.data = data
}

func (c *CouchbaseObject) SetId(id gocb.Cas) {
	c.objectId = id
}

func (c *CouchbaseObject) SetExpiry(time uint32) {
	c.expiry = time
}
