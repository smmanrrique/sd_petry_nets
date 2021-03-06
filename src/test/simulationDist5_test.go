package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	dcs "sd_petry_nets/src/distconssim"
	u "sd_petry_nets/src/utils"

	"gopkg.in/ini.v1"
)

var path5 string
var environment5 string
var subNetNames5 []string
var subNetIDS5 []string
var connect5 u.Connections

func init() {
	testing.Init()
	gob.Register(&u.Message{})
	gob.Register(&dcs.EventDist{})
	gob.Register(dcs.IndGlobalTrans(0))
	gob.Register(dcs.TypeClock(0))
	gob.Register(&dcs.LefsDist{})
	gob.Register(&dcs.TransitionConstant{})
	gob.Register(&dcs.TransitionList{})

	// Loading configuration file
	cfg, err := ini.Load("../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment5 = cfg.Section("general").Key("environment").String()
	subNetNames5 = strings.Split(cfg.Section(environment5).Key("subNetName5").String(), ",")
	subNetIDS5 = strings.Split(cfg.Section(environment5).Key("subNetID5").String(), ",")
	connect5 = u.NewConnec(subNetIDS5)
}

func Test5Dist(t *testing.T) {
	println("------------------------------- ESTOY TestSSHDist5 ---------------------------------------")
	for i, ip := range subNetIDS5 {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])

		println(pathTest+subNetNames5[i], ip, addr)

		go u.ExcecuteSSH(pathTest+subNetNames5[i], connection)
	}

	time.Sleep(80 * time.Second)
}

func TestSubNet51(t *testing.T) {

	IDSubNet := connect5.GetConnection(0)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet51.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(IDSubNet.GetIDSubRed(), IDSubNet.GetIp(), IDSubNet.GetPort(), IDSubNet.GetIds())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T0
			dcs.TransitionDist{
				IDGlobal:       0,
				IDLocal:        0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{-1, -1},
					dcs.TransitionConstant{-3, -1},
					dcs.TransitionConstant{-5, -1},
					dcs.TransitionConstant{-7, -1},
				},
			},
			// T9
			dcs.TransitionDist{
				IDGlobal:       9,
				IDLocal:        1,
				IiValorLef:     4,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 4},
					dcs.TransitionConstant{0, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			2: connect5.GetConnection(1),
			4: connect5.GetConnection(2),
			6: connect5.GetConnection(3),
			8: connect5.GetConnection(4),
		},
		Post: dcs.Incidence{
			1: connect5.GetConnection(1),
			3: connect5.GetConnection(2),
			5: connect5.GetConnection(3),
			7: connect5.GetConnection(4),
		},
	}
	// log.Println(IDSubNet)
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(8 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet52(t *testing.T) {

	IDSubNet := connect5.GetConnection(1)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet52.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(IDSubNet.GetIDSubRed(), IDSubNet.GetIp(), IDSubNet.GetPort(), IDSubNet.GetIds())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T1
			dcs.TransitionDist{
				IDGlobal:       1,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T2
			dcs.TransitionDist{
				IDGlobal:       2,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(8 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet53(t *testing.T) {

	IDSubNet := connect5.GetConnection(2)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet53.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(IDSubNet.GetIDSubRed(), IDSubNet.GetIp(), IDSubNet.GetPort(), IDSubNet.GetIds())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T3
			dcs.TransitionDist{
				IDGlobal:       3,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T4
			dcs.TransitionDist{
				IDGlobal:       4,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(8 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet54(t *testing.T) {

	IDSubNet := connect5.GetConnection(3)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet54.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(IDSubNet.GetIDSubRed(), IDSubNet.GetIp(), IDSubNet.GetPort(), IDSubNet.GetIds())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T5
			dcs.TransitionDist{
				IDGlobal:       5,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T6
			dcs.TransitionDist{
				IDGlobal:       6,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(8 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet55(t *testing.T) {

	IDSubNet := connect5.GetConnection(4)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet55.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(IDSubNet.GetIDSubRed(), IDSubNet.GetIp(), IDSubNet.GetPort(), IDSubNet.GetIds())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T7
			dcs.TransitionDist{
				IDGlobal:       7,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T8
			dcs.TransitionDist{
				IDGlobal:       8,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(8 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}
