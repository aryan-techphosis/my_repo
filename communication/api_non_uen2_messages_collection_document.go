// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package communication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/omec-project/amf/logger"
)

// NonUeN2MessageTransfer - Namf_Communication Non UE N2 Message Transfer service Operation
func HTTPNonUeN2MessageTransfer(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Message Transfer is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
