// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenterdatacentercompliancestats

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterDatacenterComplianceStatsV1Client represents a custom type interface
type VmwareVcenterDatacenterComplianceStatsV1Client interface {
    //  Returns a representation of the compliance statistics of the specified data center.
    ReadVmwareVcenterDatacenterComplianceStats(
        vcenterId string, 
        datacenterId string)(
        *models.ReadVMwareDatacenterStatsResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterDatacenterComplianceStatsV1 returns VmwareVcenterDatacenterComplianceStatsV1Client
func NewVmwareVcenterDatacenterComplianceStatsV1(config config.Config) VmwareVcenterDatacenterComplianceStatsV1Client{
    client := new(VmwareVcenterDatacenterComplianceStatsV1)
    client.config = config
    return client
}
