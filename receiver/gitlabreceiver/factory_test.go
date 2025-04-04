// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gitlabreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitlabreceiver/internal/metadata"
)

func TestCreateTracesReceiver(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")

	tReceiver, err := factory.CreateTraces(context.Background(), receivertest.NewNopSettings(metadata.Type), cfg, consumertest.NewNop())
	assert.NoError(t, err)
	assert.NotNil(t, tReceiver, "traces receiver creation failed")
}
