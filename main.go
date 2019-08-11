package main

import (
	"./controller"
	"html/template"
	"log"
	"net/http"
)

func View() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()

		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			err = tpl.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	http.HandleFunc("/user/login", controller.Login)
	http.HandleFunc("/user/register", controller.Register)
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/createcommunity", controller.Createcommunity)
	http.HandleFunc("/chat", controller.Chat)
	http.HandleFunc("/attach/upload", controller.Upload)

	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	View()
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
