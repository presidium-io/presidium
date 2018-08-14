package presidio

import (
	"fmt"
	"os"

	message_types "github.com/Microsoft/presidio-genproto/golang"
	"github.com/Microsoft/presidio/pkg/cache"
	"github.com/Microsoft/presidio/pkg/cache/redis"
	log "github.com/Microsoft/presidio/pkg/logger"
	"github.com/Microsoft/presidio/pkg/rpc"
)

//SetupAnalyzerService GRPC connection
func SetupAnalyzerService() *message_types.AnalyzeServiceClient {
	analyzerSvcAddress := os.Getenv("ANALYZER_SVC_ADDRESS")
	if analyzerSvcAddress == "" {
		log.Fatal("analyzer service address is empty")
	}

	analyzeService, err := rpc.SetupAnalyzerService(analyzerSvcAddress)
	if err != nil {
		log.Fatal("Connection to analyzer service failed %q", err)
	}

	return analyzeService
}

//SetupAnoymizerService GRPC connection
func SetupAnoymizerService() *message_types.AnonymizeServiceClient {

	anonymizerSvcAddress := os.Getenv("ANONYMIZER_SVC_ADDRESS")
	if anonymizerSvcAddress == "" {
		log.Fatal("anonymizer service address is empty")
	}

	anonymizeService, err := rpc.SetupAnonymizeService(anonymizerSvcAddress)
	if err != nil {
		log.Fatal("Connection to anonymizer service failed %q", err)
	}
	return anonymizeService
}

//SetupDatasinkService GRPC connection
func SetupDatasinkService() *message_types.DatasinkServiceClient {
	address := "localhost"
	grpcPort := os.Getenv("DATASINK_GRPC_PORT")
	datasinkService, err := rpc.SetupDatasinkService(fmt.Sprintf("%s:%s", address, grpcPort))
	if err != nil {
		log.Fatal("Connection to datasink service failed %q", err)
	}

	return datasinkService
}

//SetupCache  Redis cache
func SetupCache() cache.Cache {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatal("redis address is empty")
	}

	cache := redis.New(
		redisUrl,
		"", // no password set
		0,  // use default DB
	)
	return cache
}
