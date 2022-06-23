# Operational Improvements

## Resilience
Resilience could be improved by increasing the number of replicas in a deployment.  Configuring the deployment to also have a `priorityClass` would be helpful to prioritize pod scheduling within the cluster. The `stock-ticker` pods can also be assigned the `Guaranteed` QoS class by configuring equal values for `cpu` limits and requests, as well as `memory` limits and requests.

## Reliability
Reliability could be improved by introducing an externally-shared cache.  As the number of replicas grows, the individual pod cache may not be sufficient to avoid rate-limiting by the Alpha Vantage API.  Robust readiness and liveness probes can be used to automatically ensure `stock-ticker` stays in an available state.

## Monitoring
Monitoring could be improved by introducing an instrumentation package, and exposing a metrics endpoint.  Prometheus provides an [open-source instrumentation library](https://github.com/prometheus/client_golang) to implement HTTP metrics.  Afterwards, metrics can be exported to the multitude of observability platforms to take useful actions.

## Scalability
Scalability could be improved by deploying a metrics server to collect container resource metrics.  The metrics server can then expose resource metrics for horizonal and/or vertical pod autoscaling.  Given the stateless nature and fast deployment time of this simple application, both forms of autoscaling are viable.  Deployed within a cloud provider, things like `cluster-autoscaler` can be used to leverage native compute scalability.
