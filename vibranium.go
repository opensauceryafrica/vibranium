/*
	package vibranium

maybe the most versatile validation library for all golang applications.
*/
package vibranium

// Any generates a schema object that matches any data type.
func Any() *AnySchema {
	a := new(AnySchema)
	a.source = a
	return a
}
