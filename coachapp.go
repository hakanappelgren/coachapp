package main

import(
  "net/http"
  "html/template"
  "os"
  "path"
  "fmt"
  "math/rand"
)

//global variables
var Israndom bool = true
var Masterindex int = 0

type Question struct {
  Thequestion string
}

func GetIndex(thelength int, userandomindex bool, currentindex int) int {
	// return an integer [0, thelength-1]
	var index int = 0
	  if userandomindex {
        index = rand.Intn(thelength)
    } else {
    	if currentindex < thelength-1 {
    		index = currentindex + 1
    	} else {
    		index = 0
    	}
    }
  	return index
}

func GetQuestion (currentindex int, userandomindex bool) (string, int) {
	questionlist := [...]string{"Ny fråga 1", "Nästan ny fråga 2", "fråga 3"}
	index := GetIndex(len(questionlist), userandomindex, currentindex)
	return questionlist[index], index
}

func main() {

  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  fmt.Println(port)

  // handle the specific requests
  http.HandleFunc("/coachquestion", ShowQuestion)

  // handle all other requests.
  // direct to the index.html file in "public" directory
  http.Handle("/", http.FileServer(http.Dir("public")))
  // start the server and listen to port. "nil" will make the server run until stopped.
  http.ListenAndServe(":"+port, nil)
}

  func ShowAbout (rw http.ResponseWriter, r *http.Request) {
// @ToDo
}

  func ShowQuestion(rw http.ResponseWriter, r *http.Request) {
    myquestion, newIndex := GetQuestion(Masterindex, Israndom)
    Masterindex = newIndex
    fmt.Println("Masterindex: ", Masterindex)

    fp := path.Join("public", "question.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(rw, Question{myquestion}); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }

}
