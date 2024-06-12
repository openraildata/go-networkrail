// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.0

package smart

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rickb777/enumeration/v3/enum"
)

// AllStepTypes lists all 7 values in order.
var AllStepTypes = []StepType{
	Between, From, To, IntermediateFirst,
	Clearout, Interpose, Intermediate,
}

// AllStepTypeEnums lists all 7 values in order.
var AllStepTypeEnums = enum.IntEnums{
	Between, From, To, IntermediateFirst,
	Clearout, Interpose, Intermediate,
}

const (
	steptypeEnumStrings = "BetweenFromToIntermediateFirstClearoutInterposeIntermediate"
	steptypeJSONStrings = "BFTDCIE"
)

var (
	steptypeEnumIndex = [...]uint16{0, 7, 11, 13, 30, 38, 47, 59}
	steptypeJSONIndex = [...]uint16{0, 1, 2, 3, 4, 5, 6, 7}
)

// Ordinal returns the ordinal number of a StepType. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v StepType) Ordinal() int {
	switch v {
	case Between:
		return 0
	case From:
		return 1
	case To:
		return 2
	case IntermediateFirst:
		return 3
	case Clearout:
		return 4
	case Interpose:
		return 5
	case Intermediate:
		return 6
	}
	return -1
}

// String returns the literal string representation of a StepType, which is
// the same as the const identifier but without prefix or suffix.
func (v StepType) String() string {
	o := v.Ordinal()
	return v.toString(o, steptypeEnumStrings, steptypeEnumIndex[:])
}

func (v StepType) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllStepTypes) {
		return fmt.Sprintf("StepType(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a StepType is one of the defined constants.
func (v StepType) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v StepType) Int() int {
	return int(v)
}

var invalidStepTypeValue = func() StepType {
	var v StepType
	for {
		if !slices.Contains(AllStepTypes, v) {
			return v
		}
		v++
	} // AllStepTypes is a finite set so loop will terminate eventually
}()

// StepTypeOf returns a StepType based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid StepType is returned.
func StepTypeOf(v int) StepType {
	if 0 <= v && v < len(AllStepTypes) {
		return AllStepTypes[v]
	}
	return invalidStepTypeValue
}

// Parse parses a string to find the corresponding StepType, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsStepType.
//
// Usage Example
//
//	v := new(StepType)
//	err := v.Parse(s)
//	...  etc
func (v *StepType) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := steptypeTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *StepType) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = StepType(num)
		return v.IsValid()
	}
	return false
}

func (v *StepType) parseFallback(in, s string) error {
	if v.parseString(s, steptypeEnumStrings, steptypeEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised steptype")
}

func (v *StepType) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllStepTypes[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// steptypeTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var steptypeTransformInput = func(in string) string {
	return in
}

// AsStepType parses a string to find the corresponding StepType, accepting either one of the string values or
// a number. The input representation is determined by steptypeMarshalTextRep. It wraps Parse.
func AsStepType(s string) (StepType, error) {
	var v = new(StepType)
	err := v.Parse(s)
	return *v, err
}

// MustParseStepType is similar to AsStepType except that it panics on error.
func MustParseStepType(s string) StepType {
	v, err := AsStepType(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v StepType) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, steptypeJSONStrings, steptypeJSONIndex[:])
}

func (v StepType) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v StepType) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v StepType) invalidError() error {
	return fmt.Errorf("%d is not a valid steptype", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v StepType) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, steptypeJSONStrings, steptypeJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *StepType) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *StepType) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := steptypeTransformInput(in)

	if v.parseString(s, steptypeJSONStrings, steptypeJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, steptypeEnumStrings, steptypeEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised steptype")
}

// steptypeMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var steptypeMarshalNumber = func(v StepType) string {
	return strconv.FormatInt(int64(v), 10)
}
