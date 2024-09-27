package main

import (
	"bufio"
	"bytes"
	"cmd-gram-cli/models"
	"cmd-gram-cli/view"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

// const for URL-templates and other values
const (
	loginEndpoint    = "/api/login"
	signinEndpoint   = "/api/sign-in"
	newChatEndpoint  = "/api/new-chat"
	allChatsEndpoint = "/api/%d/chats"
	websocketOrigin  = "http://localhost/"
	successStatus    = '2' // need to define response status
)

var ip = flag.String("ip", "0.0.0.0:8080", "Input an IP to connect")

var client = &http.Client{}
var User = &models.User{}
var chatMap = map[int]*models.Chat{}
var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)
var PATH string

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	PATH = path[:34]

	f, err := os.Open(fmt.Sprintf("%s\\text\\gopher.txt", PATH))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var body = make([]byte, 64)
	for {
		n, err := f.Read(body)
		if err == io.EOF {
			break
		}
		fmt.Print(string(body[:n]))
	}
	f.Close()

	fmt.Print("\n\n")
	handleHelp()

}

func main() {
	flag.Parse()

	scanner.Split(bufio.ScanLines)

	for {
		scanner.Scan()
		commandLine := scanner.Text()

		commandArgs := strings.Split(commandLine, " ")

		switch commandArgs[0] {
		case "/login":
			handleLogin()

		case "/reg":
			handleSignin()

		case "/new-chat":
			handleNewChat()

		case "/all-chats":
			handleAllChats()

		case "/open-chat":
			handleOpenChat(commandArgs)
		case "/help":
			handleHelp()
		case "":
		default:
			fmt.Println("Undefined command")
		}
	}
}

// authorisation check
func isAuthenticated() bool {
	return User.Email != ""
}

func handleLogin() {
	if isAuthenticated() {
		fmt.Println("You are already logged in.")
		return
	}

	url := fmt.Sprintf("http://%s%s", *ip, loginEndpoint)
	method := http.MethodPost

	jsonData, err := login()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := sendReq(jsonData, method, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	handleResponse(resp, &User, "Welcome")
}

func handleSignin() {
	if isAuthenticated() {
		fmt.Println("You are already logged in.")
		return
	}

	url := fmt.Sprintf("http://%s%s", *ip, signinEndpoint)
	method := http.MethodPost

	jsonData, err := reg()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := sendReq(jsonData, method, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	handleResponse(resp, &User, "Welcome")
}

func handleNewChat() {
	if !isAuthenticated() {
		fmt.Println("Please log into your account.")
		return
	}

	url := fmt.Sprintf("http://%s%s", *ip, newChatEndpoint)
	method := http.MethodPost

	jsonData, err := newChat()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := sendReq(jsonData, method, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	handleResponse(resp, &models.Chat{}, "Chat created")
}

func handleAllChats() {
	if !isAuthenticated() {
		fmt.Println("Please log into your account.")
		return
	}

	url := fmt.Sprintf("http://%s"+allChatsEndpoint, *ip, User.ID)
	method := http.MethodGet

	resp, err := sendReq([]byte{}, method, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var chats map[string][]models.Chat
	handleResponse(resp, &chats, "")
	for i, chat := range chats["chat"] {
		fmt.Println(i+1, chat.Name)
		chatMap[i+1] = &chat
	}
}

func handleOpenChat(commandArgs []string) {
	if !isAuthenticated() {
		fmt.Println("Please log into your account.")
		return
	}
	if len(commandArgs) < 2 {
		fmt.Println("You didn`t choose the chat")
		return
	}
	chatIndex, err := strconv.Atoi(commandArgs[1])
	if err != nil {
		fmt.Println("Invalid argument.")
		return
	}

	url := fmt.Sprintf("ws://%s/api/%d/chats/%d", *ip, User.ID, chatIndex)
	conn, err := websocket.Dial(url, "tcp", websocketOrigin)
	if err != nil {
		fmt.Println(err)
		return
	}

	messages := []*models.MessageDTO{}
	err = websocket.JSON.Receive(conn, &messages)
	if err != nil {
		fmt.Println("Failed to receive messages:", err)
		err = startMessaging(chatIndex, conn)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	for _, message := range messages {
		view.Messages(message, User)
	}
	fmt.Println("/exit - to exit")
	err = startMessaging(chatIndex, conn)
	if err != nil {
		fmt.Println(err)
	}
}

func startMessaging(cid int, conn *websocket.Conn) error {
	go func(c *websocket.Conn) {
		msg := &models.MessageDTO{}
		for {
			if !c.IsServerConn() {
				return
			}
			websocket.JSON.Receive(conn, msg)

			view.Messages(msg, User)
		}
	}(conn)

	for {
		msg := &models.MessageDTO{UserID: User.ID, ChatID: cid, Time: time.Now()}

		scanner.Scan()

		b := scanner.Text()
		if b == "" {
			continue
		}
		msg.Body = b

		jsonData, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		_, err = conn.Write(jsonData)
		if err != nil {
			return err
		}

		if b == "/exit" {
			break
		}
	}
	return nil
}

func handleResponse(resp *http.Response, target interface{}, successMessage string) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	if resp.Status[0] != successStatus {
		var apiError models.Error
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			fmt.Println("Failed to unmarshal error response:", err)
			return
		}
		fmt.Println(apiError.Text)
	} else {
		err = json.Unmarshal(body, target)
		if err != nil {
			fmt.Println("Failed to unmarshal success response:", err)
			return
		}
		fmt.Println(successMessage)
	}
}

func sendReq(jsonData []byte, meth string, url string) (*http.Response, error) {
	if jsonData == nil && meth != http.MethodGet {
		return nil, fmt.Errorf("data is nil")
	}
	req, err := http.NewRequest(meth, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("somthing went wrong, check your connection %v", err)

	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can`t do req due to error: %v", err)
	}

	return resp, nil
}

func login() ([]byte, error) {
	gotU := &models.UserDTO{}

	fmt.Println("email: ")
	fmt.Scan(&gotU.Email)
	fmt.Println("password: ")
	fmt.Scan(&gotU.EncryptedPassword)

	jsonData, err := json.Marshal(gotU)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}

func reg() ([]byte, error) {
	gotU := &models.UserDTO{}

	fmt.Println("email: ")
	fmt.Scan(&gotU.Email)
	fmt.Println("password: ")
	fmt.Scan(&gotU.EncryptedPassword)

	jsonData, err := json.Marshal(gotU)
	if err != nil {
		return nil, err
	}

	return jsonData, err
}

func newChat() ([]byte, error) {
	var u1 = &models.UserDTO{Email: User.Email}
	var u2 = &models.UserDTO{}

	fmt.Println("Intput the user")
	fmt.Scan(&u2.Email)

	jsonData, err := json.Marshal([]*models.UserDTO{u1, u2})
	if err != nil {
		return nil, err
	}

	return jsonData, err
}

func handleHelp() {
	f, err := os.Open(fmt.Sprintf("%s/text/help.txt", PATH))
	if err != nil {
		fmt.Println(err)
		return
	}
	body := make([]byte, 128)
	for {
		n, err := f.Read(body)
		if err == io.EOF {
			break
		}
		fmt.Print(string(body[:n]))
	}
	f.Close()
	fmt.Print("\n\n")
}
