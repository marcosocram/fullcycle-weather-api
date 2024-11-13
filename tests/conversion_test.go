package tests

import (
	"github.com/marcosocram/fullcycle-weather-api/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	assert.InDelta(t, 77.0, pkg.CelsiusToFahrenheit(25), 0.1)
	assert.InDelta(t, 32.0, pkg.CelsiusToFahrenheit(0), 0.1)
	assert.InDelta(t, -40.0, pkg.CelsiusToFahrenheit(-40), 0.1)
}

func TestCelsiusToKelvin(t *testing.T) {
	assert.InDelta(t, 298.15, pkg.CelsiusToKelvin(25), 0.1)
	assert.InDelta(t, 273.15, pkg.CelsiusToKelvin(0), 0.1)
	assert.InDelta(t, 233.15, pkg.CelsiusToKelvin(-40), 0.1)
}
