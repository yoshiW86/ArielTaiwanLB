package models

import (
	db "github.com/yoshiW86/ArielTaiwanLB/database" 
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	Sn         int    
	UserName   string 
	UserLineID string 
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SQLDB.Exec("INSERT INTO person(user_name, user_lineid) VALUES (?, ?)", p.UserName, p.UserLineID)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	log.Println("@GetPerson=======")
	persons = make([]Person, 0)
	rows, err := db.SQLDB.Query("SELECT sn FROM person where = ?", p.UserName)
	log.Println("rows:",rows,"err:",err)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer rows.Close()


	for rows.Next() {
		var person Person
		rows.Scan(person.Sn, person.UserName, person.UserLineID)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
