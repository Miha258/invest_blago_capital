package service

import (
	"invest_blango_criptal_backend/models"
	"invest_blango_criptal_backend/repository"
)



type InvestitionsService struct {
	repo repository.Investitions
}


func NewInvestitionsService(repo repository.Investitions) *InvestitionsService {
	return &InvestitionsService{repo: repo}
}


func (s *InvestitionsService) CreateInvestition(investition models.Investition) (int, error) {
	return s.repo.CreateInvestition(investition)
}