package collections

import "errors"

// Error for when a call a Sequence Get or Update uses an index outside
// the bounds of the sequence
var ErrIndexOutOfRange = errors.New("index out of range")

// Error for when we attempt to create a vector larger than the
// maximum vector sizes
var ErrVectorTooLarge = errors.New("vector to large")

// Error for when we attempt to access 'current' on an iterator that
// has already been exhausted.
var ErrIterationOutOfRange = errors.New("iteration out of range")

// Error for when NewRange is called with a end value lower than the start value
var ErrInvalidRangeBounds = errors.New("invalid range bounds, end less than start")

// Error for when a negative integer is passed to Skip
var ErrInvalidSkipArgument = errors.New("count in Skip must be non-negative")

// Error for when a negative integer is passed to take
var ErrInvalidTakeArgument = errors.New("count in Skip must be non-negative")
