package adt

import "adtExample/lib"

type ofUnion interface {
	lib.TypeA | lib.TypeB
}

type union interface {
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

type Atd struct {
	v union
}

// Of

func Of[T ofUnion](union T) *Atd {
	var t interface{}
	t = union
	switch v := t.(type) {
	case lib.TypeA:
		return &Atd{TypeAContainer{v}}
	case lib.TypeB:
		return &Atd{TypeBContainer{v}}
	default:
		panic("called Of with disallowed type")
	}
}

// Do

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

func (e doChainExecutor) Exec(of *Atd) {
	switch container := of.v.(type) {
	case TypeAContainer:
		e.fns.doFnTypeA(container.v)
	case TypeBContainer:
		e.fns.doFnTypeB(container.v)
	}
}

func Do() *doChainA {
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

func (mapper mapChainMapper[T]) Map(of *Atd) T {
	switch container := of.v.(type) {
	case TypeAContainer:
		return mapper.fns.mapFnTypeA(container.v)
	case TypeBContainer:
		return mapper.fns.mapFnTypeB(container.v)
	default:
		panic("called Of with disallowed type")
	}
}

func Mapper[T any]() *mapChainA[T] {
	return &mapChainA[T]{mapFns[T]{}}
}
