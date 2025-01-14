/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"net/http"

	"magma/lte/cloud/go/lte"
	"magma/lte/cloud/go/services/policydb/obsidian/models"
	"magma/orc8r/cloud/go/errors"
	"magma/orc8r/cloud/go/obsidian/handlers"
	"magma/orc8r/cloud/go/services/configurator"

	"github.com/labstack/echo"
)

func listBaseNames(c echo.Context) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	baseNames, err := configurator.ListEntityKeys(networkID, lte.BaseNameEntityType)
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, baseNames)
}

func createBaseName(c echo.Context) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	bnr := new(models.BaseNameRecord)
	if err := c.Bind(bnr); err != nil {
		return handlers.HttpError(err, http.StatusBadRequest)
	}

	_, err := configurator.CreateEntity(networkID, bnr.ToEntity())
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, string(bnr.Name))
}

func getBaseName(c echo.Context) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	baseName := getBaseNameParam(c)
	if len(baseName) == 0 {
		return baseNameHTTPErr()
	}

	ret, err := configurator.LoadEntity(
		networkID,
		lte.BaseNameEntityType,
		baseName,
		configurator.EntityLoadCriteria{LoadAssocsFromThis: true},
	)
	if err == errors.ErrNotFound {
		return handlers.HttpError(err, http.StatusNotFound)
	}
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, (&models.BaseNameRecord{}).FromEntity(ret).RuleNames)
}

func updateBaseName(c echo.Context) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	baseName := getBaseNameParam(c)
	if len(baseName) == 0 {
		return baseNameHTTPErr()
	}

	ruleNames := models.RuleNames{}
	if err := c.Bind(&ruleNames); err != nil {
		return handlers.HttpError(err, http.StatusBadRequest)
	}
	_, err := configurator.UpdateEntity(
		networkID,
		configurator.EntityUpdateCriteria{
			Type:              lte.BaseNameEntityType,
			Key:               baseName,
			AssociationsToSet: ruleNames.ToAssocs(),
		},
	)
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteBaseName(c echo.Context) error {
	networkID, nerr := handlers.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	baseName := getBaseNameParam(c)
	if len(baseName) == 0 {
		return baseNameHTTPErr()
	}

	err := configurator.DeleteEntity(networkID, lte.BaseNameEntityType, baseName)
	if err != nil {
		return handlers.HttpError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
