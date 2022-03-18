package helper

import (
	"reflect"
	"strconv"
	"time"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.uber.org/zap"
)

type JsonTime string
type StrFloat float64

var CustomRegistry *bsoncodec.Registry

// CreateCustomRegistry
// Custom BSON converter. Converts time.Time to UTC + 3 string
// https://web.archive.org/web/20220212160824/https://linggar.asia/?p=651
func CreateCustomRegistry() *bsoncodec.RegistryBuilder {
	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)

	customTimeType := reflect.TypeOf(JsonTime(""))
	customStrFloat := reflect.TypeOf(StrFloat(0))
	rb.RegisterTypeDecoder(
		customTimeType,
		bsoncodec.ValueDecoderFunc(func(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
			read, err := vr.ReadDateTime()
			if err != nil {
				return err
			}

			// Convert to UTC+3 and remove milliseconds
			convertedTimeStr := time.UnixMilli(read).In(configs.Server.Config.Location).Truncate(1000 * time.Millisecond).Format(time.RFC3339)
			val.SetString(convertedTimeStr)
			return nil
		}),
	)
	rb.RegisterTypeDecoder(
		customStrFloat,
		bsoncodec.ValueDecoderFunc(func(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
			read, err := vr.ReadString()
			if err != nil {
				return err
			}

			result, err := strconv.ParseFloat(read, 64)
			if err != nil {
				zap.L().Error("strconv.ParseFloat", zap.Error(err))
				val.SetFloat(0)
			}
			val.SetFloat(result)
			return nil
		}),
	)
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}
