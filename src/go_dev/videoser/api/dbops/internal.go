package dbops

import (
	"database/sql"
	"go_dev/videoser/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions (session_id,TTL,
		login_name) VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(`SELECT login_name,ttl FROM sessions
	WHERE session_id = ?`)
	if err != nil {
		return nil, err
	}
	var un string
	var ttl string
	err = stmtOut.QueryRow(sid).Scan(&un, &ttl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var ttlint int64

	if ttlint, err = strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = ttlint
		ss.Username = un
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil

}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrive sessions error:%s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id:%s,ttl:%d", id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtOut.Exec(sid)
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}
