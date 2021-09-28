package public

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"

	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

func SendImage(filename, bot_token string) {
	bot := wxworkbot.New(bot_token)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(data)
	md5Str, _ := HashFileMd5(filename)
	// 图片消息
	image := wxworkbot.Image{
		Base64: base64Str,
		MD5:    md5Str,
	}
	err = bot.Send(image)
	if err != nil {
		log.Fatal(err)
	}
}

func HashFileMd5(filePath string) (md5Str string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}
