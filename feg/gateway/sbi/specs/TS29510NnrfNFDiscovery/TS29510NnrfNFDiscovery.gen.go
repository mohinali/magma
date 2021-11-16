// Package TS29510NnrfNFDiscovery provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29510NnrfNFDiscovery

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29503NudmSDM"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29510NnrfNFManagement"
	externalRef2 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// NFProfile defines model for NFProfile.
type NFProfile struct {
	AmfInfo                          *externalRef1.AmfInfo                           `json:"amfInfo,omitempty"`
	AusfInfo                         *externalRef1.AusfInfo                          `json:"ausfInfo,omitempty"`
	BsfInfo                          *externalRef1.BsfInfo                           `json:"bsfInfo,omitempty"`
	Capacity                         *int                                            `json:"capacity,omitempty"`
	ChfInfo                          *externalRef1.ChfInfo                           `json:"chfInfo,omitempty"`
	CustomInfo                       *map[string]interface{}                         `json:"customInfo,omitempty"`
	DefaultNotificationSubscriptions *[]externalRef1.DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`
	Fqdn                             *externalRef1.Fqdn                              `json:"fqdn,omitempty"`
	Ipv4Addresses                    *[]externalRef2.Ipv4Addr                        `json:"ipv4Addresses,omitempty"`
	Ipv6Addresses                    *[]externalRef2.Ipv6Addr                        `json:"ipv6Addresses,omitempty"`
	Load                             *int                                            `json:"load,omitempty"`
	Locality                         *string                                         `json:"locality,omitempty"`
	NfInstanceId                     externalRef2.NfInstanceId                       `json:"nfInstanceId"`
	NfServicePersistence             *bool                                           `json:"nfServicePersistence,omitempty"`
	NfServices                       *[]NFService                                    `json:"nfServices,omitempty"`
	NfStatus                         externalRef1.NFStatus                           `json:"nfStatus"`
	NfType                           externalRef1.NFType                             `json:"nfType"`
	NsiList                          *[]string                                       `json:"nsiList,omitempty"`
	PcfInfo                          *externalRef1.PcfInfo                           `json:"pcfInfo,omitempty"`
	PerPlmnSnssaiList                *[]externalRef1.PlmnSnssai                      `json:"perPlmnSnssaiList,omitempty"`
	PlmnList                         *[]externalRef2.PlmnId                          `json:"plmnList,omitempty"`
	Priority                         *int                                            `json:"priority,omitempty"`
	RecoveryTime                     *externalRef2.DateTime                          `json:"recoveryTime,omitempty"`
	SNssais                          *[]externalRef2.Snssai                          `json:"sNssais,omitempty"`
	SmfInfo                          *externalRef1.SmfInfo                           `json:"smfInfo,omitempty"`
	UdmInfo                          *externalRef1.UdmInfo                           `json:"udmInfo,omitempty"`
	UdrInfo                          *externalRef1.UdrInfo                           `json:"udrInfo,omitempty"`
	UpfInfo                          *externalRef1.UpfInfo                           `json:"upfInfo,omitempty"`
}

// NFService defines model for NFService.
type NFService struct {
	ApiPrefix                        *string                                         `json:"apiPrefix,omitempty"`
	Capacity                         *int                                            `json:"capacity,omitempty"`
	ChfServiceInfo                   *externalRef1.ChfServiceInfo                    `json:"chfServiceInfo,omitempty"`
	DefaultNotificationSubscriptions *[]externalRef1.DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`
	Fqdn                             *externalRef1.Fqdn                              `json:"fqdn,omitempty"`
	IpEndPoints                      *[]externalRef1.IpEndPoint                      `json:"ipEndPoints,omitempty"`
	Load                             *int                                            `json:"load,omitempty"`
	NfServiceStatus                  externalRef1.NFServiceStatus                    `json:"nfServiceStatus"`
	Priority                         *int                                            `json:"priority,omitempty"`
	RecoveryTime                     *externalRef2.DateTime                          `json:"recoveryTime,omitempty"`
	Scheme                           externalRef2.UriScheme                          `json:"scheme"`
	ServiceInstanceId                string                                          `json:"serviceInstanceId"`
	ServiceName                      externalRef1.ServiceName                        `json:"serviceName"`
	SupportedFeatures                *externalRef2.SupportedFeatures                 `json:"supportedFeatures,omitempty"`
	Versions                         []externalRef1.NFServiceVersion                 `json:"versions"`
}

// SearchResult defines model for SearchResult.
type SearchResult struct {
	NfInstances          []NFProfile                     `json:"nfInstances"`
	NrfSupportedFeatures *externalRef2.SupportedFeatures `json:"nrfSupportedFeatures,omitempty"`
	ValidityPeriod       *int                            `json:"validityPeriod,omitempty"`
}

