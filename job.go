package gobber

// Job is the interface that defines getting a job ready to work.
type Job interface {
	// Payload provides the payload to put the job into a work pipeline.
	Payload() []byte
}
