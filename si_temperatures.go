package main

type SITemp struct {
	temp float64
	si   string
}

type SITempConversion interface {
	CelsiusToFahrenheit(temperature float64) SITemp
	CelsiusToKelvin(temperature float64) SITemp
	FahrenheitToCelsius(temperature float64) SITemp
	FahrenheitToKelvin(temperature float64) SITemp
	KelvinToCelsius(temperature float64) SITemp
	KelvinToFahrenheit(temperature float64) SITemp
	SameFormat(temperature float64, format string) SITemp
}

func (s SITemp) CelsiusToFahrenheit(temperature float64) SITemp {
	//°F = (°C x 1.8) + 32
	s.si = "°F"
	s.temp = (temperature * 1.8) + 32
	return s
}

func (s SITemp) CelsiusToKelvin(temperature float64) SITemp {
	s.si = "°K"
	s.temp = temperature + 273.15
	return s
}

func (s SITemp) FahrenheitToCelsius(temperature float64) SITemp {
	//(°F - 32) x {\frac  {5}{9}}
	s.si = "°C"
	s.temp = ((temperature - 32) * 5) / 9
	return s
}

func (s SITemp) FahrenheitToKelvin(temperature float64) SITemp {
	s.si = "°K"
	s.temp = (temperature-32)/1.8 + 273.15
	//(Fahrenheit - 32) / 1.8 + 273.15
	return s
}

func (s SITemp) KelvinToCelsius(temperature float64) SITemp {
	s.si = "°C"
	s.temp = temperature - 273.15
	//T (K) - 273,15
	return s
}

func (s SITemp) KelvinToFahrenheit(temperature float64) SITemp {
	s.si = "°F"
	s.temp = 1.8*(temperature-273.15) + 32
	//1.8*(K-273) + 32
	return s
}

func (s SITemp) SameFormat(temperature float64, format string) SITemp {
	s.si = "°" + format
	s.temp = 1.0 * temperature
	return s
}