// SearchNFInstancesParams defines parameters for SearchNFInstances.
type SearchNFInstancesParams struct {
	// Type of the target NF
	TargetNfType externalRef1.NFType `json:"target-nf-type"`

	// Type of the requester NF
	RequesterNfType externalRef1.NFType `json:"requester-nf-type"`

	// Names of the services offered by the NF
	ServiceNames *[]externalRef1.ServiceName `json:"service-names,omitempty"`

	// FQDN of the requester NF
	RequesterNfInstanceFqdn *externalRef1.Fqdn `json:"requester-nf-instance-fqdn,omitempty"`

	// Id of the PLMN of the target NF
	TargetPlmnList *[]externalRef2.PlmnId `json:"target-plmn-list,omitempty"`

	// Id of the PLMN where the NF issuing the Discovery request is located
	RequesterPlmnList *[]externalRef2.PlmnId `json:"requester-plmn-list,omitempty"`

	// Identity of the NF instance being discovered
	TargetNfInstanceId *externalRef2.NfInstanceId `json:"target-nf-instance-id,omitempty"`

	// FQDN of the NF instance being discovered
	TargetNfFqdn *externalRef1.Fqdn `json:"target-nf-fqdn,omitempty"`

	// Uri of the home NRF
	HnrfUri *externalRef2.Uri `json:"hnrf-uri,omitempty"`

	// Slice info of the target NF
	Snssais *[]externalRef2.Snssai `json:"snssais,omitempty"`

	// Slice info of the requester NF
	RequesterSnssais *[]externalRef2.Snssai `json:"requester-snssais,omitempty"`

	// PLMN specific Slice info of the target NF
	PlmnSpecificSnssaiList *[]externalRef1.PlmnSnssai `json:"plmn-specific-snssai-list,omitempty"`

	// Dnn supported by the BSF, SMF or UPF
	Dnn *externalRef2.Dnn `json:"dnn,omitempty"`

	// NSI IDs that are served by the services being discovered
	NsiList        *[]string `json:"nsi-list,omitempty"`
	SmfServingArea *string   `json:"smf-serving-area,omitempty"`

	// Tracking Area Identity
	Tai *externalRef2.Tai `json:"tai,omitempty"`

	// AMF Region Identity
	AmfRegionId *externalRef2.AmfRegionId `json:"amf-region-id,omitempty"`

	// AMF Set Identity
	AmfSetId *externalRef2.AmfSetId `json:"amf-set-id,omitempty"`

	// Guami used to search for an appropriate AMF
	Guami *externalRef2.Guami `json:"guami,omitempty"`

	// SUPI of the user
	Supi *externalRef2.Supi `json:"supi,omitempty"`

	// IPv4 address of the UE
	UeIpv4Address *externalRef2.Ipv4Addr `json:"ue-ipv4-address,omitempty"`

	// IP domain of the UE, which supported by BSF
	IpDomain *string `json:"ip-domain,omitempty"`

	// IPv6 prefix of the UE
	UeIpv6Prefix *externalRef2.Ipv6Prefix `json:"ue-ipv6-prefix,omitempty"`

	// Combined PGW-C and SMF or a standalone SMF
	PgwInd *bool `json:"pgw-ind,omitempty"`

	// PGW FQDN of a combined PGW-C and SMF
	Pgw *externalRef1.Fqdn `json:"pgw,omitempty"`

	// GPSI of the user
	Gpsi *externalRef2.Gpsi `json:"gpsi,omitempty"`

	// external group identifier of the user
	ExternalGroupIdentity *externalRef0.ExtGroupId `json:"external-group-identity,omitempty"`

	// data set supported by the NF
	DataSet *externalRef1.DataSetId `json:"data-set,omitempty"`

	// routing indicator in SUCI
	RoutingIndicator *string `json:"routing-indicator,omitempty"`

	// Group IDs of the NFs being discovered
	GroupIdList *[]externalRef2.NfGroupId `json:"group-id-list,omitempty"`

	// Data network access identifiers of the NFs being discovered
	DnaiList *[]externalRef2.Dnai `json:"dnai-list,omitempty"`

	// list of PDU Session Type required to be supported by the target NF
	PduSessionTypes *[]externalRef2.PduSessionType `json:"pdu-session-types,omitempty"`

	// Features required to be supported by the target NF
	SupportedFeatures *externalRef2.SupportedFeatures `json:"supported-features,omitempty"`

	// UPF supporting interworking with EPS or not
	UpfIwkEpsInd *bool `json:"upf-iwk-eps-ind,omitempty"`

	// PLMN ID supported by a CHF
	ChfSupportedPlmn *externalRef2.PlmnId `json:"chf-supported-plmn,omitempty"`

	// preferred target NF location
	PreferredLocality *string `json:"preferred-locality,omitempty"`

	// AccessType supported by the target NF
	AccessType *externalRef2.AccessType `json:"access-type,omitempty"`

	// Maximum number of NFProfiles to return in the response
	Limit *int `json:"limit,omitempty"`

	// Features required to be supported by the target NF
	RequiredFeatures *[]externalRef2.SupportedFeatures `json:"required-features,omitempty"`

	// the complex query condition expression
	ComplexQuery *externalRef2.ComplexQuery `json:"complex-query,omitempty"`

	// Maximum payload size of the response expressed in kilo octets
	MaxPayloadSize *int `json:"max-payload-size,omitempty"`

	// Validator for conditional requests, as described in IETF RFC 7232, 3.2
	IfNoneMatch *string `json:"If-None-Match,omitempty"`
}
