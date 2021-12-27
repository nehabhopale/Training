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
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","count(*) as COUNT")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("no of employees",emp.COUNT)
}
func(e *Emp)GetSum(){
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","sum(sal) as SUM")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("sum of salaries of employees",emp.SUM)
}
func(e *Emp)GetAvg(){
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp = []repo.QueryProcessor{repo.Select("employees","avg(sal) as AVG")}
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("avg salaries of employees",emp.AVG)

}
func(e *Emp)GetValue(){
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT"))
	queryp=append(queryp,repo.Select("employees","sum(sal) as SUM"))
	queryp=append(queryp,repo.Select("employees","avg(sal) as AVG"))
	e.Repo.GetAll(uow, &emp, queryp)
	fmt.Println("no of employees",emp.COUNT)
	fmt.Println("sum of salaries of employees",emp.SUM)
	fmt.Println("avg salaries of employees",emp.AVG)

}
func (e *Emp)GetDepHeadCount(){
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
	//select COUNT(*) as c ,deptnofrom emp WHERE DEPTNO = 10 OR DEPTNO = 20 group by DEPTNO HAVING COUNT(*) >= 2 ORDER BY COUNT(*)Desc
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees","count(*) as COUNT,deptno as DEPTNO"))
	queryp=append(queryp,repo.Filter("DEPTNO = ? OR DEPTNO = ?",10,20))
	queryp=append(queryp,repo.Group("DEPTNO"))
	queryp=append(queryp,repo.Having("COUNT(*) >= 2"))
	queryp=append(queryp,repo.Order("COUNT(*) DESC",true))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("count is",emp.COUNT,"for dept no",emp.DEPTNO)
	}

}
func(e *Emp)GetName(){
	//SELECT ename,dname,job from dept left join emp on emp.DEPTNO = dept.DEPTNO;
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("dept","employees.ename as EMPNAME,dname as DEPNAME,employees.job as JOB"))
	queryp=append(queryp,repo.Joins("left join employees on employees.DEPTNO = dept.DEPTNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("empname is->",emp.EMPNAME,"having department->",emp.DEPNAME,"with job->",emp.JOB)
	}
}
func(e *Emp)GetCountByDname(){
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
	//display all the departments , with employees if any (if no emps then display null)
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("dept","employees.ename as EMPNAME,dept.dname as DEPNAME"))
	queryp=append(queryp,repo.Joins("LEFT OUTER JOIN employees ON employees.DEPTNO=dept.DEPTNO"))
	
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("employee->",emp.EMPNAME,"for dept name",emp.DEPNAME)
	}
}
func(e *Emp)GetDepWithNoEmp(){
	//display the departments where there are no employees
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
	//display the empname and their bossnames
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
	// display all the empnames and boss names if any (if no boss display null)

	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("employees emp1","emp1.ename as EMPNAME ,emp2.ename as BOSSNAME"))
	queryp=append(queryp,repo.Joins("LEFT JOIN employees emp2 ON emp1.MGR=emp2.EMPNO"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("employee->",emp.EMPNAME,"Boss->",emp.BOSSNAME)
	}
}
func(e *Emp)DispAllName(){
// display ename,deptname and bossname .
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
// display the countries there no states
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
	uow := repo.NewUnitOfWork(e.DB, true)
	var emp []model.Employee
	var queryp []repo.QueryProcessor
	queryp=append(queryp,repo.Select("countries","countries.country_name as CNAME,locations.state_province as STATEP,regions.region_name as RNAME"))
	queryp=append(queryp,repo.Joins("INNER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID"))
	queryp=append(queryp,repo.Joins("INNER JOIN REGIONS ON COUNTRIES.REGION_ID=REGIONS.REGION_ID"))
	e.Repo.GetAll(uow, &emp, queryp)
	for _,emp:=range(emp){
		fmt.Println("country name->",emp.CNAME,"state provience->",emp.STATEP,"rname->",emp.RNAME)
	}
}	
func (e *Emp) InsertSwabhavLocation() {
	//Make an insert of swabhav Techlabs in location/state tables map to india and asia.
	uow := repo.NewUnitOfWork(e.DB, true)
	e.DB.AutoMigrate(&model.Location{})
	Location := model.NewLocation(2, "mum", "8796", "Mumbai", "Maharashtra","IN")
	e.Repo.Add(uow, &Location)
}

func(e *Emp)GetData(){
	//Filter details based on mumbai location
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
	//create a foo table and insert values of different data
	unit := repo.NewUnitOfWork(e.DB, true)

	e.DB.AutoMigrate(&model.Foo{})
	foo1:=model.NewFoo("neha","lname")
	e.Repo.Add(unit, foo1)
	foo2 :=model.NewFoo("pooja","b")
	e.Repo.Add(unit, foo2)
}
