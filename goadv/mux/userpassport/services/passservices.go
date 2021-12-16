package services

import (
	"fmt"
	repo"pass/repository"

	"pass/model"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PassportService struct {
	Repo repo.Repository

}

func NewPassportService(Repo repo.Repository) *PassportService {
	return &PassportService{
		Repo: Repo,
		
	}
}

func(p *PassportService) GetPassports(db *gorm.DB, out *[]model.Passport, preloadAssociations []string) {
	uow:=repo.NewUnitOfWork(db,true)
	err := p.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		fmt.Println("Error in getting all user ---->", err)
	}
	uow.Commit()
	
}
func (p *PassportService)GetPassportFromId(db *gorm.DB, out interface{}, ID uuid.UUID, preloadAssociations []string)  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=p.Repo.Get(uow,out,ID,preloadAssociations,"pass_id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		fmt.Println("error while getting user from id---->",err)
	}
	uow.Commit()

}

func (p *PassportService) UpdatePassport(db *gorm.DB, entity model.Passport) { //becaz db.model(&User{})
	uow:=repo.NewUnitOfWork(db,false)
	err:=p.Repo.Update(uow,entity)
	if err!=nil{
		fmt.Println("err while updating passport--->",err)
		uow.Complete()
	}
	uow.Commit()


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