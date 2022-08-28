package dao

import (
	"fmt"

	"github.com/go-dora/go-common/lib"
	"github.com/luxpo/gateway/public"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http_rule" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp_rule" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc_rule" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
}

// GenServiceAddr 生成 ServiceAddr
func (serviceDetail *ServiceDetail) GenServiceAddr() string {
	// 1. http后缀接入 clusterIP+clusterPort+path
	// 2. http域名接入 domain
	// 3. tcp、grpc接入 clusterIP+servicePort
	serviceAddr := "unknown"
	clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
	clusterPort := lib.GetStringConf("base.cluster.cluster_port")
	clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

	switch true {
	case serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HTTPRule.NeedHttps == 1:
		serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterSSLPort, serviceDetail.HTTPRule.Rule)
	case serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HTTPRule.NeedHttps == 1:
		serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterPort, serviceDetail.HTTPRule.Rule)
	case serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypeDomain:
		serviceAddr = serviceDetail.HTTPRule.Rule
	case serviceDetail.Info.LoadType == public.LoadTypeTCP:
		serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TCPRule.Port)
	case serviceDetail.Info.LoadType == public.LoadTypeGRPC:
		serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
	}

	return serviceAddr
}
