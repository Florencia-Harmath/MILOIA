package models

import (
    "gorm.io/gorm"
)

type Message struct {
    gorm.Model
    Content   string `gorm:"type:text;not null"`
    UserID    uint   `gorm:"not null"`
    User      RegisterUser   `gorm:"foreignKey:UserID"`
    RoomID    string `gorm:"not null"`
}

func CreateMessage(db *gorm.DB, content string, userID uint, roomID string) (*Message, error) {
    message := &Message{
        Content: content,
        UserID:  userID,
        RoomID:  roomID,
    }
    if err := db.Create(message).Error; err != nil {
        return nil, err
    }
    return message, nil
}

func GetMessagesByRoomID(db *gorm.DB, roomID string) ([]Message, error) {
    var messages []Message
    if err := db.Where("room_id = ?", roomID).Order("created_at asc").Find(&messages).Error; err != nil {
        return nil, err
    }
    return messages, nil
}
