package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"

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

// type Text struct {
// 	Text Data `json:"text"`
// }

// Group 项目组
type Group struct {
	Projectnames   []string
	Projectpaths   []string
	Authors        []string
	Channel        string
	Username       string
	WebhookURL     string
	FormatComment  bool
	FormatFileList bool
}

//Config 配置文件
type Config struct {
	APPName string `default:"app name"`
	// Mattermost struct {
	// 	WebhookURL string
	// 	Iconurl    string `default:"https://svn.apache.org/repos/asf/subversion/trunk/notes/logo/256-colour/subversion_logo-200x173.png"`
	// 	Username   string `default:"Svnbot"`
	// 	Channel    string `default:"Svn"`
	// }
	VerifyCert bool
	Groups     map[string]Group
}

var conf Config // Config 配置文件
var (
	rev         string // text 发送文本
	author      string
	projectPath string
	sendType    string // path author
	comments    string // commit comments
	filelist    string // commit file list
	sysType     string //系统类型
)

func main() {
	realConf := flag.String("conf", "../conf/config.toml", "conf file")
	// flag.BoolVar(&conf.Groups.Group.FormatComment, "formatcomment", true, "FormatComment")
	flag.StringVar(&rev, "rev", "", "text")
	flag.StringVar(&author, "author", "default", "username")
	flag.StringVar(&projectPath, "projectpath", "", "svn project path")
	flag.StringVar(&sendType, "sendtype", "author", "send type eg. path author")
	flag.StringVar(&comments, "comments", "", "commit comments")
	flag.StringVar(&filelist, "filelist", "", "commit file list")
	flag.Parse()

	configor.Load(&conf, *realConf)
	fmt.Printf("config: %#v", conf)
	fmt.Printf("rev :%s", rev)
	fmt.Printf("comments :%s", comments)
	fmt.Printf("filelist :%s", filelist)
	sysType := runtime.GOOS

	if sysType == "linux" {
		// LINUX系统
	}

	if sysType == "windows" {
		// windows系统
	}
	//tfilelist := " A   444.txt A   5555.txt D   tes111.txt U   test.txt A   test1.txt A   trunk/proxy/666.txt A   trunk/web/666.txt U   trunk/web/aaa.txt A   trunk/web_custom/333333.txt"
	if sendType == "path" {
		PostSvnByPath(&conf)
	} else { //其他全部按照人员发
		PostSvnByAuthor(&conf)
	}
}

// 格式化  [___] sad多行 [___] 2.多行 [___] 3.卡里的机房 [___] 5大框架房 [___] 5会计暗坑 [___] 5卡就看见了 [___] 7拉
func formatComments(commentstr string) string {
	result := strings.ReplaceAll(commentstr, "[___]", "\n")
	return result
}

// 格式化  [___] sad多行 [___] 2.多行 [___] 3.卡里的机房 [___] 5大框架房 [___] 5会计暗坑 [___] 5卡就看见了 [___] 7拉
func notFormatComments(commentstr string) string {
	result := strings.ReplaceAll(commentstr, "[___]", " ")
	return result
}

// PostSvnByPath svn by path
func PostSvnByPath(conf *Config) {
	fmt.Println(" post svn by path start")
	tmpComments := ""
	tmpFilelist := ""
	for key, value := range conf.Groups {
		fmt.Println("key=", key, "value=", value)
		curChannel := value.Channel
		curUsername := value.Username
		webhook := value.WebhookURL
		isExist := inArray(projectPath, value.Projectpaths)
		if sysType == "windows" {
			if value.FormatComment {
				tmpComments = formatComments(comments)
			} else {
				tmpComments = notFormatComments(comments)
			}
			if value.FormatFileList {
				tmpFilelist = formatComments(filelist)
			} else {
				tmpFilelist = notFormatComments(filelist)
			}
		} else {
			tmpComments = comments
			tmpFilelist = filelist
		}
		// if 全路径匹配就发送
		if isExist {
			SendMessageToMattermost(curUsername, curChannel, webhook, tmpComments, tmpFilelist)
		} else {
			isExist := likeArray(projectPath, value.Projectnames)
			//如果 全路径不匹配，但是路径包含项目名，也发送
			if isExist {
				SendMessageToMattermost(curUsername, curChannel, webhook, tmpComments, tmpFilelist)
			}
		}
	}
	fmt.Println(" post svn by author end")
	return
}

// PostSvnByAuthor svn by author
func PostSvnByAuthor(conf *Config) {
	fmt.Println(" post svn by author start")
	tmpComments := ""
	tmpFilelist := ""
	for key, value := range conf.Groups {
		fmt.Println("key=", key, "value=", value)
		curUsername := value.Username
		curChannel := value.Channel
		webhook := value.WebhookURL
		isExist := inArray(author, value.Authors)
		if isExist {
			if sysType == "windows" {
				if value.FormatComment {
					tmpComments = formatComments(comments)
				} else {
					tmpComments = notFormatComments(comments)
				}
				if value.FormatFileList {
					tmpFilelist = formatComments(filelist)
				} else {
					tmpFilelist = notFormatComments(filelist)
				}
			} else {
				tmpComments = comments
				tmpFilelist = filelist
			}
			SendMessageToMattermost(curUsername, curChannel, webhook, tmpComments, tmpFilelist)
		}
	}
	fmt.Println(" post svn by author end")
	return
}
func inArray(input string, arrays []string) bool {
	length := len(arrays)
	for i := 0; i < length; i++ {
		item := arrays[i]
		if item == input {
			return true
		}
	}
	return false
}
func likeArray(key string, arrays []string) bool {
	length := len(arrays)
	for i := 0; i < length; i++ {
		item := arrays[i]
		// 数组中如果都存在
		if strings.Index(key, item) != -1 { //存在子串
			return true
		}
	}
	return false
}

