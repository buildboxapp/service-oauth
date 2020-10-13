package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func proxyping(c *gin.Context) {
	// ping response
	name := "ru"
	ver := "ru"
	prefix := conf["domain"]
	if prefix != "" {
		appOpts.StaticPrefix = prefix + "/"
		prefar := strings.Split(prefix, "/")
		name = prefar[0]
		if len(prefar) > 1 {
			ver = prefar[1]
		}
	}
	pid := fmt.Sprintf("%d:%s", os.Getpid(), conf["data-uid"])

	pong := Pong{
		{
			name,
			ver,
			appOpts.Port,
			pid,
		},
	}

	c.JSON(200, pong)
}

func info(c *gin.Context) {
	// get embed buttons
	btnsurl := fmt.Sprintf("%s/%s?embed=1", appOpts.Host, appOpts.StaticPrefix)
	fmt.Print("Get embed buttons from ", btnsurl, "... ")
	resp, err := http.Get(btnsurl)
	embed_btns := ""
	if err == nil {
		defer resp.Body.Close()

		if resp.StatusCode < 400 {
			b, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				embed_btns = string(b)
				fmt.Println("OK")
			} else {
				fmt.Println("")
				fmt.Println("Error read embed buttons:", err)
			}
		} else {
			fmt.Println("status", resp.StatusCode)
		}
	} else {
		fmt.Println("")
		fmt.Println("Error get embed buttons:", err)
	}

	for _, lng := range (*gincon)(c).ParseUserLangs() {
		if lng.Lang == "" {
			continue
		}

		switch strings.ToLower(lng.Lang)[:2] {
		case "ru":
			c.HTML(200, "info-ru.tpl", gin.H{
				"vendors": strings.Join(vendorslist, ","),
				"embedbtns": embed_btns,
				"prefix": appOpts.StaticPrefix,
				"host": appOpts.Host,
			})
			return
		case "en":
			c.HTML(200, "info-en.tpl", gin.H{
				"vendors": strings.Join(vendorslist, ","),
				"embedbtns": embed_btns,
				"prefix": appOpts.StaticPrefix,
				"host": appOpts.Host,
			})
			return
		}
	}
	c.HTML(200, "info-en.tpl", gin.H{
		"vendors": strings.Join(vendorslist, ","),
		"embedbtns": embed_btns,
		"prefix": appOpts.StaticPrefix,
		"host": appOpts.Host,
	})
}

func loginform(c *gin.Context) {
	redirecturl := c.Query("redirect")
	vendors := c.Query("vendors")
	btntheme := c.Query("theme")
	embed := c.Query("embed")
	vendors_ar := []string{}
	if vendors != "" {
		vendors_ar = strings.Split(vendors, ",")
	} else {
		vendors_ar = vendorslist
	}

	if redirecturl != "" && !regexp.MustCompile(`^https?://([\pL\d_.-]+\.\pL{2,}|[\pL\d_]+)(:\d+)?(/[\pL\d_.%/?&=:+-]+[\pL\d_/]|/)?$`).MatchString(redirecturl) {
		// wrong callback format
		c.String(400, "Wrong redirect param format"+ifs(appOpts.Test, ": "+redirecturl, ""))
		return
	}

	if !regexp.MustCompile(`^[\w_.-]+$`).MatchString(btntheme) {
		btntheme = ""
	}
	btnthemeclass := ifs(btntheme != "", fmt.Sprintf("bboauth-%s", btntheme), "")

	buttonparams := make(buttons, 0)
	for _, vendor := range vendors_ar {
		if !arrayContains(vendorslist, vendor) {
			// not valid vendor
			continue
		}

		conf := authconfs[vendor]
		state := fmt.Sprintf(`%s|%s`, vendor, redirecturl)
		link := conf.AuthCodeURL(state)

		buttonparams = append(buttonparams, button{
			Link:      link,
			Vendor:    vendor,
			Title:     authvendors[vendor].Title,
			AClass:    btnthemeclass,
			IconClass: authvendors[vendor].ButtonClass,
			Redirect:  redirecturl,
			Prefix:    appOpts.StaticPrefix,
		})
	}

	if len(buttonparams) == 0 {
		c.String(500, "No valid vendors")
		return
	}

	if embed != "" {
		// embedded
		c.HTML(200, "embedded.tpl", gin.H{"buttons": buttonparams})
		return
	}

	if len(vendors_ar) == 1 {
		// сразу редирект, когда 1 вендор
		c.HTML(200, "onevendor.tpl", gin.H{"buttons": buttonparams})
	} else {
		c.HTML(200, "loginpage.tpl", gin.H{"buttons": buttonparams})
	}
}

