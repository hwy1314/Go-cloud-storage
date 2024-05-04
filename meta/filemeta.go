package meta

import (
	mydb "FileStore-Server/db"
)

// FileMeta: �ļ�Ԫ��Ϣ�ṹ
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// fileMetas: ���ļ���ϣֵΪkey���ļ�Ԫ��Ϣ�ṹ��Ϊvalue
var fileMetas map[string]FileMeta

func init()  {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: ����/�����ļ�Ԫ��Ϣ
func UpdateFileMeta(fMeta FileMeta)  {
	fileMetas[fMeta.FileSha1] = fMeta
}

// UpdateFileMetaDB: ����/����Ԫ��Ϣ��mysql��
func UpdateFileMetaDB(fmeta FileMeta) bool{
	return mydb.OnFileUploadFinished(fmeta.FileSha1,fmeta.FileName,fmeta.FileSize,fmeta.Location)
}

// UpdateFileMetaDB: ��mysql�л�ȡ�ļ�Ԫ��Ϣ
func GetFileMetaDB(fileSha1 string) (FileMeta,error){
	tfile,err := mydb.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{},err
	}

	fmeta := FileMeta{
		FileSha1:tfile.FileHash,
		FileName:tfile.FileName.String,
		FileSize:tfile.FileSize.Int64,
		Location:tfile.FileAddr.String,
	}

	return fmeta,nil
}

// GetFileMeta: ͨ��sha1��ȡ�ļ�Ԫ��Ϣ����
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// RemoveFileMeta: ɾ��Ԫ��Ϣ
func RemoveFileMeta(fileSha1 string)  {
	delete(fileMetas,fileSha1)
}