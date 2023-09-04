package models

type Message struct {
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreatTime  string `json:"creat_time"`
}
