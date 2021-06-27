package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type BattleStage string

const (
	AnyStage        BattleStage = "any"
	Reef            BattleStage = "0"
	Musselforge     BattleStage = "1"
	Mainstage       BattleStage = "2"
	Shipyard        BattleStage = "3"
	Inkblot         BattleStage = "4"
	Humpback        BattleStage = "5"
	Manta           BattleStage = "6"
	Port            BattleStage = "7"
	Moray           BattleStage = "8"
	Canal           BattleStage = "9"
	Kelp            BattleStage = "10"
	Skatepark       BattleStage = "11"
	Shellendorf     BattleStage = "12"
	Mako            BattleStage = "13"
	Walleye         BattleStage = "14"
	Mall            BattleStage = "15"
	Camp            BattleStage = "16"
	Pit             BattleStage = "17"
	Arena           BattleStage = "18"
	Hotel           BattleStage = "19"
	World           BattleStage = "20"
	Games           BattleStage = "21"
	Skipper         BattleStage = "22"
	Windmill        BattleStage = "100"
	Wayslide        BattleStage = "101"
	Secret          BattleStage = "102"
	Goosponge       BattleStage = "103"
	Cannon          BattleStage = "105"
	Glass           BattleStage = "106"
	Fancy           BattleStage = "107"
	Grapplink       BattleStage = "108"
	Zappy           BattleStage = "109"
	Bunker          BattleStage = "110"
	Balance         BattleStage = "111"
	Switches        BattleStage = "112"
	Valley          BattleStage = "113"
	Twins           BattleStage = "114"
	Chillin         BattleStage = "115"
	Gusher          BattleStage = "116"
	Maze            BattleStage = "117"
	Flooders        BattleStage = "118"
	SplatInOurZones BattleStage = "119"
	Spreading       BattleStage = "120"
	Bridge          BattleStage = "121"
	Chronicles      BattleStage = "122"
	Furler          BattleStage = "123"
	Diaries         BattleStage = "124"
	Shifty          BattleStage = "9999"
)

func (bs *BattleStage) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type BS BattleStage
	r := (*BS)(bs)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *bs {
	case AnyStage, Reef, Musselforge, Mainstage, Shipyard, Inkblot, Humpback, Manta, Port, Moray, Canal, Kelp,
		Skatepark, Shellendorf, Mako, Walleye, Mall, Camp, Pit, Arena, Hotel, World, Games, Skipper, Windmill, Wayslide,
		Secret, Goosponge, Cannon, Glass, Fancy, Grapplink, Zappy, Bunker, Balance, Switches, Valley, Twins, Chillin,
		Gusher, Maze, Flooders, SplatInOurZones, Spreading, Bridge, Chronicles, Furler, Diaries, Shifty:
		return nil
	}
	return errors.New("Invalid BattleStage. Got: " + fmt.Sprint(*bs))
}

func (bs BattleStage) GetDisplay(printer *message.Printer) string {
	switch bs {
	case AnyStage:
		return printer.Sprintf("Any Stage")
	case Reef:
		return printer.Sprintf("The Reef")
	case Musselforge:
		return printer.Sprintf("Musselforge Fitness")
	case Mainstage:
		return printer.Sprintf("Starfish Mainstage")
	case Shipyard:
		return printer.Sprintf("Sturgeon Shipyard")
	case Inkblot:
		return printer.Sprintf("Inkblot Art Academy")
	case Humpback:
		return printer.Sprintf("Humpback Pump Track")
	case Manta:
		return printer.Sprintf("Manta Maria")
	case Port:
		return printer.Sprintf("Port Mackerel")
	case Moray:
		return printer.Sprintf("Moray Towers")
	case Canal:
		return printer.Sprintf("Snapper Canal")
	case Kelp:
		return printer.Sprintf("Kelp Dome")
	case Skatepark:
		return printer.Sprintf("Blackbelly Skatepark")
	case Shellendorf:
		return printer.Sprintf("Shellendorf Institute")
	case Mako:
		return printer.Sprintf("MakoMart")
	case Walleye:
		return printer.Sprintf("Walleye Warehouse")
	case Mall:
		return printer.Sprintf("Arowana Mall")
	case Camp:
		return printer.Sprintf("Camp Triggerfish")
	case Pit:
		return printer.Sprintf("Piranha Pit")
	case Arena:
		return printer.Sprintf("Goby Arena")
	case Hotel:
		return printer.Sprintf("New Albacore Hotel")
	case World:
		return printer.Sprintf("Wahoo World")
	case Games:
		return printer.Sprintf("Ancho-V Games")
	case Skipper:
		return printer.Sprintf("Skipper Pavillion")
	case Windmill:
		return printer.Sprintf("Windmill House on the Pearlie")
	case Wayslide:
		return printer.Sprintf("Wayslide Cool")
	case Secret:
		return printer.Sprintf("The Secret of S.P.L.A.T.")
	case Goosponge:
		return printer.Sprintf("Goosponge")
	case Cannon:
		return printer.Sprintf("Cannon Fire Pearl")
	case Glass:
		return printer.Sprintf("Zone of Glass")
	case Fancy:
		return printer.Sprintf("Fancy Spew")
	case Grapplink:
		return printer.Sprintf("Grapplink Girl")
	case Zappy:
		return printer.Sprintf("Zappy Longshocking")
	case Bunker:
		return printer.Sprintf("The Bunker Games")
	case Balance:
		return printer.Sprintf("A Swiftly Tilting Balance")
	case Switches:
		return printer.Sprintf("The Switches")
	case Valley:
		return printer.Sprintf("Sweet Valley Tentacles")
	case Twins:
		return printer.Sprintf("The Bouncey Twins")
	case Chillin:
		return printer.Sprintf("Railway Chillin'")
	case Gusher:
		return printer.Sprintf("Gusher Towns")
	case Maze:
		return printer.Sprintf("The Maze Dasher")
	case Flooders:
		return printer.Sprintf("Flooders in the Attic")
	case SplatInOurZones:
		return printer.Sprintf("The Splat in Our Zones")
	case Spreading:
		return printer.Sprintf("The Ink is Spreading")
	case Bridge:
		return printer.Sprintf("Bridge to Tentaswitchia")
	case Chronicles:
		return printer.Sprintf("The Chronicles of Rolonium")
	case Furler:
		return printer.Sprintf("Furler in the Ashes")
	case Diaries:
		return printer.Sprintf("MC.Princess Diaries")
	case Shifty:
		return printer.Sprintf("Shifty Station")
	}
	return ""
}

func GetStages() []BattleStage {
	return []BattleStage{
		AnyStage, Reef, Musselforge, Mainstage, Shipyard, Inkblot, Humpback, Manta, Port, Moray, Canal, Kelp,
		Skatepark, Shellendorf, Mako, Walleye, Mall, Camp, Pit, Arena, Hotel, World, Games, Skipper, Windmill, Wayslide,
		Secret, Goosponge, Cannon, Glass, Fancy, Grapplink, Zappy, Bunker, Balance, Switches, Valley, Twins, Chillin,
		Gusher, Maze, Flooders, SplatInOurZones, Spreading, Bridge, Chronicles, Furler, Diaries, Shifty,
	}
}
