package jedlik

import (
	"fmt"
	"math"
)

func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	laps := math.Ceil(float64(trackDistance) / float64(c.speed))
	return c.battery >= int(laps)*c.batteryDrain
}
