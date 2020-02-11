package main

import (
	"flag"
	"fmt"

	"mattermost-plugin-svn/matterhook"

	"github.com/jinzhu/configor"
)

// // Item this text Data
// type Data struct {
// 	Text     string `json:"text"`
// 	Username string `json:"username"`
// 	Channel  string `json:"channel"`
// 	IconURL  string `json:"icon_url"`
// }

// //send text Text
// type Text struct {
// 	Text Data `json:"text"`
// }

//Config 配置文件
type Config struct {
	APPName    string `default:"app name"`
	Mattermost struct {
		WebhookURL string
		Iconurl    string `default:"https://svn.apache.org/repos/asf/subversion/trunk/notes/logo/256-colour/subversion_logo-200x173.png"`
		Username   string `default:"Svnbot"`
		Channel    string `default:"Svn"`
	}
	VerifyCert bool
}

var conf Config
var text string

func main() {
	realConf := flag.String("conf", "../conf/config.toml", "conf file")
	//flag.StringVar(&conf, "conf", "../conf/config.toml", "conf file")
	// flag.StringVar(&conf.Mattermost.Username, "username", "", "username")
	flag.StringVar(&text, "text", "123", "text")
	flag.Parse()

	configor.Load(&conf, *realConf)
	fmt.Printf("config: %#v", conf)
	PostSvn(text, &conf)
}

// PostSvn is post data to mattermost hook url
func PostSvn(_text string, conf *Config) {
	fmt.Println(" exe post svn")
	fmt.Println(conf.APPName)
	fmt.Println(conf.Mattermost.WebhookURL)
	fmt.Println(conf.Mattermost.Iconurl)
	fmt.Println(conf.Mattermost.Username)
	fmt.Println(conf.Mattermost.Channel)
	fmt.Println(conf.VerifyCert)
	//提交信息到mattermost
	fmt.Println(" Mattermost POST method, posts text to the Mattermost incoming webhook URL")

	//appName := conf.APPName
	webhookURL := conf.Mattermost.WebhookURL
	iconURL := conf.Mattermost.Iconurl
	userName := conf.Mattermost.Username
	channel := conf.Mattermost.Channel
	//verifyCert := conf.VerifyCert

	message := matterhook.Message{
		Text:      _text,
		Username:  userName,
		Channel:   channel,
		IconEmoji: iconURL,
	}
	// err := matterhook.Send(webhookURL, message, "utizeyqoppb3xjqwpmbjdz761c")
	err := matterhook.Send(webhookURL, message, "")
	if err != nil {
		fmt.Println(err)
	}

	// data := url.Values{}
	// data.Set("text", _text)
	// data.Set("username", userName)
	// data.Set("icon_url", iconURL)
	// data.Set("channel", channel)

	// // postForm形式
	// // data := make(url.Values)
	// // data["text"] = []string{_text}
	// // data["username"] = []string{userName}
	// // data["icon_url"] = []string{iconURL}
	// // data["channel"] = []string{channel}
	// // fmt.Println(data)
	// // resp, err := http.PostForm(webhookURL, data)
	// // if err != nil {
	// // 	fmt.Println(err.Error())
	// // 	return
	// // }
	// // defer resp.Body.Close()
	// // fmt.Println("post send success")

	// // sendData := bytes.NewBuffer(jsonStr)
	// // buf := bytes.NewBuffer(jsonStr)
	// // if err != nil {
	// // 	fmt.Printf("request Encode err %s", err)
	// // 	return
	// // }
	// // //ingore ssl verify
	// // fmt.Println("values  =", data)

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: verifyCert},
	// }
	// client := &http.Client{Transport: tr}
	// // req, err := http.NewRequest("POST", webhookURL, strings.NewReader(data.Encode()))
	// req, err := http.NewRequest("POST", webhookURL, strings.NewReader(data.Encode()))
	// if err != nil {
	// 	fmt.Printf("request post err %s", err)
	// 	return
	// }
	// //req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Printf("response  err %s", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// fmt.Println("res ==== ", resp)
	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("response read err %s", err)
	// 	return
	// }
	// fmt.Printf("response data:%v\n", string(respBody))
}
