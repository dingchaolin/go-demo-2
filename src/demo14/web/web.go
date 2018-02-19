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
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

var (
	store = sessions.NewFilesystemStore("sessions")//session的存储位置
)

func render( w http.ResponseWriter, name string, data interface{}){
	path := filepath.Join("web/template", name + ".tpl")
	tpl, err := template.ParseFiles(path)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
}
func Login( w http.ResponseWriter, r *http.Request){

	render(w, "login", nil )
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
		http.Redirect(w, r, "hello", 302)
	}else{
		fmt.Fprintf(w,"user:%s,password:%s login error!", user, passwd)
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

//main函数启动之前调用一次
func init(){}

func Start(){
	{
		db, err := sql.Open("sqlite3", "web.db")
		if err != nil {
			log.Fatal(err)
		}
		err =  db.Ping()
	}
	//先挂载 再启动
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err =  db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/checklogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))//阻塞式调用 后面的代码不会执行
}