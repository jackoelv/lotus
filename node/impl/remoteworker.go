package impl

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-jsonrpc/auth"
	sectorstorage "github.com/filecoin-project/sector-storage"
	"github.com/filecoin-project/specs-actors/actors/abi"
	storage2 "github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
)

type remoteWorker struct {
	api.WorkerAPI
	closer jsonrpc.ClientCloser
}

func (r *remoteWorker) NewSector(ctx context.Context, sector abi.SectorID) error {
	log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go NewSector")
	return xerrors.New("unsupported")
}

func (r *remoteWorker) AddPiece(ctx context.Context, sector abi.SectorID, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData storage2.Data) (abi.PieceInfo, error) {
	log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go AddPiece")
	nodeApi := r.WorkerAPI
	abipiece, err := nodeApi.RemoteAddPiece(ctx, sector, pieceSizes, newPieceSize)
	if err != nil {
		log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go nodeApi.RemoteAddPiece:%s", err)
		return abipiece, err
	}

	return abipiece, nil
}
func (r *remoteWorker) RemoteAddPiece(ctx context.Context, sector abi.SectorID, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize) (abi.PieceInfo, error) {
	log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go RemoteAddPiece")
	nodeApi := r.WorkerAPI
	abipiece, err := nodeApi.RemoteAddPiece(ctx, sector, pieceSizes, newPieceSize)
	if err != nil {
		log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go RemoteAddPiece nodeApi.RemoteAddPiece:%s", err)
		return abipiece, err
	}

	return abipiece, nil
}

func connectRemoteWorker(ctx context.Context, fa api.Common, url string) (*remoteWorker, error) {
	token, err := fa.AuthNew(ctx, []auth.Permission{"admin"})
	if err != nil {
		return nil, xerrors.Errorf("creating auth token for remote connection: %w", err)
	}

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+string(token))

	wapi, closer, err := client.NewWorkerRPC(url, headers)
	if err != nil {
		return nil, xerrors.Errorf("creating jsonrpc client: %w", err)
	}
	// log.Warnf("jackoelvAddpiecetest:lotus/node/imple/remoteworker.go wapi:%s", wapi)
	return &remoteWorker{wapi, closer}, nil
}

func (r *remoteWorker) Close() error {
	r.closer()
	return nil
}

var _ sectorstorage.Worker = &remoteWorker{}
