# TS3 Exporter
![](https://github.com/hikhvar/ts3exporter/workflows/tests/badge.svg) ![](https://github.com/hikhvar/ts3exporter/workflows/release/badge.svg)

This exporter uses the server query protocol to provide prometheus metrics.

## Build
The build is tested with go version 1.14. Since the build uses new standard error formatting, it requires at least 1.13.
```bash
go build
```

## Usage
``` 
# ./ts3exporter -h
Usage of ./ts3exporter:
  -listen string
    	listen address of the exporter (default ":9189")
  -password string
    	the serverquery password of the ts3exporter
  -remote string
    	remote address of server query port (default "localhost:10011")
  -user string
    	the serverquery user of the ts3exporter (default "serveradmin")
```

## Examples:
```bash
# curl localhost:9189/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.14"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 781560
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 781560
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 3635
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 292
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.436808e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 781560
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.5536e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.277952e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 3785
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 6.5536e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6813952e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 4077
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 1736
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 19584
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 526541
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 294912
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 294912
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.1125e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 6
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.01
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1024
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 9.498624e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.58427726029e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.037832192e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP ts3_exporter_data_model_refresh_errors_total Errors encountered while updating the internal server model
# TYPE ts3_exporter_data_model_refresh_errors_total counter
ts3_exporter_data_model_refresh_errors_total 0
# HELP ts3_serverinfo_bytes_received_total total received bytes
# TYPE ts3_serverinfo_bytes_received_total counter
ts3_serverinfo_bytes_received_total{virtualserver="Gute Stube"} 23096
# HELP ts3_serverinfo_bytes_send_total total send bytes
# TYPE ts3_serverinfo_bytes_send_total counter
ts3_serverinfo_bytes_send_total{virtualserver="Gute Stube"} 23236
# HELP ts3_serverinfo_channels_online number of online channels
# TYPE ts3_serverinfo_channels_online gauge
ts3_serverinfo_channels_online{virtualserver="Gute Stube"} 6
# HELP ts3_serverinfo_client_connections currently established client connections
# TYPE ts3_serverinfo_client_connections gauge
ts3_serverinfo_client_connections{virtualserver="Gute Stube"} 1
# HELP ts3_serverinfo_clients_online number of currently online clients
# TYPE ts3_serverinfo_clients_online gauge
ts3_serverinfo_clients_online{virtualserver="Gute Stube"} 1
# HELP ts3_serverinfo_control_bytes_received_total total received bytes for control traffic
# TYPE ts3_serverinfo_control_bytes_received_total counter
ts3_serverinfo_control_bytes_received_total{virtualserver="Gute Stube"} 4182
# HELP ts3_serverinfo_control_bytes_sent_total total sent bytes for control traffic
# TYPE ts3_serverinfo_control_bytes_sent_total counter
ts3_serverinfo_control_bytes_sent_total{virtualserver="Gute Stube"} 4182
# HELP ts3_serverinfo_file_transfer_bytes_received_total total received bytes for file transfers
# TYPE ts3_serverinfo_file_transfer_bytes_received_total counter
ts3_serverinfo_file_transfer_bytes_received_total{virtualserver="Gute Stube"} 0
# HELP ts3_serverinfo_file_transfer_bytes_sent_total total sent bytes for file transfers
# TYPE ts3_serverinfo_file_transfer_bytes_sent_total counter
ts3_serverinfo_file_transfer_bytes_sent_total{virtualserver="Gute Stube"} 0
# HELP ts3_serverinfo_keepalive_bytes_received_total total received bytes for keepalive traffic
# TYPE ts3_serverinfo_keepalive_bytes_received_total counter
ts3_serverinfo_keepalive_bytes_received_total{virtualserver="Gute Stube"} 15413
# HELP ts3_serverinfo_keepalive_bytes_sent_total total send bytes for keepalive traffic
# TYPE ts3_serverinfo_keepalive_bytes_sent_total counter
ts3_serverinfo_keepalive_bytes_sent_total{virtualserver="Gute Stube"} 15047
# HELP ts3_serverinfo_max_clients maximal number of allowed clients
# TYPE ts3_serverinfo_max_clients gauge
ts3_serverinfo_max_clients{virtualserver="Gute Stube"} 32
# HELP ts3_serverinfo_max_download_bandwidth maximal bandwidth available for downloads
# TYPE ts3_serverinfo_max_download_bandwidth gauge
ts3_serverinfo_max_download_bandwidth{virtualserver="Gute Stube"} 1.8446744073709552e+19
# HELP ts3_serverinfo_max_upload_bandwidth maximal bandwidth available for uploads
# TYPE ts3_serverinfo_max_upload_bandwidth gauge
ts3_serverinfo_max_upload_bandwidth{virtualserver="Gute Stube"} 1.8446744073709552e+19
# HELP ts3_serverinfo_online is the virtualserver online
# TYPE ts3_serverinfo_online gauge
ts3_serverinfo_online{virtualserver="Gute Stube"} 1
# HELP ts3_serverinfo_query_client_connections currently established query client connections
# TYPE ts3_serverinfo_query_client_connections gauge
ts3_serverinfo_query_client_connections{virtualserver="Gute Stube"} 1
# HELP ts3_serverinfo_speech_bytes_received_total total received bytes for speech traffic
# TYPE ts3_serverinfo_speech_bytes_received_total counter
ts3_serverinfo_speech_bytes_received_total{virtualserver="Gute Stube"} 3501
# HELP ts3_serverinfo_speech_bytes_sent_total total sent bytes for speech traffic
# TYPE ts3_serverinfo_speech_bytes_sent_total counter
ts3_serverinfo_speech_bytes_sent_total{virtualserver="Gute Stube"} 0
# HELP ts3_serverinfo_uptime uptime of the virtual server
# TYPE ts3_serverinfo_uptime counter
ts3_serverinfo_uptime{virtualserver="Gute Stube"} 88180
```