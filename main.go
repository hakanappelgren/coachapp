package main

import(
  "net/http"
  "os"
  //"math.rand"
)

func main() {
  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  // handle the specific requests
  http.HandleFunc("/coachquestion", ShowQuestion)
  //  http.HandleFunc("/coachquestion", GenerateMarkdown)
  // handle all other requests
  //ToDo: have to change this action and serve the start page insted och the FileServer stuff
  http.Handle("/", http.FileServer(http.Dir("public")))
  // start the server
  http.ListenAndServe(":"+port, nil)
  // http.ListenAndServe(":8080", nil)
  // http.ListenAndServe(":8080", http.xxx(http.Dir(".")))
}
  func ShowQuestion(rw http.ResponseWriter, r *http.Request) {
//    questionlist := []byte {'f', 'a'}
//    questionlist := []byte ("min första fråga")
    questionlist := [...]string{"fråga 1", "fråga 2"}
    // index = rand.int31n(2)
    thequestion := questionlist[1]
//    questionlist := "min första fråga"

//    thecoachquestion := questionlist [1]
//    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
//    rw.Write(questionlist[1])
    rw.Write([]byte(thequestion))
}
