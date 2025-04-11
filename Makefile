# Makefile for Azure Linux OpenTelemetry Collector

.PHONY: azl-otelcol

azl-otelcol:
	mkdir -p bin
	go build -o bin/azl-otelcol $(MODFLAGS) -ldflags="$(LDFLAGS)" -trimpath -tags "$(BUILDTAGS)" .

clean:
	rm -f bin/azl-otelcol
