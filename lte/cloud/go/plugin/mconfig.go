/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package plugin

import (
	"fmt"
	"sort"

	"magma/lte/cloud/go/lte"
	"magma/lte/cloud/go/protos/mconfig"
	"magma/lte/cloud/go/services/cellular/obsidian/models"
	merrors "magma/orc8r/cloud/go/errors"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/configurator"
	models2 "magma/orc8r/cloud/go/services/dnsd/obsidian/models"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Builder struct{}

func (*Builder) Build(networkID string, gatewayID string, graph configurator.EntityGraph, network configurator.Network, mconfigOut map[string]proto.Message) error {
	// we only build an mconfig if cellular network and gateway configs exist
	inwConfig, found := network.Configs[lte.CellularNetworkType]
	if !found || inwConfig == nil {
		return nil
	}
	cellularNwConfig := inwConfig.(*models.NetworkCellularConfigs)

	cellGW, err := graph.GetEntity(lte.CellularGatewayType, gatewayID)
	if err == merrors.ErrNotFound {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}
	if cellGW.Config == nil {
		return nil
	}
	cellularGwConfig := cellGW.Config.(*models.GatewayCellularConfigs)

	if err := validateConfigs(cellularNwConfig, cellularGwConfig); err != nil {
		return err
	}

	enodebs, err := graph.GetAllChildrenOfType(cellGW, lte.CellularEnodebType)
	if err != nil {
		return errors.WithStack(err)
	}

	gwRan := cellularGwConfig.Ran
	gwEpc := cellularGwConfig.Epc
	gwNonEpsService := cellularGwConfig.NonEpsService
	nwRan := cellularNwConfig.Ran
	nwEpc := cellularNwConfig.Epc
	nonEPSServiceMconfig := getNonEPSServiceMconfigFields(gwNonEpsService)

	pipelineDServices, err := getPipelineDServicesConfig(nwEpc.NetworkServices)
	if err != nil {
		return errors.WithStack(err)
	}

	enbConfigsBySerial := getEnodebConfigsBySerial(enodebs)

	vals := map[string]proto.Message{
		"enodebd": &mconfig.EnodebD{
			LogLevel:               protos.LogLevel_INFO,
			Pci:                    int32(gwRan.Pci),
			Earfcndl:               int32(nwRan.Earfcndl),
			FddConfig:              getFddConfig(nwRan.FddConfig),
			TddConfig:              getTddConfig(nwRan.TddConfig),
			SubframeAssignment:     int32(nwRan.SubframeAssignment),
			SpecialSubframePattern: int32(nwRan.SpecialSubframePattern),
			BandwidthMhz:           int32(nwRan.BandwidthMhz),
			AllowEnodebTransmit:    gwRan.TransmitEnabled,
			Tac:                    int32(nwEpc.Tac),
			PlmnidList:             fmt.Sprintf("%s%s", nwEpc.Mcc, nwEpc.Mnc),
			CsfbRat:                nonEPSServiceMconfig.csfbRat,
			Arfcn_2G:               nonEPSServiceMconfig.arfcn_2g,
			EnbConfigsBySerial:     enbConfigsBySerial,
		},
		"mobilityd": &mconfig.MobilityD{
			LogLevel: protos.LogLevel_INFO,
			IpBlock:  gwEpc.IPBlock,
		},
		"mme": &mconfig.MME{
			LogLevel:                 protos.LogLevel_INFO,
			Mcc:                      nwEpc.Mcc,
			Mnc:                      nwEpc.Mnc,
			Tac:                      int32(nwEpc.Tac),
			MmeCode:                  1,
			MmeGid:                   1,
			EnableDnsCaching:         shouldEnableDNSCaching(network),
			NonEpsServiceControl:     nonEPSServiceMconfig.nonEpsServiceControl,
			CsfbMcc:                  nonEPSServiceMconfig.csfbMcc,
			CsfbMnc:                  nonEPSServiceMconfig.csfbMnc,
			Lac:                      nonEPSServiceMconfig.lac,
			RelayEnabled:             nwEpc.RelayEnabled,
			CloudSubscriberdbEnabled: nwEpc.CloudSubscriberdbEnabled,
			AttachedEnodebTacs:       getEnodebTacs(enbConfigsBySerial),
		},
		"pipelined": &mconfig.PipelineD{
			LogLevel:      protos.LogLevel_INFO,
			UeIpBlock:     gwEpc.IPBlock,
			NatEnabled:    gwEpc.NatEnabled,
			DefaultRuleId: nwEpc.DefaultRuleID,
			RelayEnabled:  nwEpc.RelayEnabled,
			Services:      pipelineDServices,
		},
		"subscriberdb": &mconfig.SubscriberDB{
			LogLevel:     protos.LogLevel_INFO,
			LteAuthOp:    nwEpc.LteAuthOp,
			LteAuthAmf:   nwEpc.LteAuthAmf,
			SubProfiles:  getSubProfiles(nwEpc),
			RelayEnabled: nwEpc.RelayEnabled,
		},
		"policydb": &mconfig.PolicyDB{
			LogLevel: protos.LogLevel_INFO,
		},
		"sessiond": &mconfig.SessionD{
			LogLevel:     protos.LogLevel_INFO,
			RelayEnabled: nwEpc.RelayEnabled,
		},
	}
	for k, v := range vals {
		mconfigOut[k] = v
	}
	return nil
}

func validateConfigs(nwConfig *models.NetworkCellularConfigs, gwConfig *models.GatewayCellularConfigs) error {
	if nwConfig == nil {
		return errors.New("Cellular network config is nil")
	}
	if gwConfig == nil {
		return errors.New("Cellular gateway config is nil")
	}

	if gwConfig.Ran == nil {
		return errors.New("Gateway RAN config is nil")
	}
	if gwConfig.Epc == nil {
		return errors.New("Gateway EPC config is nil")
	}
	if nwConfig.Ran == nil {
		return errors.New("Network RAN config is nil")
	}
	if nwConfig.Epc == nil {
		return errors.New("Network EPC config is nil")
	}
	return nil
}

func shouldEnableDNSCaching(network configurator.Network) bool {
	idnsConfig, found := network.Configs[orc8r.DnsdNetworkType]
	if !found || idnsConfig == nil {
		return false
	}
	return idnsConfig.(*models2.NetworkDNSConfig).EnableCaching
}

type nonEPSServiceMconfigFields struct {
	csfbRat              mconfig.EnodebD_CSFBRat
	arfcn_2g             []int32
	nonEpsServiceControl mconfig.MME_NonEPSServiceControl
	csfbMcc              string
	csfbMnc              string
	lac                  int32
}

func getNonEPSServiceMconfigFields(gwNonEpsService *models.GatewayNonEpsServiceConfigs) nonEPSServiceMconfigFields {
	if gwNonEpsService == nil {
		return nonEPSServiceMconfigFields{
			csfbRat:              mconfig.EnodebD_CSFBRAT_2G,
			arfcn_2g:             []int32{},
			nonEpsServiceControl: mconfig.MME_NON_EPS_SERVICE_CONTROL_OFF,
			csfbMcc:              "",
			csfbMnc:              "",
			lac:                  1,
		}
	} else {
		arfcn2g := make([]int32, 0, len(gwNonEpsService.Arfcn2g))
		for _, i := range gwNonEpsService.Arfcn2g {
			arfcn2g = append(arfcn2g, int32(i))
		}

		return nonEPSServiceMconfigFields{
			csfbRat:              mconfig.EnodebD_CSFBRat(gwNonEpsService.CsfbRat),
			arfcn_2g:             arfcn2g,
			nonEpsServiceControl: mconfig.MME_NonEPSServiceControl(gwNonEpsService.NonEpsServiceControl),
			csfbMcc:              gwNonEpsService.CsfbMcc,
			csfbMnc:              gwNonEpsService.CsfbMnc,
			lac:                  int32(gwNonEpsService.Lac),
		}
	}
}

var networkServicesByName = map[string]mconfig.PipelineD_NetworkServices{
	"metering":           mconfig.PipelineD_METERING,
	"dpi":                mconfig.PipelineD_DPI,
	"policy_enforcement": mconfig.PipelineD_ENFORCEMENT,
}

// move this out of this package eventually
func getPipelineDServicesConfig(networkServices []string) ([]mconfig.PipelineD_NetworkServices, error) {
	if networkServices == nil || len(networkServices) == 0 {
		return []mconfig.PipelineD_NetworkServices{
			mconfig.PipelineD_METERING,
			mconfig.PipelineD_DPI,
			mconfig.PipelineD_ENFORCEMENT,
		}, nil
	}
	apps := make([]mconfig.PipelineD_NetworkServices, 0, len(networkServices))
	for _, service := range networkServices {
		mc, found := networkServicesByName[service]
		if !found {
			return nil, errors.Errorf("unknown network service name %s", service)
		}
		apps = append(apps, mc)
	}
	return apps, nil
}

func getFddConfig(fddConfig *models.NetworkRanConfigsFddConfig) *mconfig.EnodebD_FDDConfig {
	if fddConfig == nil {
		return nil
	}
	return &mconfig.EnodebD_FDDConfig{
		Earfcndl: int32(fddConfig.Earfcndl),
		Earfcnul: int32(fddConfig.Earfcnul),
	}
}

func getTddConfig(tddConfig *models.NetworkRanConfigsTddConfig) *mconfig.EnodebD_TDDConfig {
	if tddConfig == nil {
		return nil
	}

	return &mconfig.EnodebD_TDDConfig{
		Earfcndl:               int32(tddConfig.Earfcndl),
		SubframeAssignment:     int32(tddConfig.SubframeAssignment),
		SpecialSubframePattern: int32(tddConfig.SpecialSubframePattern),
	}
}

func getEnodebConfigsBySerial(enodebs []configurator.NetworkEntity) map[string]*mconfig.EnodebD_EnodebConfig {
	ret := make(map[string]*mconfig.EnodebD_EnodebConfig, len(enodebs))
	for _, ent := range enodebs {
		serial := ent.Key
		ienbConfig := ent.Config
		if ienbConfig == nil {
			glog.Errorf("enb with serial %s is missing config", serial)
		}

		cellularEnbConfig := ienbConfig.(*models.NetworkEnodebConfigs)
		ret[serial] = &mconfig.EnodebD_EnodebConfig{
			Earfcndl:               int32(cellularEnbConfig.Earfcndl),
			SubframeAssignment:     int32(cellularEnbConfig.SubframeAssignment),
			SpecialSubframePattern: int32(cellularEnbConfig.SpecialSubframePattern),
			Pci:                    int32(cellularEnbConfig.Pci),
			TransmitEnabled:        cellularEnbConfig.TransmitEnabled,
			DeviceClass:            cellularEnbConfig.DeviceClass,
			BandwidthMhz:           int32(cellularEnbConfig.BandwidthMhz),
			Tac:                    int32(cellularEnbConfig.Tac),
			CellId:                 int32(cellularEnbConfig.CellID),
		}
	}
	return ret
}

func getEnodebTacs(enbConfigsBySerial map[string]*mconfig.EnodebD_EnodebConfig) []int32 {
	ret := make([]int32, 0, len(enbConfigsBySerial))
	for _, enbConfig := range enbConfigsBySerial {
		ret = append(ret, enbConfig.Tac)
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i] < ret[j] })
	return ret
}

func getSubProfiles(epc *models.NetworkEpcConfigs) map[string]*mconfig.SubscriberDB_SubscriptionProfile {
	if epc.SubProfiles == nil {
		return map[string]*mconfig.SubscriberDB_SubscriptionProfile{}
	}

	ret := map[string]*mconfig.SubscriberDB_SubscriptionProfile{}
	for name, profile := range epc.SubProfiles {
		ret[name] = &mconfig.SubscriberDB_SubscriptionProfile{
			MaxUlBitRate: profile.MaxUlBitRate,
			MaxDlBitRate: profile.MaxDlBitRate,
		}
	}
	return ret
}
