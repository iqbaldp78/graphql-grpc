package graph

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mbizmarket/dmp/graphql/graph/model"
	"github.com/mbizmarket/dmp/graphql/helper"

	pbRfqs "github.com/mbizmarket/dmp/graphql/proto/pb/rfqs"
)

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.Users, error) {
	var output []*model.Users

	res, err := r.ResolverClientRadiance.GetAllUsers(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println("err nya ", err)
		return output, err
	}

	for _, v := range res.Data {
		out := &model.Users{
			ID:                   int(v.ID),
			CompanyID:            int(v.CompanyID),
			Email:                v.Email,
			Password:             v.Password,
			VerifiedEmail:        helper.IntToPointerInt(v.VerifiedEmail),
			Status:               &v.Status,
			Name:                 v.Name,
			Phone:                &v.Phone,
			Mobile:               &v.Mobile,
			JobTitle:             &v.JobTitle,
			ImageStorageID:       helper.IntToPointerInt(v.ImageStorageID),
			IsNotifPaymentReturn: helper.IntToPointerInt(v.IsNotifPaymentReturn),
			Campaign:             helper.IntToPointerInt(v.Campaign),
			IDToken:              &v.IDToken,
			LastLogin:            helper.TimestampProtoToPointerTime(v.LastLogin),
			CreatedAt:            helper.TimestampProtoToPointerTime(v.CreatedAt),
			UpdatedAt:            helper.TimestampProtoToPointerTime(v.UpdatedAt),
		}
		output = append(output, out)
	}
	return output, nil
}

func (r *queryResolver) GetUserNRfqs(ctx context.Context) ([]*model.Users, error) {
	var output []*model.Users
	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)
	defer cancel()

	res, err := r.ResolverClientRadiance.GetAllUsers(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println("err nya radiance ", err)
		return output, err
	}

	for _, v := range res.Data {
		rfqs, err := r.ResolverClientRadiance.GetUserNRfqs(ctx, &pbRfqs.Req{ID: v.CompanyID})

		if err != nil {
			log.Println("err nya GetUserNRfqs", err)
			return output, err
		}
		out := &model.Users{
			ID:                   int(v.ID),
			CompanyID:            int(v.CompanyID),
			Email:                v.Email,
			Password:             v.Password,
			VerifiedEmail:        helper.IntToPointerInt(v.VerifiedEmail),
			Status:               &v.Status,
			Name:                 v.Name,
			Phone:                &v.Phone,
			Mobile:               &v.Mobile,
			JobTitle:             &v.JobTitle,
			ImageStorageID:       helper.IntToPointerInt(v.ImageStorageID),
			IsNotifPaymentReturn: helper.IntToPointerInt(v.IsNotifPaymentReturn),
			Campaign:             helper.IntToPointerInt(v.Campaign),
			IDToken:              &v.IDToken,
			LastLogin:            helper.TimestampProtoToPointerTime(v.LastLogin),
			CreatedAt:            helper.TimestampProtoToPointerTime(v.CreatedAt),
			UpdatedAt:            helper.TimestampProtoToPointerTime(v.UpdatedAt),
		}

		var outputRfqs []*model.Rfq
		for _, val := range rfqs.Data {
			tempRfqs := &model.Rfq{
				ID:              int(val.ID),
				TransactionID:   int(val.TransactionId),
				RequestedBy:     int(val.RequestedBy),
				CreatedBy:       int(val.CreatedBy),
				PaymentType:     int(val.PaymentType),
				PaymentDuration: int(val.PaymentDuration),
				CompanyBuyerID:  int(val.CompanyBuyerId),
				CompanySellerID: int(val.CompanySellerId),
				RfqNo:           val.RfqNo,
				ReferenceNo:     val.ReferenceNo,
				Notes:           val.Notes,
				NoteForSeller:   val.NoteForSeller,
				SubTotal:        float64(val.SubTotal),
				TaxBasis:        float64(val.TaxBasis),
				Ppn:             float64(val.Ppn),
				Pph:             float64(val.Pph),
				Rounding:        float64(val.Rounding),
				GrandTotal:      float64(val.GrandTotal),
				Status:          int(val.Status),
				StatusReason:    val.StatusReason,
				ExpiredAt:       helper.TimestampProtoToPointerTime(val.ExpiredAt),
				CreatedAt:       helper.TimestampProtoToPointerTime(val.CreatedAt),
				UpdatedAt:       helper.TimestampProtoToPointerTime(val.UpdatedAt),
				DeletedAt:       helper.TimestampProtoToPointerTime(val.DeletedAt),
			}
			outputRfqs = append(outputRfqs, tempRfqs)
		}
		out.Rfqs = outputRfqs
		output = append(output, out)
	}
	return output, nil
}
