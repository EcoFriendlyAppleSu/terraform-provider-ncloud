/*
 * server
 *
 * <br/>https://ncloud.beta-apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:08:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeletePortForwardingRulesResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

	// 포트포워딩설정번호
PortForwardingConfigurationNo *string `json:"portForwardingConfigurationNo,omitempty"`

	// 포트포워딩공인IP
PortForwardingPublicIp *string `json:"portForwardingPublicIp,omitempty"`

	// ZONE
Zone *Zone `json:"zone,omitempty"`

	// 인터넷회선구분
InternetLineType *CommonCode `json:"internetLineType,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

PortForwardingRuleList []*PortForwardingRule `json:"portForwardingRuleList,omitempty"`
}
