package adt_test

import (
	"adtExample/adt"
	"adtExample/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecutingWithTwoTypes(t *testing.T) {
	typeACalled := false
	typeBCalled := false

	do := adt.Do().
		WithTypeA(func(typeA lib.TypeA) { typeACalled = true }).
		WithTypeB(func(typeB lib.TypeB) { typeBCalled = true })

	do.Exec(adt.Of(lib.TypeA{}))

	require.Equal(t, typeACalled, true)
	require.Equal(t, typeBCalled, false)

	typeACalled = false
	typeBCalled = false

	do.Exec(adt.Of(lib.TypeB{}))

	require.Equal(t, typeACalled, false)
	require.Equal(t, typeBCalled, true)
}
