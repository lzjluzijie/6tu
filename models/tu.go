package models

import (
	"errors"
	"log"
	"time"
)

type Tu struct {
	ID     int64 `xorm:"notnull pk autoincr"`
	Name   string
	Size   int64
	MD5    string `xorm:"'md5'"`
	SHA256 string `xorm:"'sha256'"`

	IP         string
	DeleteCode string

	Requests uint64

	Width  int
	Height int

	OneDriveFolderID   string
	OneDriveID         string
	OneDriveURL        string
	OneDriveWebPID     string `xorm:"'one_drive_webp_id'"`
	OneDriveWebPURL    string `xorm:"'one_drive_webp_url'"`
	OneDriveFHDID      string `xorm:"'one_drive_fhd_id'"`
	OneDriveFHDURL     string `xorm:"'one_drive_fhd_url'"`
	OneDriveFHDWebPID  string `xorm:"'one_drive_fhd_webp_id'"`
	OneDriveFHDWebPURL string `xorm:"'one_drive_fhd_webp_url'"`

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func InsertTu(tu *Tu) (err error) {
	_, err = x.Insert(tu)
	return
}

func UpdateTu(tu *Tu) (err error) {
	_, err = x.ID(tu.ID).Update(tu)
	return
}

func DeleteTu(tu *Tu) (err error) {
	_, err = x.ID(tu.ID).Delete(tu)
	return
}

func DeleteByCode(dc string) (err error) {
	tu := &Tu{DeleteCode: dc}
	has, err := x.Get(tu)
	if err != nil {
		return
	}

	if !has {
		err = errors.New("not found")
		return
	}

	_, err = x.ID(tu.ID).Delete(tu)
	return
}

func GetTu(t *Tu) (has bool, tu *Tu, err error) {
	tu = &Tu{}

	if t.ID != 0 {
		has, err = x.ID(t.ID).Get(tu)
	} else {
		has, err = x.Get(tu)
	}

	if err != nil {
		return
	}

	if !has {
		return
	}

	go func() {
		tu.Requests++
		_, err := x.Update(tu)
		if err != nil {
			log.Println(err.Error())
		}
	}()
	return
}

func GetUploadHistory(ip string) (t []*Tu, err error) {
	t = make([]*Tu, 0)
	//todo fix timezone, this is -2 hour if server is CST
	err = x.Where("ip = ? and created_at > ?", ip, time.Now().Add(-10*time.Hour)).Find(&t)
	return
}
