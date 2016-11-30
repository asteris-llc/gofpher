package coretypes

// UnitType is a base unit type
type UnitType struct{}

// Unit returns a basic unit
func Unit() UnitType {
	return UnitType{}
}
