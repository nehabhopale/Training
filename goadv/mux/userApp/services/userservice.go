package services

import ("pass/model"
repo"pass/repository"
"github.com/jinzhu/gorm"
"fmt"
uuid"github.com/satori/go.uuid"
"golang.org/x/crypto/bcrypt"
"github.com/rs/zerolog")


type UserService struct{
	Repo repo.Repository
	DB *gorm.DB
	Logger *zerolog.Logger

}
func NewUserService(Repo repo.Repository,db *gorm.DB,logger *zerolog.Logger ) *UserService{
	return &UserService{
		Repo:Repo,
		Logger:logger,
		DB:db,
	}
}
func (u *UserService)AddUser(db *gorm.DB, user *model.User)error {
	uow:=repo.NewUnitOfWork(db,false)
	var courses []model.Course
	for _, course := range user.Courses {
		var c model.Course
		var queryp []repo.QueryProcessor
		fmt.Println("course name is ",course.CourseName)
		queryp = append(queryp, repo.Filter("course_name=?", course.CourseName))
		u.Repo.GetFirst(uow, &c, queryp)
		courses = append(courses, c)
	}
	user.Courses = courses
	err := u.Repo.Add(uow, user)
	if err != nil {
		uow.Complete()
		return (err)
	}
	u.Logger.Info().Msg("add users ")
	uow.Commit()
	return nil
}

//all users with pagination
func(u *UserService) GetAllUsers(db *gorm.DB, out *[]model.User,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(db,true)

	var queryp [] repo.QueryProcessor
	var count int
	var preload =[]string{"Passport","Courses","Hobbies"}
	queryp = append(queryp, repo.PreloadAssociations(preload))
	if limit != 0 {
		queryp = append(queryp, repo.Paginate(limit, offset, &count))
	}
	fmt.Println(count)
	err := u.Repo.GetAllTenant(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return err
	}
	u.Logger.Info().Msg("Get all users with pagination")
	uow.Commit()
	return nil
}
func(u *UserService) GetUsers(db *gorm.DB, out *[]model.User, preloadAssociations []string)error {
	uow:=repo.NewUnitOfWork(db,true)
	
	err := u.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		return  err
	}
	u.Logger.Info().Msg("Get all users ")
	uow.Commit()
	return nil
	
	
}

func (u *UserService)GetUserFromId(db *gorm.DB, out interface{}, ID uuid.UUID, preloadAssociations []string)error  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=u.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	uow.Commit()
	u.Logger.Info().Msg("Get users from id ")
	return nil

}
func (u *UserService) GetUserFromEmail(out *model.User, email string) error {
	uow:=repo.NewUnitOfWork(u.DB,true)
	var queryp []repo.QueryProcessor
	preload:=[]string{"Passport","Courses"}
	queryp = append(queryp, repo.PreloadAssociations(preload))
	queryp = append(queryp, repo.Filter("email=?", email))
	err := u.Repo.GetFirst(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	u.Logger.Info().Msg("Get  users from email")
	return nil
}

func (u *UserService) GetPasswordFromEmail(email string) (string,bool){
	var user model.User
	err:=u.GetUserFromEmail(&user,email)
	if err!=nil{
		return "" ,false
	}
	return user.Password,true
}
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (u *UserService) UpdateUser(db *gorm.DB, entity model.User) error{ //becaz db.model(&User{})
	uow:=repo.NewUnitOfWork(db,false)
	err:=u.Repo.Update(uow,entity)
	if err!=nil{
		
		uow.Complete()
		return err
	}
	u.Logger.Info().Msg("updating users ")
	uow.Commit()
	return nil

}

func  (u *UserService)DeleteUser(db *gorm.DB, entity model.User) error{
	uow:=repo.NewUnitOfWork(db,false)
	err:=u.Repo.Delete(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil

}

func (U *UserService) GetUsersCount(db *gorm.DB) int {
	uow:=repo.NewUnitOfWork(db,true)
	var users []model.User
	var count int
	db.Debug().Model(&users).Count(&count)
	// if err != nil {
	// 	fmt.Println("count err",err)
	// 	uow.Complete()
	// 	return -1
	// }
	uow.Commit()
	fmt.Println(count)
	return count
}