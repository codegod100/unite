package federation

import (
	"net/url"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func NewNote() vocab.ActivityStreamsNote {
	note := streams.NewActivityStreamsNote()
	id, _ := url.Parse("https://example.com/some/path/to/this/note")
	idProperty := streams.NewJSONLDIdProperty()
	idProperty.Set(id)

	// Set the `id` property on our Note.
	note.SetJSONLDId(idProperty)
	return note
}
