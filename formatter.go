package gandalff

import (
	"fmt"
	"math"
)

const (
	cutoffDelta                = 3.14159265358979323846 * 1e-6 // pi * 10^-6 used in rounding last digits close to 5
	defaultDecimalDigits       = 11
	defaultThreshold           = -8
	defaultScientificThreshold = 9
	defaultMaxDigits           = 11
	defaultMovingDigits        = 3
	defaultNanText             = "NA"
	defaultInfText             = "Inf"
)

type NumericFormatter struct {
	decimalDigits       int    // The number of digits to print after the decimal point.
	threshold           int    // The number of digits to print before the decimal point.
	scientificThreshold int    // The number of digits to print before switching to scientific notation.
	maxDigits           int    // The maximum number of digits to print.
	movingDigits        int    // The number of digits to print before the decimal point for very small numbers.
	nanText             string // The text to print for NaNs.
	infText             string // The text to print for Infs.
	exponentialFormat   bool   // Whether to use scientific notation.
	useLipGloss         bool   // Whether to use lipgloss.
}

func NewNumericFormatter() *NumericFormatter {
	return &NumericFormatter{
		decimalDigits:       defaultDecimalDigits,
		threshold:           defaultThreshold,
		scientificThreshold: defaultScientificThreshold,
		maxDigits:           defaultMaxDigits,
		movingDigits:        defaultMovingDigits,
		nanText:             defaultNanText,
		infText:             defaultInfText,
		exponentialFormat:   false,
		useLipGloss:         false,
	}
}

func (f *NumericFormatter) Push(num float64) {
	if math.IsNaN(num) || math.IsInf(num, 0) {
		return
	}

	absNum := math.Abs(num)
	if f.exponentialFormat || absNum <= 1e-50 {
		return
	}

	// Set exponential format if very large.
	if absNum >= math.Pow(10, float64(f.scientificThreshold))-0.5-0.1*cutoffDelta { // 999 999 999.4999...
		f.exponentialFormat = true
		return
	}

	// Check if very small, here more care is needed because of rounding.
	signif, exponent := sigAndExp(num)

	// If the number would be rounded up, there might be less digits needed / no switch to exp format.
	// To check this, determine the number of digits that would be needed for this number:
	digits := int(math.Min(float64(f.maxDigits), math.Max(0, float64(f.movingDigits-exponent))))

	// If this is 0, no need to check anything, otherwise find out where the rounding digit is related to the exponent.
	if digits > 0 && (math.Abs(signif) >= 10-math.Pow(0.1, math.Min(float64(f.maxDigits+exponent), float64(f.movingDigits))+1)*(5+cutoffDelta)) { // this is for rounding issues
		signif = 1 // not used but will keep for consistency
		exponent++
	}
	if exponent < 2-f.scientificThreshold { // -7
		f.exponentialFormat = true
		return
	} else if exponent > f.threshold && f.decimalDigits > 0 {
		f.decimalDigits = int(math.Min(float64(f.maxDigits), math.Max(0, float64(f.movingDigits-exponent))))
		f.threshold = exponent
	}
}

func (f *NumericFormatter) Format(num float64) string {
	if math.IsInf(num, 0) {
		return f.infText
	}
	if math.IsNaN(num) {
		return f.nanText
	}

	// Very small numbers, which are treated as zero, are formatted as 0.
	if math.Abs(num) <= 1e-50 {
		return "0"
	}

	signif, exponent := sigAndExp(num)
	if f.exponentialFormat {
		if math.Abs(signif) >= 9.9995-1e-4*cutoffDelta { // to avoid printing "10.000 x 10^exp" when signif rounded up, eg. 9.99999
			if signif < 0 {
				signif = -1.0
			} else {
				signif = 1.0
			}
			exponent++
		}
		return fmt.Sprintf("%se%d", fmt.Sprintf("%.*f", 3, signif), exponent)
	} else if exponent < -f.decimalDigits && f.decimalDigits != f.maxDigits {
		minNeededDigits := -exponent - 1
		if math.Abs(signif) >= 9.5-cutoffDelta { // if the last digit would be 9, check if rounded up or down
			minNeededDigits++
		}
		return fmt.Sprintf("%.*f", minNeededDigits, num)
	} else {
		return fmt.Sprintf("%.*f", f.decimalDigits, num)
	}
}

func sigAndExp(num float64) (float64, int) {
	exponent := int(math.Floor(math.Log10(math.Abs(num))))
	signif := num * math.Pow(10, -float64(exponent))
	if math.Abs(signif) >= 10 { // extra check for possible rounding wierdness
		exponent++
		signif /= 10
	}

	return signif, exponent
}