func auth(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	vendor := ""
	redirecturl := ""

	fmt.Println("Rendering auth form...")

	if state != "" {
		state_ar := strings.Split(state, "|")
		if len(state_ar) > 0 {
			vendor = state_ar[0]
		}
		if len(state_ar) > 1 {
			redirecturl = state_ar[1]
		}
	}

	if vendor == "" || code == "" || !arrayContains(vendorslist, vendor) {
		fmt.Println("Bad request")
		c.AbortWithStatus(500)
		return
	}

	// get token
	conf := authconfs[vendor]
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Error get token:", err)
		c.AbortWithStatus(500)
		return
	}
	fmt.Println("Get token OK")

	// copy profile url for modify (vk)
	getprofileurl := authvendors[vendor].GetProfile

	// fix profile request for vendors
	if vendor == "vk" {
		//fix for vk: get user id
		switch tok.Extra("user_id").(type) {
		case float64:
			uid := int(tok.Extra("user_id").(float64))
			getprofileurl += fmt.Sprintf("&user_id=%d&access_token=%s", uid, tok.AccessToken)
		default:
			fmt.Println("Error get user id")
			c.AbortWithStatus(500)
			return
		}
	}
	if vendor == "mailru" {
		getprofileurl += fmt.Sprintf("?access_token=%s", tok.AccessToken)
	}

	// get profile
	client := conf.Client(oauth2.NoContext, tok)
	resp, err := client.Get(getprofileurl)
	if err != nil {
		fmt.Println("Error get profile:", err)
		c.AbortWithStatus(500)
		return
	}
	fmt.Println("Get profile OK")
	defer resp.Body.Close()

	// read profile data
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read profile:", err)
		c.AbortWithStatus(500)
		return
	}

	if appOpts.Test {
		fmt.Println("Raw profile: ", string(response))
	}

	var userData UserProfile

	// cast data to unification
	switch vendor {
	case "google":
		tmpUserData := GoogleData{}
		err = json.Unmarshal(response, &tmpUserData)
		userData = UserProfile(tmpUserData)
	case "yandex":
		tmpUserData := YandexData{}
		err = json.Unmarshal(response, &tmpUserData)
		userData = UserProfile(tmpUserData)
		if userData.Img != "" && !regexp.MustCompile(`(?i)^(https?|ftp)://`).MatchString(userData.Img) {
			userData.Img = "https://avatars.yandex.net/get-yapic/" + userData.Img + "/islands-200"
		}
	case "vk":
		tmpUserData := struct {
			Data []VKData `json:"response"`
		}{
			[]VKData{
				{},
			},
		}
		err = json.Unmarshal(response, &tmpUserData)
		if err == nil && len(tmpUserData.Data) > 0 {
			// convert fields:
			userData = UserProfile(tmpUserData.Data[0])
			// convert id
			userData.ID = strconv.Itoa(userData.IDInt)
			// convert gender
			genders := map[int]string{
				1: "female",
				2: "male",
			}
			userData.Gender = genders[userData.GenderInt]
		}
	case "mailru":
		tmpUserData := MailRuData{}
		err = json.Unmarshal(response, &tmpUserData)
		userData = UserProfile(tmpUserData)
		// convert gender
		genders := map[string]string{
			"f": "female",
			"m": "male",
		}
		userData.Gender = genders[userData.Gender]
	default:
		fmt.Println("Error parsing profile: unknown vendor")
		c.AbortWithStatus(500)
		return
	}
	if err != nil {
		fmt.Println("Error deserialize profile:", err)
		c.AbortWithStatus(500)
		return
	}

	// add vendor
	userData.Vendor = vendor
	// add full name
	if userData.Name == "" {
		userData.Name = fmt.Sprintf("%s %s", userData.FName, userData.LName)
	}
	// format birth date
	if userData.Birth != "" && !regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(userData.Birth) {
		userData.Birth = regexp.MustCompile(`\b(\d)\b`).ReplaceAllString(userData.Birth, "0$1")
		userData.Birth = regexp.MustCompile(`(\d{2})\.(\d{2})\.(\d{4})`).ReplaceAllString(
			userData.Birth, "$3-$2-$1")
	}

	fmt.Println("Auth OK")
	//c.AbortWithStatusJSON(200, userData)

	fmt.Println("Sending data...")
	udatastr, err := json.Marshal(userData)
	if err != nil {
		fmt.Println("Error serialize user data")
		c.String(500, "Internal error")
		return
	}

	if sendresulttofront {
		// send to front

		fields := url.Values{
			"suser": {string(udatastr)},
		}
		params := fields.Encode()

		if redirecturl == "" {
			c.HTML(400, "authok.tpl", gin.H{"error": "Не передан URL для перехода"})
		} else {
			c.HTML(200, "authok.tpl", gin.H{
				"frontcallback": strings.ReplaceAll(redirecturl, "'", `\'`),
				"frontcallbackparams": strings.ReplaceAll(params, "'", `\'`),
				"clearparams": "suser",
			})
		}
	} else {
		// send to back

		fields := url.Values{
			"suser": {string(udatastr)},
			"ref":   {redirecturl},
		}
		params := fields.Encode()

		uri := ifs(appOpts.DataCallback[:1] == "/", appOpts.Host+appOpts.DataCallback, appOpts.DataCallback)
		urldata := fmt.Sprintf("%s?%s", uri, params)

		if appOpts.Test {
			fmt.Println("Sending data: ", urldata)
		}

		res, err := http.Get(urldata)
		if err != nil {
			fmt.Println("Error send user data to app")
			c.String(500, "Error send data to application")
			return
		}

		defer res.Body.Close()

		if res.StatusCode < 400 {
			fmt.Println("Send user data OK", ifs(appOpts.Test, fmt.Sprintf("(status %d)", res.StatusCode), ""))
			c.HTML(200, "authok.tpl", nil)
			if appOpts.Test {
				b, _ := ioutil.ReadAll(res.Body)
				fmt.Println("Send user data app response:", string(b))
			}
		} else {
			fmt.Println("Send user data response code:", res.StatusCode)
			c.String(500, "Error send data to application")
		}
	}
}

func reqauthresult(c *gin.Context) {

}
