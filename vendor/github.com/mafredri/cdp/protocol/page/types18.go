// +build !go1.9

// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol"
	"github.com/mafredri/cdp/protocol/network"
)

// ResourceType Resource type as it was perceived by the rendering engine.
//type ResourceType string

// ResourceType as enums.
const (
	ResourceTypeNotSet             protocol.PageResourceType = ""
	ResourceTypeDocument           protocol.PageResourceType = "Document"
	ResourceTypeStylesheet         protocol.PageResourceType = "Stylesheet"
	ResourceTypeImage              protocol.PageResourceType = "Image"
	ResourceTypeMedia              protocol.PageResourceType = "Media"
	ResourceTypeFont               protocol.PageResourceType = "Font"
	ResourceTypeScript             protocol.PageResourceType = "Script"
	ResourceTypeTextTrack          protocol.PageResourceType = "TextTrack"
	ResourceTypeXHR                protocol.PageResourceType = "XHR"
	ResourceTypeFetch              protocol.PageResourceType = "Fetch"
	ResourceTypeEventSource        protocol.PageResourceType = "EventSource"
	ResourceTypeWebSocket          protocol.PageResourceType = "WebSocket"
	ResourceTypeManifest           protocol.PageResourceType = "Manifest"
	ResourceTypeSignedExchange     protocol.PageResourceType = "SignedExchange"
	ResourceTypePing               protocol.PageResourceType = "Ping"
	ResourceTypeCSPViolationReport protocol.PageResourceType = "CSPViolationReport"
	ResourceTypeOther              protocol.PageResourceType = "Other"
)

// FrameID Unique frame identifier.
//type FrameID string

// Frame Information about the Frame on the page.
type Frame struct {
	ID             protocol.PageFrameID  `json:"id"`                 // Frame unique identifier.
	ParentID       *protocol.PageFrameID `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderID       network.LoaderID      `json:"loaderId"`           // Identifier of the loader associated with this frame.
	Name           *string               `json:"name,omitempty"`     // Frame's name as specified in the tag.
	URL            string                `json:"url"`                // Frame document's URL.
	SecurityOrigin string                `json:"securityOrigin"`     // Frame document's security origin.
	MimeType       string                `json:"mimeType"`           // Frame document's mimeType as determined by the browser.
	// UnreachableURL If the frame failed to load, this contains the URL
	// that could not be loaded.
	//
	// Note: This property is experimental.
	UnreachableURL *string `json:"unreachableUrl,omitempty"`
}

// FrameResource Information about the Resource on the page.
//
// Note: This type is experimental.
type FrameResource struct {
	URL          string                    `json:"url"`                    // Resource URL.
	Type         protocol.PageResourceType `json:"type"`                   // Type of this resource.
	MimeType     string                    `json:"mimeType"`               // Resource mimeType as determined by the browser.
	LastModified network.TimeSinceEpoch    `json:"lastModified,omitempty"` // last-modified timestamp as reported by server.
	ContentSize  *float64                  `json:"contentSize,omitempty"`  // Resource content size.
	Failed       *bool                     `json:"failed,omitempty"`       // True if the resource failed to load.
	Canceled     *bool                     `json:"canceled,omitempty"`     // True if the resource was canceled during loading.
}