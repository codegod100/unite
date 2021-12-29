package federation

import (
	"context"
	"net/url"

	"github.com/go-fed/activity/streams/vocab"
)

type Database struct {
}

func init() {
	// s2sActor := pub.NewFederatingActor(cb, fp, db, clock)
}

func (Database) Lock(c context.Context, id *url.URL) error {
	return nil
}

func (Database) Unlock(c context.Context, id *url.URL) error {
	return nil
}

func (Database) Get(c context.Context, id *url.URL) (value vocab.Type, err error) {
	return nil, nil
}

func (Database) Create(c context.Context, asType vocab.Type) error {
	return nil
}

func (Database) Update(c context.Context, asType vocab.Type) error {
	return nil
}
