package service

import (
	"javne_nabavke_back/model"
	"javne_nabavke_back/repository"
	"log"
)

type ProcurementService struct {
	l    *log.Logger
	repo repository.INabavkaRepo
}

func NewProcurementService(l *log.Logger, repo repository.INabavkaRepo) *ProcurementService {
	return &ProcurementService{
		l:    l,
		repo: repo,
	}
}
func (s *ProcurementService) SaveProcurement(procurement *model.Procurement) error {
	return nil
}
func (s *ProcurementService) SaveProcurementPlan(procurement *model.ProcurementPlan) error {

	return nil
}
func (s *ProcurementService) GetProcurements(procurement model.Procurement) []model.Procurement {
	return nil
}
