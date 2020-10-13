package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type gincon gin.Context

func (c *gincon) ParseUserLangs() []LangQ {
	langstr := strings.Join(c.Request.Header["Accept-Language"], ",")
	return parseAcceptLanguage(langstr)
}

type LangQ struct {
	Lang string
	Q    float64
}

func parseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := LangQ{langQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			lq := LangQ{langQ[0], q}
			lqs = append(lqs, lq)
		}
	}
	return lqs
}

func vendorsConfInit() VendorsConfMap {
	authconfs := make(VendorsConfMap)

	for _, vendor := range vendorslist {
		authconfs[vendor] = &oauth2.Config{
			ClientID:     authvendors[vendor].ID,
			ClientSecret: authvendors[vendor].Pass,
			RedirectURL:  authcallback,
			Scopes:       authvendors[vendor].Scopes,
			Endpoint:     authvendors[vendor].Endpoint,
		}
	}

	return authconfs
}

func parseAppOpts() (opts AppOpts) {
	if len(os.Args) > 1 && os.Args[1] != "" && os.Args[1][:1] != "-" {
		opts.Command = os.Args[1]
	}

	flag.IntVar(&opts.Port, "p", defaultport, "Custom app port")
	flag.BoolVar(&opts.Test, "t", false, "Test mode")
	flag.StringVar(&opts.Host, "h", defaulthost, "Host with proto [+ port]")
	flag.StringVar(&opts.ConfFile, "c", "", "Config file path")
	flag.StringVar(&opts.DataCallback, "d", defaultdatacallback, "Data callback")
	flag.Parse()

	return
}

func readConfig() (conf Config) {
	if appOpts.ConfFile != "" {
		fpath := ""

		if !strings.Contains(appOpts.ConfFile, "/") {
			fpath += "../../ini/"
			fpath += appOpts.ConfFile
			if !strings.Contains(appOpts.ConfFile, ".json") {
				fpath += ".json"
			}
			appOpts.ConfFile = curDir() + "/" + fpath
		}

		err := readJSONfile(appOpts.ConfFile, &conf)
		if err != nil {
			fmt.Println("Error read config:", err)
			os.Exit(1)
		}
	}

	return
}

func setup() {
	// автоматическая настройка портов
	addressProxy := conf["address_proxy_pointsrc"]
	portInterval := conf["port_auto_interval"]
	port := conf["port_service"]

	if appOpts.Port != defaultport {
		// set in console param
		return
	}

	if port != "" {
		var err error
		appOpts.Port, err = strconv.Atoi(port)
		if err != nil {
			fmt.Println("Error port format in config")
			os.Exit(1)
		}
	}

	if port == "" && addressProxy != "" && portInterval != "" {
		// запрашиваем порт у указанного прокси-сервера
		fmt.Println("Get port from proxy...")
		proxy_url := addressProxy + "port?interval=" + portInterval
		res, err := http.Get(proxy_url)
		if err != nil {
			fmt.Println("Error get port from proxy")
			os.Exit(1)
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error read proxy response")
			os.Exit(1)
		}

		if appOpts.Test {
			fmt.Println("Proxy get port response:", string(b))
		}

		bbres := BBResponse{}
		err = json.Unmarshal(b, &bbres)
		if err != nil {
			fmt.Println("Error parsing BB response:", err)
			os.Exit(1)
		}

		switch bbres.Data.(type) {
		case string:
			port = bbres.Data.(string)
			var err error
			appOpts.Port, err = strconv.Atoi(port)
			if err != nil {
				fmt.Println("Error port format from proxy")
				os.Exit(1)
			}

		default:
			fmt.Println("Error BB response format")
			os.Exit(1)
		}
	}

	fmt.Println("App port:", appOpts.Port)
}
