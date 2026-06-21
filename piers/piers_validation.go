package piers

import (
	"reflect"

	"github.com/tudorhulban/hxerrors"
)

func ValidatePiers(piers any) error {
	// 1. Safe guard against a completely untyped nil
	if piers == nil {
		return hxerrors.ErrValidation{
			Caller: "ValidatePiers",
			Issue: hxerrors.ErrNilInput{
				InputName: "piers",
			},
		}
	}

	piersType := reflect.TypeOf(piers)
	piersValue := reflect.ValueOf(piers)

	// 2. Loop to handle nested pointers (e.g., **MyStruct) and handle typed nils safely
	for piersType.Kind() == reflect.Ptr {
		if piersValue.IsNil() {
			return hxerrors.ErrValidation{
				Caller: "ValidatePiers",
				Issue: hxerrors.ErrNilInput{
					InputName: "piers",
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
						Caller: "ValidatePiers",
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
			Caller: "ValidatePiers",
			Issue: hxerrors.ErrInvalidInput{
				InputName: "piers",
			},
		}
	}

	return nil
}
