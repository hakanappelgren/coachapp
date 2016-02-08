package main

import(
  "net/http"
  "os"
)

func main() {
  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  // handle the specific requests
  http.HandleFunc("/coachquestion", ShowQuestsion)
  //  http.HandleFunc("/coachquestion", GenerateMarkdown)
  // handle all other requests
  http.Handle("/", http.FileServer(http.Dir("public")))
  // start the server
  http.ListenAndServe(":"+port, nil)
  // http.ListenAndServe(":8080", nil)
  // http.ListenAndServe(":8080", http.xxx(http.Dir(".")))
}
  func ShowQuestsion(rw http.ResponseWriter, r *http.Request) {
    questionlist := []byte {'f', 'a'}
//    questionlist := []byte {'första frågan', 'andra frågan'}

//    thecoachquestion := questionlist [1]
//    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(questionlist[1])
//    rw.Write([]byte(thecoachquestion))
}
