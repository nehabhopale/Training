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

func (p *PassportService) GetPassportByUserId(db *gorm.DB,out *model.Passport, userId uuid.UUID) error {
	uow:=repo.NewUnitOfWork(db,true)
	var queryp []repo.QueryProcessor
	queryp = append(queryp, repo.Filter("id=?", userId))
	err := p.Repo.GetFirst(uow, out, queryp)
	fmt.Println(out)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}