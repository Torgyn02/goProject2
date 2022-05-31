package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	//"strconv"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	//"strconv"
	"text/template"
)

type Employee struct {
	EmployeeId  string `json:"employeeId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Adress      string `json:"adress"`
	PhoneNumber string `json:"phoneNumber"`
	TrariffRate string `json:"tariffRate"`
}
type Client struct {
	ClientId    string `json:"clientId"`
	Address     string `json:"address"`
	CompanyName string `json:"companyName"`
	Country     string `Json:"country"`
	PhoneNumber string `json:"phoneNumber"`
	ClientName  string `json:"clientName"`
	ProjectDes  string `json:"projectDes"`
}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
type Payment struct {
	PaymentId  string `json:"paymentId"`
	ProjectId  string `json:"projectId"`
	Ammount    string `json:"ammount"`
	CreditCard string `json:"creditCard"`
	ClientIDP  string `json:"clientIDP"`
}
type Project struct {
	ProjectId   string `json:"projectId"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	ProjectName string `json:"projectName"`
	Calculation string `json:"calculation"`
	EndDate     string `json:"endDate"`
	ClientIdPr  string `json:"clientIdPr"`
	PaymentIdP  string `json:"paymentIdP"`
}
type TimeCard struct {
	TimeCardId      string `json:"timeCardId"`
	EmployeeIdT     string `json:"employeeIdT"`
	DateIssue       string `json:"dateIssue"`
	DatePerform     string `json:"datePerform"`
	ProjectIdT      string `json:"projectIdT"`
	WorkDescription string `json:"workDescription"`
}

type MiddleData struct {
	User     User
	Projects []Project
	Employee []Employee
	Client   []Client
	Payment  []Payment
	TimeCard []TimeCard
}

type MangaMiddleData struct {
	User User
}

var user User

// var projects []Project
func selectff() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Query(fmt.Sprintf("SELECT project_id,description,start_date,project_name,calculation,end_date,client_id,payment_id  FROM `projects` "))

	for res.Next() {
		var tempProject Project
		err = res.Scan(&tempProject.ProjectId, &tempProject.Description, &tempProject.StartDate, &tempProject.ProjectName, &tempProject.Calculation, &tempProject.EndDate, &tempProject.ClientIdPr, &tempProject.PaymentIdP)

	}
	if err != nil {
		log.Fatal(err)
	}
	res.Close()

	resC, err := db.Query(fmt.Sprintf("SELECT  * FROM `employees` "))

	for res.Next() {
		var tempProject Project
		err = resC.Scan(&tempProject)

	}
	if err != nil {
		log.Fatal(err)
	}
	resP, err := db.Query(fmt.Sprintf("SELECT  * FROM `payment` "))

	for res.Next() {
		var tempProject Project
		err = resP.Scan(&tempProject)

	}
	if err != nil {
		log.Fatal(err)
	}
	resPr, err := db.Query(fmt.Sprintf("SELECT  * FROM `projects` "))

	for res.Next() {
		var tempProject Project
		err = resPr.Scan(&tempProject)

	}
	if err != nil {
		log.Fatal(err)
	}
	resT, err := db.Query(fmt.Sprintf("SELECT  * FROM `timecard` "))

	for res.Next() {
		var tempProject Project
		err = resT.Scan(&tempProject)

	}
	if err != nil {
		log.Fatal(err)
	}

	res.Close()

}

