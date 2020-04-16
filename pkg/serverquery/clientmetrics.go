package serverquery

import "sync"

type ClientMetrics struct {
	sync.RWMutex
	successful int
	failed     int
}

func (c *ClientMetrics) CountFailure() {
	c.Lock()
	defer c.Unlock()
	c.failed = c.failed + 1
}

func (c *ClientMetrics) CountSuccess() {
	c.Lock()
	defer c.Unlock()
	c.successful = c.successful + 1
}

func (c *ClientMetrics) Success() int {
	c.RLock()
	defer c.RUnlock()
	return c.successful
}

func (c *ClientMetrics) Failed() int {
	c.RLock()
	defer c.RUnlock()
	return c.failed
}
