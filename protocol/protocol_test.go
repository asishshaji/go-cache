package protocol

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSetCommand(t *testing.T) {

	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}

	r := bytes.NewReader(cmd.Bytes())
	pCmd, _ := ParseCommand(r)

	assert.Equal(t, cmd, pCmd)

}

func TestParseGetCommand(t *testing.T) {

	cmd := &CommandGet{
		Key: []byte("Foo"),
	}

	r := bytes.NewReader(cmd.Bytes())
	pCmd, _ := ParseCommand(r)

	assert.Equal(t, cmd, pCmd)

}

func BenchmarkParseCommand(b *testing.B) {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}

	for i := 0; i < b.N; i++ {
		bufReader := bytes.NewBuffer(cmd.Bytes())
		ParseCommand(bufReader)
	}
}
