package repository

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/mbizmarket/dmp/aghanim/models"
	"github.com/mbizmarket/dmp/aghanim/proto/rfqs/pb"
)

//Service --
type Service struct {
	DB *gorm.DB
}

//GetRfqsByUser --
func (s *Service) GetRfqsByUser(ctx context.Context, in *pb.Req, out *pb.RespData) error {
	log.Println("GetRfqsByUser received .....")
	rfqs := []models.Rfqs{}
	pbRfqs := []*pb.Rfqs{}

	// s.DB.Debug().Limit(30).Where("company_seller_id = ?", in.ID).Find(&rfqs)
	s.DB.Debug().Where("company_seller_id = ?", in.ID).Find(&rfqs)
	copier.Copy(&pbRfqs, &rfqs)

	for i := range rfqs {
		if rfqs[i].DeletedAt == nil {
			pbRfqs[i].DeletedAt = nil
		} else {
			pbRfqs[i].DeletedAt, _ = ptypes.TimestampProto(*rfqs[i].DeletedAt)
		}

		pbRfqs[i].CreatedAt, _ = ptypes.TimestampProto(rfqs[i].CreatedAt)
		pbRfqs[i].UpdatedAt, _ = ptypes.TimestampProto(rfqs[i].UpdatedAt)
		pbRfqs[i].ExpiredAt, _ = ptypes.TimestampProto(rfqs[i].ExpiredAt)

	}
	out.Data = pbRfqs
	return nil
}
