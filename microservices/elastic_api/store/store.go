package store

import (
	"strings"
	"github.com/elastic/go-elasticsearch/v8"
)

type Store struct {
	db           *elasticsearch.Client
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
		  "http://localhost:9200",
		},
	  }
	  
	 db, err := elasticsearch.NewClient(cfg)
	 if err != nil{
		return err
	 }

	 s.db = db
	 return nil
}

func (s *Store) Get(phrase string) (string, error) {
	res, err := s.db.Search(
		s.db.Search.WithBody(strings.NewReader(`{
		  "query": {
			"query_string": {
				"query": "title:\"Sience\""
			}
		  }
		}`)),
		s.db.Search.WithPretty(),
	)
	if err != nil {
		return err.Error(), err
	}

	return res.String(), nil
}