package crondescriptor

import "fmt"

//  Error for cases when something is missing
type missingFielError struct {
	field string
}

func (e missingFielError) Error() string {
	return fmt.Sprintf("field '%v' not found", e.field)
}
func NewmissingFielError(field string) error {
	return missingFielError{
		field: field,
	}
}

// Error for cases when something has wrong format
type formatError error

// Error for cases when wrong argument is passed
type wrongArgumentError error
