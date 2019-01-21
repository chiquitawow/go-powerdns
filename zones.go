package powerdns

import "context"

// https://doc.powerdns.com/authoritative/http-api/zone.html
type ZoneService service

// Server represents server object
type Zone struct {
	// Optional field for local policy hooks
	Account string `json:"account,omitempty"`
	// Whether or not the zone will be rectified on data changes via the API.
	APIRectify bool `json:"api_rectify,omitempty"`
	// Whether or not this zone is DNSSEC signed
	DNSSec bool `json:"dnssec,omitempty"`
	// Obaque zone ID assigned by the server, should not be interpreted by the
	// application.
	ID string `json:"id,omitempty"`
	// Zone kind ("Native", "Master", or "Slave")
	Kind string `json:"kind,omitempty"`
	// List of IPs configured as a master for this zone.
	Masters []string `json:"masters,omitempty"`
	// Name of the zone (e.g. “example.com.”) MUST have a trailing dot
	Name string `json:"name,omitempty"`
	// MAY be sent in client bodies during creation
	Nameservers []string `json:"nameservers,omitempty"`
	// The SOA serial notifications have been sent out for.
	NotifiedSerial string `json:"notified_serial,omitempty"`
	// Whether or not the zone uses NSEC3 narrow
	NSEC3Narrow bool `json:"nsec3narrow,omitempty"`
	// The NSEC3PARAM record
	NSEC3Param string `json:"nsec3param,omitempty"`
	// Whether or not the zone is pre-signed.
	Presigned bool `json:"presigned,omitempty"`
	//RRSets in this zone
	RRSets []RRSet
	// The SOA serial number
	Serial string `json:"serial,omitempty"`
	// The SOA-EDIT metadata item
	SOAEdit string `json:"soa_edit,omitempty"`
	// The SOA-EDIT-API metadate item
	SOAEditAPI string `json:"soa_edit_api,omitempty"`
	// The id of the TSIG keys used for master operation in this zone.
	TSIGMasterKeyIDs []string `json:"master_tsig_key_ids"`
	// The id of the TSIG keys used for slave operation in this zone
	TSIGSlaveKeyIDs []string `json:"slave_tsig_key_ids,omitempty"`
	// API endpoint for this zone
	URL string `json:"url,omitempty"`
	// MAY contain a BIND-style zone file when creating a zone
	Zone string `json:"zone,omitempty"`
}

// RRSet defines a Resource Record Set (all records with the same name and
// type)
type RRSet struct {
	// ChangeType MUST be added when updating the RRSet. Must be REPLACE or
	// DELETE.
	ChangeType string `json:"change_type,omitempty"`
	// List of Comment
	Comments []Comment
	// Name for the record set (e.g." www.powerdns.com.")
	Name string `json:"name,omitempty"`
	// All records in the RRSet.
	Records []Record
	// DNS TTL of the records, in seconds.
	TTL int `json:"ttl,omitempty"`
	// Type of record ("A", "PTR", "MX", etc)
	RRType string `json:"type,omitempty"`
}

// Record defines the RREntry object.
type Record struct {
	// The content of this record (IP).
	Content string `json:"content,omitempty"`
	// Whether or not this record is disabled.
	Disabled bool `json:"disabled,omitempty"`
	// If set to true, the server will find the matching reverse zone and create
	// a PTR there.
	SetPTR bool `json:"set_ptr,omitempty"`
}

// Comment defines a comment about a RRSet.
type Comment struct {
	// Name of an account that added the comment.
	Account string `json:"account,omitempty"`
	// The actual comment.
	Content string `json:"content,omitempty"`
	// Timestamp of the last change to the comment.
	ModifiedAt string `json:"modified_at,omitempty"`
}

// List returns all Zones in a server
func (s *ZoneService) List(ctx context.Context) ([]Zone, *Response, error) {
	req, err := s.client.NewRequest("GET", "servers/localhost/zones", nil)
	if err != nil {
		return nil, nil, err
	}

	var zz []Zone
	resp, err := s.client.Do(ctx, req, &zz)
	if err != nil {
		return nil, resp, err
	}
	return zz, resp, nil
}
