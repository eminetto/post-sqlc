package person

import "context"

// ID person ID
type ID int

// Person defines a person
type Person struct {
	ID       ID
	Name     string
	LastName string
}

// UseCase defines the domain use case
type UseCase interface {
	Get(ctx context.Context, id ID) (*Person, error)
	Search(ctx context.Context, query string) ([]*Person, error)
	List(ctx context.Context) ([]*Person, error)
	Create(ctx context.Context, firstName, lastName string) (ID, error)
	Update(ctx context.Context, e *Person) error
	Delete(ctx context.Context, id ID) error
}
