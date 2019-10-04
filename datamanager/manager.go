package datamanager

import (
	"encoding/json"
	"handbook/config"
	"handbook/entity"
	"handbook/utils"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 加载数据的版本标识
var gDataVersion string = ""

// 格式化后的数据
var gCards []entity.Card
var gShortPages []entity.ShortPage
var gHandBooks []entity.HandBook

// Load 加载数据内容
func Load() {

	// 修改才重新加载
	version := getDataVersion(config.DataDirPath)
	log.Println("Data Version: ", version)
	if version == gDataVersion {
		log.Println("Data Already Loaded")
		return
	}
	log.Println("Reload Data")

	var newCards []entity.Card
	var newShortPages []entity.ShortPage
	var newHandBooks []entity.HandBook

	// 遍历目录下所有子目录
	pathStack := utils.Stack{config.DataDirPath}
	for !pathStack.IsEmpty() {
		curDirPath := pathStack.Pop()
		// 检查当前是否是一个知识区块--是否有名为config.SecDescFile的文件
		exist, _ := utils.IsFileExit(curDirPath + config.SecDescFile)
		if exist {
			log.Println("Parse K Section: ", curDirPath)
			desc := parseDesc(curDirPath)
			newCards = append(newCards, parseCards(curDirPath, desc)...)
			//secShortPages := ParseShortPages(curDirPath)
			//secHandBooks := ParseHandBooks(curDirPath)
		} else {
			_, dirs, _ := utils.GetFilesAndDirsNotRecursive(curDirPath)
			for _, dir := range dirs {
				pathStack.Push(dir.Path)
			}
		}
	}

	// 最后成功后 再更新数据
	gCards = newCards
	gShortPages = newShortPages
	gHandBooks = newHandBooks
	gDataVersion = version
}

func getDataVersion(dirpath string) string {
	f, err := os.Open(dirpath + config.DataVersionFile)
	if err != nil {
		log.Printf("Open File Fail [Err:%s]", err.Error())
		return ""
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(bytes))
}

func parseDesc(dirpath string) (desc entity.Desc) {
	f, err := os.Open(dirpath + config.SecDescFile)
	if err != nil {
		log.Printf("Open file failed [Err:%s]", err.Error())
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&desc)
	if err != nil {
		log.Printf("Decoder failed [Err:%s]", err.Error())
	}
	return desc
}

func parseCards(dirpath string, desc entity.Desc) (cards []entity.Card) {
	files, _, err := utils.GetFilesAndDirsNotRecursive(dirpath + config.CardsDirPath)
	if err != nil {
		log.Println("Get Cards Files Fail ", err.Error())
		return nil
	}

	for _, file := range files {
		card := entity.Card{Desc: desc, Content: file.Content}
		cards = append(cards, card)
	}

	log.Println("Get Cards: ", cards)

	return cards
}

// GetRomdonCards 随机获取几张卡片
func GetRomdonCards() (cards []entity.Card) {
	// 对外函数均先自行加载一遍内容
	Load()

	randomIdx := utils.GenRandNoRepeated(0, len(gCards), config.RandomCardsNumByOneQuery)
	if randomIdx == nil {
		cards = append(cards, gCards...)
	} else {
		for _, idx := range randomIdx {
			cards = append(cards, gCards[idx])
		}
	}

	return cards
}
