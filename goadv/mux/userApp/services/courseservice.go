package services

import ("pass/model"
repo"pass/repository"
"github.com/jinzhu/gorm"
uuid"github.com/satori/go.uuid"
"github.com/rs/zerolog")

type CourseService struct {
	Repo repo.Repository
	Logger *zerolog.Logger

}

func NewCourseService(Repo repo.Repository,logger *zerolog.Logger) *CourseService {
	return &CourseService{
		Repo: Repo,
		Logger :logger,
	}
}

func (c *CourseService) AddCourse(db *gorm.DB, course *model.Course) error {
	uow:=repo.NewUnitOfWork(db,false)
	err := c.Repo.Add(uow, course)
	if err != nil {
		return err
	}
	c.Logger.Info().Msg("add courses")
	uow.Commit()
	return nil
	
}

func(c *CourseService) GetAllCourses(db *gorm.DB, out *[]model.Course,limit int,offset int )error {
	uow:=repo.NewUnitOfWork(db,true)

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

func (c *CourseService) GetCourseFromId(db *gorm.DB,out *model.Course, ID uuid.UUID, preloadAssociations []string)error  {
	uow:=repo.NewUnitOfWork(db,true)
	err:=c.Repo.Get(uow,out,ID,preloadAssociations,"id")
	if err!=nil{
		uow.Complete()		//complete will rollback operation
		return err
	}
	c.Logger.Info().Msg("get courses from id ")
	uow.Commit()
	return nil
}

func (c *CourseService) UpdateCourse(db *gorm.DB,entity model.Course) error{
	uow:=repo.NewUnitOfWork(db,false)
	err:=c.Repo.Update(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	c.Logger.Info().Msg("update courses")
	uow.Commit()
	return nil
}

func (c *CourseService) DeleteCourse(db *gorm.DB,entity model.Course)error  {
	uow:=repo.NewUnitOfWork(db,false)
	err:=c.Repo.Delete(uow,entity)
	if err!=nil{
		uow.Complete()
		return err
	}
	c.Logger.Info().Msg("delete courses")
	uow.Commit()
	return nil
}

