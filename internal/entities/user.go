package entities

type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    Email     string `json:"email"`
    Password  string `json:"pass"`
}

type UserTelegram struct {
    UserID  string `json:"user_id"`
    TUserID string `json:"tuser_id"`
    ChatID  string `json:"chat_id"`
}