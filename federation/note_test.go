package federation

import (
	"testing"

	"github.com/fuck-capitalism/unite/common"
)

func TestNewNote(t *testing.T) {
	note := NewNote()
	common.Log(note.JSONLDContext())
}
