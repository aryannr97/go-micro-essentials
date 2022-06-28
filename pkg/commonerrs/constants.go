package commonerrs

import "errors"

var (
	// ErrDataMissing is used when there is no content to render as json response
	ErrDataMissing = errors.New("data not found while preparing response body")
)
