package gandalff

import (
	"fmt"
	"math"

	"github.com/charmbracelet/lipgloss"
)

type Formatter interface {
	Push(val any)
	Format(val any) string
}

const (
	cutoffDelta                = 3.14159265358979323846 * 1e-6 // pi * 10^-6 used in rounding last digits close to 5
	defaultDecimalDigits       = 11
	defaultThreshold           = -8
	defaultScientificThreshold = 9
	defaultMaxDigits           = 10
	defaultMovingDigits        = 3
	defaultNaText              = NA_TEXT
	defaultInfText             = INF_TEXT
)

type NumericFormatter struct {
	decimalDigits       int    // The number of digits to print after the decimal point.
	threshold           int    // The number of digits to print before the decimal point.
	scientificThreshold int    // The number of digits to print before switching to scientific notation.
	maxDigits           int    // The maximum number of digits to print.
	maxWidth            int    // Maximum width required to print the number.
	movingDigits        int    // The number of digits to print before the decimal point for very small numbers.
	naText              string // The text to print for NaNs.
	infText             string // The text to print for Infs.
	useExpFormat        bool   // Whether to use scientific notation.
	hasNegative         bool   // Whether negative numbers are present.
	useLipGloss         bool   // Whether to use lipgloss.
	justifyLeft         bool   // Whether to justify left.

	styleBold   lipgloss.Style
	styleItalic lipgloss.Style
	styleNa     lipgloss.Style
	styleNum    lipgloss.Style
	styleNumNeg lipgloss.Style
	styleNone   lipgloss.Style
}

func NewNumericFormatter() *NumericFormatter {
	return &NumericFormatter{
		decimalDigits:       defaultDecimalDigits,
		threshold:           defaultThreshold,
		scientificThreshold: defaultScientificThreshold,
		maxDigits:           defaultMaxDigits,
		maxWidth:            0,
		movingDigits:        defaultMovingDigits,
		naText:              defaultNaText,
		infText:             defaultInfText,
		useExpFormat:        false,
		hasNegative:         false,
		useLipGloss:         false,
		justifyLeft:         false,

		styleBold:   lipgloss.NewStyle().Bold(true),
		styleItalic: lipgloss.NewStyle().Italic(true),
		styleNa:     lipgloss.NewStyle().Bold(true).Copy().Foreground(lipgloss.Color("#c00020")),
		styleNum:    lipgloss.NewStyle().Foreground(lipgloss.Color("#0080c0")),
		styleNumNeg: lipgloss.NewStyle().Foreground(lipgloss.Color("#c04000")),
		styleNone:   lipgloss.NewStyle(),
	}
}

func (f *NumericFormatter) SetDecimalDigits(decimalDigits int) *NumericFormatter {
	f.decimalDigits = decimalDigits
	return f
}

func (f *NumericFormatter) SetThreshold(threshold int) *NumericFormatter {
	f.threshold = threshold
	return f
}

func (f *NumericFormatter) SetScientificThreshold(scientificThreshold int) *NumericFormatter {
	f.scientificThreshold = scientificThreshold
	return f
}

func (f *NumericFormatter) SetMaxDigits(maxDigits int) *NumericFormatter {
	f.maxDigits = maxDigits
	return f
}

func (f *NumericFormatter) SetMovingDigits(movingDigits int) *NumericFormatter {
	f.movingDigits = movingDigits
	return f
}

func (f *NumericFormatter) SetNaText(naText string) *NumericFormatter {
	f.naText = naText
	return f
}

func (f *NumericFormatter) SetInfText(infText string) *NumericFormatter {
	f.infText = infText
	return f
}

func (f *NumericFormatter) SetUseLipGloss(useLipGloss bool) *NumericFormatter {
	f.useLipGloss = useLipGloss
	return f
}

