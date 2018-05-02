package main

import (
	"fmt"
)

//typedef
type	Celsius		float64
type	Fahrenheit	float64

//const test var
const (
	AbsoluteZeroC	Celsius	=	-273.15	
	FreezingC	Celsius	=	0
	BoilingC	Celsius	=	100
)

func	CToF	(c Celsius)	Fahrenheit	{
	return Fahrenheit(c*9/5 + 32)
}

func	FToC	(f Fahrenheit)	Celsius		{
	return Celsius((f-32)*5/9)
}

func main() {
	fmt.Println("AbsoluteZeroC->F", CToF(AbsoluteZeroC))
	fmt.Println("FreezingC->F", CToF(FreezingC))
	fmt.Println("BoilingC->F", CToF(BoilingC))
	fmt.Println("0F->C", FToC(0))
}
