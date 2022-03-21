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

func TestMapWithTwoTypes(t *testing.T) {
	mapper := adt.Mapper[string]().
		WithTypeA(func(typeA lib.TypeA) string { return "type-a" }).
		WithTypeB(func(typeB lib.TypeB) string { return "type-b" })

	typeAMapped := mapper.Map(adt.Of(lib.TypeA{}))
	typeBMapped := mapper.Map(adt.Of(lib.TypeB{}))

	require.Equal(t, "type-a", typeAMapped)
	require.Equal(t, "type-b", typeBMapped)
}
