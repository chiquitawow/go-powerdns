package powerdns

import (
	"context"
)

// https://doc.powerdns.com/authoritative/http-api/server.html
type ServerService service

// Server represents server object
type Server struct {
	ConfigUrl  string `json:"config_url,omitempty"`
	DaemonType string `json:"daemon_type,omitempty"`
	ID         string `json:"id,omitempty"`
	ServerType string `json:"type,omitempty"`
	URL        string `json:"url,omitempty"`
	Version    string `json:"version,omitempty"`
	ZonesURL   string `json:"zones_url,omitempty"`
}

// GetServers
func (s *ServerService) Get(ctx context.Context) ([]Server, *Response, error) {
	req, err := s.client.NewRequest("GET", "servers", nil)
	if err != nil {
		return nil, nil, err
	}

	var srvs []Server
	resp, err := s.client.Do(ctx, req, &srvs)
	if err != nil {
		return nil, resp, err
	}
	return srvs, resp, nil
}
