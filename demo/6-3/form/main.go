package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)  //设置访问的路由
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		fmt.Println(err)
	}
}

func loginHandler(rep http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Fprintf(rep, "err %v", err)
			return
		}
		err = t.Execute(rep, nil)
		if err != nil {
			fmt.Fprintf(rep, "err %v", err)
		}
	} else {
		fmt.Println(" req.Form:", req.Form)        //req.Form: map[]
		fmt.Println(" req.Form:", req.Form == nil) // req.Form: true
		fmt.Println(" req.PostForm:", req.PostForm)
		fmt.Println("FormValue:", req.FormValue("username"))        //FormValue: admin
		fmt.Println("PostFormValue", req.PostFormValue("username")) //PostFormValue admin
		fmt.Println(" req.PostForm:", req.PostForm)
		fmt.Println("FormValue:param:", req.FormValue("param"))        //FormValue:param: 123
		fmt.Println("PostFormValue:param", req.PostFormValue("param")) //PostFormValue:param
		fmt.Println(" req.Form:", req.Form)                            // req.Form: map[param:[123] password:[123] username:[admin]]

	}
}
