// Code generated by cdpgen. DO NOT EDIT.

package tethering

import (
	"github.com/mafredri/cdp/rpcc"
)

// AcceptedClient is a client for Accepted events. Informs that port was
// successfully bound and got a specified connection id.
type AcceptedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*AcceptedReply, error)
	rpcc.Stream
}

// AcceptedReply is the reply for Accepted events.
type AcceptedReply struct {
	Port         int    `json:"port"`         // Port number that was successfully bound.
	ConnectionID string `json:"connectionId"` // Connection id to be used.
}