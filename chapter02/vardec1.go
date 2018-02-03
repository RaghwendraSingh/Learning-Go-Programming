package main

import (
	"fmt"
)

var Name, desc string
var radius int32
var mass float64
var active bool
var satellites []string

func main1() {
	Name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(Name)
	fmt.Println(desc)
	fmt.Println("Radis (km), ", radius)
	fmt.Println("Mass (kg), ", mass)
	fmt.Println("Satellites", satellites)
}
