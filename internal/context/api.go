package context

import "github.com/jwhittle933/streamline/media/mpeg/box"

type Contexter interface {
	Act()
	Respond(action string, handler func(boxed box.Boxed))
}

// Context is intended to be passed down through the initial read process -
// when a file is consumed - and used for upward and downward communication.
// An action can be triggered from anywhere in the structure tree and any
// structure that holds a reference to the Context can respond to the action
//
// Only a single Context should be created. Each structure that needs
// it should keep the reference to it.
type Context struct {
	Command    chan string
	CommandOut chan string
	Status     chan string
}
