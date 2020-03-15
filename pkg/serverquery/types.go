package serverquery

// An Executor executes the given serverQuery command and return the result
type Executor interface {
	Exec(cmd string) ([]Result, error)
}
