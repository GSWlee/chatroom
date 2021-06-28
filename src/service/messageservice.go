package service

import (
	"chatroom/src/dataorm"
	"sort"
	"strconv"
)

type HistorySlice []dataorm.History

func (s HistorySlice) Len() int      { return len(s) }
func (s HistorySlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s HistorySlice) Less(i, j int) bool {
	return s[i].Time.Before(s[j].Time)
}

type messageService interface {
	Send(userid int, roomid int, message string) error
	Retrieve(roomid int) ([]dataorm.History, error)
}

type MessageService struct {
}

//@function: 发送信息
//@param1: userid: 用户ID
//@param2: roomid: 房间ID
//@param3: message: 信息内容
//@return：如果删除失败，返回错误信息

func (s MessageService) Send(userid int, roomid int, message string) error {
	history := dataorm.History{
		Userid: userid,
		Roomid: roomid,
		Data:   message,
	}
	if err := dataorm.Insert(history); err != nil {
		return err
	}
	return nil
}

//@function: 接收信息
//@param2: roomid: 房间ID
//@return：如果删除失败，返回错误信息,如果成功，返回History数组,按信息时间从早到晚排序

func (s MessageService) Retrieve(roomid int) ([]dataorm.History, error) {
	values, err := dataorm.Query("History", nil, []string{"roomid"}, []string{strconv.Itoa(roomid)})
	if err != nil {
		return nil, err
	}
	historys, _ := values.([]dataorm.History)
	history := HistorySlice(historys)
	sort.Sort(history)
	historys = []dataorm.History(history)
	return historys, nil
}
