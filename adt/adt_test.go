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

	do := adt.UnionExecutor().
		WithTypeA(func(typeA lib.TypeA) { typeACalled = true }).
		WithTypeB(func(typeB lib.TypeB) { typeBCalled = true })

	do.Exec(adt.UnionOf(lib.TypeA{}))

	require.Equal(t, typeACalled, true)
	require.Equal(t, typeBCalled, false)

	typeACalled = false
	typeBCalled = false

	do.Exec(adt.UnionOf(lib.TypeB{}))

	require.Equal(t, typeACalled, false)
	require.Equal(t, typeBCalled, true)
}

func TestMapWithTwoTypes(t *testing.T) {
	stringMapper := adt.UnionMapper[string]().
		WithTypeA(func(typeA lib.TypeA) string { return "type-a" }).
		WithTypeB(func(typeB lib.TypeB) string { return "type-b" })

	typeAMappedString := stringMapper.Map(adt.UnionOf(lib.TypeA{}))
	typeBMappedString := stringMapper.Map(adt.UnionOf(lib.TypeB{}))

	require.Equal(t, "type-a", typeAMappedString)
	require.Equal(t, "type-b", typeBMappedString)

	intMapper := adt.UnionMapper[int]().
		WithTypeA(func(typeA lib.TypeA) int { return 1000 }).
		WithTypeB(func(typeB lib.TypeB) int { return 2000 })

	typeAMappedInt := intMapper.Map(adt.UnionOf(lib.TypeA{}))
	typeBMappedInt := intMapper.Map(adt.UnionOf(lib.TypeB{}))

	require.Equal(t, 1000, typeAMappedInt)
	require.Equal(t, 2000, typeBMappedInt)
}
