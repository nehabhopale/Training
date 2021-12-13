package services

import ("test/model"
repo"test/repo"
"github.com/jinzhu/gorm"
"fmt"
uuid"github.com/satori/go.uuid")


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
func (u *UserService)GetUser(db *gorm.DB){
	uow:=repo.NewUnitOfWork(db,true)
	query:=repo.Filter("user_name=?","pooja")
	queries:=[]repo.QueryProcessor{}
	queries=append(queries,query)
	reqAssociations:=[]string{"Courses","Hobbies"}
	preAssociations:=repo.PreloadAssociations(reqAssociations)
	queries=append(queries,preAssociations)
	var user model.User
	err:=u.Repo.GetFirst(uow,&user,queries)
	if err!=nil{
		fmt.Println(err)
		uow.Complete()
	}
	uow.Commit()
	

}
func (u *UserService)GetUserFromId(db *gorm.DB, out interface{}, ID uuid.UUID, preloadAssociations []string)  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=u.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		fmt.Println("error while getting user from id---->",err)
	}
	uow.Commit()

}

func (u *UserService) UpdateUser(db *gorm.DB, entity model.User) { //becaz db.model(&User{})
	uow:=repo.NewUnitOfWork(db,false)
	err:=u.Repo.Update(uow,entity)
	if err!=nil{
		fmt.Println("err while updating user--->",err)
		uow.Complete()
	}
	uow.Commit()


}

func  (u *UserService)DeleteUser(db *gorm.DB, entity model.User) {
	uow:=repo.NewUnitOfWork(db,false)
	err:=u.Repo.Delete(uow,entity)
	if err!=nil{
		fmt.Println("err while deleting user")
		uow.Complete()
	}
	uow.Commit()

}

