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

// AllEvents lists all 4 values in order.
var AllEvents = []Event{
	ArriveUp, DepartUp, ArriveDown, DepartDown,
}

// AllEventEnums lists all 4 values in order.
var AllEventEnums = enum.IntEnums{
	ArriveUp, DepartUp, ArriveDown, DepartDown,
}

const (
	eventEnumStrings = "ArriveUpDepartUpArriveDownDepartDown"
	eventJSONStrings = "ABCD"
)

var (
	eventEnumIndex = [...]uint16{0, 8, 16, 26, 36}
	eventJSONIndex = [...]uint16{0, 1, 2, 3, 4}
)

// Ordinal returns the ordinal number of a Event. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Event) Ordinal() int {
	switch v {
	case ArriveUp:
		return 0
	case DepartUp:
		return 1
	case ArriveDown:
		return 2
	case DepartDown:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Event, which is
// the same as the const identifier but without prefix or suffix.
func (v Event) String() string {
	o := v.Ordinal()
	return v.toString(o, eventEnumStrings, eventEnumIndex[:])
}

func (v Event) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllEvents) {
		return fmt.Sprintf("Event(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Event is one of the defined constants.
func (v Event) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Event) Int() int {
	return int(v)
}

var invalidEventValue = func() Event {
	var v Event
	for {
		if !slices.Contains(AllEvents, v) {
			return v
		}
		v++
	} // AllEvents is a finite set so loop will terminate eventually
}()

// EventOf returns a Event based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Event is returned.
func EventOf(v int) Event {
	if 0 <= v && v < len(AllEvents) {
		return AllEvents[v]
	}
	return invalidEventValue
}

// Parse parses a string to find the corresponding Event, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsEvent.
//
// Usage Example
//
//	v := new(Event)
//	err := v.Parse(s)
//	...  etc
func (v *Event) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := eventTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Event) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Event(num)
		return v.IsValid()
	}
	return false
}

func (v *Event) parseFallback(in, s string) error {
	if v.parseString(s, eventEnumStrings, eventEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised event")
}

func (v *Event) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllEvents[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// eventTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var eventTransformInput = func(in string) string {
	return in
}

// AsEvent parses a string to find the corresponding Event, accepting either one of the string values or
// a number. The input representation is determined by eventMarshalTextRep. It wraps Parse.
func AsEvent(s string) (Event, error) {
	var v = new(Event)
	err := v.Parse(s)
	return *v, err
}

// MustParseEvent is similar to AsEvent except that it panics on error.
func MustParseEvent(s string) Event {
	v, err := AsEvent(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Event) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, eventJSONStrings, eventJSONIndex[:])
}

func (v Event) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Event) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Event) invalidError() error {
	return fmt.Errorf("%d is not a valid event", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Event) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, eventJSONStrings, eventJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Event) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Event) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := eventTransformInput(in)

	if v.parseString(s, eventJSONStrings, eventJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, eventEnumStrings, eventEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised event")
}

// eventMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var eventMarshalNumber = func(v Event) string {
	return strconv.FormatInt(int64(v), 10)
}
