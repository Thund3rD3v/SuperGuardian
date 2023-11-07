package database

import (
	"fmt"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"gorm.io/gorm"
)

func MemberExists(db *gorm.DB, id string) bool {
	var member structs.Member
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		fmt.Println("Error Checking If Member Exists:", err.Error())
		return false
	}

	return true
}

func GetMember(db *gorm.DB, id string) (structs.Member, error) {
	var member structs.Member
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		return member, err
	}

	return member, nil
}

func GetMemberByCursor(db *gorm.DB, cursor int, amount int) ([]structs.Member, error) {
	var members []structs.Member

	if err := db.Offset(cursor).Order("level DESC").Limit(amount).Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func GetMembersCount(db *gorm.DB) int64 {
	var count int64
	db.Model(structs.Member{}).Count(&count)
	return count
}

func CreateMember(db *gorm.DB, id string, level int, xp int) error {
	member := structs.Member{
		Id:            id,
		Level:         level,
		Xp:            xp,
		MessageSentAt: time.Now().UnixMilli(),
	}

	if err := db.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMember(db *gorm.DB, member *structs.Member) error {
	if err := db.Save(member).Error; err != nil {
		return err
	}

	return nil
}
