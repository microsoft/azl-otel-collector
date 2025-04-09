package smartdata

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("smartdata")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval.String(),
	}
}

func createTracesReceiver(_ context.Context, params receiver.Settings, baseCfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	logger := params.Logger
	smartdataConfig := baseCfg.(*Config)

	smartdataReceiver := &smartdataReceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       smartdataConfig,
	}

	return smartdataReceiver, nil
}

// NewFactory creates a factory for the smartdata receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, component.StabilityLevelAlpha))
}
