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

// func InitializeVariables {
  //@ToDo read text file and initialize variables

  // Read a slice from a jason file
  // import (
  //     "encoding/json"
  //     "fmt"
  //     "io/ioutil"
  // )

  //     data, err := ioutil.ReadFile("data.json")
  //     if err != nil {
  //         fmt.Println(err)
  //         return
  //     }
  //     fmt.Print("data:  ",string(data))
  //     var slice []string
  //     err = json.Unmarshal(data, &slice)
  //     if err != nil {
  //         fmt.Println(err)
  //         return
  //     }
  //     fmt.Printf("slice: %q\n",slice)
  //}


  // f, err := os.Open("data/list.txt") // For read access.
  // if err != nil {
  //     // Failed to open file, log / handle error
  //     log.Fatal("Open Filename: ", err)
  //     panic(err)
  //     fmt.Println(err)
  //     return
  // }
  // defer f.Close()
  // Here you may read from f
// }

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
	questionlist := [...]string{"Vad är ditt nästa steg?", "Vad är det värsta som kan hända?", "Hur ser ditt drömmål ut?", "Hur känner du inför uppgiften?", "Kan du utveckla?"}
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
