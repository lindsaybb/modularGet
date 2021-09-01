package main

import (
        "fmt"
        "flag"
        "net/http"
        "io/ioutil"
        "crypto/tls"
)

var (
        helpFlag = flag.Bool("h", false, "Show this help")
        hostFlag = flag.String("ip", "10.5.100.10", "OLT Management IP")
        autoFlag = flag.Bool("a", false, "Run all example endpoints")
)

const usage = "'modularGet' [options] <endpoints>"

var endpointList = []string{
        "msanSwInfo",
        "msanSwComponentTable",
        "msanLicenseGlobal",
        "msanCpu",
        "msanCpuDetailTable",
        "msanMemory",
        "msanDiskTable",
        "msanSystem",
        "msanSntpGlobal",
        "msanSnmpGlobal",
        "msanIgmpSnoopingGlobal",
        "msanDhcpRaGlobal",
        "msanAcsGlobal",
        "msanNetworkGlobal",
        "msanFtpServerGlobal",
        "msanOltGlobal",
        "msanErrorDisablePortTable",
        "msanDiagnosticsTempTable",
        "msanDiagnosticsExtAlrSensorTable",
        "msanOnuUpgradeTable",
        "msanOnuUpgradeHwTypeTable",
        "msanOnuUpgradeServerTable",
        "msanOnuDefaultConfigFileTable",
        "msanOnuCfgAttributeTable",
        "msanOltPortTable",
        "msanOltOnuPortStatTable",
        "msanServiceProfileTable",
        "msanServiceFlowProfileTable",
        "msanVlanProfileTable",
        "msanMulticastProfileTable",
        "msanSecurityProfileTable",
        "msanOnuFlowProfileTable",
        "msanOnuTcontProfileTable",
        "msanOnuVlanProfileTable",
        "msanOnuVlanProfileRuleTable",
        "msanOnuMulticastProfileTable",
        "msanL2cpProfileTable",
        "msanOnuBlackListTable",
        "msanOnuCfgTable",
        "msanServicePortProfileTable",
        "msanOnuInfoTable",
}

func help() {
        fmt.Println("Example Endpoints:")
        for _, e := range endpointList {
                fmt.Println(e)
        }
}

func main() {
        flag.Parse()
        if *helpFlag || (flag.NArg() < 1 && !*autoFlag) {
                fmt.Println(usage)
                flag.PrintDefaults()
                help()
                return
        }
        var endpoints []string
        if *autoFlag {
                endpoints = endpointList
        } else {
                endpoints = flag.Args()
        }
        for _, e := range endpoints {
                tr := &http.Transport{
                        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
                }
                client := &http.Client{Transport: tr}
                url := fmt.Sprintf("https://%s/restconf/data/ISKRATEL-MSAN-MIB:ISKRATEL-MSAN-MIB/%s", *hostFlag, e)
                req, err := http.NewRequest("GET", url, nil)
                if err != nil {
                        fmt.Println(err)
                }
                req.Header.Set("Cookie", "session=em+protection-user=admin&em+protection-pw=admin")
                fmt.Printf(">> GET Request %s\n", url)
                resp, err := client.Do(req)
                if err != nil {
                        fmt.Println(err)
                }
                defer resp.Body.Close()
                data, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(string(data))
        }
}
