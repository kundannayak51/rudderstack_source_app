package sourcetemplate

import (
	"context"
	"errors"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/source_template"
	"github.com/rudderstack_source_app/repository/user"
)

type SourceTemplateService struct {
	sourceTemplateRepo source_template.SourceTemplateRepo
	userRepo           user.UserRepo
}

func NewService(sourceTemplateRepo source_template.SourceTemplateRepo, userRepo user.UserRepo) *SourceTemplateService {
	return &SourceTemplateService{
		sourceTemplateRepo: sourceTemplateRepo,
		userRepo:           userRepo,
	}
}

func (s *SourceTemplateService) AddSourceTemplate(ctx context.Context, template entity.SourceTemplate, userId int64) (string, error) {
	//Check if user is Admin
	user, err := s.userRepo.GetUserByUserId(ctx, userId)
	if err == nil {
		return "", err
	}
	if !user.IsAdmin {
		return "", errors.New("Admin can only add Source Template")
	}

	// Check if template type already exists
	if _, err := s.sourceTemplateRepo.GetSourceTemplateByType(ctx, template.Type); err == nil {
		return "", errors.New("template type already exists")
	}

	// Insert new template into database
	id, err := s.sourceTemplateRepo.InsertSourceTemplate(ctx, &template)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *SourceTemplateService) GetSourceTemplateByType(ctx context.Context, templateType string) (*entity.SourceTemplate, error) {
	return s.sourceTemplateRepo.GetSourceTemplateByType(ctx, templateType)
}
