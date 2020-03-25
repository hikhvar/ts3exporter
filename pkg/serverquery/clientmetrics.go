package serverquery

type ClientMetrics struct {
	successful int
	failed     int
}

func (c *ClientMetrics) CountFailure() {
	c.failed = c.failed + 1
}

func (c *ClientMetrics) CountSuccess() {
	c.successful = c.successful + 1
}

func (c *ClientMetrics) Success() int {
	return c.successful
}

func (c *ClientMetrics) Failed() int {
	return c.failed
}
