package main
//Dependency Inversion:=high-level logic should not depend on its low-level implementations(concret implemention). 
import (
	"fmt"
)

type DBConn interface {
	Query() interface{}
}
type UsersRepository struct {
	db DBConn
	//db MySQL //Violates DIP because it is directly dependent on concrete implementation(MySQL) rather than interface
}
type MySQL struct{}

/*func (db MySQL) QueryMySQLDB() map[string]int { //it is specifically for Myysql not in generalized form
	return map[string]string{
		"abcde": 1,
		"abc": 2,
	}
}*/
func (db MySQL) Query() interface{} {
	return map[string]int{
		"abcde": 1,
		"abc": 2,
	}
}

type PostgreSQL struct{}

/*func (db MySQL) QueryPostgreSQL() []string {
	return []string{"neha", "pooja"}
}*/
func (db PostgreSQL) Query() interface{} {
	return []string{"neha", "pooja"}
}
func main() {
	/*db := MySQL{}
	users := UsersRepository{db: db}
	fmt.Println(users.db.QueryMySQLDB())*/

	dbPSQL := PostgreSQL{}
	user := UsersRepository{db: dbPSQL}
	fmt.Println(user.db.Query())
}