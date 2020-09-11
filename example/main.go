package main

import (
	linebotcf "github.com/Tony-Liou/linebot-cloud-function"
)

func main() {
	targetID := "Replace this!!!"
	linebotcf.PushMessage("Push messsage", targetID)
}
