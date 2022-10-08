package main

import (

    
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

    //"log"
    "fmt"

	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	

)

type Link struct {
	Url string
	ChatID int64
}

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func LinkHandler( c *gin.Context) {

	var link Link

	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
		return

	}

	if err := linkSender(link.Url, link.ChatID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}



	c.JSON(http.StatusOK, link)
	

}


type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func ResponseBot(chatID int64) error {

	viper.SetConfigFile("ENV")
	viper.ReadInConfig()


	viper.AutomaticEnv()

	
	token := fmt.Sprint(viper.Get("TOKEN"))

	strChatId :=strconv.FormatInt(chatID, 10)

	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text: strChatId,
	}
	
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil    

}

func Handler(c *gin.Context) {
	

	body := &webhookReqBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
		return

	}

	
	if err := ResponseBot(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

func linkSender(url string, chatID int64) error {

	viper.SetConfigFile("ENV")
	viper.ReadInConfig()


	viper.AutomaticEnv()

	
	token := fmt.Sprint(viper.Get("TOKEN"))

	

	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text: url,
	}
	
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil    

}
