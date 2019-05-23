package models

import (
	db "github.com/yoshiW86/ArielTaiwanLB/database" 
	// _ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	Sn         int    
	UserName   string 
	UserLineID string 
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(user_name, user_lineid) VALUES (?, ?)", p.UserName, p.UserLineID)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err = rs.LastInsertId()
	return
}


//////////////working/////////////
// func (p *Person) GetPersons() (persons []Person, err error) {
// 	log.Println("@GetPersons=======")
// 	persons = make([]Person, 0)
// 	rows, err := db.SqlDB.Query("SELECT sn FROM person where user_name= ?", p.UserName)
// 	log.Println("rows:",rows,"err:",err)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}
// 	defer rows.Close()


// 	for rows.Next() {
// 		var person Person
// 		rows.Scan(&person.Sn)
// 		persons = append(persons, person)
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}
// 	return
// }

// get a person
func (p *Person) HadAUser() bool {
	rs := db.SqlDB.QueryRow("SELECT sn FROM person WHERE user_lineid = ?", p.UserLineID)
	if err := rs.Scan(&p.Sn); nil != err{
		log.Fatal(err)
	}

	log.Println("sn:", p.Sn)
	//add another columns

	if 0 < p.Sn { return true } 
	return false

}
