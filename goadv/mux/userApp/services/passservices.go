package services

import (
	"fmt"
	repo"pass/repository"

	"pass/model"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
)

type PassportService struct {
	Repo repo.Repository
	Logger *zerolog.Logger
}

func NewPassportService(Repo repo.Repository,logger *zerolog.Logger) *PassportService {
	return &PassportService{
		Repo: Repo,
		Logger:logger,
		
	}
}
func(p *PassportService) GetAllPassports(db *gorm.DB, out *[]model.Passport,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(db,true)

	var queryp [] repo.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repo.Paginate(limit, offset, &count))
	}
	//fmt.Println(count)
	err := p.Repo.GetAllTenant(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return err
	}
	p.Logger.Info().Msg("get all passports with pagination  ")
	uow.Commit()
	return nil
}
func(p *PassportService) GetPassports(db *gorm.DB, out *[]model.Passport, preloadAssociations []string) error{
	uow:=repo.NewUnitOfWork(db,true)
	err := p.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		return  err
	}
	p.Logger.Info().Msg("get all passports ")
	uow.Commit()
	
	return nil
	
}
func (p *PassportService)GetPassportFromId(db *gorm.DB, out interface{}, ID uuid.UUID, preloadAssociations []string)error  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=p.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	p.Logger.Info().Msg("get passports from its id ")
	uow.Commit()
	return nil

}

func (p *PassportService) UpdatePassport(db *gorm.DB, entity model.Passport) error{ //becaz db.model(&User{})
	uow:=repo.NewUnitOfWork(db,false)
	err:=p.Repo.Update(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	p.Logger.Info().Msg("update passports ")
	uow.Commit()
	return nil

}
func (p *PassportService) GetPassportByUserId(db *gorm.DB,out *model.Passport, userId uuid.UUID) error {
	uow:=repo.NewUnitOfWork(db,true)
	var queryp []repo.QueryProcessor
	queryp = append(queryp, repo.Filter("uid=?", userId))
	err := p.Repo.GetFirst(uow, out, queryp)
	fmt.Println(out)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}
func (p *PassportService) DeletePassport(db *gorm.DB,passId uuid.UUID) error {
	unit := repo.NewUnitOfWork(db, false)
	DeletePassport := model.Passport{Base: model.Base{ID: passId}}
	err := p.Repo.Delete(unit, &DeletePassport)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}