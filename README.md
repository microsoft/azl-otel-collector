# Azure Linux OpenTelemetry Collector (AZL OTEL Collector)

The **AZL OTEL Collector** is a custom OpenTelemetry Collector distribution maintained by the Azure Linux team. It is designed to provide observability and telemetry collection tailored for Azure Linux environments, with a curated set of receivers, processors, and exporters optimized for Azure workloads.

## Components List

**Receivers:**
- otlpreceiver
- filelogreceiver
- hostmetricsreceiver
- journaldreceiver
- smartdatareceiver

**Processors:**
- batchprocessor
- memorylimiterprocessor
- attributesprocessor
- filterprocessor
- resourcedetectionprocessor

**Exporters:**
- debugexporter
- otlpexporter
- fileexporter
- azuredataexplorerexporter

**Providers:**
- envprovider
- fileprovider
- httpprovider
- httpsprovider
- yamlprovider


## Building the Collector

1. **Build the Collector:**  
   Use the Makefile to build the collector and manage dependencies.
    ```bash
    make azl-otelcol
    ```
    This will create a binary named `azl-otelcol` in the bin directory.
2. **Run the Collector:**
    You can run the collector binary with an OTEL configuration file. A default configuration file is provided in the [`config`](./config) directory, or you can create your own.
      ```bash
      ./azl-otelcol --config <path_to_otel_config_file>
      ```

## Updating the Collector
The AZL OTEL Collector is based on the OpenTelemetry Collector and uses the OTEL Collector Builder (OCB) tool for updates. To update the collector components to a newer upstream version, run the following command:
```bash
make run-ocb OCB_VERSION=<version>
```
This will download the OCB binary and update the collector components to the specified version. The `OCB_VERSION` should be a valid version tag from the OpenTelemetry Collector GitHub repository.
For example, to update to version `0.124.0`, run:
```bash
make run-ocb OCB_VERSION=0.124.0
```
The OCB tool uses the builder-config.yaml file to determine which components and versions to include in the collector. This file gets updated as part of the update command.

## Support

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
