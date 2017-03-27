package main

import (
        "log"
        "github.com/BurntSushi/toml"
        "github.com/packethost/packngo"
)

type DeviceConfig struct {
        HostName string `toml:"host_name"`
        Plan string `toml:"plan"`
        OS string `toml:"os"`
        Facility string `toml:"facility"`
        ItamaeRecipe string `toml:"itamae_recipe"`
}

type Config struct {
        ApiToken string `toml:"api_token"`
        ProjectID string `toml:"project_id"`
        Device DeviceConfig `toml:"device"`
}

func main() {
        var conf Config
        if _, err := toml.DecodeFile("pakkun.toml", &conf); err != nil {
                log.Fatal(err)
        }
        packetClient := packngo.NewClient("", conf.ApiToken, nil)
        deviceCreateRequest := packngo.DeviceCreateRequest{}
        deviceCreateRequest.HostName = conf.Device.HostName
        deviceCreateRequest.ProjectID = conf.ProjectID
        deviceCreateRequest.Plan = conf.Device.Plan
        deviceCreateRequest.OS = conf.Device.OS
        deviceCreateRequest.Facility = conf.Device.Facility
        deviceCreateRequest.BillingCycle = "hourly"
        _, _, err := packetClient.Devices.Create(&deviceCreateRequest)
        if err != nil {
                log.Fatal(err)
        }

}
