package actor

// Message needs to be implemented by actor messages
// Apply should make use of the message's fields and the current state to
//  produce a new one
type Message[St any] interface {
	Apply(St) St
}

func NewTyped[St any](initial St) Actor[Message[St]] {
	return NewFromReducer(initial, func(msg Message[St], state St) St {
		return msg.Apply(state)
	})
}
