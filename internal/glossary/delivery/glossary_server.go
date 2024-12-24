package delivery

import (
	"context"
	"fmt"
	"github.com/arvaliullin/go-restful-service/internal/domain"
	"github.com/arvaliullin/go-restful-service/internal/glossary/usecase"
	"github.com/arvaliullin/go-restful-service/pkg/glossary"
)

// GlossaryServer реализует интерфейс gRPC сервиса
type GlossaryServer struct {
	glossary.UnimplementedGlossaryServiceServer
	usecase *usecase.GlossaryUsecase
}

// NewGlossaryServer создает новый gRPC сервер
func NewGlossaryServer(uc *usecase.GlossaryUsecase) *GlossaryServer {
	return &GlossaryServer{usecase: uc}
}

func (s *GlossaryServer) GetAllTerms(ctx context.Context, in *glossary.Empty) (*glossary.GlossaryTerms, error) {
	terms, err := s.usecase.GetAllTerms()
	if err != nil {
		return nil, err
	}

	grpcTerms := make([]*glossary.GlossaryTerm, len(terms))
	for i, t := range terms {
		grpcTerms[i] = &glossary.GlossaryTerm{
			Id:          t.ID,
			Term:        t.Term,
			Description: t.Description,
		}
	}
	return &glossary.GlossaryTerms{Terms: grpcTerms}, nil
}

func (s *GlossaryServer) GetTermByTerm(ctx context.Context, in *glossary.GlossaryTermRequest) (*glossary.GlossaryTerm, error) {
	term, err := s.usecase.GetTerm(in.Term)
	if err != nil {
		return nil, fmt.Errorf("термин не найден: %v", err)
	}
	return &glossary.GlossaryTerm{
		Id:          term.ID,
		Term:        term.Term,
		Description: term.Description,
	}, nil
}

func (s *GlossaryServer) CreateTerm(ctx context.Context, in *glossary.GlossaryTerm) (*glossary.GlossaryTerm, error) {
	domainTerm := domain.GlossaryTerm{
		ID:          in.Id,
		Term:        in.Term,
		Description: in.Description,
	}
	if err := s.usecase.AddTerm(domainTerm); err != nil {
		return nil, fmt.Errorf("ошибка при создании термина: %v", err)
	}
	return in, nil
}

func (s *GlossaryServer) UpdateTerm(ctx context.Context, in *glossary.GlossaryTerm) (*glossary.GlossaryTerm, error) {
	domainTerm := domain.GlossaryTerm{
		ID:          in.Id,
		Term:        in.Term,
		Description: in.Description,
	}
	if err := s.usecase.UpdateTerm(domainTerm); err != nil {
		return nil, fmt.Errorf("ошибка при обновлении термина: %v", err)
	}
	return in, nil
}

func (s *GlossaryServer) DeleteTerm(ctx context.Context, in *glossary.GlossaryTermRequest) (*glossary.Status, error) {
	if err := s.usecase.DeleteTerm(in.Term); err != nil {
		return nil, fmt.Errorf("ошибка при удалении термина: %v", err)
	}
	return &glossary.Status{Message: "Термин успешно удален"}, nil
}
