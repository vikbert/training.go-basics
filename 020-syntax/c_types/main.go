package main

import "fmt"

type Centimeter float32
type Meter float32

func (m *Meter) toCentimeter() Centimeter {
	return Centimeter(float32(*m) * 100)
}

func (cm *Centimeter) toMeter() Meter {
	return Meter(float32(*cm) / 100)
}

func main() {
	fmt.Println("Types:")

	lengthCentimeter := Centimeter(140.0)
	var lengthMeter = centimeterToMeter(lengthCentimeter)

	fmt.Printf("Type of lengthCentimeter: %T\n", lengthCentimeter)

	// Ausgabe: "140 cm equals 1.40 m"
	fmt.Printf("%.2f cm equals %.2f m", lengthCentimeter, lengthMeter)
}

func centimeterToMeter(cm Centimeter) Meter {
	return cm.toMeter()
}
