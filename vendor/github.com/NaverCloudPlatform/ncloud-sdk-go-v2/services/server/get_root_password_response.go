/*
 * server
 *
 * <br/>https://ncloud.beta-apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:08:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetRootPasswordResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

RootPassword *string `json:"rootPassword,omitempty"`
}
