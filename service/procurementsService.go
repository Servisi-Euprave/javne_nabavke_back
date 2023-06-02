package service

import (
	"javne_nabavke_back/model"
	"javne_nabavke_back/repository"
	"log"
)

type ProcurementService struct {
	l         *log.Logger
	repo      repository.INabavkaRepo
	repoOffer repository.IOfferRepo
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

	plans, err := s.repo.GetProcurementPlans(companyPiB)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (s *ProcurementService) GetCompanyProcurements(companyPiB string) ([]*model.Procurement, error) {
	s.l.Println("Procurement Service - Get company procurements")

	proc, err := s.repo.GetCompanyProcurements(companyPiB)
	if err != nil {
		return nil, err
	}
	return proc, nil

}
func (s *ProcurementService) DeclareWinner(companyPiB string, id string) error {
	s.l.Println("Service Procurement - Declare winner")

	err := s.repo.DeclareWinner(companyPiB, id)
	if err != nil {
		return err
	}
	return nil

}

func (s *ProcurementService) GetProcurementAndWinningOffer() ([]*model.ProcurementWithWinnerOffer, error) {
	s.l.Println("Service Procurement - Get procurements and winning offer")

	procurementsWithWinner, err := s.repo.GetWinnerWithProc()
	if err != nil {
		return nil, err
	}
	return procurementsWithWinner, nil
}

func (s *ProcurementService) GetAllFinishedProcurements() ([]*model.Procurement, error) {
	s.l.Println("Service Procurement - Get procurements and winning offer")
	procurementsWithWinner, err := s.repo.GetAllProcurements()
	if err != nil {
		return nil, err
	}
	return procurementsWithWinner, nil

}
