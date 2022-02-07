package actor_test

import (
	"fmt"
	"testing"

	"github.com/rviscarra/go-actor/actor"
	"github.com/stretchr/testify/assert"
)

type radderAdd struct {
	value int
}

type radderGet struct {
	reply chan int
}

func TestReducerActor(t *testing.T) {

	adder := actor.NewFromReducer(0, func(raw interface{}, sum int) int {
		switch msg := raw.(type) {
		case radderAdd:
			return msg.value + sum
		case radderGet:
			msg.reply <- sum
			return sum
		default:
			panic(fmt.Errorf("unsupported message %T", raw))
		}
	})
	adder.Send(radderAdd{value: 1})
	adder.Send(radderAdd{value: 2})
	adder.Send(radderAdd{value: 3})
	get := radderGet{reply: make(chan int, 1)}
	adder.Send(get)

	actual := <-get.reply
	assert.Equal(t, 6, actual)
}
