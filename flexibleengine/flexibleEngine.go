package flexibleengine

import (
	"fmt"

	"zabbix.com/pkg/plugin"
	"zabbix.com/plugins/flexibleengine/ecs"
	"zabbix.com/plugins/flexibleengine/evs"
	"zabbix.com/plugins/flexibleengine/nat"
	"zabbix.com/plugins/flexibleengine/rds"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	switch key {
	case "flexibleengine.nat.connections":
		result, err = nat.CalculConnection(params, "snat_connection")
		return
	case "flexibleengine.ecs.cpu":
		result, err = ecs.CalculCPU(params, "cpu_util")
		return
	case "flexibleengine.ecs.disk.free":
		result, err = ecs.CalculDisk(params, "SlAsH_disk_free")
		return
	case "flexibleengine.ecs.disk.used":
		result, err = ecs.CalculDisk(params, "SlAsH_disk_usedPercent")
		return
	case "flexibleengine.ecs.diskio.read":
		result, err = ecs.CalculDiskIO(params, "disk_read_bytes_rate")
		return
	case "flexibleengine.ecs.diskio.write":
		result, err = ecs.CalculDiskIO(params, "disk_write_bytes_rate")
		return
	case "flexibleengine.ecs.diskio.requestread":
		result, err = ecs.CalculDiskIO(params, "disk_read_requests_rate")
		return
	case "flexibleengine.ecs.diskio.requestwrite":
		result, err = ecs.CalculDiskIO(params, "disk_write_requests_rate")
		return
	case "flexibleengine.ecs.health":
		result, err = ecs.CalculHealth(params)
		return
	case "flexibleengine.ecs.memory":
		result, err = ecs.CalculMemory(params, "mem_usedPercent")
		return
	case "flexibleengine.ecs.network.incominginband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_rate_inband")
		return
	case "flexibleengine.ecs.network.outgoinginband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_rate_inband")
		return
	case "flexibleengine.ecs.network.incomingoutband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.network.outgoingoutband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.proc.total":
		result, err = ecs.CalculProcess(params, "proc_total_count")
		return
	case "flexibleengine.ecs.proc.running":
		result, err = ecs.CalculProcess(params, "proc_running_count")
		return
	case "flexibleengine.ecs.proc.zombie":
		result, err = ecs.CalculProcess(params, "proc_zombie_count")
		return
	case "flexibleengine.ecs.proc.sleeping":
		result, err = ecs.CalculProcess(params, "proc_sleeping_count")
		return
	case "flexibleengine.ecs.proc.idle":
		result, err = ecs.CalculProcess(params, "proc_idle_count")
		return
	case "flexibleengine.ecs.status":
		result, err = ecs.CalculStatus(params)
		return
	case "flexibleengine.rds.connections":
		result, err = rds.CalculConnection(params)
		return
	case "flexibleengine.rds.cpu":
		result, err = rds.CalculCPU(params)
		return
	case "flexibleengine.rds.diskio.read":
		result, err = rds.CalculDiskIO(params, "rds049_disk_read_throughput")
		return
	case "flexibleengine.rds.diskio.write":
		result, err = rds.CalculDiskIO(params, "rds050_disk_write_throughput")
		return
	case "flexibleengine.rds.health":
		result, err = rds.CalculHealth(params)
		return
	case "flexibleengine.rds.memory":
		result, err = rds.CalculMemory(params, "rds002_mem_util")
		return
	case "flexibleengine.rds.network.input":
		result, err = rds.CalculNetwork(params, "rds004_bytes_in")
		return
	case "flexibleengine.rds.network.output":
		result, err = rds.CalculNetwork(params, "rds005_bytes_out")
		return
	case "flexibleengine.rds.storage":
		result, err = rds.CalculStorage(params, "rds039_disk_util")
		return
	case "flexibleengine.rds.transaction":
		result, err = rds.CalculTransaction(params, "rds009_tps")
		return
	case "flexibleengine.rds.querie":
		result, err = rds.CalculQuerie(params, "rds008_qps")
		return
	case "flexibleengine.rds.querie.delete":
		result, err = rds.CalculQuerie(params, "rds028_comdml_del_count")
		return
	case "flexibleengine.rds.querie.insert":
		result, err = rds.CalculQuerie(params, "rds029_comdml_ins_count")
		return
	case "flexibleengine.rds.querie.insertselect":
		result, err = rds.CalculQuerie(params, "rds030_comdml_ins_sel_count")
		return
	case "flexibleengine.rds.querie.replace":
		result, err = rds.CalculQuerie(params, "rds031_comdml_rep_count")
		return
	case "flexibleengine.rds.querie.replaceselection":
		result, err = rds.CalculQuerie(params, "rds032_comdml_rep_sel_count")
		return
	case "flexibleengine.rds.querie.select":
		result, err = rds.CalculQuerie(params, "rds033_comdml_sel_count")
		return
	case "flexibleengine.rds.querie.update":
		result, err = rds.CalculQuerie(params, "rds034_comdml_upd_count")
		return
	case "flexibleengine.evs.diskio.read":
		result, err = evs.CalculDiskIO(params, "disk_device_read_bytes_rate")
		return
	case "flexibleengine.evs.diskio.write":
		result, err = evs.CalculDiskIO(params, "disk_device_write_bytes_rate")
		return
	case "flexibleengine.evs.diskio.requestread":
		result, err = evs.CalculDiskIO(params, "disk_device_read_requests_rate")
		return
	case "flexibleengine.evs.diskio.requestwrite":
		result, err = evs.CalculDiskIO(params, "disk_device_write_requests_rate")
		return
	case "flexibleengine.evs.diskio.readoperation":
		result, err = evs.CalculDiskIO(params, "disk_device_read_bytes_per_operation")
		return
	case "flexibleengine.evs.diskio.readawait":
		result, err = evs.CalculDiskIO(params, "disk_device_read_await")
		return
	case "flexibleengine.evs.diskio.queuelength":
		result, err = evs.CalculDiskIO(params, "disk_device_queue_length")
		return
	case "flexibleengine.evs.diskio.ioutil":
		result, err = evs.CalculDiskIO(params, "disk_device_io_util")
		return
	case "flexibleengine.evs.diskio.iosvctm":
		result, err = evs.CalculDiskIO(params, "disk_device_io_svctm")
		return
	case "flexibleengine.evs.status":
		result, err = evs.CalculStatus(params)
		return
	default:
		return nil, fmt.Errorf("Invalid KEY")
	}

}

