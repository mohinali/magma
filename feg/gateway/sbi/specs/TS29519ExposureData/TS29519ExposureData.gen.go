// Package TS29519ExposureData provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29519ExposureData

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29518NamfCommunication"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29518NamfEventExposure"
	externalRef2 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

// AccessAndMobilityData defines model for AccessAndMobilityData.
type AccessAndMobilityData struct {
	AccessType           *externalRef2.AccessType     `json:"accessType,omitempty"`
	ConnStates           *[]externalRef1.CmInfo       `json:"connStates,omitempty"`
	ConnStatesTs         *externalRef2.DateTime       `json:"connStatesTs,omitempty"`
	CurrentPlmn          *externalRef2.PlmnId         `json:"currentPlmn,omitempty"`
	CurrentPlmnTs        *externalRef2.DateTime       `json:"currentPlmnTs,omitempty"`
	Location             *externalRef2.UserLocation   `json:"location,omitempty"`
	LocationTs           *externalRef2.DateTime       `json:"locationTs,omitempty"`
	RatType              *[]externalRef2.RatType      `json:"ratType,omitempty"`
	RatTypesTs           *externalRef2.DateTime       `json:"ratTypesTs,omitempty"`
	ReachabilityStatus   *externalRef1.UeReachability `json:"reachabilityStatus,omitempty"`
	ReachabilityStatusTs *externalRef2.DateTime       `json:"reachabilityStatusTs,omitempty"`
	RegStates            *[]externalRef1.RmInfo       `json:"regStates,omitempty"`
	RegStatesTs          *externalRef2.DateTime       `json:"regStatesTs,omitempty"`

	// True  The serving PLMN of the UE is different from the HPLMN of the UE; False  The serving PLMN of the UE is the HPLMN of the UE.
	RoamingStatus      *bool                    `json:"roamingStatus,omitempty"`
	RoamingStatusTs    *externalRef2.DateTime   `json:"roamingStatusTs,omitempty"`
	SmsOverNasStatus   *externalRef0.SmsSupport `json:"smsOverNasStatus,omitempty"`
	SmsOverNasStatusTs *externalRef2.DateTime   `json:"smsOverNasStatusTs,omitempty"`
	TimeZone           *externalRef2.TimeZone   `json:"timeZone,omitempty"`
	TimeZoneTs         *externalRef2.DateTime   `json:"timeZoneTs,omitempty"`
}

// ExposureDataChangeNotification defines model for ExposureDataChangeNotification.
type ExposureDataChangeNotification struct {
	AccessAndMobilityData    *AccessAndMobilityData      `json:"accessAndMobilityData,omitempty"`
	DelResources             *[]externalRef2.Uri         `json:"delResources,omitempty"`
	PduSessionManagementData *[]PduSessionManagementData `json:"pduSessionManagementData,omitempty"`
	UeId                     *externalRef2.VarUeId       `json:"ueId,omitempty"`
}

