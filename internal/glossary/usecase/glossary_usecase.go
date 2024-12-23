package usecase

import (
	"github.com/arvaliullin/go-restful-service/internal/domain"
	"github.com/arvaliullin/go-restful-service/internal/glossary/repository"
)

type GlossaryUsecase struct {
	repo *repository.GlossaryRepository
}

func NewGlossaryUsecase(repo *repository.GlossaryRepository) *GlossaryUsecase {
	return &GlossaryUsecase{repo: repo}
}

func (uc *GlossaryUsecase) GetAllTerms() ([]domain.GlossaryTerm, error) {
	return uc.repo.GetAll()
}

func (uc *GlossaryUsecase) GetTerm(term string) (*domain.GlossaryTerm, error) {
	return uc.repo.GetByTerm(term)
}

func (uc *GlossaryUsecase) AddTerm(term domain.GlossaryTerm) error {
	return uc.repo.Create(term)
}

func (uc *GlossaryUsecase) UpdateTerm(term domain.GlossaryTerm) error {
	return uc.repo.Update(term)
}

func (uc *GlossaryUsecase) DeleteTerm(term string) error {
	return uc.repo.Delete(term)
}
