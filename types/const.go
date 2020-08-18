package types

type Units string

const UnitDegrees Units = "degrees"
const UnitRadians Units = "radians"
const UnitMiles Units = "miles"
const UnitKilometers Units = "kilometers"
const UnitNauticalmiles Units = "nauticalmiles"

// Factors
const FactorEarthRadius = 6371008.8
const FactorCentimeters = FactorEarthRadius * 100
const FactorDegrees = FactorEarthRadius / 111325
const FactorFeet = FactorEarthRadius * 3.28084
const FactorInches = FactorEarthRadius * 39.370
const FactorKilometers = FactorEarthRadius / 1000
const FactorMeters = FactorEarthRadius
const FactorMiles = FactorEarthRadius / 1609.344
const FactorMillimeters = FactorEarthRadius * 1000
const FactorNauticalmiles = FactorEarthRadius / 1852
const FactorRadians = 1
const FactorYards = FactorEarthRadius / 1.0936
