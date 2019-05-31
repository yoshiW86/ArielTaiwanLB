package models

import (
	db "github.com/yoshiW86/ArielTaiwanLB/database" 
	"log"
	"time"
)

//TimeSheet is for recording emp's working time
type TimeSheet struct {		//db column name
	Sn	int    				//sn
	EmpNo	int 			//emp_no
	StartTime	time.Time 	//start_time
	EndTime time.Time 		//end_time
	ActiveStatus int 		//active_status
	DDay time.Time			//d_day
}


//ClockInNOut will collect user info to check theic working time
func ClockInNOut(p *Person) (id int64, err error) {
	//clock in	
	rs, err := db.SqlDB.Exec("INSERT INTO timesheet(emp_no, activity_status, d_day) VALUES (?, ?, CURRENT_DATE())", p.GetUserSN(), 0)
	//activity_status defalut 0 is on_duty
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err = rs.LastInsertId()
	return
}

//GetSN let userLineID and target time put into this function, we can get S/N in timesheet.
func GetSN(p *Person, target time.Time) int{
	log.Println("target time:", target)
	rs := db.SqlDB.QueryRow("select sn from timesheet where d_day = ? and emp_no = ?", target, p.GetUserSN())
	var sn int
	err := rs.Scan(sn)
	if nil != err {
		log.Fatal(err)
	}
	log.Println("timesheet S/N:", sn)
	return sn
}

//HasARecord will check is there already a record? 
func HasARecord(sn int)bool{
	return 0 < sn
}