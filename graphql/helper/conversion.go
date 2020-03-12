package helper

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

//IntToPointerInt --
func IntToPointerInt(v interface{}) *int {
	switch v.(type) {
	case int64:
		val := int(v.(int64))
		return &val
	case int32:
		val := int(v.(int32))
		return &val
	default:
		log.Println("Type undifined ..")
		log.Fatal("Type undifined ..")
	}
	return nil
}

//TimestampProtoToPointerTime --
func TimestampProtoToPointerTime(val *timestamp.Timestamp) *time.Time {
	v, _ := ptypes.Timestamp(val)
	return &v
}
