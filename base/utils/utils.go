package utils

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/alecthomas/template"
)

// parse body
func ParseBody(r *http.Request, body interface{}) {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		json.NewDecoder(r.Body).Decode(&body)
	}
}

// application Json
func Json(str string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, str)
	}
}

// get template dir
func getTemplatesDir() string {
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = "local"
	}
	return env
}

// parse templete
func ParseTemplate(path string) (*template.Template, error) {
	//reads all files
	log.Println("Path", path)
	p, er := template.New("mustache").Delims("<<", ">>").ParseGlob(getTemplatesDir() + "static" + "/templates/[a-z]*.mustache")
	log.Println("Error:----", er)
	return template.Must(p, er), er
}

// is prime number
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func checkPrime(value int, ch chan int) {
	if isPrime(value) {
		// send prime number
		ch <- value
	} else {
		// send the value zero when not a prime number
		ch <- 0
	}
}

// prime number
func PrimeNumbers(n int) (primes []int) {
	ch := make(chan int, n)
	for i := 1; i <= n; i++ {
		go checkPrime(i, ch)
	}

	for i := 0; i < n; i++ {
		val := <-ch
		if val != 0 {
			primes = append(primes, val)
		}
	}
	return
}

// end of month
func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

// remove all the html tags
func RemoveTags(html string) string {
	re, _ := regexp.Compile("<[^>/]+></[^>]+>")
	rep := re.ReplaceAllString(html, "")
	if rep != html {
		return RemoveTags(rep)
	}
	return rep
}

// word count
func WordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

// remove all the special character and numbers
func RemoveSpecialCharacters(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, " ")
}
