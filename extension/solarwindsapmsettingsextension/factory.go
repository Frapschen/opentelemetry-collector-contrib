// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solarwindsapmsettingsextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/solarwindsapmsettingsextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/solarwindsapmsettingsextension/internal/metadata"
)

const (
	DefaultInterval = "1m"
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: DefaultInterval,
	}
}

func createExtension(_ context.Context, settings extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newSolarwindsApmSettingsExtension(cfg.(*Config), settings.Logger)
}

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}
