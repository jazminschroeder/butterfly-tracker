package butterflytracker

import(
	"context"
	"errors"
)
//ErrAlreadyExists: Butterfly name already used
var (
	ErrAlreadyExists   = errors.New("already exists")
)

// Service interface
type Service interface{
	Track(ctx context.Context, b Butterfly) error
}

// ButterflyTracker structure
type ButterflyTracker struct{

}

// Track implementation
func(ButterflyTracker) Track(ctx context.Context, b Butterfly) error{
	if b.Name == "Henry" {
		return ErrAlreadyExists
	}
	return nil
}