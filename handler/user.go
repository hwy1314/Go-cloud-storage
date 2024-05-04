package handler

import (
	"FileStore-Server/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	dblayer "FileStore-Server/db"
	"time"
)

const (
	//���ڼ��ܵ���ֵ(�Զ���)
	pwdSalt = "*#890"
)

//SignupHandler : �����û�ע������
func SignupHandler(w http.ResponseWriter,r *http.Request) {
	//GET: ҳ������
	if r.Method == http.MethodGet {
		data,err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	//POST: ��������
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		passwd := r.Form.Get("password")

		if len(username) < 3 || len(passwd) < 5 {
			w.Write([]byte("Invalid Parameter"))
			return
		}

		//��������м��μ�ȡSha1ֵ����
		encPasswd := util.Sha1([]byte(passwd+pwdSalt))
		ok := dblayer.UserSignUp(username,encPasswd)
		if ok {
			w.Write([]byte("SUCCESS"))
		}else {
			w.Write([]byte("FAILED"))
		}
	}
}

//SignInHandler: ��¼�ӿ�
func SignInHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//1.У���û���������
	encPasswd := util.Sha1([]byte(password + pwdSalt))
	pwdChecked := dblayer.UserSignin(username,encPasswd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	//2.�����û�����ƾ֤Token
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username,token)
	if !upRes {
		fmt.Println("���·���ƾ֤ʧ��")
		w.Write([]byte("FAILED"))
		return
	}

	//3.��װƾ֤����Ӧ��Ϣ���ͻ���
	resp := util.RespMsg{
		Code:0,
		Msg:"OK",
		Data: struct {
			Location	string
			Username	string
			Token		string
		}{
			Location:"http://"+r.Host+"/static/view/home.html",
			Username:username,
			Token:token,
		},
	}
	w.Write(resp.JSONBytes())
}

//UserInfoHandler: ��ѯ�û���Ϣ
func UserInfoHandler(w http.ResponseWriter,r *http.Request) {
	//1.�����������
	r.ParseForm()
	username := r.Form.Get("username")

	//���¹����Ѿ�������������
	////2.��֤Token�Ƿ���Ч
	//isValidToken := IsTokenValid(token)
	//if !isValidToken {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	//3.��ѯ�û���Ϣ
	user,err := dblayer.GetUserInfo(username)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//4.��װ����Ӧ�û�����
	resp := util.RespMsg{
		Code:0,
		Msg:"OK",
		Data:user,
	}
	w.Write(resp.JSONBytes())
}

//GenToken: ����40λtoken
func GenToken(username string) string {
	//40λtoken: md5(username + timestamp(ʱ���) + token_salt) + timestamp[:8]
	ts := fmt.Sprint("%x",time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username+ts+"_tokensalt"))
	return tokenPrefix + ts[:8]
}

//IsTokenValid: ��֤token�Ƿ�ʧЧ
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		fmt.Println("token�����쳣")
		return false
	}
	return true
}