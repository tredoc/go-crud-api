package validator

import "regexp"

var (
	CantBeEmpty        = "can't be empty"
	OnlyLatinLetters   = "must contain only latin letters"
	OnlyInThePast      = "publish date can't be in future"
	CantBeLessThanOne  = "can't be less than 1"
	CantBeBiggerThan5k = "must be less than 5000"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{make(map[string]string)}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key string, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key string, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
