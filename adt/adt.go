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

func Do() *doChainA {
	return &doChainA{doFns{}}
}

func (e doChainExecutor) Exec(of *Atd) {
	switch container := of.v.(type) {
	case TypeAContainer:
		e.fns.doFnTypeA(container.v)
	case TypeBContainer:
		e.fns.doFnTypeB(container.v)
	}
}

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
