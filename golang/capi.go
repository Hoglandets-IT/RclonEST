
package main

import "C"

import (
	"encoding/json"
	"github.com/hoglandets-it/rclonest/golang/endpoints"
)

func Unmarshal(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}

//func RunGeneric(credentials string, jsondata string, endpoint interface{}) (string, string) {
//    Unmarshal(jsondata, &endpoint)
    
//    cl := ClientFromJson(credentials)
//    res, err := cl.Run(endpoint)
//    if (err != nil) {
//        return "", string(err.Error())
//    }
//    return string(res), ""
//}


func RunGeneric(credentials string, jsondata string, path string) (string, string) {
	var creds map[string]string
	json.Unmarshal([]byte(credentials), &creds)

    //jsVal := Remarshal(jsondata, &endpoint)


	// // x := reflect.ValueOf(&endpoint)
	// // fmt.Println(x)

	if creds["url"] == "" || creds["username"] == "" || creds["password"] == "" {
		return "", "Missing credentials"
	}

	client := NewClient(creds["url"], creds["username"], creds["password"])
	res, err := client.RunC()
	// if err != nil {
	// 	return "", string(err.Error())
	// }
	// return string(res), ""

}


//export RunBackendCommand
func RunBackendCommand (credentials string, json_data string) (string, string) {
    var cmBackendCommand endpoints.BackendCommand
    json.Unmarshal([]byte(json_data), &cmBackendCommand)
    
    return RunGeneric(credentials, json_data, "backend/command")
}


//export RunCacheExpire
func RunCacheExpire (credentials string, json_data string) (string, string) {
    var cmCacheExpire endpoints.CacheExpire
    json.Unmarshal([]byte(json_data), &cmCacheExpire)
    
    return RunGeneric(credentials, json_data, "cache/expire")
}


//export RunCacheFetch
func RunCacheFetch (credentials string, json_data string) (string, string) {
    var cmCacheFetch endpoints.CacheFetch
    json.Unmarshal([]byte(json_data), &cmCacheFetch)
    
    return RunGeneric(credentials, json_data, "cache/fetch")
}


//export RunCacheStats
func RunCacheStats (credentials string, json_data string) (string, string) {
    var cmCacheStats endpoints.CacheStats
    json.Unmarshal([]byte(json_data), &cmCacheStats)
    
    return RunGeneric(credentials, json_data, "cache/stats")
}


//export RunConfigCreate
func RunConfigCreate (credentials string, json_data string) (string, string) {
    var cmConfigCreate endpoints.ConfigCreate
    json.Unmarshal([]byte(json_data), &cmConfigCreate)
    
    return RunGeneric(credentials, json_data, "config/create")
}


//export RunConfigDelete
func RunConfigDelete (credentials string, json_data string) (string, string) {
    var cmConfigDelete endpoints.ConfigDelete
    json.Unmarshal([]byte(json_data), &cmConfigDelete)
    
    return RunGeneric(credentials, json_data, "config/delete")
}


//export RunConfigDump
func RunConfigDump (credentials string, json_data string) (string, string) {
    var cmConfigDump endpoints.ConfigDump
    json.Unmarshal([]byte(json_data), &cmConfigDump)
    
    return RunGeneric(credentials, json_data, "config/dump")
}


//export RunConfigGet
func RunConfigGet (credentials string, json_data string) (string, string) {
    var cmConfigGet endpoints.ConfigGet
    json.Unmarshal([]byte(json_data), &cmConfigGet)
    
    return RunGeneric(credentials, json_data, "config/get")
}


//export RunConfigListremotes
func RunConfigListremotes (credentials string, json_data string) (string, string) {
    var cmConfigListremotes endpoints.ConfigListremotes
    json.Unmarshal([]byte(json_data), &cmConfigListremotes)
    
    return RunGeneric(credentials, json_data, "config/listremotes")
}


//export RunConfigPassword
func RunConfigPassword (credentials string, json_data string) (string, string) {
    var cmConfigPassword endpoints.ConfigPassword
    json.Unmarshal([]byte(json_data), &cmConfigPassword)
    
    return RunGeneric(credentials, json_data, "config/password")
}


//export RunConfigProviders
func RunConfigProviders (credentials string, json_data string) (string, string) {
    var cmConfigProviders endpoints.ConfigProviders
    json.Unmarshal([]byte(json_data), &cmConfigProviders)
    
    return RunGeneric(credentials, json_data, "config/providers")
}


