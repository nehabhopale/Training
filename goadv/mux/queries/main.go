package main
import ("github.com/jinzhu/gorm"
"fmt"
_"github.com/jinzhu/gorm/dialects/mysql"
repo"query/repository"
"query/query"
)
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/exercise?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	repo1 := repo.NewRepository()
	emp:=query.NewEmp(repo1,db)
	emp.GetEmp()
	emp.GetSum()
	emp.GetAvg()
	emp.GetValue()
	emp.GetDepHeadCount()
	emp.GetJobHeadCount()
	emp.GetDepJobHeadCount()
	emp.GetDeptEmpWithCondition()
	emp.GetName()
	emp.GetCountByDname()
	 emp.GetCountByDJname()
	emp.GetAllDepEmp()
	emp.GetDepWithNoEmp()
	emp.DispEmpBoss()
	emp.DispAllEmpBoss()
	 emp.DispAllName()
	emp.GetRegionsNoCountry()
	emp.GetCountryNoState()
	 emp.DispRCSName()
	emp.GetData()
	emp.InsertSwabhavLocation()
	emp.InsertFoo()

}