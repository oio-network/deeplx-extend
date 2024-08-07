package task

import (
	"net"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/oio-network/deeplx-extend/api/deeplx/v1"
	"github.com/oio-network/deeplx-extend/app/deeplx/pkgs/msg"
)

const LogTaskCreateAccessLog = "logTask.CreateAccessLog"

type AccessLogParams struct {
	RemoteAddr string
	CreatedAt  *timestamppb.Timestamp
}

func (t *logTask) RegisterLogTask(srv MachineryServer) error {
	if err := srv.HandleFunc(LogTaskCreateAccessLog, t.CreateAccessLog); err != nil {
		return err
	}

	return nil
}

func (t *logTask) CreateAccessLog(b []byte) error {
	params := &AccessLogParams{}
	if err := msg.Unmarshal(b, params); err != nil {
		t.log.Error(err)
		return err
	}

	log := &v1.AccessLog{CreatedAt: params.CreatedAt}
	log.Ip, _, _ = net.SplitHostPort(params.RemoteAddr)

	record, err := t.mmdb.Country(net.ParseIP(log.Ip))
	if err == nil {
		log.CountryName = record.Country.Names["en"]
		log.CountryCode = record.Country.IsoCode
	}

	if err = t.sink.Append(log); err != nil {
		t.log.Error(err)
		return err
	}

	return nil
}