//export RunConfigSetpath
func RunConfigSetpath (credentials string, json_data string) (string, string) {
    var cmConfigSetpath endpoints.ConfigSetpath
    json.Unmarshal([]byte(json_data), &cmConfigSetpath)
    
    return RunGeneric(credentials, json_data, "config/setpath")
}


//export RunConfigUpdate
func RunConfigUpdate (credentials string, json_data string) (string, string) {
    var cmConfigUpdate endpoints.ConfigUpdate
    json.Unmarshal([]byte(json_data), &cmConfigUpdate)
    
    return RunGeneric(credentials, json_data, "config/update")
}


//export RunCoreBwlimit
func RunCoreBwlimit (credentials string, json_data string) (string, string) {
    var cmCoreBwlimit endpoints.CoreBwlimit
    json.Unmarshal([]byte(json_data), &cmCoreBwlimit)
    
    return RunGeneric(credentials, json_data, "core/bwlimit")
}


//export RunCoreCommand
func RunCoreCommand (credentials string, json_data string) (string, string) {
    var cmCoreCommand endpoints.CoreCommand
    json.Unmarshal([]byte(json_data), &cmCoreCommand)
    
    return RunGeneric(credentials, json_data, "core/command")
}


//export RunCoreGc
func RunCoreGc (credentials string, json_data string) (string, string) {
    var cmCoreGc endpoints.CoreGc
    json.Unmarshal([]byte(json_data), &cmCoreGc)
    
    return RunGeneric(credentials, json_data, "core/gc")
}


//export RunCoreGroupList
func RunCoreGroupList (credentials string, json_data string) (string, string) {
    var cmCoreGroupList endpoints.CoreGroupList
    json.Unmarshal([]byte(json_data), &cmCoreGroupList)
    
    return RunGeneric(credentials, json_data, "core/group-list")
}


//export RunCoreMemstats
func RunCoreMemstats (credentials string, json_data string) (string, string) {
    var cmCoreMemstats endpoints.CoreMemstats
    json.Unmarshal([]byte(json_data), &cmCoreMemstats)
    
    return RunGeneric(credentials, json_data, "core/memstats")
}


//export RunCoreObscure
func RunCoreObscure (credentials string, json_data string) (string, string) {
    var cmCoreObscure endpoints.CoreObscure
    json.Unmarshal([]byte(json_data), &cmCoreObscure)
    
    return RunGeneric(credentials, json_data, "core/obscure")
}


//export RunCorePid
func RunCorePid (credentials string, json_data string) (string, string) {
    var cmCorePid endpoints.CorePid
    json.Unmarshal([]byte(json_data), &cmCorePid)
    
    return RunGeneric(credentials, json_data, "core/pid")
}


//export RunCoreQuit
func RunCoreQuit (credentials string, json_data string) (string, string) {
    var cmCoreQuit endpoints.CoreQuit
    json.Unmarshal([]byte(json_data), &cmCoreQuit)
    
    return RunGeneric(credentials, json_data, "core/quit")
}


//export RunCoreStats
func RunCoreStats (credentials string, json_data string) (string, string) {
    var cmCoreStats endpoints.CoreStats
    json.Unmarshal([]byte(json_data), &cmCoreStats)
    
    return RunGeneric(credentials, json_data, "core/stats")
}


//export RunCoreStatsDelete
func RunCoreStatsDelete (credentials string, json_data string) (string, string) {
    var cmCoreStatsDelete endpoints.CoreStatsDelete
    json.Unmarshal([]byte(json_data), &cmCoreStatsDelete)
    
    return RunGeneric(credentials, json_data, "core/stats-delete")
}


//export RunCoreStatsReset
func RunCoreStatsReset (credentials string, json_data string) (string, string) {
    var cmCoreStatsReset endpoints.CoreStatsReset
    json.Unmarshal([]byte(json_data), &cmCoreStatsReset)
    
    return RunGeneric(credentials, json_data, "core/stats-reset")
}


//export RunCoreTransferred
func RunCoreTransferred (credentials string, json_data string) (string, string) {
    var cmCoreTransferred endpoints.CoreTransferred
    json.Unmarshal([]byte(json_data), &cmCoreTransferred)
    
    return RunGeneric(credentials, json_data, "core/transferred")
}


//export RunCoreVersion
func RunCoreVersion (credentials string, json_data string) (string, string) {
    var cmCoreVersion endpoints.CoreVersion
    json.Unmarshal([]byte(json_data), &cmCoreVersion)
    
    return RunGeneric(credentials, json_data, "core/version")
}


