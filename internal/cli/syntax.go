package cli

import (
	"fmt"
	"strings"

	"github.com/apex/log"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"

	"github.com/little-angry-clouds/particle/internal/config"
)

func Syntax(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var stringErrors []string

	v := validator.New()

	// Translate the error to something that is more or less intelligible
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	_ = v.RegisterTranslation("eq", trans, func(ut ut.Translator) error {
		return ut.Add("eq", "{0} has an incorrect value", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eq", fe.StructNamespace())
		return t
	})

	allErrors := v.Struct(configuration)

	if allErrors != nil {
		// Translate all the errors and join them in one big error
		for _, e := range allErrors.(validator.ValidationErrors) {
			stringErrors = append(stringErrors, e.Translate(trans))
		}

		err = fmt.Errorf(strings.Join(stringErrors, ", "))
	}

	return err
}
