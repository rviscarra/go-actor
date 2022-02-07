package actor

type Actor[Msg any] interface {
	Send(msg Msg)
}

type getProj[St, Ret any] struct {
	reply      chan Ret
	projection func(St) Ret
}

func (a *getProj[St, _]) Apply(state St) St {
	a.reply <- a.projection(state)
	return state
}

// GetAsync applies the projection func to the actor's state and returns the
// result asynchronously
func GetAsync[St, Ret any](a Actor[Message[St]], projection func(St) Ret) <-chan Ret {
	reply := make(chan Ret, 1)
	gp := &getProj[St, Ret]{reply: reply, projection: projection}
	a.Send(gp)
	return reply
}

// Get applies the projection func to the actor's state and returns the result
func Get[St, Ret any](a Actor[Message[St]], projection func(St) Ret) Ret {
	reply := GetAsync(a, projection)
	return <-reply
}
