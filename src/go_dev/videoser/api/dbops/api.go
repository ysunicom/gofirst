package dbops

import (
	"database/sql"
	"go_dev/videoser/api/defs"
	"go_dev/videoser/api/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	stmOut.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("Delete user error:%s", err)
		return err
	}
	_, err = stmDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	stmDel.Close()
	return nil

}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("JAN 2 2006,15:04:05")

	stmIns, err := dbConn.Prepare(`INSERT INTO video_info 
	(id,author_id,name,display_ctime) VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmIns.Close()
	return res, nil

}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmOut, err := dbConn.Prepare(`SELECT author_id,name,
	display_ctime FROM video_info WHERE id = ? `)
	if err != nil {
		return nil, err
	}
	var aid int
	var dct string
	var name string

	err = stmOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmDel, err := dbConn.Prepare(`DELETE FROM video_info
	 WHERE id = ?`)
	if err != nil {
		log.Printf("DELETE user error:%s", err)
		return err
	}
	_, err = stmDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmDel.Close()

	return nil

}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare(`INSERT INTO comments (id,
		video_id,author_id,content) VALUES(?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content FROM comments
	 INNER JOIN users ON comments.author_id = users.id
	 WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil

}
