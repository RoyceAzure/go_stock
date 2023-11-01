package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grcpGateWayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "usesr-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	mtda := &MetaData{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %+v\n", md)
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtda.UserAgent = userAgents[0]
		}
		if userAgents := md.Get(grcpGateWayUserAgentHeader); len(userAgents) > 0 {
			mtda.UserAgent = userAgents[0]
		}
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtda.UserAgent = clientIPs[0]
		}

		if p, ok := peer.FromContext(ctx); ok {
			mtda.ClientIP = p.Addr.String()
		}
	}
	return mtda
}
