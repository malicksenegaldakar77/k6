// Code generated by "enumer -type=ResponseType -transform=title-lower -json -text -trimprefix ResponseType -output response_type_gen.go"; DO NOT EDIT.

package httpext

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _ResponseTypeName = "textbinaryarrayBuffernone"

var _ResponseTypeIndex = [...]uint8{0, 4, 10, 21, 25}

const _ResponseTypeLowerName = "textbinaryarraybuffernone"

func (i ResponseType) String() string {
	if i >= ResponseType(len(_ResponseTypeIndex)-1) {
		return fmt.Sprintf("ResponseType(%d)", i)
	}
	return _ResponseTypeName[_ResponseTypeIndex[i]:_ResponseTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResponseTypeNoOp() {
	var x [1]struct{}
	_ = x[ResponseTypeText-(0)]
	_ = x[ResponseTypeBinary-(1)]
	_ = x[ResponseTypeArrayBuffer-(2)]
	_ = x[ResponseTypeNone-(3)]
}

var _ResponseTypeValues = []ResponseType{ResponseTypeText, ResponseTypeBinary, ResponseTypeArrayBuffer, ResponseTypeNone}

var _ResponseTypeNameToValueMap = map[string]ResponseType{
	_ResponseTypeName[0:4]:        ResponseTypeText,
	_ResponseTypeLowerName[0:4]:   ResponseTypeText,
	_ResponseTypeName[4:10]:       ResponseTypeBinary,
	_ResponseTypeLowerName[4:10]:  ResponseTypeBinary,
	_ResponseTypeName[10:21]:      ResponseTypeArrayBuffer,
	_ResponseTypeLowerName[10:21]: ResponseTypeArrayBuffer,
	_ResponseTypeName[21:25]:      ResponseTypeNone,
	_ResponseTypeLowerName[21:25]: ResponseTypeNone,
}

var _ResponseTypeNames = []string{
	_ResponseTypeName[0:4],
	_ResponseTypeName[4:10],
	_ResponseTypeName[10:21],
	_ResponseTypeName[21:25],
}

// ResponseTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResponseTypeString(s string) (ResponseType, error) {
	if val, ok := _ResponseTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _ResponseTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResponseType values", s)
}

// ResponseTypeValues returns all values of the enum
func ResponseTypeValues() []ResponseType {
	return _ResponseTypeValues
}

// ResponseTypeStrings returns a slice of all String values of the enum
func ResponseTypeStrings() []string {
	strs := make([]string, len(_ResponseTypeNames))
	copy(strs, _ResponseTypeNames)
	return strs
}

// IsAResponseType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResponseType) IsAResponseType() bool {
	for _, v := range _ResponseTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ResponseType
func (i ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseType
func (i *ResponseType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ResponseType should be a string, got %s", data)
	}

	var err error
	*i, err = ResponseTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for ResponseType
func (i ResponseType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ResponseType
func (i *ResponseType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ResponseTypeString(string(text))
	return err
}
