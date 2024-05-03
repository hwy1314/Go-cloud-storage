package db

import (
	mydb "FileStore-Server/db/mysql"
	"fmt"
)

//User: �û���model
type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

//UserSignUp: ͨ���û������������ע��
func UserSignUp(username string,passwd string)bool {
	//Ԥ����sql���
	stmt,err := mydb.DBConn().Prepare(
			"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("Failed to insert, err: "+err.Error())
		return false
	}

	//�������ݿ����
	ret,err := stmt.Exec(username,passwd)
	if err != nil {
		fmt.Println("Failed to insert, err: "+err.Error())
	}

	//�ж�ע����������
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

//UserSignin: ��֤��¼
func UserSignin(username string,encpwd string) bool {
	//Ԥ����sql���
	stmt,err := mydb.DBConn().Prepare(
		"select * from tbl_user where user_name = ? limit 1")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	//�������ݿ��ѯ�����������
	rows,err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Printf("username:%s not found: \n",username)
		return false
	}

	//����ѯ���ת��Ϊmap����
	pRows := mydb.ParseRows(rows)
	//�жϲ�ѯ��������Լ������Ƿ�ƥ��
	if len(pRows)>0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

//UpdateToken: ˢ���û���¼token
func UpdateToken(username string,token string) bool {
	stmt,err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	_,err = stmt.Query(username,token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

//GetUserInfo: ��ѯ�û���Ϣ
func GetUserInfo(username string) (User,error) {
	user := User{}

	stmt,err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")
	defer stmt.Close()
	if err != nil {
		return user,err
	}

	//ִ�в�ѯ
	err = stmt.QueryRow(username).Scan(&user.Username,&user.SignupAt)
	if err != nil {
		return user,err
	}

	return user,nil
}