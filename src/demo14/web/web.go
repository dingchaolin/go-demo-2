package web

import (
	"log"
	"net/http"
	"io"
	"html/template"
	"path/filepath"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"

	"database/sql"
	//"github.com/jmoiron/sqlx"
)

var (
	store = sessions.NewFilesystemStore("sessions")//session的存储位置
	db *sql.DB
)

func render( w http.ResponseWriter, name string, data interface{}){
	path := filepath.Join("web/template", name + ".tpl")
	tpl, err := template.ParseFiles(path)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
}
func Login( w http.ResponseWriter, r *http.Request){

	render(w, "login", nil )
	//render(w, "login", "password error" )
}

func CheckLogin( w http.ResponseWriter, r *http.Request){
    // get post 都可以
	r.ParseForm()// 不区分 get post
	user := r.FormValue("user")
	passwd := r.FormValue("password")

	if user  == "admin" && passwd == "admin"{
		//fmt.Fprintf(w, "login ok")
		//设置cookie
		//cookie := &http.Cookie{
		//	Name:"user",
		//	Value: user,
		//	MaxAge: 10,//单位秒
		//}
		//http.SetCookie(w, cookie )

		//session
		//session := sessions.NewSession(store, "web")
		//session.Values["user"]= user
		//session.Save( r, w )
		http.Redirect(w, r, "/list", 302)

	}else{
		//fmt.Fprintf(w,"user:%s,password:%s login error!", user, passwd)
		render(w, "login", "password error" )
	}

}


func Hello(w http.ResponseWriter, r *http.Request){
	//获取cookie
	//_, err := r.Cookie("user")
	//if err != nil {
	//	http.Redirect(w, r, "/login", 302)
	//	return
	//}

	//获取session
	//session, _ := store.Get(r, "web")
	//_, ok := session.Values["user"]
	//if !ok {
	//		http.Redirect(w, r, "/login", 302)
	//		return
	//}
	io.WriteString(w, "hello http")
}

func List(w http.ResponseWriter, r *http.Request){
	type MyUser struct{
		Id int
		Name string
		Note string
	}
	list := []MyUser{
		{1,"dcl","i am admin"},
		{2,"ys","i am girl"},
	}
	render(w, "list", list )
}

func Delete( w http.ResponseWriter, r *http.Request){
	// get post 都可以
	r.ParseForm()// 不区分 get post
	id := r.FormValue("id")

		fmt.Fprintf(w,"id:%s delete success!", id)

}

//main函数启动之前调用一次
func init(){}

func Start(){
	/*
	{
		dbx, errx := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
		if errx != nil {
			log.Fatal(errx)
		}

		type MyUser struct{
			Id int
			Name string
			Password string
			Note string
			Isadmin bool
		}

		var users []MyUser
		errx = dbx.Select(&users, "select * from user")
		if errx != nil {
			log.Fatal(errx)
		}
		fmt.Print( "users===============",users )


		var user MyUser
		errx = dbx.Get(&user, "select * from user WHERE name = ?", "ys")
		if errx != nil {
			log.Fatal(errx)
		}
		fmt.Print( "user===============",user, "id=====",user.Id )
		return
	}
	var err error
	//{
	//	db, err := sql.Open("sqlite3", "web.db")//会自动产生一个web.db的文件
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	err =  db.Ping()
	//}
	//先挂载 再启动
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err =  db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	row:= db.QueryRow("select CURRENT_USER()")
	if err != nil {
		log.Fatal(err)
	}
	var user string
	row.Scan(&user)
	log.Print(user)


	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var (
		id       int
		name     string
		passwd   string
		note     string
		isadmin  int
	)
	for rows.Next(){
		rows.Scan(&id, &name, &passwd, &note, &isadmin)
		log.Print( id, name, passwd, note, isadmin)
	}
	return
*/

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/checklogin", CheckLogin)
	http.HandleFunc("/list", List)
	http.HandleFunc("/delete", Delete)

	log.Fatal(http.ListenAndServe(":8090", nil))//阻塞式调用 后面的代码不会执行
}

// db.QueryRow 不用close
/*
row := db.QueryRow("select password from user where name = ?", user)
var dbpass string
err := row.Scan(&dbpass)
if err == sql.ErrNoRows{
	http.Redirect(w, r, "/login", 302)
	return
}

md5.Sum([]byte(passwd), dbpass)

Query 是查询数据
Exec 增加 删除 修改 数据的

res, err := db.Exec("insert into user values(NULL, ?, ?, ?, ?)", name, passwd, note, 1)
if err != nil{
	http.Error(w, err.Error(), 500)
    return
}
res.LastInsertId()//最后插入的id
res.RowsAffected()//影响的行数


stmt, err := db.Prepare("insert into user values(NULL, ?, ?, ?, ?)")
stmt.Exec(name, passwd, note, 1)
stmt.Exec(name, passwd, note, 2)

事物
tx, err := db.Begin()
tx.Exec(...)//处理业务
tx.Commit()//提交
tx.RollBack()//回滚



SetMaxIdleConns 最大空闲连接数
比如压力大的时候 有100个连接  当压力过去了 不能把所有的连接都释放 要保留一些常驻空闲连接

SetMaxOpenConns 最大连接数

 */