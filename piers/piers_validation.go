package piers

import (
	"reflect"

	"github.com/tudorhulban/hxerrors"
)

func ValidateDependencies(dependencies any) error {
	// 1. Safe guard against a completely untyped nil
	if dependencies == nil {
		return hxerrors.ErrValidation{
			Caller: "ValidateDependencies",
			Issue: hxerrors.ErrNilInput{
				InputName: "dependencies",
			},
		}
	}

	piersType := reflect.TypeOf(dependencies)
	piersValue := reflect.ValueOf(dependencies)

	// 2. Loop to handle nested pointers (e.g., **MyStruct) and handle typed nils safely
	for piersType.Kind() == reflect.Ptr {
		if piersValue.IsNil() {
			return hxerrors.ErrValidation{
				Caller: "ValidateDependencies",
				Issue: hxerrors.ErrNilInput{
					InputName: "dependencies",
				},
			}
		}

		piersType = piersType.Elem()
		piersValue = piersValue.Elem()
	}

	switch piersType.Kind() {
	case reflect.Struct:
		for fieldIndex := 0; fieldIndex < piersType.NumField(); fieldIndex++ {
			field := piersType.Field(fieldIndex)

			// 3. Skip unexported fields to prevent panics on fieldValue.IsNil()
			if !field.IsExported() {
				continue
			}

			fieldValue := piersValue.Field(fieldIndex)

			switch field.Type.Kind() {
			case reflect.Ptr, reflect.Interface:
				if fieldValue.IsNil() {
					return hxerrors.ErrValidation{
						Caller: "ValidateDependencies",
						Issue: hxerrors.ErrNilInput{
							InputName: field.Name,
						},
					}
				}

			default:
				continue
			}
		}

	default:
		return hxerrors.ErrValidation{
			Caller: "ValidateDependencies",
			Issue: hxerrors.ErrInvalidInput{
				InputName: "dependencies",
			},
		}
	}

	return nil
}
