package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
    return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(spentOvenTime int) int {
    return (40 - spentOvenTime)
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
    return (2*layers)
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers, spentOvenTime int) int {
    return ((2*layers) + spentOvenTime)
}