func (f *NumericFormatter) Push(val any) {
	var num float64
	switch val.(type) {
	case int:
		num = float64(val.(int))
	case int64:
		num = float64(val.(int64))
	case float64:
		num = val.(float64)
	default:
		return
	}

	if math.IsNaN(num) {
		f.maxWidth = int(math.Max(float64(f.maxWidth), float64(len(f.naText))))
		return
	}

	if math.IsInf(num, 1) {
		f.maxWidth = int(math.Max(float64(f.maxWidth), float64(len(f.infText))))
		return
	}

	if math.IsInf(num, -1) {
		f.maxWidth = int(math.Max(float64(f.maxWidth), float64(len(f.infText)+1)))
		f.hasNegative = true
		return
	}

	absNum := math.Abs(num)
	if f.useExpFormat || absNum <= 1e-50 {
		return
	}

	// Set exponential format if very large.
	if absNum >= math.Pow(10, float64(f.scientificThreshold))-0.5-0.1*cutoffDelta { // 999 999 999.4999...
		f.useExpFormat = true
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
		f.useExpFormat = true

	} else if exponent > f.threshold && f.decimalDigits > 0 {
		f.decimalDigits = int(math.Min(float64(f.maxDigits), math.Max(0, float64(f.movingDigits-exponent))))
		f.threshold = exponent
	}
}

func (f *NumericFormatter) Format(val any) string {
	var num float64
	switch val.(type) {
	case int:
		num = float64(val.(int))
	case int64:
		num = float64(val.(int64))
	case float64:
		num = val.(float64)
	default:
		return f.naText
	}

	if math.IsInf(num, -1) {
		return f.render(fmt.Sprintf("-%s", f.infText), f.styleNumNeg)
	}
	if math.IsInf(num, 1) {
		return f.render(f.infText, f.styleNum)
	}
	if math.IsNaN(num) {
		return f.render(f.naText, f.styleNa)
	}

	// Very small numbers, which are treated as zero, are formatted as 0.
	if math.Abs(num) <= 1e-50 {
		return f.render("0", f.styleNum)
	}

	signif, exponent := sigAndExp(num)
	if f.useExpFormat {
		if math.Abs(signif) >= 9.9995-1e-4*cutoffDelta { // to avoid printing "10.000 x 10^exp" when signif rounded up, eg. 9.99999
			if signif < 0 {
				signif = -1.0
			} else {
				signif = 1.0
			}
			exponent++
		}

		return f.render(fmt.Sprintf("%se%d", fmt.Sprintf("%.*f", 3, signif), exponent), f.styleNum)
	}

	if exponent < -f.decimalDigits && f.decimalDigits != f.maxDigits {
		minNeededDigits := -exponent - 1
		if math.Abs(signif) >= 9.5-cutoffDelta { // if the last digit would be 9, check if rounded up or down
			minNeededDigits++
		}

		return f.render(fmt.Sprintf("%.*f", minNeededDigits, num), f.styleNum)
	}

	return f.render(fmt.Sprintf("%.*f", f.decimalDigits, num), f.styleNum)
}

func (f *NumericFormatter) render(s string, style lipgloss.Style) string {
	if f.useLipGloss {
		if f.justifyLeft {
			return style.Render(fmt.Sprintf("%-*s", f.maxDigits, s))
		} else {
			return style.Render(fmt.Sprintf("%*s", f.maxDigits, s))
		}
	}

	if f.justifyLeft {
		return fmt.Sprintf("%-*s", f.maxDigits, s)
	} else {
		return fmt.Sprintf("%*s", f.maxDigits, s)
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

type StringFormatter struct {
	useLipGloss bool
	lengths     []int

	styleString lipgloss.Style
}

func NewStringFormatter() *StringFormatter {
	return &StringFormatter{
		useLipGloss: false,
		lengths:     make([]int, 0),
		styleString: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#98AFC7")),
	}
}

func (f *StringFormatter) SetUseLipGloss(useLipGloss bool) *StringFormatter {
	f.useLipGloss = useLipGloss
	return f
}

func (f *StringFormatter) Push(val any) {
	if s, ok := val.(string); ok {
		for i := 0; i < len(f.lengths); i++ {
			if len(s) > f.lengths[i] {
				f.lengths = append(f.lengths, 0)
				copy(f.lengths[i+1:], f.lengths[i:])
				f.lengths[i] = len(s)
				return
			}
		}
	}
}

func (f *StringFormatter) Format(val any) string {
	if s, ok := val.(string); ok {
		if f.useLipGloss {
			return f.styleString.Render(fmt.Sprintf("%-*s", f.lengths[0], s))
		} else {
			return fmt.Sprintf("%-*s", f.lengths[0], s)
		}
	}
	return NA_TEXT
}
