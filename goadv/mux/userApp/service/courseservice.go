package service

import ("pass/model"
repo"pass/repository"
"github.com/jinzhu/gorm"
uuid"github.com/satori/go.uuid"
"github.com/rs/zerolog"
"fmt")

type CourseService struct {
	Repo repo.Repository
	Logger *zerolog.Logger
	DB *gorm.DB

}

func NewCourseService(Repo repo.Repository,logger *zerolog.Logger,DB *gorm.DB) *CourseService {
	return &CourseService{
		Repo: Repo,
		Logger :logger,
		DB:DB,
	}
}

func (c *CourseService) AddCourse(course *model.Course) error {
	uow:=repo.NewUnitOfWork(c.DB,false)
	err := c.Repo.Add(uow, course)
	if err != nil {
		return err
	}
	err2:=c.DB.Debug().Model(course).Association("Users").Error
	if err2!=nil{
		fmt.Println("error in association------>",err2)
	}
	c.Logger.Info().Msg("add courses")
	uow.Commit()
	return nil
	
}
func (c *CourseService) GetCourseFromName(out *model.Course, name string) (string,error) {
	uow:=repo.NewUnitOfWork(c.DB,true)
	var queryp []repo.QueryProcessor
	 var preload []string
	queryp = append(queryp, repo.PreloadAssociations(preload))
	queryp = append(queryp, repo.Filter("course_name=?", name))
	err := c.Repo.GetFirst(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return "",err
	}
	uow.Commit()
	c.Logger.Info().Msg("Get course from name ")
	return out.CourseName,nil
}
func (c *CourseService) GetCourseFromNamePrize(out *model.Course, name string,prize int ) (string,int,error) {
	uow:=repo.NewUnitOfWork(c.DB,true)
	var queryp []repo.QueryProcessor
	 var preload []string
	queryp = append(queryp, repo.PreloadAssociations(preload))
	queryp = append(queryp, repo.Filter("course_name=?", name))
	queryp = append(queryp, repo.Filter("prize=?", prize))
	err := c.Repo.GetFirst(uow, out, queryp)
	if err != nil {
		uow.Complete()
		fmt.Println(err)
		return "",-1,err
	}
	uow.Commit()
	c.Logger.Info().Msg("Get course from name ")
	return out.CourseName,out.Prize,nil
}


func(c *CourseService) GetAllCourses(out *[]model.Course,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(c.DB,true)

	var queryp [] repo.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repo.Paginate(limit, offset, &count))
	}
	//fmt.Println(count)
	err := c.Repo.GetAllTenant(uow, out, queryp)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (c *CourseService) GetCourseFromId(out *model.Course, ID uuid.UUID, preloadAssociations []string)error  {
	uow:=repo.NewUnitOfWork(c.DB,true)
	err:=c.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	c.Logger.Info().Msg("get courses from id ")
	uow.Commit()
	return nil
}

func (c *CourseService) UpdateCourse(entity model.Course) error{
	uow:=repo.NewUnitOfWork(c.DB,false)
	err:=c.Repo.Update(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	c.Logger.Info().Msg("update courses")
	uow.Commit()
	return nil
}



func (c *CourseService) DeleteCourse(entity model.Course)error  {
	uow:=repo.NewUnitOfWork(c.DB,false)
	err1:=c.Repo.Delete(uow,entity)
	if err1!=nil{
		uow.Complete()
		return err1
	}
	c.Logger.Info().Msg("delete courses")
	uow.Commit()
	return nil
}
func (c*CourseService)CheckCourse(id uuid.UUID)bool {
	var courses model.Course
	var str1 []string
	err1:=(c.GetCourseFromId(&courses,id,str1))
	if err1!=nil{
		return false
	}
	return true
}

