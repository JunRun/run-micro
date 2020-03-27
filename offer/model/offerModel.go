/**
 *
 * @Description: offer 实体类
 * @Version: 1.0.0
 * @Date: 2020-03-27 10:16
 */
package model

import (
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/logger"
	"time"
)

type Offer struct {
	Id          string
	Name        string
	Url         string
	PostBackUrl string
	AffName     string
	Date        time.Time `gorm:"DEFAULT:now()"`
}

func (o *Offer) GetOfferById(id string) *Offer {
	off := &Offer{}
	gDB.Where("id = ?", id).Find(off)
	return off
}

func (o *Offer) SetOffer(offer *Offer) error {
	logger.Info(offer.Id)
	if offer.Id == "" {
		offer.Id = uuid.New().String()
		return gDB.Save(offer).Error
	} else {
		gDB.Debug()
		return gDB.Update(offer).Error
	}

}
