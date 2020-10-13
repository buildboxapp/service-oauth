package main

import "golang.org/x/oauth2"

const (
	defaultport         = 8525
	defaulthost         = "https://buildbox.app"
	defaultdatacallback = "/buildbox/gui/login/auth"
)

var (
	appOpts   AppOpts
	authconfs VendorsConfMap
)

type AppOpts struct {
	Host         string
	Port         int
	Dir          string
	ConfFile     string
	Command      string
	DataCallback string
	StaticPrefix string
	Test         bool
}

// proxy ping response
type Pong []struct {
	Name    string
	Version string
	Port    int
	Pid     string
}
type Config map[string]string
type BBResponse struct {
	Data interface{} `json:"data"`
}

type buttons []button
type button struct {
	Link      string
	Vendor    string
	Title     string
	AClass    string
	IconClass string
	Redirect  string
	Prefix    string
}

type VendorsConfMap map[string]*oauth2.Config
type VendorsMap map[string]struct {
	Title       string
	ID          string
	Pass        string
	ButtonClass string
	GetProfile  string
	Endpoint    oauth2.Endpoint
	Scopes      []string
}

// common profile struct
type UserProfile struct {
	Vendor    string `json:"network"`
	ID        string `json:"uid"`
	IDInt     int    `json:"-"`
	Name      string `json:"name"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"email"`
	Img       string `json:"photo"`
	Login     string `json:"login"`
	Nick      string `json:"nick"`
	Gender    string `json:"sex"`
	GenderInt int    `json:"-"`
	Lang      string `json:"lang"`
	Birth     string `json:"birth_date"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

// vendors structs
type GoogleData struct {
	Vendor    string
	ID        string `json:"id"`
	IDInt     int
	Name      string `json:"name"`
	FName     string `json:"given_name"`
	LName     string `json:"family_name"`
	Email     string `json:"email"`
	Img       string `json:"picture"`
	Login     string
	Nick      string
	Gender    string `json:"gender"`
	GenderInt int
	Lang      string `json:"locale"`
	Birth     string
	City      string
	Country   string
}
type YandexData struct {
	Vendor    string
	ID        string `json:"id"`
	IDInt     int
	Name      string `json:"real_name"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"default_email"`
	Img       string `json:"default_avatar_id"`
	Login     string `json:"login"`
	Nick      string `json:"display_name"`
	Gender    string `json:"sex"`
	GenderInt int
	Lang      string
	Birth     string `json:"birthday"`
	City      string
	Country   string
}
type VKData struct {
	Vendor    string
	ID        string
	IDInt     int `json:"id"`
	Name      string
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string
	Img       string `json:"photo_max_orig"`
	Login     string
	Nick      string `json:"screen_name"`
	Gender    string
	GenderInt int `json:"sex"`
	Lang      string
	Birth     string
	City      string
	Country   string
}
type MailRuData struct {
	Vendor    string
	ID        string `json:"id"`
	IDInt     int
	Name      string `json:"name"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"email"`
	Img       string `json:"image"`
	Login     string
	Nick      string `json:"nickname"`
	Gender    string `json:"gender"`
	GenderInt int
	Lang      string `json:"locale"`
	Birth     string `json:"birthday"`
	City      string
	Country   string
}
