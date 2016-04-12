package couchbase

import "errors"

// Nickel is a pseudonym for N1ql which is Couchbase's query language.
type NickelQuery struct {
	query string
}

func NewNickelQuery(query string) *NickelQuery {
	return &NickelQuery{
		query: query,
	}
}

func (n *NickelQuery) Execute() (error, []*CouchbaseObject) {
	return errors.New("not implemented"), nil
}
