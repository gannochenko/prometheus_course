# Prometheus course

Linux metrics:

~~~
# Memory Usage Above 60%:
# 69 < (100 - job:node_memory_Mem_bytes:available) < 75

# CPU Usage High:
# 100 - (avg by(instance) (irate(node_cpu_seconds_total{mode="idle"}[1m])) * 100) > 80

# Free disk space less than 30%
# (sum by (instance) (node_filesystem_free_bytes) / sum by (instance) (node_filesystem_size_bytes)) * 100 < 30
~~~

## Routing tree editor

https://prometheus.io/webtools/alerting/routing-tree-editor/
