package clients

import (
	"mymachine707/config"
	"mymachine707/protogen/eCommerce"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Category eCommerce.CategoryServiceClient
	Product  eCommerce.ProductServiceClient

	conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCategory, err := grpc.Dial(cfg.CategoryServiceGrpcHost+cfg.CategoryServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	category := eCommerce.NewCategoryServiceClient(connCategory)

	connProduct, err := grpc.Dial(cfg.ProductServiceGrpcHost+cfg.ProductServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	product := eCommerce.NewProductServiceClient(connProduct)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Category: category,
		Product:  product,
		conns:    append(conns, connCategory, connProduct),
	}, nil
}

// Close for disconnection connection another serevice
func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
