package ui

// Callback is callback function on event
type Callback func(el Element)

// mouseCallbacks keep callback functions
type mouseCallbacks map[MouseEvent]Callback

// SetMouseCallback set callback function to element. When set nil, it means delete
func (cs mouseCallbacks) SetMouseCallback(e MouseEvent, c Callback) {
	if c == nil {
		delete(cs, e)
	} else {
		cs[e] = c
	}
}
