package enumutils

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/inumericconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
)

func (mapper *ReverseByteEnumMap) Get(value byte) ibyteconstant.IByteConstant {
	return mapper.Map[value]
}

func ParseByteConstant(value byte, valueType ibyteconstant.IByteConstant) ibyteconstant.IByteConstant {
	mapper := NewReverseByteEnumMap(valueType)
	return mapper.Get(value)
}

func IsDefined(valueType ibyteconstant.IByteConstant, value byte) bool {
	return ParseByteConstant(value, valueType) != nil
}

func ParseStringConstant(valueType istringconstant.IStringConstant, value string) istringconstant.IStringConstant {
	mapper := NewReverseStringEnumMap(valueType)
	return mapper.Get(value)
}

func ParseNumericConstant(valueType inumericconstant.INumericConstant, value int) inumericconstant.INumericConstant {
	mapper := NewReverseIntEnumMap(valueType)
	return mapper.Get(value)
}
