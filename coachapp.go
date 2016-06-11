package main

import(
  "net/http"
  "html/template"
  "os"
  "path"
  "fmt"
  //"math.rand"
)

//global variables
var qindex int = 0

type Question struct {
  Thequestion string
}

func GetIndex(x int) int {
  var ind int = x
  if ind == 0 {
    return 1
  } else {
    return 0
  }


//    if qindex < 1 {
//      fmt.Printf("qindex < 1, dvs =0 \n")
//      qindex = qindex + 1
//    } else {
//      fmt.Printf("qindex > 0, dvs =1 \n")
//      qindex = 0
//    }
//    tempindex = qindex
//   return tempindex
}


func main() {
  fmt.Printf("början av main.\n")
  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  fmt.Println(port)


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
    questionlist := [...]string{"Ny fråga 1", "NY fråga 2"}
    // index = rand.int31n(2)
    //This section is only a quick and dirty to get index iterating
    //Replace by function returning index
    if qindex < 1 {
      fmt.Printf("qindex < 1, dvs =0 \n")
      qindex = qindex + 1
    } else {
      fmt.Printf("qindex > 0, dvs =1 \n")
      qindex = 0
    }

    myquestion := Question{questionlist[qindex]}
//    myquestion := Question{questionlist[GetIndex]}
//        rw.Write([]byte("fråga 1"))

    fp := path.Join("public", "question.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(rw, myquestion); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }

}
