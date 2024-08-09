package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameDataEnc = "data_enc"

// DataEnc mapped from table <data_enc>
type DataEnc struct {
	ID                uint           `gorm:"primary"`
	CreatedAt         time.Time      `gorm:"column:create_at"`
	UpdatedAt         time.Time      `gorm:"column:update_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:delete_at"`                                                        // 删除时间
	FileName          string         `gorm:"column:file_name;comment:文件名称" json:"file_name"`                          // 文件名称
	FileSize          int64          `gorm:"column:file_size;comment:文件大小" json:"file_size"`                          // 文件大小
	FileMd5           string         `gorm:"column:file_md5;comment:文件hash" json:"file_md5"`                          // 文件hash
	FileKey           string         `gorm:"column:file_key;comment:对称密钥" json:"file_key"`                            // 对称密钥
	FileAlgo          string         `gorm:"column:file_algo;comment:加密算法" json:"file_algo"`                          // 加密算法
	FileEncryptedHash string         `gorm:"column:file_encrypted_hash;comment:加密之后的hash" json:"file_encrypted_hash"` // 加密之后的hash
	FileURL           string         `gorm:"column:file_url;comment:可以进行文件下载的URL" json:"file_url"`                    // 可以进行文件下载的URL
	NodeID            string         `gorm:"column:node_id;comment:节点ID" json:"node_id"`                              // 节点ID
}

// TableName DataEnc's table name
func (*DataEnc) TableName() string {
	return TableNameDataEnc
}

var DataEncQuery *DataEncDao

type DataEncDao struct {
	db *gorm.DB
}

func NewDataEncDao(db *gorm.DB) {
	DataEncQuery = &DataEncDao{db: db}
}
func (de *DataEncDao) Create(dta *DataEnc) error {
	return de.db.Model(&DataEnc{}).Create(dta).Error
}

func (de *DataEncDao) Exist(encHash string) (*DataEnc, error) {
	var dtae DataEnc
	if err := de.db.Model(&DataEnc{}).Where("file_encrypted_hash = ?", encHash).First(&dtae).Error; err != nil {
		return nil, err
	}
	return &dtae, nil
}

func (de *DataEncDao) ExistRawByDataAid(aid string, fileKey string) (*DataEnc, error) {
	var dtae DataEnc
	if err := de.db.Model(&DataEnc{}).Where("file_md5  =?", aid).First(&dtae).Error; err != nil {
		return nil, err
	}
	return &dtae, nil
}
