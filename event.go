package gosync

// Event is asynchronous event.
// Event is a little faster than sync.WaitGroup
type Event chan struct{}

// NewEvent makes asynchronous event
func NewEvent() Event {
	return make(Event)
}

// Done triggers the event.
// It can be called only once.
func (ev Event) Done() {
	close(ev)
}

// Wait waits when the event happens.
func (ev Event) Wait() {
	<-ev
}
