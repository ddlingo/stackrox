// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	bleve "github.com/blevesearch/bleve"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

type Indexer interface {
	AddImageIntegration(imageintegration *storage.ImageIntegration) error
	AddImageIntegrations(imageintegrations []*storage.ImageIntegration) error
	Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error)
	DeleteImageIntegration(id string) error
	DeleteImageIntegrations(ids []string) error
	MarkInitialIndexingComplete() error
	NeedsInitialIndexing() (bool, error)
	Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	return &indexerImpl{index: index}
}