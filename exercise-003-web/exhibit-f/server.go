package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"syscall"
)

type View struct {
	LastSignup    string
	UserDirectory map[string]int
}

var homeT = template.Must(template.ParseFiles("exhibit-f/home.html"))
var lastSignup string
var userDirectory map[string]int = make(map[string]int)

func home(w http.ResponseWriter, r *http.Request) {
	v := View{LastSignup: lastSignup, UserDirectory: userDirectory}
	homeT.Execute(w, &v)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")

	if _, ok := userDirectory[username]; ok {
		userDirectory[username]++
	} else {
		userDirectory[username] = 1
	}

	lastSignup = username
	http.Redirect(w, r, "/home", 302)
}

func saveList() {
	var writer *bufio.Writer
	if _, err := os.Stat("user_list.txt"); os.IsNotExist(err) {
		userList, err := os.Create("user_list.txt")
		if err != nil {
			fmt.Println("Unable to create directory file")
			return
		}
		writer = bufio.NewWriter(userList)
		defer userList.Close()
	} else {
		userList, err := os.Open("user_list.txt")
		if err != nil {
			fmt.Println("Unable to open saved directory file")
		}
		writer = bufio.NewWriter(userList)
		defer userList.Close()
	}

	for user, count := range userDirectory {
		fmt.Fprintln(writer, user, count)
	}

	writer.Flush()
}

func main() {
	//runtime.GOMAXPROCS(2)
	if _, err := os.Stat("user_list.txt"); err == nil {
		userList, err := os.Open("user_list.txt")
		if err != nil {
			fmt.Println("Unable to open saved directory file")
		}
		defer userList.Close()

		//should technically scan for buffer overflow
		scanner := bufio.NewScanner(userList)
		for scanner.Scan() {
			entry := regexp.MustCompile(`\s+`).Split(scanner.Text(), -1)
			userDirectory[entry[0]], err = strconv.Atoi(entry[1])
			if err != nil {
				fmt.Println("Unable to load saved directory file")
			}
		}
	}

	interruptsChannel := make(chan os.Signal, 2)
	signal.Notify(interruptsChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptsChannel
		saveList()
		os.Exit(1)
	}()

	runtime.Gosched()

	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
