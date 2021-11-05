package event

// Generated by https://quicktype.io

// MESSAGE_CREATED / DIRECT_MESSAGE_CREATED イベントリクエストボディスキーマ
type MessageEvent struct {
	EventTime string  `json:"eventTime"`
	Message   Message `json:"message"`
}

// メッセージスキーマ
type Message struct {
	ID        string     `json:"id"`
	User      User       `json:"user"`
	ChannelID string     `json:"channelId"`
	Text      string     `json:"text"`
	PlainText string     `json:"plainText"`
	Embedded  []Embedded `json:"embedded"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

// 埋め込み要素スキーマ
type Embedded struct {
	Raw  string `json:"raw"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

// ユーザースキーマ
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IconID      string `json:"iconId"`
	Bot         bool   `json:"bot"`
}

// メッセージテキスト 取得メソッド
func (r *MessageEvent) GetText() string {
	return r.Message.PlainText
}

// チャンネルID 取得メソッド
func (r *MessageEvent) GetChannelID() string {
	return r.Message.ChannelID
}

// ユーザーID 取得メソッド
func (r *MessageEvent) GetUserID() string {
	return r.Message.User.Name
}
