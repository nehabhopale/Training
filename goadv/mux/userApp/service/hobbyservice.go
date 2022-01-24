package service

import (
	"pass/model"
	repo "pass/repository"

	"github.com/jinzhu/gorm"

	//"fmt"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

type HobbyService struct {
	Repo   repo.Repository
	Logger *zerolog.Logger
	DB     *gorm.DB
}

func NewHobbyService(Repo repo.Repository, logger *zerolog.Logger, DB *gorm.DB) *HobbyService {
	return &HobbyService{
		Repo:   Repo,
		Logger: logger,
		DB:     DB,
	}
}

func (h *HobbyService) GetHobbies(out *[]model.Hobby, limit int, offset int) error {
	uow := repo.NewUnitOfWork(h.DB, true)

	var queryp []repo.QueryProcessor
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
func (h *HobbyService) GetHobbyByUserId(out *[]model.Hobby, userId uuid.UUID) error {
	unit := repo.NewUnitOfWork(h.DB, true)
	var queryp []repo.QueryProcessor
	queryp = append(queryp, repo.Filter("user_id=?", userId))
	err := h.Repo.Getall(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	h.Logger.Info().Interface("UserId-", userId).Msg("Get hobby by UserId")
	unit.Commit()
	return nil
}
func (h *HobbyService) AddHobby(entity model.Hobby) error {
	unit := repo.NewUnitOfWork(h.DB, false)
	err := h.Repo.Add(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	h.Logger.Info().Interface("hobby-", entity).Msg("Add Hobby")
	unit.Commit()
	return nil
}
func (h *HobbyService) GetHobbyFromId(out *model.Hobby, ID uuid.UUID, preloadAssociations []string) error {
	uow := repo.NewUnitOfWork(h.DB, true)
	err := h.Repo.Get(uow, out, ID, preloadAssociations, "id")
	if err != nil {
		uow.Complete() //complete will rollback operation
		return err
	}
	h.Logger.Info().Msg("Get all hobbies from id ")
	uow.Commit()
	return nil
}

func (h *HobbyService) UpdateHobby(entity model.Hobby) error {
	uow := repo.NewUnitOfWork(h.DB, false)
	err1 := h.Repo.Update(uow, entity)
	if err1 != nil {
		uow.Complete()
		return err1
	}
	h.Logger.Info().Msg("update hobbies")
	uow.Commit()
	return nil
}

func (h *HobbyService) DeleteHobby(entity model.Hobby) error {
	uow := repo.NewUnitOfWork(h.DB, false)
	err1 := h.Repo.Delete(uow, entity)
	if err1 != nil {
		uow.Complete()
		return err1
	}
	h.Logger.Info().Msg("delete hobbies ")
	uow.Commit()
	return nil
}

func (h *HobbyService) CheckHobby(id uuid.UUID) bool {
	var hobbies model.Hobby
	var str1 []string
	err1 := (h.GetHobbyFromId(&hobbies, id, str1))
	if err1 != nil {
		return false
	}
	return true
}
