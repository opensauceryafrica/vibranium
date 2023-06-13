package vibranium

// Wakanda is the interface that all vibranium types must implement.
type Wakanda interface {
	// Validate summons the power of vibranium to shield your application from invalidity present in the incoming data.
	Validate(interface{}, Plane) error

	// Tribe describes the type of vibranium being used.
	Tribe() string

	// Source exposes the root vibranium type.
	Source() Wakanda
}

// Messages is a map of strings that holds custom error messages for vabranium rules.
type Messages map[string]string

// Plane is a struct that holds the configuration for a vibranium validation.
type Plane struct {
	// AllowUnknown is a bool that determines if unknown fields should be allowed in validations involving object types. Defaults to false.
	AllowUnknown bool
	// AbortEarly is a bool that determines if the validation should stop on the first error. Defaults to true.
	AbortEarly bool
}

const (
	// AnyTribe describes a vibranium type that can be used to validate any type of data.
	AnyTribe = "any"

	// ArrayTribe describes a vibranium type that can be used to validate slices.
	ArrayTribe = "array"

	// ObjectTribe describes a vibranium type that can be used to validate objects.
	ObjectTribe = "object"

	// StringTribe describes a vibranium type that can be used to validate strings.
	StringTribe = "string"

	// NumberTribe describes a vibranium type that can be used to validate integers.
	NumberTribe = "number"

	// BooleanTribe describes a vibranium type that can be used to validate booleans.
	BooleanTribe = "boolean"

	// DateTribe describes a vibranium type that can be used to validate date-time
	DateTribe = "date"
)
