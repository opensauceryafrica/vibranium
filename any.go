package vibranium

type Any struct {
	// source is the root vibranium type.
	source Wakanda
	// required is a bool pointer that determines if the field is required.
	required bool
	// forbidden is a bool pointer that determines if the field is forbidden.
	forbidden bool
	// description is a string pointer that holds a description of the field.
	description string
	// allow is a slice of interface that holds the allowed values for the field.
	allow Array
	// only is a bool that determines if the field should only allow the values specified in the allow slice.
	only bool
	// deny is a slice of interface that holds the disallowed values for the field.
	deny Array
	// valid is a slice of interface that holds the only allowed values for the field.
	valid Array
	// messages is a map of strings that holds custom error messages for validators.
	messages map[string]string
}

// Source exposes the root vibranium type.
func (a *Any) Source() Wakanda {
	return a.source
}

// Tribe describes the type of vibranium being used.
func (a *Any) Tribe() string {
	return AnyTribe
}

// Required marks the field as required. Defaults to false.
func (a *Any) Required() *Any {
	a.required = true
	return a
}

// Optional marks the field as optional. This is the default behavior.
func (a *Any) Optional() *Any {
	a.required = false
	return a
}

// Forbidden marks the field as forbidden. This means that the field must not be included in the incoming data. Defaults to false.
func (a *Any) Forbidden() *Any {
	a.forbidden = true
	return a
}

// Description adds a description to the field.
func (a *Any) Description(description string) *Any {
	a.description = description
	return a
}

// Allow sets the allowed values for the field. These values are validated against the incoming data before any other validation rules are applied. These values are in addition to any other permitted values. To create an exclusive list of values, use the Valid() method.
func (a *Any) Allow(allow ...interface{}) *Any {
	a.allow = append(a.allow, allow...)
	return a
}

// Only sets the only option to true. This means that the field should only allow the values specified in the allow slice. Defaults to false.
// Alternatively, you can use the Valid() method to achieve the same result.
func (a *Any) Only() *Any {
	a.only = true
	return a
}

// Deny sets the denied values for the field.
func (a *Any) Deny(deny ...interface{}) *Any {
	a.deny = append(a.deny, deny...)
	return a
}

// Valid sets the only allowed values for the field.
func (a *Any) Valid(valid ...interface{}) *Any {
	a.valid = append(a.valid, valid...)
	return a
}

// Messages sets the custom error messages for different validation rules and terminates the rule chain.
func (a *Any) Messages(messages Messages) {
	a.messages = messages
}

// Validate summons the power of vibranium to shield your application from invalidity present in the incoming data using the current vibranium type.
func (a *Any) Validate(value interface{}, options Plane) error {

	// error to be returned
	err := Kill{
		Details: []Monger{},
	}

	// check if the field is forbidden
	if a.forbidden {
		err.Details = append(err.Details, Monger{

			Message: getMessages(Herb{
				Messages: a.messages,
				Tribe:    AnyTribe,
				Rule:     "unknown",
				Field:    "value",
			}),
			Path:  "",
			Value: value,
			Kind:  a.Tribe() + ".forbidden",
		})
		if options.AbortEarly {
			return err
		}
	}

	// check if the field is required
	if a.required {
		if value == nil {
			err.Details = append(err.Details, Monger{
				Message: getMessages(Herb{
					Messages: a.messages,
					Tribe:    AnyTribe,
					Rule:     "required",
					Field:    "value",
				}),
				Path:  "",
				Value: value,
				Kind:  "required",
			})
			if options.AbortEarly {
				return err
			}
		}
	}

	// check if the field has an only rule set
	if a.only {
		// check if the value is in the allow slice
		if !a.allow.Includes(value) {
			err.Details = append(err.Details, Monger{
				Message: getMessages(Herb{
					Messages: a.messages,
					Tribe:    AnyTribe,
					Rule:     "only",
					Field:    "value",
					Valid:    a.allow,
				}),
				Path:  "",
				Value: value,
				Kind:  "only",
			})
			if options.AbortEarly {
				return err
			}
		}
	}

	// return error, if any
	if err.Error() != "" {
		return err
	}

	return nil
}
