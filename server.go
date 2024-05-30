package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/index.html")
}

func conversion(w http.ResponseWriter, r *http.Request) {
	initialFormat := r.URL.Query().Get("initialFormat")
	finalFormat := r.URL.Query().Get("finalFormat")
	temp := r.URL.Query().Get("temp")
	tempFloat, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		http.ServeFile(w, r, "pages/error.html")
	}

	valueToReturn := SITemp{}
	switch initialFormat + finalFormat {
	case "CF":
		valueToReturn = valueToReturn.CelsiusToFahrenheit(tempFloat)
		break
	case "CK":
		valueToReturn = valueToReturn.CelsiusToKelvin(tempFloat)
		break
	case "FC":
		valueToReturn = valueToReturn.FahrenheitToCelsius(tempFloat)
		break
	case "FK":
		valueToReturn = valueToReturn.FahrenheitToKelvin(tempFloat)
		break
	case "KC":
		valueToReturn = valueToReturn.KelvinToCelsius(tempFloat)
		break
	case "KF":
		valueToReturn = valueToReturn.KelvinToFahrenheit(tempFloat)
		break
	default:
		valueToReturn = valueToReturn.SameFormat(tempFloat, initialFormat)
		break
	}
	w.Write([]byte(fmt.Sprintf(`<h2>%.2f %s</h2>`, valueToReturn.temp, valueToReturn.si)))

}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/conversion", conversion)
	http.ListenAndServe(":8080", nil)

}
