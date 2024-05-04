package handler

import (
	"FileStore-Server/db"
	"FileStore-Server/meta"
	"FileStore-Server/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//UploadHandler: �ļ��ϴ�·��
func UploadHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method == "GET" {
		//�����ϴ��ļ���htmlҳ��
		data,err := ioutil.ReadFile("./static/view/index.html")

		if err != nil {
			io.WriteString(w,fmt.Sprint("internel server error: err: %s",err.Error()))
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method == "POST" {
		//�����ļ������洢������Ŀ¼

		//��ȡ���ϴ����ļ�������
		file,head,err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			fmt.Printf("Failed to get data, err: %s\n",err.Error())
			return
		}

		//�����ļ�Ԫ��Ϣʵ��
		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"./static/tempFiles/"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		//���������ļ�
		localFile,err := os.Create(fileMeta.Location)
		defer localFile.Close()
		if err != nil {
			fmt.Printf("Failed to create file, err: %s\n",err.Error())
			return
		}

		//�����ļ���Ϣ�������ļ�
		fileMeta.FileSize,err = io.Copy(localFile,file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err: %s\n",err.Error())
			return
		}

		//�����ļ���ϣֵ
		localFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(localFile)
		//���ļ�Ԫ��Ϣ��ӵ�mysql��
		meta.UpdateFileMetaDB(fileMeta)

		//�����û��ļ����¼
		r.ParseForm()
		username := r.Form.Get("username")
		ok := db.OnUserFileUploadFinished(username,fileMeta.FileSha1,fileMeta.FileName,fileMeta.FileSize)
		if ok {
			http.Redirect(w,r,"/static/view/home.html",http.StatusFound)
		}else {
			w.Write([]byte("Upload Failed"))
		}

		//log
		fmt.Printf("Your file's meta is: %s\n",fileMeta)

		//�ض������ϴ��ɹ�ҳ��
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

//UploadSucHandler: �ϴ��ɹ�
func UploadSucHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Upload finished!")
}

//GetFileMetaHandler: ͨ���ļ�hashֵ����ȡ�ļ�Ԫ��Ϣ
func GetFileMetaHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	//��ȡhashֵ����ͨ�����ѯ�ļ�Ԫ��Ϣ
	filehash := r.Form["filehash"][0]
	fMeta,err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//json��ʽ��metaʵ��
	data,err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//DownloadHandler: �����ļ���ϣֵ�����ļ�
func DownloadHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	//��ȡ�ļ�hashֵ������ȡԪ��Ϣ
	fileSha1 := r.Form.Get("filehash")
	fm,err := meta.GetFileMetaDB(fileSha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//�����ļ�Ԫ��Ϣ�򿪱����ļ�
	f,err := os.Open(fm.Location)
	defer f.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//�������ļ������ڴ�
	data,err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//������Ӧͷ��д������
	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("content-disposition","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

//FileMetaUpdateHandler: �����ļ�Ԫ��Ϣ
func FileMetaUpdateHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	w.WriteHeader(http.StatusOK)
	data,err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//FileDeleteHandler: ɾ���ļ�
func FileDeleteHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	//ɾ�������ļ�
	fMeta := meta.GetFileMeta(fileSha1)
	err := os.Remove(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//ɾ��Ԫ��Ϣ
	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
}