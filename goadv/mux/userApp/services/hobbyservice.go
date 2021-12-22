package services

import ("pass/model"
repo"pass/repository"
"github.com/jinzhu/gorm"

uuid"github.com/satori/go.uuid"
"github.com/rs/zerolog")

type HobbyService struct {
	Repo repo.Repository
	Logger *zerolog.Logger
	DB *gorm.DB

}

func NewHobbyService(Repo repo.Repository,logger *zerolog.Logger,DB *gorm.DB) *HobbyService {
	return &HobbyService{
		Repo: Repo,
		Logger :logger,
		DB:DB,
	}
}


func(h *HobbyService) GetHobbies(out *[]model.Hobby,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(h.DB,true)

	var queryp [] repo.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repo.Paginate(limit, offset, &count))
	}
	//fmt.Println(count)
	err := h.Repo.GetAllTenant(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return err
	}
	h.Logger.Info().Msg("Get all hobbies with pagination")
	uow.Commit()
	return nil
}

func (h *HobbyService) GetHobbyFromId(out *model.Hobby, ID uuid.UUID, preloadAssociations []string) error {
	uow:=repo.NewUnitOfWork(h.DB,true)
	err:=h.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	h.Logger.Info().Msg("Get all hobbies from id ")
	uow.Commit()
	return nil
}

func (h *HobbyService) UpdateHobby(entity model.Hobby)error {
	uow:=repo.NewUnitOfWork(h.DB,false)
	err:=h.Repo.Update(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	h.Logger.Info().Msg("update hobbies")
	uow.Commit()
	return nil
}

func (h *HobbyService) DeleteHobby(entity model.Hobby) error {
	uow:=repo.NewUnitOfWork(h.DB,false)
	err:=h.Repo.Delete(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	h.Logger.Info().Msg("delete hobbies ")
	uow.Commit()
	return nil 
}

