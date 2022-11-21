package tukorea

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func loginService(body tukoreaLogin) bool {
	// struct to json
	pbytes, _ := json.Marshal(body)
	buff := bytes.NewBuffer(pbytes)

	// post request
	res, err := http.Post("https://eclass.tukorea.ac.kr/ilos/lo/login.acl", "application/json", buff)

	// err
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}
	responseString := string(responseBytes)

	// result err
	if !strings.Contains(responseString, `document.location.href="https://eclass.tukorea.ac.kr/ilos/main/member/login_form.acl"`) {
		log.Fatal("잘못된 접근")
		return false
	}
	if strings.Contains(responseString, "로그인 정보가 일치하지 않습니다.") {
		log.Fatal("로그인 정보가 일치하지 않습니다.")
		return false
	}

	return true
}

func getTimetableService() []timetable {
	// post request
	res, err := http.Get("https://eclass.tukorea.ac.kr/ilos/lo/login.acl")

	// err
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// parsing
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var _timetable []timetable
	// Find the review items
	doc.Find("#contentsIndex > div.index-leftarea02 > div:nth-child(2) > ol").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the subject
		tukoreaSubject := s.Find("em.sub_open").Text()
		_timetable = append(_timetable, timetable{subject: tukoreaSubject})
	})

	return _timetable
}