//export RunDebugSetBlockProfileRate
func RunDebugSetBlockProfileRate (credentials string, json_data string) (string, string) {
    var cmDebugSetBlockProfileRate endpoints.DebugSetBlockProfileRate
    json.Unmarshal([]byte(json_data), &cmDebugSetBlockProfileRate)
    
    return RunGeneric(credentials, json_data, "debug/set-block-profile-rate")
}


//export RunDebugSetGcPercent
func RunDebugSetGcPercent (credentials string, json_data string) (string, string) {
    var cmDebugSetGcPercent endpoints.DebugSetGcPercent
    json.Unmarshal([]byte(json_data), &cmDebugSetGcPercent)
    
    return RunGeneric(credentials, json_data, "debug/set-gc-percent")
}


//export RunDebugSetMutexProfileFraction
func RunDebugSetMutexProfileFraction (credentials string, json_data string) (string, string) {
    var cmDebugSetMutexProfileFraction endpoints.DebugSetMutexProfileFraction
    json.Unmarshal([]byte(json_data), &cmDebugSetMutexProfileFraction)
    
    return RunGeneric(credentials, json_data, "debug/set-mutex-profile-fraction")
}


//export RunDebugSetSoftMemoryLimit
func RunDebugSetSoftMemoryLimit (credentials string, json_data string) (string, string) {
    var cmDebugSetSoftMemoryLimit endpoints.DebugSetSoftMemoryLimit
    json.Unmarshal([]byte(json_data), &cmDebugSetSoftMemoryLimit)
    
    return RunGeneric(credentials, json_data, "debug/set-soft-memory-limit")
}


//export RunFscacheClear
func RunFscacheClear (credentials string, json_data string) (string, string) {
    var cmFscacheClear endpoints.FscacheClear
    json.Unmarshal([]byte(json_data), &cmFscacheClear)
    
    return RunGeneric(credentials, json_data, "fscache/clear")
}


//export RunFscacheEntries
func RunFscacheEntries (credentials string, json_data string) (string, string) {
    var cmFscacheEntries endpoints.FscacheEntries
    json.Unmarshal([]byte(json_data), &cmFscacheEntries)
    
    return RunGeneric(credentials, json_data, "fscache/entries")
}


//export RunJobList
func RunJobList (credentials string, json_data string) (string, string) {
    var cmJobList endpoints.JobList
    json.Unmarshal([]byte(json_data), &cmJobList)
    
    return RunGeneric(credentials, json_data, "job/list")
}


//export RunJobStatus
func RunJobStatus (credentials string, json_data string) (string, string) {
    var cmJobStatus endpoints.JobStatus
    json.Unmarshal([]byte(json_data), &cmJobStatus)
    
    return RunGeneric(credentials, json_data, "job/status")
}


//export RunJobStop
func RunJobStop (credentials string, json_data string) (string, string) {
    var cmJobStop endpoints.JobStop
    json.Unmarshal([]byte(json_data), &cmJobStop)
    
    return RunGeneric(credentials, json_data, "job/stop")
}


//export RunJobStopgroup
func RunJobStopgroup (credentials string, json_data string) (string, string) {
    var cmJobStopgroup endpoints.JobStopgroup
    json.Unmarshal([]byte(json_data), &cmJobStopgroup)
    
    return RunGeneric(credentials, json_data, "job/stopgroup")
}


//export RunMountListmounts
func RunMountListmounts (credentials string, json_data string) (string, string) {
    var cmMountListmounts endpoints.MountListmounts
    json.Unmarshal([]byte(json_data), &cmMountListmounts)
    
    return RunGeneric(credentials, json_data, "mount/listmounts")
}


//export RunMountMount
func RunMountMount (credentials string, json_data string) (string, string) {
    var cmMountMount endpoints.MountMount
    json.Unmarshal([]byte(json_data), &cmMountMount)
    
    return RunGeneric(credentials, json_data, "mount/mount")
}


//export RunMountTypes
func RunMountTypes (credentials string, json_data string) (string, string) {
    var cmMountTypes endpoints.MountTypes
    json.Unmarshal([]byte(json_data), &cmMountTypes)
    
    return RunGeneric(credentials, json_data, "mount/types")
}


