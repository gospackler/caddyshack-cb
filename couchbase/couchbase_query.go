package couchbase

import "github.com/couchbase/gocb"

// Nickel is a pseudonym for N1ql which is Couchbase's query language.
type NickelQuery struct {
	query  string
	bucket *gocb.Bucket
}

func NewNickelQuery(query string, bucket *gocb.Bucket) *NickelQuery {
	return &NickelQuery{
		query:  query,
		bucket: bucket,
	}
}

func (n *NickelQuery) Execute() (error, []*CouchbaseObject) {
	query := gocb.NewN1qlQuery(n.query)
	rows, err := n.bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		return err, nil
	}

	var document interface{}
	var documents []*CouchbaseObject
	for rows.Next(&document) {
		doc := &CouchbaseObject{
			data: document,
		}
		documents = append(documents, doc)
	}

	err = rows.Close()
	if err != nil {
		return err, nil
	}

	return nil, documents
}
