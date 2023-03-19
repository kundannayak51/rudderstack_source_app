package source

import (
	"context"
	"errors"
	"fmt"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/source"
	"github.com/rudderstack_source_app/repository/source_template"
	"github.com/rudderstack_source_app/repository/user"
)

type SourceService struct {
	sourceRepo         source.SourceRepo
	sourceTemplateRepo source_template.SourceTemplateRepo
	userRepo           user.UserRepo
}

func NewService(sourceRepo source.SourceRepo, sourceTemplateRepo source_template.SourceTemplateRepo, userRepo user.UserRepo) *SourceService {
	return &SourceService{
		sourceRepo:         sourceRepo,
		sourceTemplateRepo: sourceTemplateRepo,
		userRepo:           userRepo,
	}
}

func (s *SourceService) CreateSource(ctx context.Context, source entity.Source) error {
	_, err := s.userRepo.GetUserByUserId(ctx, source.UserId)
	if err != nil {
		return errors.New(fmt.Sprintf("User: %v not found", source.UserId))
	}
	return s.sourceRepo.CreateSource(ctx, &source)
}

func (s *SourceService) GetAllSources(ctx context.Context, sourceType string) (*[]entity.Source, error) {
	_, err := s.sourceTemplateRepo.GetSourceTemplateByType(ctx, sourceType)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Source Template of %v type is not present", sourceType))
	}
	sources, err := s.sourceRepo.GetAllSourcesByType(ctx, sourceType)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error in fetching sources from DB: %v", err.Error()))
	}
	return sources, nil
}
