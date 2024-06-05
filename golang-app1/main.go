package main

import (
    "fmt"
    "math"
    "net/http"
    "os"
    "strconv"
)

// Mass represents the mass density of a material
type Mass struct {
    Density float64
}

// MassVolume is an interface for calculating mass based on volume
type MassVolume interface {
    volume(dimension float64) float64
    density() float64
}

// Sphere represents a sphere with a given mass density
type Sphere struct {
    Mass
}

// Cube represents a cube with a given mass density
type Cube struct {
    Mass
}

// Implement the MassVolume interface for Sphere
func (s Sphere) volume(dimension float64) float64 {
    // Volume of a sphere: (4/3) * π * r^3
    return (4.0 / 3) * math.Pi * math.Pow(dimension, 3)
}

func (s Sphere) density() float64 {
    return s.Density
}

// Implement the MassVolume interface for Cube
func (c Cube) volume(dimension float64) float64 {
    // Volume of a cube: side^3
    return math.Pow(dimension, 3)
}

func (c Cube) density() float64 {
    return c.Density
}

func Handler(massVolume MassVolume) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if dimension, err := strconv.ParseFloat(r.URL.Query().Get("dimension"), 64); err == nil {
            weight := massVolume.density() * massVolume.volume(dimension)
            w.Write([]byte(fmt.Sprintf("%.2f", math.Round(weight*100)/100)))
            return
        }
        w.WriteHeader(http.StatusBadRequest)
    }
}

func main() {
    port, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }

    aluminiumSphere := Sphere{Mass{Density: 2.710}}
    ironCube := Cube{Mass{Density: 7.874}}

    http.HandleFunc("/aluminium/sphere", Handler(aluminiumSphere))
    http.HandleFunc("/iron/cube", Handler(ironCube))

    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        panic(err)
    }
}
