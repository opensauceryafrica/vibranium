package vibranium

import (
	"bytes"
	"text/template"
)

// Kill is an error type that holds references to errors that occur during a vabranium validation.
type Kill struct {
	// Errors is a slice of errors that occur during a vibranium validation.
	Details []Monger
}

// Monger is an error type that holds references to errors that occur during a vabranium validation.
type Monger struct {
	// Message is a string that describes the error that occurred during a vibranium validation.
	Message string
	// Path is a string that describes the path to the field that caused the error.
	Path string
	// Value is an interface that holds the value that caused the error.
	Value interface{}
	// Kind is a string that describes the rule set that raised the error that occurred during a vibranium validation.
	Kind string
}

// Error returns the first error recorded during a vibranium validation, if any.
func (k Kill) Error() string {
	if len(k.Details) > 0 {
		return k.Details[0].Message
	}
	return ""
}

// KillMonger is map of strings that holds custom error messages for vabranium rules.
var KillMonger = Messages{
	// Any
	AnyTribe + ".custom":   "{{ .Field }} failed custom validation because {{ .error.message }}",
	AnyTribe + ".default":  "{{ .Field }} threw an error when running default method",
	AnyTribe + ".failover": "{{ .Field }} threw an error when running failover method",
	AnyTribe + ".invalid":  "{{ .Field }} contains an invalid value",
	AnyTribe + ".only":     `{{ .Field }} must be {if( .Valid == 1, "", "one of ")}{{ .Valid }}`,
	// AnyTribe + ".ref":      "{{ .Field }} {{ .Arg }} references {{ :.Ref }} which {{ .Reason }}",
	AnyTribe + ".required": "{{ .Field }} is required",
	AnyTribe + ".unknown":  "{{ .Field }} is not allowed",
}

// getMessages returns the custom error message for a given rule by first checking the user supplied messages and then checking the default messages.
// It formats the error message using the field name and the rule name.
func getMessages(herb Herb) string {
	if herb.Messages != nil {
		if message, ok := herb.Messages[herb.Tribe+"."+herb.Rule]; ok {
			herb.Message = message
			return formatMessage(herb)
		}
	}
	if message, ok := KillMonger[herb.Tribe+"."+herb.Rule]; ok {
		herb.Message = message
		return formatMessage(herb)
	}
	return ""
}

type Herb struct {
	// Messages is a map of strings that holds custom error messages for vabranium rules as provided by the weilder.
	Messages Messages
	// Message is a string that describes the error that occurred during a vibranium validation.
	Message string
	// Tribe is a string that describes the type of vibranium that raised the error that occurred during a vibranium validation.
	Tribe string
	// Field is a string that describes the path to the field that caused the error.
	Field string
	// Rule is a string that describes the rule set that raised the error that occurred during a vibranium validation.
	Rule string
	// Valid is a slice of strings that holds the allowed values for the field.
	Valid Array
}

// formatMessage formats the error message using the field name and the rule name.
func formatMessage(replacement Herb) string {
	templ, err := template.New("message").Parse(replacement.Message)
	if err != nil {
		return ""
	}
	var message bytes.Buffer
	err = templ.Execute(&message, replacement)
	if err != nil {
		return ""
	}
	return message.String()
}
