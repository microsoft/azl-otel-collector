# Makefile for Azure Linux OpenTelemetry Collector

OCB_VERSION ?= 0.124.0
OCB_BIN ?= ./bin/ocb

.PHONY: azl-otelcol run-ocb

azl-otelcol:
	mkdir -p bin
	go build -o bin/azl-otelcol $(MODFLAGS) -ldflags="$(LDFLAGS)" -trimpath -tags "$(BUILDTAGS)" .

run-ocb:
	@echo "Updating OTEL component versions in builder-config.yaml to $(OCB_VERSION)..."
	@sed -i 's/0\.[0-9]\+\.[0-9]\+/$(OCB_VERSION)/g' builder-config.yaml
	@echo "Downloading OCB version $(OCB_VERSION)..."
	@mkdir -p bin
	@curl --proto '=https' --tlsv1.2 -fL -o $(OCB_BIN) \
		https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv$(OCB_VERSION)/ocb_$(OCB_VERSION)_linux_amd64
	@chmod +x $(OCB_BIN)
	@echo "Running OCB..."
	@$(OCB_BIN) --config builder-config.yaml --skip-compilation
	@rm -f main_windows.go

clean:
	rm -f bin/azl-otelcol
