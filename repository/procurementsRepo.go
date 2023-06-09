package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
	"time"
)

type ProcurementRepository struct {
	l  *log.Logger
	db *gorm.DB
}

func CreateProcurementRepository(logger *log.Logger, dbConnection *gorm.DB) *ProcurementRepository {
	repository := &ProcurementRepository{logger, dbConnection}
	return repository
}

func (n *ProcurementRepository) InsertProcurement(procurement *model.Procurement) error {
	n.l.Println("Procurement_Repo: Insert Procurement")

	uuid := uuid.NewV4().String()
	startDateStr := time.Now().Format("2006-01-02")
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		n.l.Println("Unable to parse date.", err)
	}
	procurement.StartDate = startDate
	procurement.Id = uuid
	procurement.WinnerId = "empty"
	n.l.Println(procurement.ProcurementPlanId)
	n.l.Println("Procurement_Repository_Postgres")
	createdProcurement := n.db.Create(procurement)
	errMessage := createdProcurement.Error
	if errMessage != nil {
		n.l.Println("Unable to Create procurement.", errMessage)
		return CannotCreateError("Procurement")
	}
	return nil
}
func (n *ProcurementRepository) InsertProcurementPlan(procurementPlan *model.ProcurementPlan) error {
	n.l.Println("Procurement_Plan_Repository_Postgres: Insert Procurement Plan")
	uuid := uuid.NewV4().String()

	// Set the UUID as the ProcurementPlanId field
	procurementPlan.ProcurementPlanId = uuid
	createdProcurementPlan := n.db.Create(procurementPlan)
	var errMessage = createdProcurementPlan.Error
	if createdProcurementPlan.Error != nil {
		n.l.Println("Unable to Create procurement plan.", errMessage)
		return CannotCreateError("Procurement plan")
	}
	return nil
}
func (n *ProcurementRepository) GetProcurements() ([]*model.Procurement, error) {
	n.l.Println("Procurement_repo: get procurements with end date matching")

	var procurements []*model.Procurement
	//now := time.Now()
	//lastMonthStart := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.Local)
	//lastMonthEnd := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	// Retrieve all procurements from the "procurements" table
	//now.Format("2006-01-02"),
	if err := n.db.Table("procurements").Order("start_date DESC").Where("winner_id = ?", "empty").Find(&procurements).Error; err != nil {
		return nil, err
	}

	return procurements, nil
}
func (n *ProcurementRepository) GetCompanyProcurements(procuringEntityPiB string) ([]*model.Procurement, error) {
	n.l.Println("Procurement_repo: get company procurements")

	var procurements []*model.Procurement
	if err := n.db.Table("procurements").Order("start_date DESC").Where("procuring_entity_pi_b =  ?", procuringEntityPiB).Find(&procurements).Error; err != nil {
		return nil, err
	}
	return procurements, nil
}
func (n *ProcurementRepository) GetProcurementPlans(companyPiB string) ([]*model.ProcurementPlan, error) {
	n.l.Println("Procurement_repo: get procurement plans")

	var procurementsPlans []*model.ProcurementPlan

	if err := n.db.Table("procurement_plans").Where("company_pib = ?", companyPiB).Find(&procurementsPlans).Error; err != nil {
		return nil, err
	}
	log.Println(procurementsPlans)

	return procurementsPlans, nil
}
func (n *ProcurementRepository) GetAllProcurements() ([]*model.Procurement, error) {
	n.l.Println("Procurement_repo: get All procurements")

	var procurements []*model.Procurement
	if err := n.db.Table("procurements").Order("start_date DESC").Find(&procurements).Error; err != nil {
		return nil, err
	}
	return procurements, nil
}

func (n *ProcurementRepository) DeclareWinner(procId string, offerId string) error {
	n.l.Println("Procurement_repo: declare winner")

	if err := n.db.Table("procurements").Where("id = ?", procId).Update("winner_id", offerId).Error; err != nil {
		return err
	}
	return nil
}
func (n *ProcurementRepository) GetWinnerWithProc() ([]*model.ProcurementWithWinnerOffer, error) {
	n.l.Println("Procurement_repo: get All procurements")

	procurements := []*model.ProcurementWithWinnerOffer{}
	if err := n.db.Table("procurements").Select("procurements.procuring_entity_pi_b," +
		" procurements.start_date,procurements.end_date,procurements.procurement_name,procurements.description," +
		" offers.price,offers.bidder_pib, offers.term_and_payment").Joins("inner join offers on offers.id = procurements.winner_id").Scan(&procurements).Error; err != nil {
		return nil, err
	}
	return procurements, nil
}
