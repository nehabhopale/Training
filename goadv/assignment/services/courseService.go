package services

import ("test/model"
repo"test/repo"
"github.com/jinzhu/gorm"
"fmt"
uuid"github.com/satori/go.uuid")

type CourseService struct {
	Repo repo.Repository	
}

func NewCourseService(Repo repo.Repository) *CourseService {
	return &CourseService{
		Repo: Repo,
	}
}

func (c *CourseService) AddCourse(db *gorm.DB, course *model.Course)  {
	uow:=repo.NewUnitOfWork(db,false)
	err := c.Repo.Add(uow, course)
	if err != nil {
		uow.Complete()
		fmt.Println("error while adding course",err)
	}
	uow.Commit()

	
}

func (c *CourseService) GetCourses(db *gorm.DB,	out *[]model.Course, preloadAssociations []string)  {
	uow:=repo.NewUnitOfWork(db,true)
	err := c.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		uow.Complete()
		fmt.Println("Error in getting all Courses ", err)
	}
	uow.Commit()
}

func (c *CourseService) GetCourseFromId(db *gorm.DB,out *model.Course, ID uuid.UUID, preloadAssociations []string)  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=c.Repo.Get(uow,out,ID,preloadAssociations,"course_id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		fmt.Println("error while getting course from id>-->",err)
	}
	uow.Commit()
}

func (c *CourseService) UpdateCourse(db *gorm.DB,entity model.Course) {
	uow:=repo.NewUnitOfWork(db,false)
	err:=c.Repo.Update(uow,entity)
	if err!=nil{
		fmt.Println("err while updating course")
		uow.Complete()
	}
	uow.Commit()
}

func (c *CourseService) DeleteCourse(db *gorm.DB,entity model.Course)  {
	uow:=repo.NewUnitOfWork(db,false)
	err:=c.Repo.Delete(uow,entity)
	if err!=nil{
		fmt.Println("err while deleting course",err)
		uow.Complete()
	}
	uow.Commit()
}

func (s *CourseService) GetCourseFromName(db *gorm.DB,out *model.Course, name string)  {
	
}