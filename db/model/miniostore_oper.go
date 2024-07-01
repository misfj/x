package model

import (
	"coredx/log"
	"gorm.io/gorm"
)

func GetTotalUseGB(db *gorm.DB, did string) (float64, error) {
	x := make([]Store, 0, 10)
	if db.Model(&Store{}).Where("did = ?", did).Find(&x).Error != nil {
		return 0, db.Model(&Store{}).Where("did = ?", did).Find(&x).Error
	}
	var useBytes int64
	for _, v := range x {
		useBytes += v.Size
	}
	log.Debug(useBytes)
	var useGB float64
	useGB = float64(useBytes / (1024 * 1024 * 1024))
	return useGB, nil
}
