package powerdns

import "context"

// https://doc.powerdns.com/authoritative/http-api/server.html
type ServerService service

// Server represents server object
type Server struct {
	ConfigUrl  *string `json:"config_url,omitempty"`
	DaemonType *string `json:"daemon_type,omitempty"`
	ID         *string `json:"id,omitempty"`
	Type       *string `json:"type,omitempty"`
	URL        *string `json:"url,omitempty"`
	Version    *string `json:"version,omitempty"`
	ZonesURL   *string `json:"zones_url,omitempty"`
}

func (s *ServerService) GetServers(ctx context.Context) (*Server, *Response, error) {
	req, err := s.client.NewRequest("GET", "servers", nil)
	if err != nil {
		return nil, nil, err
	}

	srv := &Server{}
	resp, err := s.client.Do(ctx, req, srv)
	if err != nil {
		return nil, resp, err
	}
	return srv, resp, nil
}
