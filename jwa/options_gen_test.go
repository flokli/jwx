// Code generated by tools/cmd/genoptions/main.go. DO NOT EDIT.

package jwa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithSymmetricAlgorithm", identSymmetricAlgorithm{}.String())
}