package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempId string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("yangshuo", "123")
	if err != nil {
		t.Errorf("Error of AddUser :%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("yangshuo")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("yangshuo", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser :%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("yangshuo")
	if err != nil {
		t.Errorf("Error of RegetUser :%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting User fail")
	}
}

func TestVideoInfoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddNewVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my_video")
	if err != nil {
		t.Errorf("Error of AddNewVide:%v", err)
	}
	tempId = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempId)
	if err != nil {
		t.Errorf("Error of GetVideo:%v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempId)
	if err != nil {
		t.Errorf("Error of DeleteVideo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempId)
	if err != nil && vi != nil {
		t.Errorf("Error of RegetVideo:%v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Prepareuser", testAddUser)
	t.Run("PrepareVideo", testAddVideoInfo)
	t.Run("AddNewComment", testAddComments)
	t.Run("ListComment", testListComment)
}

func testAddComments(t *testing.T) {
	err := AddNewComments(tempId, 1, "my_content")
	if err != nil {
		t.Errorf("Error of AddNewComments:%v", err)
	}
}

func testListComment(t *testing.T) {
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(tempId, from, to)
	if err != nil {
		t.Errorf("Error of ListComments:%v", err)
	}

	for i, v := range res {
		fmt.Printf("comment: %d, %v \n", i, v)
	}

}
