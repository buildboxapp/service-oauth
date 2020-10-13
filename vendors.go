package main

import (
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/mailru"
	"golang.org/x/oauth2/vk"
	"golang.org/x/oauth2/yandex"
)

// parameters for vendors
var authvendors = VendorsMap{
	"google": {
		Title:       "Google",
		ID:          "241226599463-ukv5s5v9imockpqt5tmboli1mpq6h366.apps.googleusercontent.com",
		Pass:        "IJjqt0IX4BgMCQqqBJZ_l2V4",
		ButtonClass: "fab fa-google",
		GetProfile:  "https://www.googleapis.com/oauth2/v1/userinfo?alt=json",
		Endpoint:    google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	},
	"yandex": {
		Title:       "Yandex",
		ID:          "d5ef7e4eb06b45c59262698cf991f8e2",
		Pass:        "6ef92c51b2974afcb7617a76997451de",
		ButtonClass: "fab fa-yandex",
		GetProfile:  "https://login.yandex.ru/info?format=json",
		Endpoint:    yandex.Endpoint,
		Scopes: []string{
			"login:email",
			"login:info",
			"login:avatar",
			"login:birthday",
		},
	},
	"vk": {
		Title:       "VK",
		ID:          "7361738",
		Pass:        "rtTzR819t89N9eH4pf0A",
		ButtonClass: "fab fa-vk",
		GetProfile: "https://api.vk.com/method/users.get" +
			"?fields=first_name,last_name,screen_name,sex,bdate,photo_max_orig,city,country" +
			"&v=5.103&lang=ru",
		Endpoint: vk.Endpoint,
		Scopes:   []string{},
	},
	"mailru": {
		Title:       "Mail.ru",
		ID:          "346c121df17f4b1f9d483e0460c576a8",
		Pass:        "2e9f7d7887ed4b36b42a443502532744",
		ButtonClass: "fas fa-at",
		GetProfile:  "https://oauth.mail.ru/userinfo",
		Endpoint:    mailru.Endpoint,
		Scopes:      []string{},
	},
}