// ExposureDataSubscription defines model for ExposureDataSubscription.
type ExposureDataSubscription struct {
	Expiry                *externalRef2.DateTime          `json:"expiry,omitempty"`
	MonitoredResourceUris []externalRef2.Uri              `json:"monitoredResourceUris"`
	NotificationUri       externalRef2.Uri                `json:"notificationUri"`
	SupportedFeatures     *externalRef2.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// PduSessionManagementData defines model for PduSessionManagementData.
type PduSessionManagementData struct {
	Dnai     *externalRef2.Dnai     `json:"dnai,omitempty"`
	DnaiTs   *externalRef2.DateTime `json:"dnaiTs,omitempty"`
	Dnn      *externalRef2.Dnn      `json:"dnn,omitempty"`
	IpAddrTs *externalRef2.DateTime `json:"ipAddrTs,omitempty"`
	Ipv4Addr *externalRef2.Ipv4Addr `json:"ipv4Addr,omitempty"`

	// UE IPv6 prefix.
	Ipv6Prefix             *[]externalRef2.Ipv6Prefix      `json:"ipv6Prefix,omitempty"`
	N6TrafficRoutingInfo   *[]externalRef2.RouteToLocation `json:"n6TrafficRoutingInfo,omitempty"`
	N6TrafficRoutingInfoTs *externalRef2.DateTime          `json:"n6TrafficRoutingInfoTs,omitempty"`
	PduSessionId           *externalRef2.PduSessionId      `json:"pduSessionId,omitempty"`

	// Possible values are - "ACTIVE" - "RELEASED"
	PduSessionStatus   *PduSessionStatus      `json:"pduSessionStatus,omitempty"`
	PduSessionStatusTs *externalRef2.DateTime `json:"pduSessionStatusTs,omitempty"`
}

// Possible values are - "ACTIVE" - "RELEASED"
type PduSessionStatus interface{}

// CreateIndividualExposureDataSubscriptionJSONBody defines parameters for CreateIndividualExposureDataSubscription.
type CreateIndividualExposureDataSubscriptionJSONBody ExposureDataSubscription

// ReplaceIndividualExposureDataSubscriptionJSONBody defines parameters for ReplaceIndividualExposureDataSubscription.
type ReplaceIndividualExposureDataSubscriptionJSONBody ExposureDataSubscription

// QueryAccessAndMobilityDataParams defines parameters for QueryAccessAndMobilityData.
type QueryAccessAndMobilityDataParams struct {
	// Supported Features
	SuppFeat *externalRef2.SupportedFeatures `json:"supp-feat,omitempty"`
}

// CreateOrReplaceAccessAndMobilityDataJSONBody defines parameters for CreateOrReplaceAccessAndMobilityData.
type CreateOrReplaceAccessAndMobilityDataJSONBody AccessAndMobilityData

// QuerySessionManagementDataParams defines parameters for QuerySessionManagementData.
type QuerySessionManagementDataParams struct {
	// IPv4 Address of the UE
	Ipv4Addr *externalRef2.Ipv4Addr `json:"ipv4-addr,omitempty"`

	// IPv6 Address Prefix of the UE
	Ipv6Prefix *externalRef2.Ipv6Prefix `json:"ipv6-prefix,omitempty"`

	// DNN of the UE
	Dnn *externalRef2.Dnn `json:"dnn,omitempty"`

	// attributes to be retrieved
	Fields *[]string `json:"fields,omitempty"`

	// Supported Features
	SuppFeat *externalRef2.SupportedFeatures `json:"supp-feat,omitempty"`
}

// CreateOrReplaceSessionManagementDataJSONBody defines parameters for CreateOrReplaceSessionManagementData.
type CreateOrReplaceSessionManagementDataJSONBody PduSessionManagementData

// CreateIndividualExposureDataSubscriptionJSONRequestBody defines body for CreateIndividualExposureDataSubscription for application/json ContentType.
type CreateIndividualExposureDataSubscriptionJSONRequestBody CreateIndividualExposureDataSubscriptionJSONBody

// ReplaceIndividualExposureDataSubscriptionJSONRequestBody defines body for ReplaceIndividualExposureDataSubscription for application/json ContentType.
type ReplaceIndividualExposureDataSubscriptionJSONRequestBody ReplaceIndividualExposureDataSubscriptionJSONBody

// CreateOrReplaceAccessAndMobilityDataJSONRequestBody defines body for CreateOrReplaceAccessAndMobilityData for application/json ContentType.
type CreateOrReplaceAccessAndMobilityDataJSONRequestBody CreateOrReplaceAccessAndMobilityDataJSONBody

// CreateOrReplaceSessionManagementDataJSONRequestBody defines body for CreateOrReplaceSessionManagementData for application/json ContentType.
type CreateOrReplaceSessionManagementDataJSONRequestBody CreateOrReplaceSessionManagementDataJSONBody
