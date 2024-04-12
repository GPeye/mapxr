package main

import (
	"runtime"
	"time"

	"mapxr/config"

	"github.com/GPeye/tapgosdk"
	"github.com/GPeye/tapgosdk/events"
	"github.com/micmonay/keybd_event"
)

type MapXR struct {
	tapManager *tapgosdk.TapManager
	tapHandler onTapped
	kb         keybd_event.KeyBonding
	config     *config.Config
}

type onTapped struct {
	//kb   *keybd_event.KeyBonding
	mpxr *MapXR
}

func (ot *onTapped) Handle(tap uint8) {
	println("tapdata: ", tap)
	if tap > 0 {
		//ot.mpxr.kb.SetKeys(ot.mpxr.config.OneHandedMap.Bindings)
		//idx := slices.IndexFunc(ot.mpxr.config.OneHandedMap.Bindings, func(c config.TapBinding) bool { return c.TapCode == tap })
		//if idx < 0 {
		//	return
		//}
		//ot.mpxr.kb.SetKeys(int(ot.mpxr.config.OneHandedMap.Bindings[idx].Action.Keypress))
		//ot.mpxr.kb.SetKeys(tapmap[tap])
		//ot.mpxr.kb.Launching()
		//ot.mpxr.kb.Clear()
	}
}

func main() {
	mapXr := MapXR{}

	c, e := config.LoadConfig()
	if e != nil {
		println(e.Error())
		return
	}

	mapXr.config = c
	println(mapXr.config.Version)
	//println(mapXr.config.OneHandedMap.Bindings[0].TapCode)
	//println(mapXr.config.OneHandedMap.Bindings[0].Keypress)
	//mapXr.config.OneHandedMap.Bindings = append(mapXr.config.OneHandedMap.Bindings, config.TapBinding{TapCode: 3, Action: config.TapAction{ActionType: config.Keypress, Keypress: 32}})
	//mapXr.config.OneHandedMap.Bindings = append(mapXr.config.OneHandedMap.Bindings, config.TapBinding{TapCode: 4, Action: config.TapAction{ActionType: config.Keypress, Keypress: 33}})
	//mapXr.config.OneHandedMap.Bindings = append(mapXr.config.OneHandedMap.Bindings, config.TapBinding{TapCode: 4, Keypress: 3})

	//fmt.Println("Hello, World!")

	//config.SaveConfig(*mapXr.config)
	//return
	//initTapMap()

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	mapXr.kb = kb

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	tm := tapgosdk.NewTapManager()
	onTapped := onTapped{&mapXr}
	events.TappedData.Register(&onTapped)
	tm.Start()

	select {}
}
