package main

import(
  "net/http"
  "html/template"
  "os"
  "path"
  //"math.rand"
)

//global variables
var qindex int = 0

type Question struct {
  Thequestion string
}

func main() {

  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  // handle the specific requests
  http.HandleFunc("/coachquestion", ShowQuestion)
  //  http.HandleFunc("/coachquestion", GenerateMarkdown)
  // handle all other requests.
  // direct to the index.html file in "public" directory
  http.Handle("/", http.FileServer(http.Dir("public")))
  // start the server and listen to port. "nil" will make the server run until stopped.
  http.ListenAndServe(":"+port, nil)
  // http.ListenAndServe(":8080", nil)
  // http.ListenAndServe(":8080", http.xxx(http.Dir(".")))
}
  func ShowAbout (rw http.ResponseWriter, r *http.Request) {

}
  func ShowQuestion(rw http.ResponseWriter, r *http.Request) {
    questionlist := [...]string{"fråga 1", "fråga 2"}
    // index = rand.int31n(2)
    if qindex < 1 {
      qindex = qindex + 1
    } else {
      qindex = 0
    }
    myquestion := Question{questionlist[qindex]}

    fp := path.Join("public", "question.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(rw, myquestion); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }

//    rw.Write([]byte(thequestion))
}
