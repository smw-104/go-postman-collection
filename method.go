package postman

// Method HTTP Method Type
type Method string

const (
	// Get HTTP method.
	Get Method = "GET"
	// Put HTTP method.
	Put = "PUT"
	// Post HTTP method.
	Post = "POST"
	// Patch HTTP method.
	Patch = "PATCH"
	// Delete HTTP method.
	Delete = "DELETE"
	// Copy HTTP method.
	Copy = "COPY"
	// Head HTTP method.
	Head = "HEAD"
	// Options HTTP method.
	Options = "OPTIONS"
	// Link HTTP method.
	Link = "LINK"
	// Unlink HTTP method.
	Unlink = "UNLINK"
	// Purge HTTP method.
	Purge = "PURGE"
	// Lock HTTP method.
	Lock = "LOCK"
	// Unlock HTTP method.
	Unlock = "UNLOCK"
	// Propfind HTTP method.
	Propfind = "PROPFIND"
	// View HTTP method.
	View = "VIEW"
)
