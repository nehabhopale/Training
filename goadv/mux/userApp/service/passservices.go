package service

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
	DB *gorm.DB
}

func NewPassportService(Repo repo.Repository,logger *zerolog.Logger,DB *gorm.DB) *PassportService {
	return &PassportService{
		Repo: Repo,
		Logger:logger,
		DB:DB,
	}
}
func(p *PassportService) GetAllPassports(out *[]model.Passport,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(p.DB,true)

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
func(p *PassportService) GetPassports(out *[]model.Passport, preloadAssociations []string) error{
	uow:=repo.NewUnitOfWork(p.DB,true)
	err := p.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		return  err
	}
	p.Logger.Info().Msg("get all passports ")
	uow.Commit()
	
	return nil
	
}
func (p *PassportService)GetPassportFromId( out interface{}, ID uuid.UUID, preloadAssociations []string)error  {
	uow:=repo.NewUnitOfWork(p.DB,true)
	err:=p.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	p.Logger.Info().Msg("get passports from its id ")
	uow.Commit()
	return nil

}

func (p *PassportService) UpdatePassport(entity model.Passport) error{ //becaz db.model(&User{})
	uow:=repo.NewUnitOfWork(p.DB,false)
	var queryp []repo.QueryProcessor
	queryp = append(queryp, repo.Filter("id=?", entity.ID))
	err:=p.Repo.GetFirst(uow, &entity, queryp)
	if err!=nil{
		fmt.Println("passport to be ted is not found")
		return err 
	}
	err1:=p.Repo.Update(uow,entity)
	if err1!=nil{
		uow.Complete()
		return err1
	}
	p.Logger.Info().Msg("update passports ")
	uow.Commit()
	return nil

}
func (p *PassportService) GetPassportByUserId(out *model.Passport, userId uuid.UUID) error {
	uow:=repo.NewUnitOfWork(p.DB,true)
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
func (p *PassportService) DeletePassport(passId uuid.UUID) error {
	uow := repo.NewUnitOfWork(p.DB, false)
	DeletedPassport := model.Passport{Base: model.Base{ID: passId}}
	err := p.Repo.Delete(uow, &DeletedPassport)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}