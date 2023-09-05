package dao

import (
	"douyin/models"
	"errors"
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"
)

type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *models.UserInfo) error {
	if userinfo == nil {
		return errors.New("空指针错误")
	}
	//DB.Where("id=?",userId).First(userinfo)
	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

func (u *UserInfoDAO) AddUserInfo(userinfo *models.UserInfo, DB *gorm.DB) error {
	if userinfo == nil {
		return errors.New("空指针错误")
	}
	return DB.Create(userinfo).Error
}

func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo models.UserInfo
	if err := DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}
func (u *UserInfoDAO) AddUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count+1 WHERE id = ?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations` (`user_info_id`,`follow_id`) VALUES (?,?)", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserInfoDAO) CancelUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count-1 WHERE id = ? AND follow_count>0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count-1 WHERE id = ? AND follower_count>0", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_info_id=? AND follow_id=?", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserInfoDAO) GetFollowListByUserId(userId int64, userList *[]*models.UserInfo) error {
	if userList == nil {
		return errors.New("空指针错误")
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].Id == 0 {
		return errors.New("用户列表为空")
	}
	return nil
}

func (u *UserInfoDAO) GetFollowerListByUserId(userId int64, userList *[]*models.UserInfo) error {
	if userList == nil {
		return errors.New("空指针错误")
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id = ? AND r.user_info_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	//if len(*userList) == 0 || (*userList)[0].Id == 0 {
	//	return ErrEmptyUserList
	//}
	return nil
}

//redis相关操作

// UpdateUserRelation 更新关注状态，state:true为点关注，false为取消关注
func (i *IndexMap) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}

// GetUserRelation 得到关注状态
func (i *IndexMap) GetUserRelation(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}

func (u *UserInfoDAO) GetFriendsListByUserId(userId int64, userList *[]*models.UserInfo) error {
	if userList == nil {
		return errors.New("空指针错误")
	}

	var err error
	if err = DB.Raw(`
    SELECT u.* 
    FROM user_infos u
    JOIN user_relations r1 ON u.id = r1.follow_id AND r1.user_info_id = ?
    JOIN user_relations r2 ON u.id = r2.user_info_id AND r2.follow_id = ?
`, userId, userId).Scan(userList).Error; err != nil {
		return err
	}

	if len(*userList) == 0 || (*userList)[0].Id == 0 {
		return errors.New("用户列表为空")
	}
	return nil
}
