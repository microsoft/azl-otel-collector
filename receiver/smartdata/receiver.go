// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package smartdata

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type smartdataReceiver struct {
	host         component.Host
	cancel       context.CancelFunc
	logger       *zap.Logger
	nextConsumer consumer.Traces
	config       *Config
}

// Start begins collecting SMART data at the configured interval
func (smartdataReceiver *smartdataReceiver) Start(ctx context.Context, host component.Host) error {
	smartdataReceiver.host = host
	ctx = context.Background()
	ctx, smartdataReceiver.cancel = context.WithCancel(ctx)

	interval, _ := time.ParseDuration(smartdataReceiver.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				smartdataReceiver.logger.Info("collecting SMART data from machine")
				err := smartdataReceiver.collectAndSendSMARTData(ctx)
				if err != nil {
					smartdataReceiver.logger.Error("failed to collect SMART data", zap.Error(err))
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

// Shutdown stops the receiver
func (smartdataReceiver *smartdataReceiver) Shutdown(ctx context.Context) error {
	if smartdataReceiver.cancel != nil {
		smartdataReceiver.cancel()
	}
	return nil
}

// collectAndSendSMARTData collects SMART data from all disks and sends it as traces
func (smartdataReceiver *smartdataReceiver) collectAndSendSMARTData(ctx context.Context) error {
	disks, err := smartdataReceiver.getMachineDisks()
	if err != nil {
		return fmt.Errorf("failed to get disk disks: %w", err)
	}

	// log the number of disks
	smartdataReceiver.logger.Info("found disks", zap.Int("diskCount", len(disks)))

	traces := ptrace.NewTraces()
	traceID := pcommon.TraceID(uuid.New())
	resourceSpans := traces.ResourceSpans().AppendEmpty()

	// Set resource attributes (common metadata)
	resource := resourceSpans.Resource()
	resource.Attributes().PutStr("service.name", "smart-data-receiver")

	// Collect machine info
	machineInfo, err := getMachineInfo()
	if err != nil {
		smartdataReceiver.logger.Warn("failed to collect machine info", zap.Error(err))
	}
	machineInfoBytes, _ := json.Marshal(machineInfo)

	scopeSpans := resourceSpans.ScopeSpans().AppendEmpty()
	scopeSpans.Scope().SetName("smart-data-collection")
	scopeSpans.Scope().SetVersion("1.0.0")
	// Collect smart data for each disk
	for _, disk := range disks {
		smartdataReceiver.logger.Debug("collecting SMART data for disk", zap.String("disk", disk))

		smartData, err := smartdataReceiver.getSMARTData(disk)
		if err != nil {
			smartdataReceiver.logger.Warn("failed to collect SMART data for disk",
				zap.String("disk", disk), zap.Error(err))
			continue
		}

		diskSpan := scopeSpans.Spans().AppendEmpty()
		diskSpan.SetName(disk + "-smart-data")
		diskSpan.SetTraceID(traceID)
		diskSpan.SetSpanID(NewSpanID())
		diskSpan.SetKind(ptrace.SpanKindInternal)

		// Set the span attributes that hold the actual data
		diskSpan.Attributes().PutStr("machine-info", string(machineInfoBytes))
		diskSpan.Attributes().PutStr("smart-data", string(smartData))
	}

	// Send the traces
	err = smartdataReceiver.nextConsumer.ConsumeTraces(ctx, traces)
	if err != nil {
		return fmt.Errorf("failed to send traces: %w", err)
	}

	smartdataReceiver.logger.Info("successfully collected and sent SMART data",
		zap.Int("diskCount", len(disks)))
	return nil
}

// getSMARTData collects SMART data for a single disk
func (smartdataReceiver *smartdataReceiver) getSMARTData(disk string) (json.RawMessage, error) {
	cmd := exec.Command("smartctl", "--all", "--json", disk)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run smartctl on %s: %w", disk, err)
	}

	return json.RawMessage(output), nil
}

// getMachineDisks returns a list of disks on the system
func (smartdataReceiver *smartdataReceiver) getMachineDisks() ([]string, error) {
	cmd := exec.Command("lsblk", "--nodeps", "--output", "NAME", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running lsblk: %w", err)
	}

	var result struct {
		BlockDevices []struct {
			Name string `json:"name"`
		} `json:"blockdevices"`
	}

	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("error parsing lsblk output: %w", err)
	}

	// log the disks found
	smartdataReceiver.logger.Debug("found disks", zap.Any("disks", result.BlockDevices))

	disks := make([]string, 0, len(result.BlockDevices))
	for _, disk := range result.BlockDevices {
		disks = append(disks, "/dev/"+disk.Name)
	}

	return disks, nil
}

// getMachineInfo retrieves machine information from the host.
func getMachineInfo() (MachineInfo, error) {
	var machineInfo MachineInfo
	readAndTrim := func(path string) (string, error) {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to read %s: %w", path, err)
		}
		return strings.TrimSpace(string(data)), nil
	}
	var err error
	machineInfo.UUID, err = readAndTrim(ProductUUIDPath)
	if err != nil {
		return machineInfo, err
	}
	machineInfo.Manufacturer, err = readAndTrim(ProductManufacturerPath)
	if err != nil {
		return machineInfo, err
	}
	machineInfo.ModelNumber, err = readAndTrim(ProductModelNumberPath)
	if err != nil {
		return machineInfo, err
	}
	return machineInfo, nil
}

func NewSpanID() pcommon.SpanID {
	var rngSeed int64
	_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed)
	randSource := rand.New(rand.NewSource(rngSeed))

	var sid [8]byte
	randSource.Read(sid[:])
	spanID := pcommon.SpanID(sid)

	return spanID
}
