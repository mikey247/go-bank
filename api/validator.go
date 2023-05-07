package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mikey247/go-bank/util"
)

// Custom curreny validator, takes in a specified "field" (currency),
// converts to string and validates with a custom function,
// in this case IsSupportedCurrency
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool{
    if currency,ok:=fieldLevel.Field().Interface().(string); ok {
		 return util.IsSupportedCurrency(currency)
	}
	return false
}