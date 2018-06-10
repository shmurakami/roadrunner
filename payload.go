package roadrunner

// Payload carries binary header and body to workers and
// back to the server.
type Payload struct {
	// Context represent payload context, might be omitted.
	Context []byte

	// body contains binary payload to be processed by worker.
	Body []byte

	// add io.Reader support for streamed requests and responses.
}

// String returns payload body as string
func (p *Payload) String() string {
	return string(p.Body)
}