//export RunMountUnmount
func RunMountUnmount (credentials string, json_data string) (string, string) {
    var cmMountUnmount endpoints.MountUnmount
    json.Unmarshal([]byte(json_data), &cmMountUnmount)
    
    return RunGeneric(credentials, json_data, "mount/unmount")
}


//export RunMountUnmountall
func RunMountUnmountall (credentials string, json_data string) (string, string) {
    var cmMountUnmountall endpoints.MountUnmountall
    json.Unmarshal([]byte(json_data), &cmMountUnmountall)
    
    return RunGeneric(credentials, json_data, "mount/unmountall")
}


//export RunOperationsAbout
func RunOperationsAbout (credentials string, json_data string) (string, string) {
    var cmOperationsAbout endpoints.OperationsAbout
    json.Unmarshal([]byte(json_data), &cmOperationsAbout)
    
    return RunGeneric(credentials, json_data, "operations/about")
}


//export RunOperationsCleanup
func RunOperationsCleanup (credentials string, json_data string) (string, string) {
    var cmOperationsCleanup endpoints.OperationsCleanup
    json.Unmarshal([]byte(json_data), &cmOperationsCleanup)
    
    return RunGeneric(credentials, json_data, "operations/cleanup")
}


//export RunOperationsCopyfile
func RunOperationsCopyfile (credentials string, json_data string) (string, string) {
    var cmOperationsCopyfile endpoints.OperationsCopyfile
    json.Unmarshal([]byte(json_data), &cmOperationsCopyfile)
    
    return RunGeneric(credentials, json_data, "operations/copyfile")
}


//export RunOperationsCopyurl
func RunOperationsCopyurl (credentials string, json_data string) (string, string) {
    var cmOperationsCopyurl endpoints.OperationsCopyurl
    json.Unmarshal([]byte(json_data), &cmOperationsCopyurl)
    
    return RunGeneric(credentials, json_data, "operations/copyurl")
}


//export RunOperationsDelete
func RunOperationsDelete (credentials string, json_data string) (string, string) {
    var cmOperationsDelete endpoints.OperationsDelete
    json.Unmarshal([]byte(json_data), &cmOperationsDelete)
    
    return RunGeneric(credentials, json_data, "operations/delete")
}


//export RunOperationsDeletefile
func RunOperationsDeletefile (credentials string, json_data string) (string, string) {
    var cmOperationsDeletefile endpoints.OperationsDeletefile
    json.Unmarshal([]byte(json_data), &cmOperationsDeletefile)
    
    return RunGeneric(credentials, json_data, "operations/deletefile")
}


//export RunOperationsFsinfo
func RunOperationsFsinfo (credentials string, json_data string) (string, string) {
    var cmOperationsFsinfo endpoints.OperationsFsinfo
    json.Unmarshal([]byte(json_data), &cmOperationsFsinfo)
    
    return RunGeneric(credentials, json_data, "operations/fsinfo")
}


//export RunOperationsList
func RunOperationsList (credentials string, json_data string) (string, string) {
    var cmOperationsList endpoints.OperationsList
    json.Unmarshal([]byte(json_data), &cmOperationsList)
    
    return RunGeneric(credentials, json_data, "operations/list")
}


//export RunOperationsMkdir
func RunOperationsMkdir (credentials string, json_data string) (string, string) {
    var cmOperationsMkdir endpoints.OperationsMkdir
    json.Unmarshal([]byte(json_data), &cmOperationsMkdir)
    
    return RunGeneric(credentials, json_data, "operations/mkdir")
}


//export RunOperationsMovefile
func RunOperationsMovefile (credentials string, json_data string) (string, string) {
    var cmOperationsMovefile endpoints.OperationsMovefile
    json.Unmarshal([]byte(json_data), &cmOperationsMovefile)
    
    return RunGeneric(credentials, json_data, "operations/movefile")
}


//export RunOperationsPubliclink
func RunOperationsPubliclink (credentials string, json_data string) (string, string) {
    var cmOperationsPubliclink endpoints.OperationsPubliclink
    json.Unmarshal([]byte(json_data), &cmOperationsPubliclink)
    
    return RunGeneric(credentials, json_data, "operations/publiclink")
}


//export RunOperationsPurge
func RunOperationsPurge (credentials string, json_data string) (string, string) {
    var cmOperationsPurge endpoints.OperationsPurge
    json.Unmarshal([]byte(json_data), &cmOperationsPurge)
    
    return RunGeneric(credentials, json_data, "operations/purge")
}


