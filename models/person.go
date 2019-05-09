package models

import (
	db "github.com/yoshiW86/ArielTaiwanLB/database" 
	// _ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Sn         int    
	UserName   string 
	UserLineID string 
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(user_name, user_lineid) VALUES (?, ?)", p.UserName, p.UserLineID)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT sn, user_name, user_lineid FROM person")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(person.Sn, person.UserName, person.UserLineID)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}