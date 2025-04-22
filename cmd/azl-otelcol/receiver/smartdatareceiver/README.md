# SMART Data Receiver

The SMART data receiver is designed to collect SMART data from physical disks on a machine at a configured interval. It utilizes the `smartctl` tool from the `smartmontools` package to gather this information. This receiver is intended to be used on baremetal machines to monitor the health and status of associated disks.

## Enabling the Receiver

The `smartdata` receiver and its interval can be enabled in the azl-otel-collector configuration file under the `receivers` section. The receiver will collect SMART data from all physical disks on the machine. See the example configuration below:

```yaml
receivers:
  smartdata:
    interval: 60m
```

## SMART Data OTEL Structure
The SMART data receiver represents each instance of data collection as a `Trace` event. The host machine information is represented as attributes of a `ResourceSpan` for the trace. Each disk is represented as a `Span`, and these spans make up the trace event. The structure of a SMART data trace event from this receiver is as follows:

```yaml
- trace
  -  trace ID = uuid.new
  -   resources
      - resourceSpan
        - resourceAttributes:
            - service.name = smart-data-receiver
            - machine.info: { json-string }
      - scopeSpan
        - name: smart-data-collection
        - version: 1.0.0
        - spans
          - span
            - name: disk + "-smart-data"
            - traceid: trace ID
            - id = uuid.new
            - kind = ptrace.SpanKindInternal
            - startTime: time.Now()
            - endTime: time.Now()
            - spanAttributes:
                - smart_data: { json-string}
        
```