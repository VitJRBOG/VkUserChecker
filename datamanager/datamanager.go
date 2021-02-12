package datamanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	govkapi "github.com/VitJRBOG/GoVkApi/v2"
	"github.com/VitJRBOG/VkUserChecker/filemanager"
)

// GetUserData получает с серверов ВК данные о пользователе
func GetUserData(cfgValues filemanager.Config, urlToUserPage string) string {
	var userData string
	userInfo := requestUserInfo(cfgValues.AccessToken, extractScreenname(urlToUserPage))
	accountCreationDate := requestAccountCreationDate(userInfo.ID)

	userData = fmt.Sprintf("Full name: %v %v\n"+
		"Birthdate: %v\n"+
		"Account creation date: %v\n",
		userInfo.FirstName, userInfo.LastName, userInfo.Birthdate, accountCreationDate)

	for _, communityInfo := range cfgValues.Communities {
		isMember := checkCommunitySubscription(cfgValues.AccessToken, userInfo.ID, communityInfo.ID)
		if isMember {
			userData = fmt.Sprintf("%vIs member of %v\n",
				userData, communityInfo.Name)
		} else {
			userData = fmt.Sprintf("%vIs NOT member of %v\n",
				userData, communityInfo.Name)
		}
	}

	return userData
}

func extractScreenname(userPageURL string) string {
	posLastSlash := strings.LastIndex(userPageURL, "/")
	if posLastSlash == -1 {
		panic(errors.New("no user's screenname found in this URL"))
	}
	return strings.ReplaceAll(userPageURL, userPageURL[0:posLastSlash+1], "")
}

// User хранит информацию о пользователе
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"bdate"`
}

func requestUserInfo(accessToken, userScreenname string) User {
	v := map[string]string{
		"user_ids": userScreenname,
		"fields":   "bdate",
		"v":        "5.126",
	}
	res, err := govkapi.Method("users.get", accessToken, v)
	if err != nil {
		panic(err.Error())
	}
	return checkBirthdate(parseUserInfo(res))
}

func parseUserInfo(res []byte) User {
	var user User
	var users []User
	users = append(users, user)
	err := json.Unmarshal(res, &users)
	if err != nil {
		panic(err.Error())
	}
	return users[0]
}

func checkBirthdate(user User) User {
	if user.Birthdate == "" {
		user.Birthdate = "NO DATA"
	}
	return user
}

// Subscription хранит информацию о статусе подписки пользователя на сообщества
type Subscription struct {
	Member int `json:"member"`
}

func checkCommunitySubscription(accessToken string, userID, communityID int) bool {
	v := map[string]string{
		"group_id": strconv.Itoa(communityID),
		"user_id":  strconv.Itoa(userID),
		"extended": "1",
		"v":        "5.126",
	}
	res, err := govkapi.Method("groups.isMember", accessToken, v)
	if err != nil {
		panic(err.Error())
	}
	subscription := parseSubscriptionInfo(res)
	if subscription.Member == 1 {
		return true
	}
	return false
}

func parseSubscriptionInfo(res []byte) Subscription {
	var subscription Subscription
	err := json.Unmarshal(res, &subscription)
	if err != nil {
		panic(err.Error())
	}
	return subscription
}

func requestAccountCreationDate(userID int) string {
	return extractAccountCreationDate(requestHTML("https://vk.com/foaf.php?id=" + strconv.Itoa(userID)))
}

func requestHTML(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err.Error())
	}

	return doc
}

func extractAccountCreationDate(doc *goquery.Document) string {

	ret, err := doc.Html()
	if err != nil {
		panic(err.Error())
	}

	tagPos := strings.Index(ret, "<ya:created dc:date=")

	return ret[tagPos+21 : tagPos+31]
}
