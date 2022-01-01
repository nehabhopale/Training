package query

import(repo"query/repository"
"query/model"
"fmt"
"github.com/jinzhu/gorm")
type Emp struct{
	Repo repo.Repository
	DB *gorm.DB
}
func NewEmp(Repo repo.Repository,DB *gorm.DB ) *Emp{
	return &Emp{
		Repo:Repo,
		DB:DB,
	}
}
func (e *Emp)GetEmp(){
	fmt.Println("display no of employees")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","count(*) as COUNT")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("no of employees",emp.COUNT)
}
func(e *Emp)GetSum(){
	fmt.Println("display sum of salaries of employees")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","sum(sal) as SUM")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("sum of salaries of employees",emp.SUM)
}
func(e *Emp)GetAvg(){
	fmt.Println("display avg salaries of employees")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","avg(sal) as AVG")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("avg salaries of employees",emp.AVG)

}
func(e *Emp)GetValue(){
	fmt.Println("display sum,avg,count of employees")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,sum(sal) as SUM,avg(sal) as AVG"))
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("no of employees",emp.COUNT)
	fmt.Println("sum of salaries of employees",emp.SUM)
	fmt.Println("avg salaries of employees",emp.AVG)

}
func (e *Emp)GetDepHeadCount(){
	fmt.Println("display the dept wise , headcount")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,deptno as DEPTNO"))
	queryp=append(queryp,repo.Group("DEPTNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("head count is",emp.COUNT,"for dept no",emp.DEPTNO)
	}
	
}
func (e *Emp)GetJobHeadCount(){
	fmt.Println("display the jobwise headcount")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,job as JOB"))
	queryp=append(queryp,repo.Group("JOB"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("head count is",emp.COUNT,"for job",emp.JOB)
	}
}
func (e *Emp)GetDepJobHeadCount(){
	fmt.Println("display dept wise ,jobwise head count")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,deptno as DEPTNO,job as JOB"))
	queryp=append(queryp,repo.Group("DEPTNO,JOB"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("head count is",emp.COUNT,"for dept no",emp.DEPTNO,"for job",emp.JOB)
	}
}

func(e *Emp)GetDeptEmpWithCondition(){
	fmt.Println("display the deptwise employees whose count greater than 2 and who are in dept 10 ,20.Sorty the result by descending order of count")
	//select COUNT(*) as c ,deptnofrom emp WHERE DEPTNO = 10 OR DEPTNO = 20 group by DEPTNO HAVING COUNT(*) >= 2 ORDER BY COUNT(*)Desc
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,deptno as DEPTNO"))
	queryp=append(queryp,repo.Filter("DEPTNO = ? OR DEPTNO = ?",10,20))
	queryp=append(queryp,repo.Group("DEPTNO"))
	queryp=append(queryp,repo.Having("COUNT(*) >= 2"))
	queryp=append(queryp,repo.Order("COUNT(*) desc",true))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("count is",emp.COUNT,"for dept no",emp.DEPTNO)
	}

}
func(e *Emp)GetName(){
	fmt.Println("display ename,deptname")
	//SELECT ename,dname,job from dept left join emp on emp.DEPTNO = dept.DEPTNO;
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("dept","employees.ename as EMPNAME,dname as DEPNAME"))
	queryp=append(queryp,repo.Joins("left join employees on employees.DEPTNO = dept.DEPTNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		if emp.EMPNAME==""{
			emp.EMPNAME="null"
		}
		if emp.DEPNAME==""{
			emp.DEPNAME="null"
		}
		fmt.Println("empname is->",emp.EMPNAME,"having department->",emp.DEPNAME)
	}
}
func(e *Emp)GetCountByDname(){
	fmt.Println("display the deptname wise count")
	//SELECT count(*),dname from emp inner join dept on emp.DEPTNO = dept.DEPTNO group by dname;
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,dept.dname as DEPNAME"))
	queryp=append(queryp,repo.Joins("Inner join dept on employees.DEPTNO = dept.DEPTNO "))
	queryp=append(queryp,repo.Group("dept.dname"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("head count is",emp.COUNT,"for dept name",emp.DEPNAME)
	}
	
}
func(e *Emp)GetCountByDJname(){
	fmt.Println("display the deptname,jobwise count")
	//SELECT count(*),dname,job from emp inner join dept on emp.DEPTNO = dept.DEPTNO group by dname,job;
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,dept.dname as DEPNAME,job as JOB"))
	queryp=append(queryp,repo.Joins("Inner join dept on employees.DEPTNO = dept.DEPTNO "))
	queryp=append(queryp,repo.Group("dept.dname,job"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("head count is",emp.COUNT,"for dept name",emp.DEPNAME,"for job",emp.JOB)
	}
	
}
func(e *Emp)GetAllDepEmp(){
	fmt.Println("display all the departments , with employees if any (if no emps then display null)")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("dept","employees.ename as EMPNAME,dept.dname as DEPNAME"))
	//queryp=append(queryp,repo.Select("dept","employees.ename ,dept.dname "))
	queryp=append(queryp,repo.Joins("LEFT OUTER JOIN employees ON employees.DEPTNO=dept.DEPTNO"))
	
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		if emp.EMPNAME ==""{
			emp.EMPNAME="null"
		}
		fmt.Println("employee->",emp.EMPNAME,"for dept name",emp.DEPNAME)
	}
}
func(e *Emp)GetDepWithNoEmp(){
	fmt.Println("display the departments where there are no employees")
	//SELECT dname FROM dept left join emp on emp.deptno=dept.deptno where emp.DEPTNO is NUll;
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("dept","dname as DEPNAME"))
	queryp=append(queryp,repo.Joins("left join employees on employees.deptno=dept.deptno"))
	queryp=append(queryp,repo.Filter("employees.DEPTNO is null"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("dept name",emp.DEPNAME)
	}

}
func(e *Emp)DispEmpBoss(){
	fmt.Println("display the empname and their bossnames")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees emp1,employees emp2 ","emp1.ename as EMPNAME ,emp2.ename as BOSSNAME"))
	queryp=append(queryp,repo.Filter("emp1.MGR=emp2.EMPNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("employee->",emp.EMPNAME,"Boss->",emp.BOSSNAME)
	}
}
func(e *Emp) DispAllEmpBoss(){
	fmt.Println(" display all the empnames and boss names if any (if no boss display null)")

	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees emp1","emp1.ename as EMPNAME ,emp2.ename as BOSSNAME"))
	queryp=append(queryp,repo.Joins("LEFT JOIN employees emp2 ON emp1.MGR=emp2.EMPNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		if emp.BOSSNAME==""{
			emp.BOSSNAME="null"
		}
		fmt.Println("employee->",emp.EMPNAME,"Boss->",emp.BOSSNAME)
	}
}
func(e *Emp)DispAllName(){
	fmt.Println(" display ename,deptname and bossname .")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees emp1,employees emp2 ","emp1.ename as EMPNAME ,dept.dname as DEPNAME,emp2.ename as BOSSNAME"))
	queryp=append(queryp,repo.Joins("INNER JOIN DEPT ON emp2.DEPTNO=DEPT.DEPTNO"))
	queryp=append(queryp,repo.Filter("emp1.MGR=emp2.EMPNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("employee->",emp.EMPNAME,"dept>",emp.DEPNAME,"Boss->",emp.BOSSNAME)
	}
}
func (e *Emp)GetRegionsNoCountry(){
	fmt.Println("display the regions there no entry for country")
	//select * from regions left join countries on regions.region_id = countries.region_id where regions.region_name is NULL
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("regions","regions.region_id as RID ,regions.region_name as RNAME"))
	queryp=append(queryp,repo.Joins("left join countries on regions.region_id= countries.region_id "))
	queryp=append(queryp,repo.Filter("regions.region_name is NULL"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("rid->",emp.RID,"name->",emp.RNAME)
	}

}
func (e *Emp)GetCountryNoState(){
	fmt.Println("display the countries there no states")
// 	select * FROM countries join locations on locations.COUNTRY_ID = countries.COUNTRY_ID 
//where locations.STATE_PROVINCE is NULL
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("countries","countries.country_id as CID ,countries.country_name as CNAME,countries.region_id as RID"))
	queryp=append(queryp,repo.Joins("inner join locations on locations.COUNTRY_ID = countries.COUNTRY_ID "))
	queryp=append(queryp,repo.Filter("locations.STATE_PROVINCE is NULL"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("rid->",emp.RID,"cname->",emp.CNAME,"cid->",emp.CID)
	}

}
func(e *Emp)DispRCSName(){
	fmt.Println("display region name,country name and state name")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("countries","countries.country_name as CNAME,locations.state_province as STATEP,regions.region_name as RNAME"))
	queryp=append(queryp,repo.Joins("INNER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID"))
	queryp=append(queryp,repo.Joins("INNER JOIN REGIONS ON COUNTRIES.REGION_ID=REGIONS.REGION_ID"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		if emp.STATEP==""{
			emp.STATEP="null"
		}
		fmt.Println("country name->",emp.CNAME,"state provience->",emp.STATEP,"rname->",emp.RNAME)
	}
}	
func (e *Emp) InsertSwabhavLocation() {
	fmt.Println("Make an insert of swabhav Techlabs in location/state tables map to india and asia.")
	uow := repo.NewUnitOfWork(e.DB, true)
	e.DB.AutoMigrate(&model.Location{})
	Location := model.NewLocation(2, "mum", "8796", "Mumbai", "Maharashtra","IN")
	e.Repo.Add(uow, &Location)
}

func(e *Emp)GetData(){
	fmt.Println("Filter details based on mumbai location")
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("locations","location_id as LID ,street_address as SADD,postal_code as PCODE,city as CITY,state_province as STATEP,country_id as CID"))
	queryp=append(queryp,repo.Filter("CITY=?", "Mumbai"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("location_id->",emp.LID,"street_address->",emp.SADD,"postal_code->",emp.PCODE,"city->",emp.CITY,"state_province->",emp.STATEP,"country_id->",emp.CID)
	}
}

func (e *Emp) InsertFoo() {
	fmt.Println("create a foo table and insert values of different data")
	unit := repo.NewUnitOfWork(e.DB, true)

	e.DB.AutoMigrate(&model.Foo{})
	foo1:=model.NewFoo("neha","lname")
	e.Repo.Add(unit, foo1)
	foo2 :=model.NewFoo("pooja","b")
	e.Repo.Add(unit, foo2)
}
// func(e *Emp) ShowIndex(){
// 	uow := repo.NewUnitOfWork(e.DB, true)
// 	var emp []model.Employee
// 	var queryp []repo.QueryProcessor
// 	queryp=append(queryp,repo.ShowIndex("show index from employees",emp))
// 	e.Repo.GetAll(uow, &emp, queryp)
// 	fmt.Println(emp)
// 	for _,emp:=range(emp){
		
// 		fmt.Println("country name->",emp.CNAME)
// 	}

// }