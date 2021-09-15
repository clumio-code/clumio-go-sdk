// Copyright (c) 2021 Clumio All Rights Reserved

package mssqlhosts

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// MssqlHostsV1Client represents a custom type interface
type MssqlHostsV1Client interface {
    //  Returns a list of hosts
    ListMssqlHostConnections(
        currentCount *int64, 
        filter *string, 
        limit *int64, 
        start *string)(
        *models.ListHcmHostsResponse,  *apiutils.APIError)
    
    //  Create a MSSQL Connection.
    CreateMssqlHostConnections(
        body *models.CreateMssqlHostConnectionsV1Request)(
        *models.CreateHcmHostResponse,  *apiutils.APIError)
    
    //  Delete the specified MSSQL host.
    DeleteMssqlHostConnections(
        embed *string, 
        body *models.DeleteMssqlHostConnectionsV1Request)(
        *models.DeleteHcmHostResponse,  *apiutils.APIError)
    
    //  Move the specified MSSQL hosts from a source Sub-Group to a destination Sub-Group.
    MoveMssqlHostConnections(
        embed *string, 
        body *models.MoveMssqlHostConnectionsV1Request)(
        *models.MoveHcmHostsResponse,  *apiutils.APIError)
    
    //  Create Edge Connector Credentials for the specified MSSQL host.
    CreateMssqlHostConnectionCredentials(
        body *models.CreateMssqlHostConnectionCredentialsV1Request)(
        *models.CreateHostECCredentialsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified host.
    ReadMssqlHostConnections(
        hostId string)(
        *models.ReadHcmHostResponse,  *apiutils.APIError)
    
    //  Returns a list of hosts
    ListMssqlHosts(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListMssqlHostsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified host.
    ReadMssqlHosts(
        hostId string)(
        *models.ReadMssqlHostResponse,  *apiutils.APIError)
    
}

// NewMssqlHostsV1 returns MssqlHostsV1Client
func NewMssqlHostsV1(config config.Config) MssqlHostsV1Client{
    client := new(MssqlHostsV1)
    client.config = config
    return client
}
