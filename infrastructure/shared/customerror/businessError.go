package customerror

import "errors"

var (
	ErrRecordNotFound = errors.New("Record not found")
	ErrRangeVal       = errors.New("The field does not supply a correct value range")
)
