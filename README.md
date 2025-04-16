# Azure Linux OpenTelemetry Collector (AZL OTEL Collector)

The **AZL OTEL Collector** is a custom OpenTelemetry Collector distribution
built by the Azure Linux team. It is designed to provide telemetry collection
tailored for Azure Linux environments and scenarios, with a curated set of
receivers, processors, and exporters from the upstream OpenTelemetry Collector
and the OpenTelemetry Collector Contrib repository. This collector is built
using the [OpenTelemetry Collector Builder
(OCB)](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder)
tool, which allows for easy customization and updates of the collector
components.

## Included Components

**Receivers:**

- [otlpreceiver](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/otlpreceiver)
- [filelogreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/filelogreceiver)
- [hostmetricsreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver)
- [journaldreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/journaldreceiver)
- [smartdatareceiver](https://github.com/microsoft/azl-otel-collector/tree/main/receiver/smartdatareceiver)

**Processors:**
- [batchprocessor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/batchprocessor)
- [memorylimiterprocessor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/memorylimiterprocessor)
- [attributesprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/attributesprocessor)
- [filterprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/filterprocessor)
- [resourcedetectionprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourcedetectionprocessor)

**Exporters:**
- [debugexporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/debugexporter)
- [otlpexporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter)
- [fileexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/fileexporter)
- [azuredataexplorerexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/azuredataexplorerexporter)


## Building the Collector

Note: Ensure you have Go installed on your machine. The version of Go should be
at least 1.23 or higher.

1. **Build the Collector:** From the root directory of the repository, run the
   following command to build the collector:
    ```bash
    make build
    ```
    This will create a binary named `azl-otelcol` in the bin directory.
2. **Run the Collector:** You can run the collector binary with an OTEL
    configuration file. A default configuration file is provided in the
    [`config`](./config) directory, or you can create your own.
      ```bash
      ./azl-otelcol --config <path_to_otel_config_file>
      ```

## Updating the Collector
The AZL OTEL Collector is based on the OpenTelemetry Collector and uses the OTEL
Collector Builder (OCB) tool for updates. To update the collector components to
a newer upstream version, run the following command:
```bash
make update-sources OCB_VERSION=<version>
```
This will download the associated version of the OCB binary and update the
collector components to the specified version. The `OCB_VERSION` should be a
valid version tag from the OpenTelemetry Collector GitHub repository. For
example, to update to version `0.124.0`, run:
```bash
make update-sources OCB_VERSION=0.124.0
```
The OCB tool uses the builder-config.yaml file to determine which components and
versions to include in the collector. This file will be updated as part of the
process.

## Support

This project welcomes contributions and suggestions.  Most contributions require
you to agree to a Contributor License Agreement (CLA) declaring that you have
the right to, and actually do, grant us the rights to use your contribution. For
details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether
you need to provide a CLA and decorate the PR appropriately (e.g., status check,
comment). Simply follow the instructions provided by the bot. You will only need
to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of
Conduct](https://opensource.microsoft.com/codeofconduct/). For more information
see the [Code of Conduct
FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact
[opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional
questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or
services. Authorized use of Microsoft trademarks or logos is subject to and must
follow [Microsoft's Trademark & Brand
Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must
not cause confusion or imply Microsoft sponsorship. Any use of third-party
trademarks or logos are subject to those third-party's policies.
