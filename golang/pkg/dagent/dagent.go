package dagent

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/dyrector-io/dyrectorio/golang/internal/grpc"
	"github.com/dyrector-io/dyrectorio/golang/pkg/dagent/config"
	"github.com/dyrector-io/dyrectorio/golang/pkg/dagent/model"
	"github.com/dyrector-io/dyrectorio/golang/pkg/dagent/update"
	"github.com/dyrector-io/dyrectorio/golang/pkg/dagent/utils"
)

func Serve(cfg *config.Configuration) {
	utils.PreflightChecks(cfg)
	log.Print("Starting dyrector.io DAgent service")

	if cfg.TraefikEnabled {
		params := model.TraefikDeployRequest{
			LogLevel: cfg.TraefikLogLevel,
			TLS:      cfg.TraefikTLS,
			AcmeMail: cfg.TraefikAcmeMail,
			Port:     cfg.TraefikPort,
			TLSPort:  cfg.TraefikTLSPort,
		}

		err := utils.ExecTraefik(context.Background(), params, cfg)
		if err != nil {
			// we wanted to start traefik, but something is not ok, thus panic!
			log.Panic().Err(err).Msg("failed to start Traefik")
		}
	}

	update.InitUpdater(cfg)

	grpcParams := grpc.GrpcTokenToConnectionParams(cfg.GrpcToken, cfg.GrpcInsecure)
	grpcContext := grpc.WithGRPCConfig(context.Background(), cfg)
	grpc.Init(grpcContext, grpcParams, &cfg.CommonConfiguration, grpc.WorkerFunctions{
		Deploy:     utils.DeployImage,
		Watch:      utils.GetContainersByNameCrux,
		Delete:     utils.DeleteContainerByName,
		SecretList: utils.SecretList,
	})
}