//export RunOperationsRmdir
func RunOperationsRmdir (credentials string, json_data string) (string, string) {
    var cmOperationsRmdir endpoints.OperationsRmdir
    json.Unmarshal([]byte(json_data), &cmOperationsRmdir)
    
    return RunGeneric(credentials, json_data, "operations/rmdir")
}


//export RunOperationsRmdirs
func RunOperationsRmdirs (credentials string, json_data string) (string, string) {
    var cmOperationsRmdirs endpoints.OperationsRmdirs
    json.Unmarshal([]byte(json_data), &cmOperationsRmdirs)
    
    return RunGeneric(credentials, json_data, "operations/rmdirs")
}


//export RunOperationsSize
func RunOperationsSize (credentials string, json_data string) (string, string) {
    var cmOperationsSize endpoints.OperationsSize
    json.Unmarshal([]byte(json_data), &cmOperationsSize)
    
    return RunGeneric(credentials, json_data, "operations/size")
}


//export RunOperationsStat
func RunOperationsStat (credentials string, json_data string) (string, string) {
    var cmOperationsStat endpoints.OperationsStat
    json.Unmarshal([]byte(json_data), &cmOperationsStat)
    
    return RunGeneric(credentials, json_data, "operations/stat")
}


//export RunOperationsUploadfile
func RunOperationsUploadfile (credentials string, json_data string) (string, string) {
    var cmOperationsUploadfile endpoints.OperationsUploadfile
    json.Unmarshal([]byte(json_data), &cmOperationsUploadfile)
    
    return RunGeneric(credentials, json_data, "operations/uploadfile")
}


//export RunOptionsBlocks
func RunOptionsBlocks (credentials string, json_data string) (string, string) {
    var cmOptionsBlocks endpoints.OptionsBlocks
    json.Unmarshal([]byte(json_data), &cmOptionsBlocks)
    
    return RunGeneric(credentials, json_data, "options/blocks")
}


//export RunOptionsGet
func RunOptionsGet (credentials string, json_data string) (string, string) {
    var cmOptionsGet endpoints.OptionsGet
    json.Unmarshal([]byte(json_data), &cmOptionsGet)
    
    return RunGeneric(credentials, json_data, "options/get")
}


//export RunOptionsLocal
func RunOptionsLocal (credentials string, json_data string) (string, string) {
    var cmOptionsLocal endpoints.OptionsLocal
    json.Unmarshal([]byte(json_data), &cmOptionsLocal)
    
    return RunGeneric(credentials, json_data, "options/local")
}


//export RunOptionsSet
func RunOptionsSet (credentials string, json_data string) (string, string) {
    var cmOptionsSet endpoints.OptionsSet
    json.Unmarshal([]byte(json_data), &cmOptionsSet)
    
    return RunGeneric(credentials, json_data, "options/set")
}


//export RunPluginsctlAddplugin
func RunPluginsctlAddplugin (credentials string, json_data string) (string, string) {
    var cmPluginsctlAddplugin endpoints.PluginsctlAddplugin
    json.Unmarshal([]byte(json_data), &cmPluginsctlAddplugin)
    
    return RunGeneric(credentials, json_data, "pluginsctl/addplugin")
}


//export RunPluginsctlGetpluginsfortype
func RunPluginsctlGetpluginsfortype (credentials string, json_data string) (string, string) {
    var cmPluginsctlGetpluginsfortype endpoints.PluginsctlGetpluginsfortype
    json.Unmarshal([]byte(json_data), &cmPluginsctlGetpluginsfortype)
    
    return RunGeneric(credentials, json_data, "pluginsctl/getpluginsfortype")
}


//export RunPluginsctlListplugins
func RunPluginsctlListplugins (credentials string, json_data string) (string, string) {
    var cmPluginsctlListplugins endpoints.PluginsctlListplugins
    json.Unmarshal([]byte(json_data), &cmPluginsctlListplugins)
    
    return RunGeneric(credentials, json_data, "pluginsctl/listplugins")
}


//export RunPluginsctlListtestplugins
func RunPluginsctlListtestplugins (credentials string, json_data string) (string, string) {
    var cmPluginsctlListtestplugins endpoints.PluginsctlListtestplugins
    json.Unmarshal([]byte(json_data), &cmPluginsctlListtestplugins)
    
    return RunGeneric(credentials, json_data, "pluginsctl/listtestplugins")
}


