package gobber

// WorkDataStore defines the connection of a data store that stores information on the work to be done.
// For example, if we define a queue as this data store, we will need to put and retrieve data from here to inform
// workers on what to work on.
type WorkDataStore interface {
	PutData(data any)
	RetrieveData() (any, error)
}
