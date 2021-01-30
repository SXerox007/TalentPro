package models

import db "talentpro/base/db/postgres"

type UserDetails struct {
	TableName struct{} `sql:"all_users" json:"-" csv:"-"`
	Id        string   `param:"id" sql:"id" json:"id"`
	LoginId   string   `param:"loginId" json:"loginId" sql:"login_id"`
	FullName  string   `param:"fullName" json:"fullName" sql:"full_name"`
	State     int      `param:"state" json:"state" sql:"state"`
}

const (
	DELETE_STATE = 2
)

// get user list
func GetUserList(limit, offset int) (data []UserDetails, err error) {
	err = db.GetPGClient().Model(&data).
		Limit(limit).
		Offset(offset).
		Order("uts desc").
		Where("state != ?", DELETE_STATE).
		Select()
	return
}

// insert user details into the all_users table
func InsertUserIntoAllUsers(data UserDetails) error {
	return db.GetPGClient().Create(&data)
}

// delete user from all_users (soft delete)
func DelteUserFromAllUsers(id string) error {
	_, err := db.GetPGClient().Model(&UserDetails{}).
		Set("state = ?", DELETE_STATE).
		Where("id = ?", id).
		Update()
	return err
}

// update user details
func UpdateUserDetails(data UserDetails) error {
	var err error
	if data.FullName != "" {
		err = UpdateUserDetail(data.Id, "full_name", data.FullName)
	}
	if data.LoginId != "" {
		err = UpdateUserDetail(data.Id, "login_id", data.LoginId)
	}
	return err
}

func UpdateUserDetail(id, setfield, setFieldValue string) error {
	_, err := db.GetPGClient().Model(&UserDetails{}).
		Set(setfield+" = ?", setFieldValue).
		Where("id = ?", id).
		Update()
	return err
}