//================= WELCOME PAGE===========================================================================
func index(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/index.html")
	if err != nil {
		log.Fatal(err)
	}

	indexMiddleData := MiddleData{
		User:     user,
		Projects: []Project{},
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//project_id,description,start_date,project_name,calculation,end_date,client_id,payment_id
	selectff()
	file, err := ioutil.ReadFile("db/project.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &indexMiddleData.Projects)
	if err != nil {
		log.Fatal(err)
	}
	fileE, err := ioutil.ReadFile("db/employee.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileE, &indexMiddleData.Employee)
	if err != nil {
		log.Fatal(err)
	}
	fileT, err := ioutil.ReadFile("db/timecard.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileT, &indexMiddleData.TimeCard)
	if err != nil {
		log.Fatal(err)
	}
	fileC, err := ioutil.ReadFile("db/clients.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileC, &indexMiddleData.Client)
	if err != nil {
		log.Fatal(err)
	}
	filePay, err := ioutil.ReadFile("db/payment.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(filePay, &indexMiddleData.Payment)
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, indexMiddleData)
}

//================= REGISTRATION =================================================================================
func signup(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/signup.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func signupFunc(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	password := request.FormValue("password")
	repassword := request.FormValue("repassword")
	type1 := "weeb"

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if password == repassword {
		insert, err := db.Query("INSERT INTO `users`(`login`,`password`,`type`) VALUES(?,?,?)", name, repassword, type1)

		if err != nil {
			panic(err)
		}
		defer insert.Close()

		file, _ := ioutil.ReadFile("db/users.json")
		var allUsers []User
		_ = json.Unmarshal(file, &allUsers)
		for _, temp := range allUsers {
			if temp.Name == name {
				http.Redirect(writer, request, "/", http.StatusSeeOther)
				return
			}
		}
		if password == "employee" {
			allUsers = append(allUsers, User{
				Name:     name,
				Password: password,
				Type:     "emp",
			})
		} else {
			allUsers = append(allUsers, User{
				Name:     name,
				Password: password,
				Type:     "weeb",
			})
		}

		file, _ = json.MarshalIndent(allUsers, "", " ")
		_ = ioutil.WriteFile("db/users.json", file, 0644)
	}

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//====================== LOG IN PAGE ===================================================================================================
func signin(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/signin.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func signinFunc(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	password := request.FormValue("password")

	file, _ := ioutil.ReadFile("db/users.json")
	var allUsers []User
	_ = json.Unmarshal(file, &allUsers)

	for _, temp := range allUsers {
		if temp.Name == name && temp.Password == password {
			user = temp
			break
		}
	}

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//========= LOG OUT PAGE ============================================================================================================
func signout(writer http.ResponseWriter, request *http.Request) {
	user = User{
		Name:     "",
		Password: "",
		Type:     "",
	}
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//========= Order page===============================================

func newOrder(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/newOrder.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func newOrderFunc(writer http.ResponseWriter, request *http.Request) {
	client := Client{
		ClientId:    request.FormValue("clientId"),
		Address:     request.FormValue("address"),
		CompanyName: request.FormValue("companyName"),
		Country:     request.FormValue("country"),
		PhoneNumber: request.FormValue("phoneNumber"),
		ClientName:  request.FormValue("clientName"),
		ProjectDes:  request.FormValue("projectDes"),
	}
	// 		ClientId:    clientId,
	// 		Address:     address,
	// 		CompanyName: companyName,
	// 		Country:     country,
	// 		PhoneNumber: phoneNumber,
	// 		ClientName:  clientName,

	file, _ := ioutil.ReadFile("db/clients.json")
	var allClients []Client
	_ = json.Unmarshal(file, &allClients)
	for _, temp := range allClients {
		if temp.ClientId == client.ClientId {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}
	}

	allClients = append(allClients, client)
	file, _ = json.MarshalIndent(allClients, "", " ")
	_ = ioutil.WriteFile("db/clients.json", file, 0644)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `client`(`client_id`,`address`,`company_name`,`country`,`phone_number`,`client_name`) VALUES(?,?,?,?,?,?)", client.ClientId, client.Address, client.CompanyName, client.Country, client.PhoneNumber, client.ClientName)

	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//=========== Employee page==========================================================
func newEmployee(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/newEmployee.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func newEmployeeFunc(writer http.ResponseWriter, request *http.Request) {

	employee := Employee{
		EmployeeId:  request.FormValue("employeeId"),
		FirstName:   request.FormValue("firstName"),
		LastName:    request.FormValue("lastName"),
		Adress:      request.FormValue("address"),
		PhoneNumber: request.FormValue("phoneNumber"),
		TrariffRate: request.FormValue("trariffRate"),
	}

	file, _ := ioutil.ReadFile("db/employee.json")
	var allEmployee []Employee
	_ = json.Unmarshal(file, &allEmployee)
	for _, temp := range allEmployee {
		if temp.EmployeeId == employee.EmployeeId {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}
	}

	allEmployee = append(allEmployee, employee)
	file, _ = json.MarshalIndent(allEmployee, "", " ")
	_ = ioutil.WriteFile("db/employee.json", file, 0644)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertE, err := db.Query("INSERT INTO `employees`(`employee_id`,`first_name`,`last_name`,`address`,`phone_number`,`tariff_rate`) VALUES(?,?,?,?,?,?)", employee.EmployeeId, employee.FirstName, employee.LastName, employee.Adress, employee.PhoneNumber, employee.TrariffRate)

	if err != nil {
		panic(err)
	}
	defer insertE.Close()

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

// func updateEmp(writer http.ResponseWriter, request *http.Request) {
// 	html, err := template.ParseFiles("htmls/updateEmp.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	html.Execute(writer, user)
// }

// func updateEmpFunc(writer http.ResponseWriter, request *http.Request) {

// 	employee := Employee{
// 		EmployeeId:  request.FormValue("employeeId"),
// 		FirstName:   request.FormValue("firstName"),
// 		LastName:    request.FormValue("lastName"),
// 		Adress:      request.FormValue("address"),
// 		PhoneNumber: request.FormValue("phoneNumber"),
// 		TrariffRate: request.FormValue("trariffRate"),
// 	}
// 	id := request.FormValue("id")
// 	file, _ := ioutil.ReadFile("db/employee.json")
// 	var allEmployee []Employee
// 	_ = json.Unmarshal(file, &allEmployee)
// 	for _, temp := range allEmployee {
// 		if temp.EmployeeId == employee.EmployeeId {
// 			http.Redirect(writer, request, "/", http.StatusSeeOther)
// 			return
// 		}
// 	}

// 	allEmployee = append(allEmployee, employee)
// 	file, _ = json.MarshalIndent(allEmployee, "", " ")
// 	_ = ioutil.WriteFile("db/employee.json", file, 0644)

// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	stmt, _ := db.Prepare(`UPDATE employees SET employee_id = ?, first_name = ?,last_name=?,address=?,phone_number=?,tariff_rate=? WHERE employee_id = ?`)
// 	res, _ := stmt.Exec(employee.EmployeeId, employee.FirstName, employee.LastName, employee.Adress, employee.PhoneNumber, employee.TrariffRate, id)
// 	//update, err := db.Query("UPDATE `employees` SET `employee_id`=?,`first_name`=?,`last_name`=?,`address`=?,`phone_number`=?,`tariff_rate`=? WHERE `employee_id`=? ", )
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	res.LastInsertId()
// 	//defer update.Close()
// 	http.Redirect(writer, request, "/", http.StatusSeeOther)
// }

//======================= Project page=================================================================
func newProject(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/newProject.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func newProjectFunc(writer http.ResponseWriter, request *http.Request) {
	project := Project{
		ProjectId:   request.FormValue("projectId"),
		Description: request.FormValue("description"),
		StartDate:   request.FormValue("startDate"),
		ProjectName: request.FormValue("projectName"),
		Calculation: request.FormValue("calculation"),
		EndDate:     request.FormValue("endDate"),
		ClientIdPr:  request.FormValue("clientIdPr"),
		PaymentIdP:  request.FormValue("paymentIdP"),
	}

	file, _ := ioutil.ReadFile("db/project.json")
	var allProjects []Project
	_ = json.Unmarshal(file, &allProjects)
	for _, temp := range allProjects {
		if temp.ProjectId == project.ProjectId {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}
	}

	allProjects = append(allProjects, project)
	file, _ = json.MarshalIndent(allProjects, "", " ")
	_ = ioutil.WriteFile("db/project.json", file, 0644)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertP, err := db.Query("INSERT INTO `projects`(`project_id`,`description`,`start_date`,`project_name`,`calculation`,`end_date`,`client_id`,`payment_id`) VALUES(?,?,?,?,?,?,?,?)", project.ProjectId, project.Description, project.StartDate, project.ProjectName, project.Calculation, project.EndDate, project.ClientIdPr, project.PaymentIdP)
	if err != nil {
		panic(err)
	}
	defer insertP.Close()

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//=================== Payment page ===========================================================
func newPayment(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/newPayment.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func newPaymentFunc(writer http.ResponseWriter, request *http.Request) {
	payment := Payment{
		PaymentId:  request.FormValue("paymentId"),
		ProjectId:  request.FormValue("projectId"),
		Ammount:    request.FormValue("ammount"),
		CreditCard: request.FormValue("creditCard"),
		ClientIDP:  request.FormValue("clientIDP"),
	}

	file, _ := ioutil.ReadFile("db/payment.json")
	var allPayment []Payment
	_ = json.Unmarshal(file, &allPayment)
	for _, temp := range allPayment {
		if temp.PaymentId == payment.PaymentId {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}
	}

	allPayment = append(allPayment, payment)
	file, _ = json.MarshalIndent(allPayment, "", " ")
	_ = ioutil.WriteFile("db/payment.json", file, 0644)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertPy, err := db.Query("INSERT INTO `payment`(`payment_id`,`project_id`,`ammount`,`credit_card`,`client_id`) VALUES(?,?,?,?,?)", payment.PaymentId, payment.ProjectId, payment.Ammount, payment.CreditCard, payment.ClientIDP)

	if err != nil {
		panic(err)
	}
	defer insertPy.Close()

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

// ================= Time Card =============================================
func newTC(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("htmls/timeCard.html")
	if err != nil {
		log.Fatal(err)
	}

	html.Execute(writer, user)
}

func newTCFunc(writer http.ResponseWriter, request *http.Request) {
	tc := TimeCard{
		TimeCardId:      request.FormValue("timeCardId"),
		EmployeeIdT:     request.FormValue("employeeIdT"),
		DateIssue:       request.FormValue("dateIssue"),
		DatePerform:     request.FormValue("datePerform"),
		ProjectIdT:      request.FormValue("projectIdT"),
		WorkDescription: request.FormValue("workDescription"),
	}

	file, _ := ioutil.ReadFile("db/timecard.json")
	var allTC []TimeCard
	_ = json.Unmarshal(file, &allTC)
	for _, temp := range allTC {
		if temp.TimeCardId == tc.TimeCardId {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}
	}

	allTC = append(allTC, tc)
	file, _ = json.MarshalIndent(allTC, "", " ")
	_ = ioutil.WriteFile("db/timecard.json", file, 0644)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/smartproject")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertTC, err := db.Query("INSERT INTO `timecard`(`time_card_id`,`employee_id`,`date_issue`,`date_performance`,`project_id`,`work_description`) VALUES(?,?,?,?,?,?)", tc.TimeCardId, tc.EmployeeIdT, tc.DateIssue, tc.DatePerform, tc.ProjectIdT, tc.WorkDescription)

	if err != nil {
		panic(err)
	}
	defer insertTC.Close()

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

//================= STATIC DIRECTORY =================================================================================================
func embeddedDirectories() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/footer/", http.StripPrefix("/footer/", http.FileServer(http.Dir("./footer/"))))
	http.Handle("/header/", http.StripPrefix("/header/", http.FileServer(http.Dir("./header/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.Handle("/db/", http.StripPrefix("/db/", http.FileServer(http.Dir("./db/"))))
}

//===================================================================================================================================
func main() {

	embeddedDirectories()
	http.HandleFunc("/", index)
	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/signup", signupFunc)
	http.HandleFunc("/signin/", signin)
	http.HandleFunc("/signin", signinFunc)
	http.HandleFunc("/signout", signout)
	http.HandleFunc("/newOrder/", newOrder)
	http.HandleFunc("/newOrder", newOrderFunc)
	http.HandleFunc("/newEmployee/", newEmployee)
	http.HandleFunc("/newEmployee", newEmployeeFunc)
	http.HandleFunc("/newProject/", newProject)
	http.HandleFunc("/newProject", newProjectFunc)
	http.HandleFunc("/newPayment/", newPayment)
	http.HandleFunc("/newPayment", newPaymentFunc)
	http.HandleFunc("/newTC/", newTC)
	http.HandleFunc("/newTC", newTCFunc)
	// http.HandleFunc("/updateEmp/", updateEmp)
	// http.HandleFunc("/updateEmp", updateEmpFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
