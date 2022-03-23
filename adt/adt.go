package adt

import "adtExample/lib"

type Union interface {
	lib.TypeA | lib.TypeB
}

type unionContainer interface {
	atd()
}

type TypeAContainer struct {
	v lib.TypeA
}

func (t TypeAContainer) atd() {}

type TypeBContainer struct {
	v lib.TypeB
}

func (t TypeBContainer) atd() {}

type TaggedUnion struct {
	v unionContainer
}

// UnionOf

func UnionOf[T Union](union T) TaggedUnion {
	var t interface{}
	t = union
	switch v := t.(type) {
	case lib.TypeA:
		return TaggedUnion{TypeAContainer{v}}
	case lib.TypeB:
		return TaggedUnion{TypeBContainer{v}}
	default:
		panic("called UnionOf with disallowed type")
	}
}

// UnionExecutor

type doFnTypeAFn func(aType lib.TypeA)
type doFnTypeBFn func(aType lib.TypeB)

type doFns struct {
	doFnTypeA doFnTypeAFn
	doFnTypeB doFnTypeBFn
}

type doChainA struct {
	doBuilder doFns
}

type doChainB struct {
	fns doFns
}

type doChainExecutor struct {
	fns doFns
}

func (chain *doChainB) WithTypeB(fn doFnTypeBFn) *doChainExecutor {
	chain.fns.doFnTypeB = fn
	return &doChainExecutor{chain.fns}
}

func (chain *doChainA) WithTypeA(fn doFnTypeAFn) *doChainB {
	chain.doBuilder.doFnTypeA = fn
	return &doChainB{chain.doBuilder}
}

func (e doChainExecutor) Exec(of TaggedUnion) {
	switch container := of.v.(type) {
	case TypeAContainer:
		e.fns.doFnTypeA(container.v)
	case TypeBContainer:
		e.fns.doFnTypeB(container.v)
	}
}

func UnionExecutor() *doChainA {
	return &doChainA{doFns{}}
}

// Map

type mapFnTypeAFn[T any] func(aType lib.TypeA) T
type mapFnTypeBFn[T any] func(aType lib.TypeB) T

type mapFns[T any] struct {
	mapFnTypeA mapFnTypeAFn[T]
	mapFnTypeB mapFnTypeBFn[T]
}

type mapChainA[T any] struct {
	fns mapFns[T]
}

type mapChainB[T any] struct {
	fns mapFns[T]
}

type mapChainMapper[T any] struct {
	fns mapFns[T]
}

func (chain mapChainA[T]) WithTypeA(fn mapFnTypeAFn[T]) *mapChainB[T] {
	chain.fns.mapFnTypeA = fn
	return &mapChainB[T]{chain.fns}
}

func (chain mapChainB[T]) WithTypeB(fn mapFnTypeBFn[T]) *mapChainMapper[T] {
	chain.fns.mapFnTypeB = fn
	return &mapChainMapper[T]{chain.fns}
}

func (mapper mapChainMapper[T]) Map(of TaggedUnion) T {
	switch container := of.v.(type) {
	case TypeAContainer:
		return mapper.fns.mapFnTypeA(container.v)
	case TypeBContainer:
		return mapper.fns.mapFnTypeB(container.v)
	default:
		panic("called UnionOf with disallowed type")
	}
}

func UnionMapper[T any]() *mapChainA[T] {
	return &mapChainA[T]{mapFns[T]{}}
}
