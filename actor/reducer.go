package actor

type reducerActor[Msg, St any] struct {
	state   St
	queue   chan Msg
	reducer func(Msg, St) St
}

func NewFromReducer[Msg, St any](initial St, reducer func(Msg, St) St) Actor[Msg] {
	ga := &reducerActor[Msg, St]{
		state:   initial,
		queue:   make(chan Msg, 1),
		reducer: reducer,
	}
	ga.start()
	return ga
}

func (a *reducerActor[_, _]) start() {
	go a.receiveLoop()
}

func (a *reducerActor[_, _]) receiveLoop() {
	for msg := range a.queue {
		a.state = a.reducer(msg, a.state)
	}
}

func (a *reducerActor[Msg, _]) Send(msg Msg) {
	a.queue <- msg
}
