package person

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eminetto/post-sqlc/person/db"
)

// Service define a service
type Service struct {
	r *db.Queries
}

// NewService cria um novo serviço. Lembre-se: receba interfaces, retorne structs ;)
func NewService(r *db.Queries) *Service {
	return &Service{
		r: r,
	}
}

// Get a person
func (s *Service) Get(ctx context.Context, id ID) (*Person, error) {
	p, err := s.r.Get(ctx, sql.NullInt32{Int32: int32(id), Valid: true})
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	return &Person{
		ID:       ID(p.ID.Int32),
		Name:     p.FirstName.String,
		LastName: p.LastName.String,
	}, nil
}

// Search person
func (s *Service) Search(ctx context.Context, query string) ([]*Person, error) {
	p, err := s.r.Search(ctx, db.SearchParams{
		FirstName: sql.NullString{
			String: query,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: query,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error searching from database: %w", err)
	}
	var people []*Person
	for _, j := range p {
		people = append(people, &Person{
			ID:       ID(j.ID.Int32),
			Name:     j.FirstName.String,
			LastName: j.LastName.String,
		})
	}
	return people, nil
}

// List person
func (s *Service) List(ctx context.Context) ([]*Person, error) {
	p, err := s.r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var people []*Person
	for _, j := range p {
		people = append(people, &Person{
			ID:       ID(j.ID.Int32),
			Name:     j.FirstName.String,
			LastName: j.LastName.String,
		})
	}
	return people, nil
}

// Create a person
func (s *Service) Create(ctx context.Context, firstName, lastName string) (ID, error) {
	result, err := s.r.Create(ctx, db.CreateParams{
		FirstName: sql.NullString{
			String: firstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: lastName,
			Valid:  true,
		},
	})
	if err != nil {
		return 0, fmt.Errorf("error creating person: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error creating person: %w", err)
	}
	return ID(id), nil
}

// Update person data
func (s *Service) Update(ctx context.Context, e *Person) error {
	err := s.r.Update(ctx, db.UpdateParams{
		FirstName: sql.NullString{
			String: e.Name,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: e.LastName,
			Valid:  true,
		},
		ID: sql.NullInt32{
			Int32: int32(e.ID),
			Valid: true,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating person: %w", err)
	}
	return nil
}

// Delete remove a person
func (s *Service) Delete(ctx context.Context, id ID) error {
	err := s.r.Delete(ctx, sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	})
	if err != nil {
		return fmt.Errorf("error removing person: %w", err)
	}
	return nil
}
