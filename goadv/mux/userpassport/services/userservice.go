package services

import ("pass/model"
repo"pass/repository"
"github.com/jinzhu/gorm"
"fmt"
)


type UserService struct{
	Repo repo.Repository
}
func NewUserService(Repo repo.Repository) *UserService{
	return &UserService{
		Repo:Repo,
	}
}
func (u *UserService)AddUser(db *gorm.DB, user *model.User) {
	uow:=repo.NewUnitOfWork(db,false)
	err := u.Repo.Add(uow, user)
	if err != nil {
		uow.Complete()
		fmt.Println(err)
	}
	uow.Commit()
}

func(u *UserService) GetUsers(db *gorm.DB, out *[]model.User, preloadAssociations []string) {
	uow:=repo.NewUnitOfWork(db,true)
	err := u.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		fmt.Println("Error in getting all user ---->", err)
	}
	uow.Commit()
	
}


func (U *UserService) GetUsersCount(db *gorm.DB) int {
	uow:=repo.NewUnitOfWork(db,true)
	var users []model.User
	var count int
	err := db.Debug().Model(&users).Count(&count)
	if err != nil {
		fmt.Println("count err",err)
		uow.Complete()
		return -1
	}
	uow.Commit()
	fmt.Println(count)
	return count
}