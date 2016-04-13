package couchbase

import "github.com/couchbase/gocb"

// Nickel is a pseudonym for N1ql which is Couchbase's query language.
type NickelQuery struct {
	query   string
	cluster *gocb.Cluster
}

func NewNickelQuery(query string, cluster *gocb.Cluster) *NickelQuery {
	return &NickelQuery{
		query:   query,
		cluster: cluster,
	}
}

func (n *NickelQuery) Execute() (error, []*CouchbaseObject) {
	query := gocb.NewN1qlQuery(n.query)
	rows, err := n.cluster.ExecuteN1qlQuery(query, nil)
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
