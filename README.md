# modularGet
Iskratel OLT utility for getting various endpoints from the Restconf API

| Flag | Description |
| ------ | ------ |
| -h  | Show this help |
| -ip | Specify the OLT Management IP (default 10.5.100.10) |
| -a  | Run all example endpoints |

# Example Call
```sh
./modularGet -ip 10.5.80.10 msanOnuCfgTable
```

# Example Output
```sh
>> GET Request https://10.5.80.10/restconf/data/ISKRATEL-MSAN-MIB:ISKRATEL-MSAN-MIB/msanOnuCfgTable
{
    "ISKRATEL-MSAN-MIB:": {
        "ISKRATEL-MSAN-MIB": {
            "msanOnuCfgTable": {
                "msanOnuCfgEntry": [
                    {
                        "msanOnuCfgIfName": "0/2/1",
                        "msanOnuCfgRegistrationid": "\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000",
                        "msanOnuCfgEnablePm": 2,
                        "msanOnuCfgSerialNumber": "ISKT76E83668",
                        "msanOnuCfgAdminState": 1,
                        "msanOnuCfgOnuDhcpMode": 1,
                        "msanOnuCfgOnuIpAddress": "10.5.81.13",
                        "msanOnuCfgOnuIPMask": "255.255.255.0",
                        "msanOnuCfgOnuDefaultGateway": "0.0.0.0",
                        "msanOnuCfgDefaultConfigFile": "",
                        "msanOnuCfgSendConfig": 2,
                        "msanOnuCfgSendConfigStatus": 6,
                        "msanOnuCfgOnuResync": 2,
                        "msanOnuCfgOnuResetFactoryDefault": 2
                    },
                    {
                        "msanOnuCfgIfName": "0/2/2",
                        "msanOnuCfgRegistrationid": "\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000",
                        "msanOnuCfgEnablePm": 2,
                        "msanOnuCfgSerialNumber": "ISKT76E83E1B",
                        "msanOnuCfgAdminState": 1,
                        "msanOnuCfgOnuDhcpMode": 1,
                        "msanOnuCfgOnuIpAddress": "10.5.81.14",
                        "msanOnuCfgOnuIPMask": "255.255.255.0",
                        "msanOnuCfgOnuDefaultGateway": "0.0.0.0",
                        "msanOnuCfgDefaultConfigFile": "",
                        "msanOnuCfgSendConfig": 2,
                        "msanOnuCfgSendConfigStatus": 6,
                        "msanOnuCfgOnuResync": 2,
                        "msanOnuCfgOnuResetFactoryDefault": 2
                    }
                ]
            }
        }
    }
}
```

# Call for Help in Selecting Valid Endpoints
```sh
./modularGet -h
```

# Help Output
```sh
'modularGet' [options] <endpoints>
  -a    Run all example endpoints
  -h    Show this help
  -ip string
        OLT Management IP (default "10.5.100.10")
Example Endpoints:
msanSwInfo
msanSwComponentTable
msanLicenseGlobal
msanCpu
msanCpuDetailTable
msanMemory
msanDiskTable
msanSystem
msanSntpGlobal
msanSnmpGlobal
msanIgmpSnoopingGlobal
msanDhcpRaGlobal
msanAcsGlobal
msanNetworkGlobal
msanFtpServerGlobal
msanOltGlobal
msanErrorDisablePortTable
msanDiagnosticsTempTable
msanDiagnosticsExtAlrSensorTable
msanOnuUpgradeTable
msanOnuUpgradeHwTypeTable
msanOnuUpgradeServerTable
msanOnuDefaultConfigFileTable
msanOnuCfgAttributeTable
msanOltPortTable
msanOltOnuPortStatTable
msanServiceProfileTable
msanServiceFlowProfileTable
msanVlanProfileTable
msanMulticastProfileTable
msanSecurityProfileTable
msanOnuFlowProfileTable
msanOnuTcontProfileTable
msanOnuVlanProfileTable
msanOnuVlanProfileRuleTable
msanOnuMulticastProfileTable
msanL2cpProfileTable
msanOnuBlackListTable
msanOnuCfgTable
msanServicePortProfileTable
msanOnuInfoTable
```
