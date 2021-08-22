package collector

import "fmt"

const namespace = "ts3"
const virtualServerLabel = "virtualserver"
const channelLabel = "channel"
const channelId = "channelId"

// fqdn generates a full qualified name of a metric. Given the subsystem and the name of the metric.
func fqdn(subsystem, name string) string {
	return fmt.Sprintf("%s_%s_%s", namespace, subsystem, name)
}
