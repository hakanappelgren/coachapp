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
 var currentindex int = 0
 var randomquestions bool = true

type Question struct {
  Thequestion string
}

func GetIndex(thelength int, randomindex bool, currentindex int) int {
	// return an integer [0, thelength-1]
	var index int = 0
	if randomindex {
        index = rand.Intn(thelength) // return an integer [0, thelength-1]
    } else {
    	if currentindex < thelength-1 {
    		index = currentindex + 1
    	} else {
    		index = 0
    	}
    }
  	return index
}

func GetQuestion (currentindex int)string {
	var index int = 0
	questionlist := [...]string{"Ny fråga 1", "Nästan ny fråga 2", "fråga 3"}
	index = GetIndex(len(questionlist), randomquestions, currentindex)
	currentindex = index
	return questionlist[index]
}

func main() {
  fmt.Printf("början av main.\n")
  port := os.Getenv("PORT")
    if port == "" {
  port = "8080"
  }
  fmt.Println(port)

	// initialize variables
// 	var qindex int = 0
// 	var randomquestions boolean = true

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
    questionlist := [...]string{"Ny fråga 1", "Nästan ny fråga 2"}
    // index = rand.Intn(6) + 1 // return an integer [1, 6
    //This section is only a quick and dirty to get index iterating
    //Replace by function returning index
    if currentindex < 1 {
      fmt.Printf("currentindex < 1, dvs =0 \n")
      currentindex = currentindex + 1
    } else {
      fmt.Printf("currentindex > 0, dvs =1 \n")
      currentindex = 0
    }

    myquestion := Question{questionlist[currentindex]}
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
