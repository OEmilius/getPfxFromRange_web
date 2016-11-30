package main

import (
	"fmt"
	"net/http"
	//"os/exec"
	"text/template"
)

var WebHostPort = ":8888"
var Templ = template.Must(template.ParseFiles("index.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, r.RequestURI, r.Method)
	//	var operator_cap []diapazon
	temp_ranges := `0106200000,0106599999
0113000000,0113299999`
	var v = struct {
		Host   string
		Data   string
		Ranges string
		Pfxs   string
	}{r.Host, "", temp_ranges, ""}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//w.Write([]byte("hello"))
		Templ.Execute(w, &v)
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err == nil {
			fmt.Println("in_text=", r.Form["in_text"])
			v.Ranges = r.Form["in_text"][0]
			if operator_cap, err := read_from_text(r.Form["in_text"][0]); err == nil {
				combined1 := FullCombine(operator_cap)
				results := list2prefix(combined1)
				v.Pfxs = fmt.Sprintf("%s", results)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				Templ.Execute(w, &v)
			} else {
				http.Error(w, "error parsing input text", 405)
			}
			//w.Write([]byte("eeend"))
			//w.Header().Set("Content-Type", "text/html; charset=utf-8")
			//Templ.Execute(w, &v)
			return
		} else {
			http.Error(w, "error ParseForm", 405)
		}
	} else {
		http.Error(w, "Method not allowed", 405)
		return
	}
}

func Send_ranges(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if err := r.ParseForm(); err == nil {
		fmt.Println("in_text=", r.Form["in_text"])
		w.Write([]byte("eeend"))
	} else {
		http.Error(w, "error ParseForm", 405)
	}
}

func WebStart() {
	http.HandleFunc("/", serveHome)
	//http.HandleFunc("/Send_ranges", Send_ranges)
	fmt.Println("ListenAndServe:" + WebHostPort)
	if err := http.ListenAndServe(WebHostPort, nil); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	fmt.Println("start")
	WebStart()
	//go WebStart()
	//_ = exec.Command("cmd", "/c", "start", "http://localhost:8888/").Start()
	//fmt.Scanln()
}
