/* SPDX-License-Identifier: Apache-2.0
* Copyright (c) 2019-2020 Intel Corporation
 */

package ngcnef

// SubscribedEvent :Identifies a UP path management event the AF requested to
// be notified of
type SubscribedEvent string

/*
// List of SubscribedEvent
const (
	//UP_PATH_CHANGE SubscribedEvent = "UP_PATH_CHANGE" >> causing lint error
	UpPathChange SubscribedEvent = "UP_PATH_CHANGE"
)
*/

// TrafficInfluSub is Traffic Influence Subscription structure
type TrafficInfluSub struct {
	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceID string `json:"afServiceId,omitempty"`
	// Identifies an application.
	AfAppID string `json:"afAppId,omitempty"`
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransID string `json:"afTransId,omitempty"`
	// Identifies data network name
	Dnn Dnn `json:"dnn,omitempty"`
	// Network slice identifier
	Snssai Snssai `json:"snssai,omitempty"` //p
	// string containing a local identifier followed by \"@\" and
	// a domain identifier.
	// Both the local identifier and the domain identifier shall be encoded as
	// strings that do not contain any \"@\" characters.
	// See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupID ExternalGroupID `json:"externalGroupId,omitempty"`
	// Identifies the requirement to be notified of the event(s).
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty"`
	//Generic Public Servie Identifiers asssociated wit the UE
	Gpsi Gpsi `json:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\"
	//notation as defined in IETF RFC 1166.
	Ipv4Addr Ipv4Addr `json:"ipv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4
	// in IETF RFC 5952.
	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`
	// string identifying mac address of UE
	MacAddr MacAddr48 `json:"macAddr,omitempty"`
	// Identifies the type of notification regarding UP path management event.
	// Possible values are:
	// EARLY - early notification of UP path reconfiguration.
	// EARLY_LATE - early and late notification of UP path reconfiguration.
	// This value shall only be present in the subscription to the
	// DNAI change event.
	// LATE - late notification of UP path reconfiguration.
	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`
	// URL where notifications shall be sent
	NotificationDestination Link `json:"notificationDestination,omitempty"`
	// Configuration used for sending notifications though web sockets
	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// URL of created subscription resource
	Self Link `json:"self,omitempty"`
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	// Settings for temporary validity of the subscription
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the
	// traffic of UE(s) located in this specific zone.
	ValidGeoZoneIDs []string `json:"validGeoZoneIds,omitempty"`
	// String identifying supported features per Traffic Influence service
	SuppFeat SupportedFeatures `json:"suppFeat,omitempty"`
	// Identifies whether an pplication can be relocated once a location of the
	// application has been selected.Set to "true" if it can be relocated;
	// otherwise set to
	// "false". Default value is "false" if omitted.
	AppReloInd bool `json:"appReloInd,omitempty"`
	//Identifies whether the AF request applies to any UE. This attribute shall
	// set to "true" if
	// applicable for any UE, otherwise, set to "false"
	AnyUeInd bool `json:"anyUeInd,omitempty"`
	// Set to true by the AF to request the NEF to send a test notification.
	//Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`
}

// TrafficInfluSubPatch Traffic Influence Subscription Patch structure
type TrafficInfluSubPatch struct {
	// Identifies whether an application can be relocated once a location of
	// the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty"`
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	// Settings for temporary validity of the subscription
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the
	// traffic of UE(s) located in this specific zone.
	ValidGeoZoneIDs []string `json:"validGeoZoneIds,omitempty"`
}

// EventNotification The UP management event notification is provided by the
// NEF to the AF through the POST method
type EventNotification struct {
	// Identifies an NEF Northbound interface transaction, generated by the AF
	AfTransID string `json:"afTransId,omitempty"`
	// Identifies the type of notification regarding UP path management event.
	DnaiChgType DnaiChangeType `json:"dnaiChgType"`
	// Identifies the N6 traffic routing information associated to the source
	// DNAI. Shall be present if the "subscribedEvent" sets to "UP_PATH_CHANGE".
	SourceTrafficRoute RouteToLocation `json:"sourceTrafficRoute,omitempty"`
	// Identifies a UP path management event the AF requested to be notified of
	SubscribedEvent SubscribedEvent `json:"subscribedEvent,omitempty"`
	// Identifies the N6 traffic routing information associated to the target
	// DNAI. Shall be present if the "subscribedEvent" sets to "UP_PATH_CHANGE".
	TargetTrafficRoute RouteToLocation `json:"targetTrafficRoute,omitempty"`
	// Identifies a user
	Gpsi Gpsi `json:"gpsi,omitempty"`
	// The IPv4 Address of the served UE for the source DNAI.
	SrcUeIpv4Addr Ipv4Addr `json:"srcUeIpv4Addr,omitempty"`
	// The Ipv6 Address Prefix of the served UE for the source DNAI.
	SrcUeIpv6Prefix Ipv6Prefix `json:"srcUeIpv6Prefix,omitempty"`
	// The IPv4 Address of the served UE for the target DNAI.
	TgtUeIpv4Addr Ipv4Addr `json:"tgtUeIpv4Addr,omitempty"`
	// The Ipv6 Address Prefix of the served UE for the target DNAI.
	TgtUeIpv6Prefix Ipv6Prefix `json:"tgtUeIpv6Prefix,omitempty"`
	// UE MAC address of the served UE
	UeMac MacAddr48 `json:"ueMac,omitempty"`
}

// TemporalValidity Indicates the time interval(s) during which the AF request
// is to be applied
type TemporalValidity struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime string `json:"startTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime string `json:"stopTime,omitempty"`
}