// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcentercompliancestats

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterComplianceStatsV1Client represents a custom type interface
type VmwareVcenterComplianceStatsV1Client interface {
    //  Returns a representation of the compliance statistics of VMs in the specified vCenter server.
    ReadVmwareVcenterComplianceStats(
        vcenterId string)(
        *models.ReadVMwareVCenterProtectionStatsResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterComplianceStatsV1 returns VmwareVcenterComplianceStatsV1Client
func NewVmwareVcenterComplianceStatsV1(config config.Config) VmwareVcenterComplianceStatsV1Client{
    client := new(VmwareVcenterComplianceStatsV1)
    client.config = config
    return client
}