func init() {
	plugin.RegisterMetrics(&impl, "FlexibleEngine",
		"flexibleengine.nat.connections", "Returns connection count.",
		"flexibleengine.ecs.cpu", "Returns CPU value.",
		"flexibleengine.ecs.disk.free", "Returns disk available space.",
		"flexibleengine.ecs.disk.used", "Returns disk usage.",
		"flexibleengine.ecs.diskio.read", "Returns disk read bytes rate.",
		"flexibleengine.ecs.diskio.write", "Returns disk write bytes rate.",
		"flexibleengine.ecs.diskio.requestread", "Returns disk read ops.",
		"flexibleengine.ecs.diskio.requestwrite", "Returns disk write ops.",
		"flexibleengine.ecs.health", "Returns health.",
		"flexibleengine.ecs.memory", "Returns memory used.",
		"flexibleengine.ecs.network.incominginband", "Returns network inband incoming bytes rate.",
		"flexibleengine.ecs.network.outgoinginband", "Returns network inband outgoing bytes rate.",
		"flexibleengine.ecs.network.incomingoutband", "Returns network outband incoming bytes rate.",
		"flexibleengine.ecs.network.outgoingoutband", "Returns network outband outgoing bytes rate.",
		"flexibleengine.ecs.proc.total", "Returns total process.",
		"flexibleengine.ecs.proc.running", "Returns running count.",
		"flexibleengine.ecs.proc.zombie", "Returns zombie process.",
		"flexibleengine.ecs.proc.sleeping", "Returns sleeping process.",
		"flexibleengine.ecs.proc.idle", "Returns idle process.",
		"flexibleengine.ecs.status", "Returns status ecs.",
		"flexibleengine.rds.connections", "Returns connection count.",
		"flexibleengine.rds.cpu", "Returns CPU value.",
		"flexibleengine.rds.diskio.read", "Returns disk read throughput.",
		"flexibleengine.rds.diskio.write", "Returns disk write throughput.",
		"flexibleengine.rds.health", "Returns health.",
		"flexibleengine.rds.memory", "Returns memory used.",
		"flexibleengine.rds.network.input", "Returns network input throughput.",
		"flexibleengine.rds.network.output", "Returns network output throughput.",
		"flexibleengine.rds.storage", "Returns storage utilization.",
		"flexibleengine.rds.transaction", "Returns transactions per second.",
		"flexibleengine.rds.querie", "Returns queries per seconde.",
		"flexibleengine.rds.querie.delete", "Returns delete statements per second.",
		"flexibleengine.rds.querie.insert", "Returns insert statements per second.",
		"flexibleengine.rds.querie.insertselect", "Returns insert/select statements per second.",
		"flexibleengine.rds.querie.replace", "Returns replace statements per second.",
		"flexibleengine.rds.querie.replaceselection", "Returns replace_selection statements per second.",
		"flexibleengine.rds.querie.select", "Returns select statements per second.",
		"flexibleengine.rds.querie.update", "Returns update statements per second.",
		"flexibleengine.evs.diskio.read", "Returns disk read bytes rate.",
		"flexibleengine.evs.diskio.write", "Returns disk write bytes rate.",
		"flexibleengine.evs.diskio.requestread", "Returns disk read ops.",
		"flexibleengine.evs.diskio.requestwrite", "Returns disk write ops.",
		"flexibleengine.evs.diskio.readoperation", "Returns avg disk read bytes per operation.",
		"flexibleengine.evs.diskio.readawait", "Returns disk read await.",
		"flexibleengine.evs.diskio.queuelength", "Returns average queue length.",
		"flexibleengine.evs.diskio.ioutil", "Returns disk I/O utilization.",
		"flexibleengine.evs.diskio.iosvctm", "Returns disk I/O service time.",
		"flexibleengine.evs.status", "Returns status evs.")
}
