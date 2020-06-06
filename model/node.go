package model

import "context"

type Node struct {
	ID        string                       `json:"id" binding:"required"`
	HostName  string                       `json:"hostname" binding:"required"`
	Endpoints map[string]*ExporterEndpoint `json:"endpoints" binding:"required"`
}

func (n Node) AddEndpoint(e ExporterEndpoint) {
	n.Endpoints[e.ID] = &e
	go e.Run(context.Background())
}
