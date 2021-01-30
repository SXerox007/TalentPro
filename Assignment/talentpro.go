package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"talentpro/Assignment/models"
	"talentpro/base/utils"
	"time"

	"github.com/gocolly/colly"
	strip "github.com/grokify/html-strip-tags-go"
)

// page struct
type Page struct {
	Title      string
	Name       string
	Email      string
	StaticHost string
	Config     string
	Json       string
}

// get word counter page
func GetWordCounterPage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := utils.ParseTemplate("talentpro-ui")
		page := &Page{
			Title:      "Word Counter",
			Config:     "config",
			Json:       "{}",
			StaticHost: "http://localhost:6011",
		}
		t.ExecuteTemplate(w, "word_counter", page)
		return
	}
}

// get prime number
func GetPrimeNumber() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("number")
		if number == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Invalid Fields", http.StatusBadRequest, nil})
			return
		}
		i, _ := strconv.Atoi(number)
		primeNumber := utils.PrimeNumbers(i)

		respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, primeNumber})
		return
	}
}

type DateDetails struct {
	Date           string `param:"Date" json:"Date"`
	LastDayOfMonth int    `param:"LastDayOfMonth" json:"LastDayOfMonth"`
}

// get last day of month
func GetLastDayOfMonth() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &DateDetails{}
		utils.ParseBody(r, body)
		if body.Date == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Invalid Fields", http.StatusBadRequest, nil})
			return
		}
		t, err := time.Parse("2006-01-02T15:04", body.Date)
		if err != nil {
			respondWithJSON(w, http.StatusConflict, CommonResponse{false, "Date format is Incorrect", http.StatusConflict, err})
			return
		}

		respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, DateDetails{
			Date:           body.Date,
			LastDayOfMonth: utils.EndOfMonth(t).Day(),
		}})
		return
	}
}

// get user list
func GetUserList() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if limit == 0 {
			// set the default limit
			limit = 10
		}
		// get users list
		if data, err := models.GetUserList(limit, offset); err != nil {
			log.Println("Error and Data", err, data)
			respondWithJSON(w, http.StatusConflict, CommonResponse{false, "Something Went Wrong", http.StatusConflict, err})
		} else {
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, data})
		}
		return
	}
}

// add user in to db
func AddUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &models.UserDetails{}
		utils.ParseBody(r, body)
		if body.LoginId == "" || body.FullName == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Missing Fields", http.StatusBadRequest, nil})
			return
		}
		// set state to 1 that means user state is active and 0 is inactive and delete state is 2
		body.State = 1
		if err := models.InsertUserIntoAllUsers(*body); err != nil {
			respondWithJSON(w, http.StatusConflict, CommonResponse{false, "Something Went Wrong", http.StatusConflict, err})
		} else {
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, nil})
		}
		return
	}
}

// delete user
func DeleteUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "UserId is missing", http.StatusBadRequest, nil})
			return
		}
		if err := models.DelteUserFromAllUsers(id); err != nil {
			respondWithJSON(w, http.StatusConflict, CommonResponse{false, "Something Went Wrong", http.StatusConflict, err})
		} else {
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Delete with Success", http.StatusOK, nil})
		}
		return
	}
}

// edit user details
func EditUserDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &models.UserDetails{}
		utils.ParseBody(r, body)
		if body.FullName == "" && body.LoginId == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Fields are missing", http.StatusBadRequest, nil})
			return
		}
		if body.Id == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "UserId is Missing", http.StatusBadRequest, nil})
			return
		}
		if err := models.UpdateUserDetails(*body); err != nil {
			respondWithJSON(w, http.StatusConflict, CommonResponse{false, "Something Went Wrong", http.StatusConflict, err})
		} else {
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Update user details with Success", http.StatusOK, nil})
		}
		return
	}
}

// get url to crawl
func GetUrlCrawl() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if url == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Url is missing", http.StatusBadRequest, nil})
			return
		}

		c := colly.NewCollector()

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		// extract status code
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
			//log.Println("Words Count:", utils.WordCount(utils.RemoveSpecialCharacters(strip.StripTags(string(r.Body)))))
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, utils.WordCount(utils.RemoveSpecialCharacters(strip.StripTags(string(r.Body))))})
		})

		c.OnError(func(r *colly.Response, err error) {
			log.Println("error:", r.StatusCode, err)
			respondWithJSON(w, r.StatusCode, CommonResponse{false, "Something went wrong", r.StatusCode, err})
		})

		c.Visit(url)

		return

	}
}
