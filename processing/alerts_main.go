package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"crypto/x509"
	"io/ioutil"
	"crypto/tls"
	"github.com/go-sql-driver/mysql"
	"time"
)



func dbConn() (db *sql.DB) {

	rootCertPool := x509.NewCertPool()
	pem, _ := ioutil.ReadFile("./certs/SSL.crt.pem")
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool, InsecureSkipVerify: true})

	dbDriver := "mysql"
	dbUser := "user"
	dbPass := "<password>"
	dbName := "tcp(emoassistdb:3306)/emodb?allowNativePasswords=true&tls=custom"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}



type Alerts struct {
	Id int `db:"idalerts"`
	Created 		string `db:"created" json:"created"`
	UserId 			string `db:"userid" json:"userid"`
	Type    		string `db:"type" json:"type"`
	Probability 	float32 `db:"probability" json:"probability"`
	Hr 				int `db:"hr" json:"hr"`
	AlertType 		string `db:"alert_type" json:"alert_type"`
	Target 			string `db:"target" json:"target"`
	Status 			string `db:"status" json:"status"`
	Data   			string `db:"data" json:"data"`
}

func getAll(c *gin.Context){
	var (
		alert Alerts
		alerts []Alerts
	)

	db := dbConn()
    defer db.Close()

    rows, err := db.Query("select idalerts, created, userid, type, probability, hr, alert_type, target, status, data from alerts;")

	if err != nil {
		fmt.Print(err.Error())
	}

	for rows.Next() {
		rows.Scan(&alert.Id, &alert.Created, &alert.UserId, &alert.Type, &alert.Probability, &alert.Hr, &alert.AlertType, &alert.Target, &alert.Status, &alert.Data)
		alerts = append(alerts, alert)
	}

	defer rows.Close()
	c.JSON(http.StatusOK, alerts)
}

func add(c *gin.Context) {

	var alerts Alerts
	c.Bind(&alerts)


	Created := time.Now()
	//Created := "123"
	UserId := alerts.UserId
	Type := alerts.Type
	Probability := alerts.Probability
	Hr := alerts.Hr
	AlertType := alerts.AlertType
	Target := alerts.Target
	Status := alerts.Status
	Data := alerts.Data

	if UserId == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please fill all mandatory field"),
		})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("insert into alerts (created, userid, type, probability, hr, alert_type, target, status, data) values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Created, UserId, Type, Probability, Hr, AlertType, Target, Status, Data)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success"),
	})
}


func addWithDelay(c *gin.Context) {

	var alerts Alerts
	c.Bind(&alerts)


	Created := time.Now()
	UserId := alerts.UserId
	Type := alerts.Type
	Probability := alerts.Probability
	Hr := alerts.Hr
	AlertType := alerts.AlertType
	Target := alerts.Target
	Status := alerts.Status
	Data := alerts.Data

	if UserId == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please fill all mandatory field"),
		})
		return
	}


	LastCreated := lastTimeCreated(UserId, Type)

	Diff := Created.Sub(LastCreated)

	if Diff.Seconds() < 30 {

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("These updates are two often. Update ignored."),
		})
		return
	}

	db := dbConn()
	defer db.Close()


	stmt, err := db.Prepare("insert into alerts (created, userid, type, probability, hr, alert_type, target, status, data) values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Created, UserId, Type, Probability, Hr, AlertType, Target, Status, Data)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success"),
	})
}




func update(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	var alerts Alerts
	c.Bind(&alerts)


	Created := time.Now()
	//Created := "123"
	UserId := alerts.UserId
	Type := alerts.Type
	Probability := alerts.Probability
	Hr := alerts.Hr
	AlertType := alerts.AlertType
	Target := alerts.Target
	Status := alerts.Status
	Data := alerts.Data

	if UserId == "" || Type == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please fill all mandatory field"),
		})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("update alerts set created= ?, type= ?, probability= ?, hr= ?, alert_type= ?, target= ?, status= ?, data= ? where idalerts= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Created, UserId, Type, Probability, Hr, AlertType, Target, Status, Data, Id)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success"),
	})
}

func getById(c *gin.Context) {
	var alerts Alerts

	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("select idalerts, created, userid, type, probability, hr, alert_type, target, status, data from alerts where idalerts = ?;", id)

	err := row.Scan(&alerts.Id, &alerts.Created, &alerts.UserId, &alerts.Type, &alerts.Probability, &alerts.Hr, &alerts.AlertType, &alerts.Target, &alerts.Status, &alerts.Data)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, alerts)
	}
}

func getLastByUserId(c *gin.Context) {
	var alerts Alerts

	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("select idalerts, created, userid, type, probability, hr, alert_type, target, status, data from alerts where userid = ? order by idalerts DESC;", id)

	err := row.Scan(&alerts.Id, &alerts.Created, &alerts.UserId, &alerts.Type, &alerts.Probability, &alerts.Hr, &alerts.AlertType, &alerts.Target, &alerts.Status, &alerts.Data)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, alerts)
	}
}


func getByUserId(c *gin.Context) {
	var (
		alert Alerts
		alerts []Alerts
	)

	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	rows, err := db.Query("select idalerts, created, userid, type, probability, hr, alert_type, target, status, data from alerts where userid = ? order by idalerts desc;", id)

	if err != nil {
		fmt.Print(err.Error())
	}


	for rows.Next() {
		rows.Scan(&alert.Id, &alert.Created, &alert.UserId, &alert.Type, &alert.Probability, &alert.Hr, &alert.AlertType, &alert.Target, &alert.Status, &alert.Data)
		alerts = append(alerts, alert)
	}

	defer rows.Close()
	c.JSON(http.StatusOK, alerts)

}

// check the last update for this particular event in DB
func lastTimeCreated(UserId, Type string) time.Time {


	var Created string
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("select created from alerts where userid = ? and type = ? order by idalerts DESC;", UserId, Type).Scan(&Created)
	if err != nil {
		fmt.Print(err.Error())
	}

	layout := "2006-01-02 15:04:05"
	LastCreated, err := time.Parse(layout, Created)

	if err != nil {
		fmt.Println(err)
	}

	return LastCreated
}

func delete(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("delete from alerts where idalerts= ?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success. Deleted alert with ID : %s", id),
	})
}


func main() {


	router := gin.Default()
	router.GET("/api/alert/:id", getById)
	router.GET("/api/useralert/:id", getByUserId)
	router.GET("/api/lastuseralert/:id", getLastByUserId)
	router.GET("/api/alerts", getAll)
	router.POST("/api/alert", add)
	router.POST("/api/alertWithCheck", addWithDelay)
	router.PUT("/api/alert/:id", update)
	router.DELETE("/api/alert/:id", delete)
	router.Run(":8001")
}
