package deployer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_get_one_element(t *testing.T) {
	w := New(".")
	itens := w.Itens()

	assert.Equal(t, 1, len(itens))
	assert.Contains(t, itens[0].Path, "example.bpmn")
}

func Test_should_get_zero_when_no_file_found(t *testing.T) {
	w := New("")

	assert.Equal(t, 0, len(w.Itens()))
}
