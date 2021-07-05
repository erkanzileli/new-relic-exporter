# new-relic-exporter

Collects application metrics from [New Relic API](https://rpm.newrelic.com/api/explore/applications) and publishes the 
collected data from `/metrics` endpoint on [Prometheus's](https://prometheus.io/) 
[GaugeVector](https://prometheus.io/docs/concepts/metric_types/#gauge) format.

Sample exported data is below. All `WebTransaction|Errors|Datastore|Memory` metrics and its values will be added like this format.
`average_call_time` is an example to show you. You can filter data with `application_name` or `metric_name`.
As you can guess, all the metric value names(`requests_per_minute`, `average_response_time`, etc.) will be added 
as a vector to Prometheus metrics like `average_call_time`.

```
average_call_time{application_name="sample-application",metric_name="Datastore/Postgres/all"} 3.04
average_call_time{application_name="sample-application",metric_name="Datastore/all"} 3.04
average_call_time{application_name="sample-application",metric_name="Datastore/getConnection"} 0.906
average_call_time{application_name="sample-application",metric_name="Datastore/operation/Postgres/insert"} 2.27
average_call_time{application_name="sample-application",metric_name="Datastore/operation/Postgres/select"} 3.56
average_call_time{application_name="sample-application",metric_name="Datastore/statement/Postgres/addresses/insert"} 1.71
average_call_time{application_name="sample-application",metric_name="Errors/WebTransaction/SpringController/addresses/{addressId} (GET)"} 12.1
average_call_time{application_name="sample-application",metric_name="Memory/Heap/Committed"} 434000
average_call_time{application_name="sample-application",metric_name="Memory/Heap/Max"} 1.54e+06
average_call_time{application_name="sample-application",metric_name="Memory/Heap/Used"} 353000
average_call_time{application_name="sample-application",metric_name="Memory/NonHeap/Max"} 0
average_call_time{application_name="sample-application",metric_name="Memory/Physical"} 842000
average_call_time{application_name="sample-application",metric_name="Memory/Used"} 504000
average_call_time{application_name="sample-application",metric_name="WebTransaction/SpringController/addresses (POST)"} 56.2
average_call_time{application_name="sample-application",metric_name="WebTransaction/SpringController/addresses (GET)"} 32
```

## Usage

Run command below with these minimum flags

````shell script
 go run . --app-id=123456789 --api-key=abcdefgh
````

### Available Flags


| flag        | Description           | Default  |
| ------------- |:-------------:| -----:|
| --addr      | metric server address | :8080 |
| --app-id      | your new-relic application id | None |
| --api-key      | your new-relic api key      |  None  |
| --interval | scrape interval      |    @every 0h0m10s |
| --personal-key | specify if your api key is personal      |    false |