//export RunPluginsctlRemoveplugin
func RunPluginsctlRemoveplugin (credentials string, json_data string) (string, string) {
    var cmPluginsctlRemoveplugin endpoints.PluginsctlRemoveplugin
    json.Unmarshal([]byte(json_data), &cmPluginsctlRemoveplugin)
    
    return RunGeneric(credentials, json_data, "pluginsctl/removeplugin")
}


//export RunPluginsctlRemovetestplugin
func RunPluginsctlRemovetestplugin (credentials string, json_data string) (string, string) {
    var cmPluginsctlRemovetestplugin endpoints.PluginsctlRemovetestplugin
    json.Unmarshal([]byte(json_data), &cmPluginsctlRemovetestplugin)
    
    return RunGeneric(credentials, json_data, "pluginsctl/removetestplugin")
}


//export RunRcError
func RunRcError (credentials string, json_data string) (string, string) {
    var cmRcError endpoints.RcError
    json.Unmarshal([]byte(json_data), &cmRcError)
    
    return RunGeneric(credentials, json_data, "rc/error")
}


//export RunRcList
func RunRcList (credentials string, json_data string) (string, string) {
    var cmRcList endpoints.RcList
    json.Unmarshal([]byte(json_data), &cmRcList)
    
    return RunGeneric(credentials, json_data, "rc/list")
}


//export RunRcNoop
func RunRcNoop (credentials string, json_data string) (string, string) {
    var cmRcNoop endpoints.RcNoop
    json.Unmarshal([]byte(json_data), &cmRcNoop)
    
    return RunGeneric(credentials, json_data, "rc/noop")
}


//export RunRcNoopauth
func RunRcNoopauth (credentials string, json_data string) (string, string) {
    var cmRcNoopauth endpoints.RcNoopauth
    json.Unmarshal([]byte(json_data), &cmRcNoopauth)
    
    return RunGeneric(credentials, json_data, "rc/noopauth")
}


//export RunSyncBisync
func RunSyncBisync (credentials string, json_data string) (string, string) {
    var cmSyncBisync endpoints.SyncBisync
    json.Unmarshal([]byte(json_data), &cmSyncBisync)
    
    return RunGeneric(credentials, json_data, "sync/bisync")
}


//export RunSyncCopy
func RunSyncCopy (credentials string, json_data string) (string, string) {
    var cmSyncCopy endpoints.SyncCopy
    json.Unmarshal([]byte(json_data), &cmSyncCopy)
    
    return RunGeneric(credentials, json_data, "sync/copy")
}


//export RunSyncMove
func RunSyncMove (credentials string, json_data string) (string, string) {
    var cmSyncMove endpoints.SyncMove
    json.Unmarshal([]byte(json_data), &cmSyncMove)
    
    return RunGeneric(credentials, json_data, "sync/move")
}


//export RunSyncSync
func RunSyncSync (credentials string, json_data string) (string, string) {
    var cmSyncSync endpoints.SyncSync
    json.Unmarshal([]byte(json_data), &cmSyncSync)
    
    return RunGeneric(credentials, json_data, "sync/sync")
}


//export RunVfsForget
func RunVfsForget (credentials string, json_data string) (string, string) {
    var cmVfsForget endpoints.VfsForget
    json.Unmarshal([]byte(json_data), &cmVfsForget)
    
    return RunGeneric(credentials, json_data, "vfs/forget")
}


//export RunVfsList
func RunVfsList (credentials string, json_data string) (string, string) {
    var cmVfsList endpoints.VfsList
    json.Unmarshal([]byte(json_data), &cmVfsList)
    
    return RunGeneric(credentials, json_data, "vfs/list")
}


//export RunVfsPollInterval
func RunVfsPollInterval (credentials string, json_data string) (string, string) {
    var cmVfsPollInterval endpoints.VfsPollInterval
    json.Unmarshal([]byte(json_data), &cmVfsPollInterval)
    
    return RunGeneric(credentials, json_data, "vfs/poll-interval")
}


//export RunVfsRefresh
func RunVfsRefresh (credentials string, json_data string) (string, string) {
    var cmVfsRefresh endpoints.VfsRefresh
    json.Unmarshal([]byte(json_data), &cmVfsRefresh)
    
    return RunGeneric(credentials, json_data, "vfs/refresh")
}


//export RunVfsStats
func RunVfsStats (credentials string, json_data string) (string, string) {
    var cmVfsStats endpoints.VfsStats
    json.Unmarshal([]byte(json_data), &cmVfsStats)
    
    return RunGeneric(credentials, json_data, "vfs/stats")
}

