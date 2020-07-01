package service

import repo "mall/internal/pkg/repository"

type CRUDService struct {
	Repo *repo.GormRepository
}

func (s *CRUDService) Exist(conditions interface{}) (bool, error) {
	return s.Repo.Exist(conditions)
}