// SendMessageToMattermost 发送消息到 mattermost
func SendMessageToMattermost(userName, channel, webhookURL, _comments, _filelist string) {
	message := matterhook.Message{
		Text: "@channel 提交人：" + author + "  **版本号：  " + rev + "**  \n**仓库地址:" +
			projectPath + "** :+1: :smile: \n**提交内容**：" + _comments + "\n",
		Username: userName,
		Channel:  channel,
	}

	att := matterhook.Attachment{
		//Fallback: "SVN Commit Log",
		Text: _filelist,
		// Color:      "#FF8000", // 黄色
		Color: "#00FF00", // 绿色
		// Fields: []matterhook.Field{
		// 	{
		// 		Title: "changed file",
		// 		Value: filelist,
		// 		Short: false,
		// 	},
		// },
		//Title: "detail:",
		//AuthorName: userName,
	}
	message.AddAttachments([]matterhook.Attachment{att})
	// err := matterhook.Send(webhookURL, message, "utizeyqoppb3xjqwpmbjdz761c")
	err := matterhook.Send(webhookURL, message, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send ok")
	return
}

// PostSvn is post data to mattermost hook url
// func PostSvn(_text string, conf *Config) {
// 	fmt.Println(" exe post svn")
// 	fmt.Println(conf.APPName)
// 	fmt.Println(conf.Mattermost.WebhookURL)
// 	fmt.Println(conf.Mattermost.Iconurl)
// 	fmt.Println(conf.Mattermost.Username)
// 	fmt.Println(conf.Mattermost.Channel)
// 	fmt.Println(conf.VerifyCert)
// 	//提交信息到mattermost
// 	fmt.Println(" Mattermost POST method, posts text to the Mattermost incoming webhook URL")

// 	//appName := conf.APPName
// 	webhookURL := conf.Mattermost.WebhookURL
// 	//iconURL := conf.Mattermost.Iconurl
// 	userName := conf.Mattermost.Username
// 	channel := conf.Mattermost.Channel
// 	//verifyCert := conf.VerifyCert

// 	message := matterhook.Message{
// 		Text:      _text,
// 		Username:  userName,
// 		Channel:   channel,
// 		IconEmoji: ":papa:",
// 	}
// 	// err := matterhook.Send(webhookURL, message, "utizeyqoppb3xjqwpmbjdz761c")
// 	err := matterhook.Send(webhookURL, message, "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// data := url.Values{}
// 	// data.Set("text", _text)
// 	// data.Set("username", userName)
// 	// data.Set("icon_url", iconURL)
// 	// data.Set("channel", channel)

// 	// // postForm形式
// 	// // data := make(url.Values)
// 	// // data["text"] = []string{_text}
// 	// // data["username"] = []string{userName}
// 	// // data["icon_url"] = []string{iconURL}
// 	// // data["channel"] = []string{channel}
// 	// // fmt.Println(data)
// 	// // resp, err := http.PostForm(webhookURL, data)
// 	// // if err != nil {
// 	// // 	fmt.Println(err.Error())
// 	// // 	return
// 	// // }
// 	// // defer resp.Body.Close()
// 	// // fmt.Println("post send success")

// 	// // sendData := bytes.NewBuffer(jsonStr)
// 	// // buf := bytes.NewBuffer(jsonStr)
// 	// // if err != nil {
// 	// // 	fmt.Printf("request Encode err %s", err)
// 	// // 	return
// 	// // }
// 	// // //ingore ssl verify
// 	// // fmt.Println("values  =", data)

// 	// tr := &http.Transport{
// 	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: verifyCert},
// 	// }
// 	// client := &http.Client{Transport: tr}
// 	// // req, err := http.NewRequest("POST", webhookURL, strings.NewReader(data.Encode()))
// 	// req, err := http.NewRequest("POST", webhookURL, strings.NewReader(data.Encode()))
// 	// if err != nil {
// 	// 	fmt.Printf("request post err %s", err)
// 	// 	return
// 	// }
// 	// //req.Header.Set("Content-Type", "application/json")
// 	// req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
// 	// resp, err := client.Do(req)
// 	// if err != nil {
// 	// 	fmt.Printf("response  err %s", err)
// 	// 	return
// 	// }
// 	// defer resp.Body.Close()
// 	// fmt.Println("res ==== ", resp)
// 	// respBody, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Printf("response read err %s", err)
// 	// 	return
// 	// }
// 	// fmt.Printf("response data:%v\n", string(respBody))
// }

// 因为返回的文件是一行 A aa.txt U bb.txt 这种形式 不好看，格式化一下
func formatChangedFilelist(filestr string) string {
	lists := strings.Split(filestr, " ")
	len := len(lists)
	result := ""
	j := 0
	for i := 0; i < len; i++ {
		item := lists[i]
		if strings.Trim(item, " ") != "" {
			if j&1 == 0 { //如果是偶数
				result = result + " " + item
			} else {
				result = result + " " + item + "\n"
			}
			j++
		}
	}
	// fmt.Println("result=", result)
	return result
}
