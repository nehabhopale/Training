package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
)
//user should be able to access courses 
type Course struct{
	Base
	Users []User `gorm:"many2many:user_courses;"`
	CourseName string

	//CourseID uuid.UUID `gorm:"primary_key;type:varchar(50);"`
}
func NewCourse(Name string) *Course{
	return &Course{
		CourseName:Name,
		// Base:Base{CreateBy:"neha",ID:uuid.NewV4(),CreateAt:time.Now()},
	}
}