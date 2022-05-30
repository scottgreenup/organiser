package term

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStandardLogger_Print(t *testing.T) {
	l := NewStandardLogger()

	expectedMessage := "I am a basic print\n"
	n, err := l.Print(expectedMessage)
	require.NoError(t, err)
	require.Equal(t, len(expectedMessage), n)
}

func TestNewStandardLogger_Printf(t *testing.T) {
	l := NewStandardLogger()

	expectedMessage := "I am a basic printf\n"
	n, err := l.Printf("I am a basic %s\n", "printf")
	require.NoError(t, err)
	require.Equal(t, len(expectedMessage), n)
}
