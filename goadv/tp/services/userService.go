package services

import ( 
	_"github.com/jinzhu/gorm/dialects/mysql"
	"test/model"
	repo"test/repo"
	uuid"github.com/satori/go.uuid")

type UserService struct {
	Repo repo.Repository
	Uow  *repo.UnitOfWork
}
func NewUser(Repo repo.Repository, Uow  *repo.UnitOfWork)*UserService{
	return &UserService{
		Repo:Repo,
		Uow:Uow,
	}
}
func (s *UserService)CreateUser()error{
	c:=model.Course{CourseName:"golang",CourseID:uuid.NewV1()}
	var c1 =[]model.Course{c}
	h1:=model.Hobby{HobbyID:uuid.NewV1(),HobbyName:"sports"}
	h2:=model.Hobby{HobbyID:uuid.NewV1(),HobbyName:"cooking"}
	var hobbies =[]model.Hobby{h1,h2}
	user:=model.User{
		ID:uuid.NewV1(),
		UserName:"neha",
		Hobbies:hobbies,
		Courses:c1,
	}
	return s.Repo.Add(s.Uow,user)
}

func (s *UserService)GetUser()error{
	var user model.User
	id1,_:=uuid.FromString("cdc3f49c-5b6c-11ec-bcdd-00155d3b1bb3")
	var str1=[]string{"courses","hobbies"}
	return s.Repo.Get(s.Uow,&user,id1,str1)
}
func (s *UserService)GetAllUser()error{
	var user []model.User
	var str1=[]string{"courses","hobbies"}
	return s.Repo.GetAll(s.Uow,&user,str1)

}