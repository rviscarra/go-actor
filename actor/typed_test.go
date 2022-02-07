package actor_test

import (
	"testing"

	"github.com/rviscarra/go-actor/actor"
	"github.com/stretchr/testify/require"
)

type adderAdd struct {
	value int
}

func (msg *adderAdd) Apply(st int) int {
	return st + msg.value
}

func add(value int) actor.Message[int] {
	return &adderAdd{value: value}
}

func identity[T any](val T) T { return val }

func TestTypedActor(t *testing.T) {

	adder := actor.NewTyped(0)
	adder.Send(add(1))
	adder.Send(add(2))
	adder.Send(add(3))

	result := actor.Get(adder, identity[int])
	require.Equal(t, 6, result)
}

func TestAddMessage(t *testing.T) {
	require.Equal(t, 0, add(0).Apply(0))
	require.Equal(t, 11, add(1).Apply(10))
}
