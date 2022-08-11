package pandas_test


import "testing"
import "strings"
import "regexp"
import "csv"

option now = () => 2030-01-01T00:00:00Z

inData =
    "
#datatype,string,long,dateTime:RFC3339,string,string,string,string,string,string,string
#group,false,false,false,false,true,true,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_time,_value,_field,_measurement,device,fstype,host,path
,,0,2018-05-22T19:53:26Z,a,used_percent,disk,disk1,apfs,host.local,/
,,0,2018-05-22T19:53:36Z,k9ngm,used_percent,disk,disk1,apfs,host.local,/
,,0,2018-05-22T19:53:46Z,b,used_percent,disk,disk1,apfs,host.local,/
,,0,2018-05-22T19:53:56Z,2COTDe,used_percent,disk,disk1,apfs,host.local,/
,,0,2018-05-22T19:54:06Z,cLnSkNMI,used_percent,disk,disk1,apfs,host.local,/
,,0,2018-05-22T19:54:16Z,13F2,used_percent,disk,disk1,apfs,host.local,/
"
outData =
    "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,string,string,string,string,string,string,string,long
#group,false,false,true,true,false,false,true,true,true,true,true,true,false
#default,_result,,,,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,device,fstype,host,path,extract0
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:53:26Z,a,used_percent,disk,disk1,apfs,host.local,/,0
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:53:36Z,k9ngm,used_percent,disk,disk1,apfs,host.local,/,0
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:53:46Z,b,used_percent,disk,disk1,apfs,host.local,/,0
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:53:56Z,2COTDe,used_percent,disk,disk1,apfs,host.local,/,1
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:54:06Z,cLnSkNMI,used_percent,disk,disk1,apfs,host.local,/,0
,,0,2018-05-22T19:53:26Z,2030-01-01T00:00:00Z,2018-05-22T19:54:16Z,13F2,used_percent,disk,disk1,apfs,host.local,/,2
"
re = regexp.compile(v: "[[:alpha:]]{1}")

testcase string_extract_index {
    got =
        csv.from(csv: inData)
            |> testing.load()
            |> range(start: 2018-05-22T19:53:26Z)
            |> map(fn: (r) => ({r with extract0: regexp.findStringIndex(r: re, v: r._value)[0]}))
    want = csv.from(csv: outData)

    testing.diff(got, want)
}
