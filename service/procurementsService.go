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
	s.l.Println("Procurement Service - saveProcurement")
	err := s.repo.InsertProcurement(procurement)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProcurementService) SaveProcurementPlan(procurementPlan *model.ProcurementPlan) error {
	s.l.Println("Procurement Service - saveProcurementPlan")
	err := s.repo.InsertProcurementPlan(procurementPlan)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProcurementService) GetProcurements() ([]*model.Procurement, error) {
	s.l.Println("Procurement Service - Get procurements")

	proc, err := s.repo.GetProcurements()
	if err != nil {
		return nil, err
	}
	return proc, nil
}
func (s *ProcurementService) GetProcurementPlans(companyPiB string) ([]*model.ProcurementPlan, error) {
	s.l.Println("Procurement Service - Get procurement plans with company PiB")

	plans, err := s.GetProcurementPlans(companyPiB)
	if err != nil {
		return nil, err
	}
	return plans, nil
}
