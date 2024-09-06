package postsort

import (
	"errors"
)

type SortResult string

const (
	MaxDimensions = 150
	MaxVolume     = 1_000_000
	MaxMass       = float32(20.0)

	// STANDARD packages (those that are not bulky or heavy) can be handled normally.
	STANDARD SortResult = "STANDARD"
	// SPECIAL packages that are either heavy or bulky can't be handled automatically.
	SPECIAL SortResult = "SPECIAL"
	// REJECTED packages that are either heavy or bulky can't be handled automatically.
	REJECTED SortResult = "REJECTED"
)

var (
	InvalidPackageError = errors.New("invalid package dimensions or mass")
)

// Sort categorizes a package based on its dimensions and mass.
// If the package is both bulky and heavy, it is REJECTED.
// If the package is either bulky or heavy, it is SPECIAL.
// If the package is neither bulky nor heavy, it is STANDARD.
func Sort(width, height, length uint, mass float32) (SortResult, error) {
	if width == 0 || height == 0 || length == 0 || mass <= 0 {
		return REJECTED, InvalidPackageError
	}

	// A package is "bulky" if its volume (Width x Height x Length)
	// is greater than or equal to 1,000,000 cm³ or when one of its dimensions is greater or equal to 150 cm.
	isBulky := width >= MaxDimensions || height >= MaxDimensions || length >= MaxDimensions || width*height*length >= MaxVolume

	// MaxMass - A package is "heavy" when its mass is greater or equal to 20 kg.
	isHeavy := mass >= MaxMass

	switch {
	case isBulky && isHeavy:
		return REJECTED, nil
	case isBulky || isHeavy:
		return SPECIAL, nil
	default:
		return STANDARD, nil
	}
}
