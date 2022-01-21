package global

import (
	"gin-vue-admin/config"
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"k8s.io/client-go/rest"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG           *zap.Logger
	GVA_K8SConfig     *rest.Config
	GVA_K8SConfigMap  map[string]*rest.Config
	GVA_K8SRestConfig *rest.Config
	GVA_K8SCurrent    string //当前集群
)

func GetK8sConfig(userName string) *rest.Config {
	value, ok := GVA_K8SConfigMap[userName]
	var gva_k8sconfig *rest.Config
	if ok {
		gva_k8sconfig = value
	} else {
		gva_k8sconfig = GVA_K8SConfigMap["defaultConfig"]
	}
	return gva_k8sconfig
}
