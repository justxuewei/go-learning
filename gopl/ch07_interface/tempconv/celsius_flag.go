package tempconv

import "fmt"

type celsiusFlag struct {
	Celsius
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
