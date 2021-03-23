# Naming Convention in Golang

Example:

    package main

    func main() {

    var max_value int // Not idiomatic in go

    var mv int // max value

        var packetsReceived int // Not ok name too long
        var b int               // OK

        const MAX_VALUE = 100 // NOT OK
        const N = 100         // Better

        var max = 100

        maxValue := 1000   // recommended
        max_values := 1000 // Not recommended

        writeToDB := true  // ok, idiomatic
        writeToDb := false // not ok
    }