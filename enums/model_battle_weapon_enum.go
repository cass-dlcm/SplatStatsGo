package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type BattleWeaponEnum string

const (
	AnyWeapon         BattleWeaponEnum = "any"
	Sploosh           BattleWeaponEnum = "0"
	NeoSploosh        BattleWeaponEnum = "1"
	Sploosh7          BattleWeaponEnum = "2"
	Jr                BattleWeaponEnum = "10"
	CustomJr          BattleWeaponEnum = "11"
	KensaJr           BattleWeaponEnum = "12"
	Splash            BattleWeaponEnum = "20"
	NeoSplash         BattleWeaponEnum = "21"
	AeroMg            BattleWeaponEnum = "30"
	AeroRg            BattleWeaponEnum = "31"
	AeroPg            BattleWeaponEnum = "32"
	Splattershot      BattleWeaponEnum = "40"
	TtekSplattershot  BattleWeaponEnum = "41"
	KensaSplattershot BattleWeaponEnum = "42"
	HeroShot          BattleWeaponEnum = "45"
	OctoShot          BattleWeaponEnum = "46"
	Point52Gal        BattleWeaponEnum = "50"
	Point52GalDeco    BattleWeaponEnum = "51"
	KensaPoint52Gal   BattleWeaponEnum = "52"
	Nzap85            BattleWeaponEnum = "60"
	Nzap89            BattleWeaponEnum = "61"
	Nzap83            BattleWeaponEnum = "62"
	Pro               BattleWeaponEnum = "70"
	ForgePro          BattleWeaponEnum = "71"
	KensaPro          BattleWeaponEnum = "72"
	Point96Gal        BattleWeaponEnum = "80"
	Point96GalDeco    BattleWeaponEnum = "81"
	Jet               BattleWeaponEnum = "90"
	CustomJet         BattleWeaponEnum = "91"
	Luna              BattleWeaponEnum = "200"
	LunaNeo           BattleWeaponEnum = "201"
	KensaLuna         BattleWeaponEnum = "202"
	Blaster           BattleWeaponEnum = "210"
	CustomBlaster     BattleWeaponEnum = "211"
	HeroBlaster       BattleWeaponEnum = "215"
	Range             BattleWeaponEnum = "220"
	CustomRange       BattleWeaponEnum = "221"
	GrimRange         BattleWeaponEnum = "222"
	Clash             BattleWeaponEnum = "230"
	ClashNeo          BattleWeaponEnum = "231"
	Rapid             BattleWeaponEnum = "240"
	RapidDeco         BattleWeaponEnum = "241"
	KensaRapid        BattleWeaponEnum = "242"
	RapidPro          BattleWeaponEnum = "250"
	RapidProDeco      BattleWeaponEnum = "251"
	L3                BattleWeaponEnum = "300"
	L3D               BattleWeaponEnum = "301"
	KensaL3           BattleWeaponEnum = "302"
	H3                BattleWeaponEnum = "310"
	H3D               BattleWeaponEnum = "311"
	CherryH3          BattleWeaponEnum = "312"
	Squeezer          BattleWeaponEnum = "400"
	FoilSqueezer      BattleWeaponEnum = "401"
	Carbon            BattleWeaponEnum = "1000"
	CarbonDeco        BattleWeaponEnum = "1001"
	roller            BattleWeaponEnum = "1010"
	KrakOnRoller      BattleWeaponEnum = "1011"
	KensaRoller       BattleWeaponEnum = "1012"
	HeroRoller        BattleWeaponEnum = "1015"
	dynamo            BattleWeaponEnum = "1020"
	GoldDynamo        BattleWeaponEnum = "1021"
	KensaDynamo       BattleWeaponEnum = "1022"
	flingza           BattleWeaponEnum = "1030"
	FoilFlingza       BattleWeaponEnum = "1031"
	inkbrush          BattleWeaponEnum = "1100"
	InkbrushNouveau   BattleWeaponEnum = "1101"
	PermanentInkbrush BattleWeaponEnum = "1102"
	octobrush         BattleWeaponEnum = "1110"
	octobrushNouveau  BattleWeaponEnum = "1111"
	KensaOctobrush    BattleWeaponEnum = "1112"
	herobrush         BattleWeaponEnum = "1115"
	squiffer          BattleWeaponEnum = "2000"
	NewSquiffer       BattleWeaponEnum = "2001"
	FreshSquiffer     BattleWeaponEnum = "2002"
	charger           BattleWeaponEnum = "2010"
	FirefinCharger    BattleWeaponEnum = "2011"
	KensaCharger      BattleWeaponEnum = "2012"
	HeroCharger       BattleWeaponEnum = "2015"
	scope             BattleWeaponEnum = "2020"
	FirefinScope      BattleWeaponEnum = "2021"
	KensaScope        BattleWeaponEnum = "2022"
	eliter            BattleWeaponEnum = "2030"
	CustomEliter      BattleWeaponEnum = "2031"
	EliterScope       BattleWeaponEnum = "2040"
	CustomEliterScope BattleWeaponEnum = "2041"
	bamboozler        BattleWeaponEnum = "2050"
	Bamboozler2       BattleWeaponEnum = "2051"
	Bamboozler3       BattleWeaponEnum = "2052"
	goo               BattleWeaponEnum = "2060"
	CustomGoo         BattleWeaponEnum = "2061"
	slosher           BattleWeaponEnum = "3000"
	SlosherDeco       BattleWeaponEnum = "3001"
	SodaSlosher       BattleWeaponEnum = "3002"
	HeroSlosher       BattleWeaponEnum = "3005"
	tri               BattleWeaponEnum = "3010"
	TriNouveau        BattleWeaponEnum = "3011"
	machine           BattleWeaponEnum = "3020"
	MachineNeo        BattleWeaponEnum = "3021"
	KensaMachine      BattleWeaponEnum = "3022"
	blob              BattleWeaponEnum = "3030"
	BlobDeco          BattleWeaponEnum = "3031"
	explosher         BattleWeaponEnum = "3040"
	CustomExplosher   BattleWeaponEnum = "3041"
	mini              BattleWeaponEnum = "4000"
	ZinkMini          BattleWeaponEnum = "4001"
	KensaMini         BattleWeaponEnum = "4002"
	heavy             BattleWeaponEnum = "4010"
	HeavyDeco         BattleWeaponEnum = "4011"
	HeavyRemix        BattleWeaponEnum = "4012"
	HeroSplatling     BattleWeaponEnum = "4015"
	hydra             BattleWeaponEnum = "4020"
	CustomHydra       BattleWeaponEnum = "4021"
	ballpoint         BattleWeaponEnum = "4030"
	BallpointNouveau  BattleWeaponEnum = "4031"
	Naut47            BattleWeaponEnum = "4040"
	Naut79            BattleWeaponEnum = "4041"
	dapple            BattleWeaponEnum = "5000"
	DappleNouveau     BattleWeaponEnum = "5001"
	ClearDapple       BattleWeaponEnum = "5002"
	dualies           BattleWeaponEnum = "5010"
	EnperryDualies    BattleWeaponEnum = "5011"
	KensaDualies      BattleWeaponEnum = "5012"
	HeroDualies       BattleWeaponEnum = "5015"
	glooga            BattleWeaponEnum = "5020"
	GloogaDeco        BattleWeaponEnum = "5021"
	KensaGlooga       BattleWeaponEnum = "5022"
	squelchers        BattleWeaponEnum = "5030"
	CustomSquelchers  BattleWeaponEnum = "5031"
	tetra             BattleWeaponEnum = "5040"
	LightTetra        BattleWeaponEnum = "5041"
	brella            BattleWeaponEnum = "6000"
	SorrellaBrella    BattleWeaponEnum = "6001"
	HeroBrella        BattleWeaponEnum = "6005"
	tenta             BattleWeaponEnum = "6010"
	TentaSorella      BattleWeaponEnum = "6011"
	TentaCamo         BattleWeaponEnum = "6012"
	Undercover        BattleWeaponEnum = "6020"
	UndercoverSorella BattleWeaponEnum = "6021"
	KensaUndercover   BattleWeaponEnum = "6022"
)

func (bwe *BattleWeaponEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type BWE BattleWeaponEnum
	r := (*BWE)(bwe)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *bwe {
	case Sploosh, NeoSploosh, Sploosh7, Jr, CustomJr, KensaJr, Splash, NeoSplash, AeroMg, AeroRg, AeroPg,
		Splattershot, TtekSplattershot, KensaSplattershot, HeroShot, OctoShot, Point52Gal, Point52GalDeco,
		KensaPoint52Gal, Nzap85, Nzap89, Nzap83, Pro, ForgePro, KensaPro, Point96Gal, Point96GalDeco,
		Jet, CustomJet, Luna, LunaNeo, KensaLuna, Blaster, CustomBlaster, HeroBlaster, Range, CustomRange,
		GrimRange, Clash, ClashNeo, Rapid, RapidDeco, KensaRapid, RapidPro, RapidProDeco, L3, L3D, KensaL3,
		H3, H3D, CherryH3, Squeezer, FoilSqueezer, Carbon, CarbonDeco, roller, KrakOnRoller, KensaRoller,
		HeroRoller, dynamo, GoldDynamo, KensaDynamo, flingza, FoilFlingza, inkbrush, InkbrushNouveau,
		PermanentInkbrush, octobrush, octobrushNouveau, KensaOctobrush, herobrush, squiffer, NewSquiffer,
		FreshSquiffer, charger, FirefinCharger, KensaCharger, HeroCharger, scope, FirefinScope, KensaScope,
		eliter, CustomEliter, EliterScope, CustomEliterScope, bamboozler, Bamboozler2, Bamboozler3, goo,
		CustomGoo, slosher, SlosherDeco, SodaSlosher, HeroSlosher, tri, TriNouveau, machine, MachineNeo,
		KensaMachine, blob, BlobDeco, explosher, CustomExplosher, mini, ZinkMini, KensaMini, heavy, HeavyDeco,
		HeavyRemix, HeroSplatling, hydra, CustomHydra, ballpoint, BallpointNouveau, Naut47, Naut79, dapple,
		DappleNouveau, ClearDapple, dualies, EnperryDualies, KensaDualies, HeroDualies, glooga, GloogaDeco,
		KensaGlooga, squelchers, CustomSquelchers, tetra, LightTetra, brella, SorrellaBrella, HeroBrella, tenta,
		TentaSorella, TentaCamo, Undercover, UndercoverSorella, KensaUndercover:
		return nil
	}
	return errors.New("Invalid BattleWeaponEnum. Got: " + fmt.Sprint(*bwe))
}

func (bwe BattleWeaponEnum) GetDisplay(printer *message.Printer) string {
	switch bwe {
	case AnyWeapon:
		return printer.Sprintf("Any Weapon")
	case Sploosh:
		return printer.Sprintf("Sploosh-o-matic")
	case NeoSploosh:
		return printer.Sprintf("Neo Sploosh-o-matic")
	case Sploosh7:
		return printer.Sprintf("Sploosh-o-matic 7")
	case Jr:
		return printer.Sprintf("Splattershot Jr.")
	case CustomJr:
		return printer.Sprintf("Custom Splattershot Jr.")
	case KensaJr:
		return printer.Sprintf("Kensa Splattershot Jr.")
	case Splash:
		return printer.Sprintf("Splash-o-matic")
	case NeoSplash:
		return printer.Sprintf("Neo Splash-o-matic")
	case AeroMg:
		return printer.Sprintf("Aerospray MG")
	case AeroRg:
		return printer.Sprintf("Aerospray RG")
	case AeroPg:
		return printer.Sprintf("Aerospray PG")
	case Splattershot:
		return printer.Sprintf("Splattershot")
	case TtekSplattershot:
		return printer.Sprintf("Tentatek Splattershot")
	case KensaSplattershot:
		return printer.Sprintf("Kensa Splattershot")
	case HeroShot:
		return printer.Sprintf("Hero Shot Replica")
	case OctoShot:
		return printer.Sprintf("Octo Shot Replica")
	case Point52Gal:
		return printer.Sprintf(".52 Gal")
	case Point52GalDeco:
		return printer.Sprintf(".52 Gal Deco")
	case KensaPoint52Gal:
		return printer.Sprintf("Kensa .52 Gal")
	case Nzap85:
		return printer.Sprintf("N-ZAP '85")
	case Nzap89:
		return printer.Sprintf("N-ZAP '89")
	case Nzap83:
		return printer.Sprintf("N-ZAP '83")
	case Pro:
		return printer.Sprintf("Splattershot Pro")
	case ForgePro:
		return printer.Sprintf("Forge Splattershot Pro")
	case KensaPro:
		return printer.Sprintf("Kensa Splattershot Pro")
	case Point96Gal:
		return printer.Sprintf(".96 Gal")
	case Point96GalDeco:
		return printer.Sprintf(".96 Gal Deco")
	case Jet:
		return printer.Sprintf("Jet Squelcher")
	case CustomJet:
		return printer.Sprintf("Custom Jet Squelcher")
	case Luna:
		return printer.Sprintf("Luna Blaster")
	case LunaNeo:
		return printer.Sprintf("Luna Blaster Neo")
	case KensaLuna:
		return printer.Sprintf("Kensa Luna Blaster")
	case Blaster:
		return printer.Sprintf("Blaster")
	case CustomBlaster:
		return printer.Sprintf("Custom Blaster")
	case HeroBlaster:
		return printer.Sprintf("Hero Blaster Replica")
	case Range:
		return printer.Sprintf("Range Blaster")
	case CustomRange:
		return printer.Sprintf("Custom Range Blaster")
	case GrimRange:
		return printer.Sprintf("Grim Range Blaster")
	case Clash:
		return printer.Sprintf("Clash Blaster")
	case ClashNeo:
		return printer.Sprintf("Clash Blaster Neo")
	case Rapid:
		return printer.Sprintf("Rapid Blaster")
	case RapidDeco:
		return printer.Sprintf("Rapid Blaster Deco")
	case KensaRapid:
		return printer.Sprintf("Kensa Rapid Blaster")
	case RapidPro:
		return printer.Sprintf("Rapid Blaster Pro")
	case RapidProDeco:
		return printer.Sprintf("Rapid Blaster Pro Deco")
	case L3:
		return printer.Sprintf("L-3 Nozzlenose")
	case L3D:
		return printer.Sprintf("L-3 Nozzlenose D")
	case KensaL3:
		return printer.Sprintf("Kensa L-3 Nozzlenose")
	case H3:
		return printer.Sprintf("H-3 Nozzlenose")
	case H3D:
		return printer.Sprintf("H-3 Nozzlenose D")
	case CherryH3:
		return printer.Sprintf("Cherry H-3 Nozzlenose")
	case Squeezer:
		return printer.Sprintf("Squeezer")
	case FoilSqueezer:
		return printer.Sprintf("Foil Squeezer")
	case Carbon:
		return printer.Sprintf("Carbon Roller")
	case CarbonDeco:
		return printer.Sprintf("Carbon Roller Deco")
	case roller:
		return printer.Sprintf("Splat Roller")
	case KrakOnRoller:
		return printer.Sprintf("Krak-On Splat Roller")
	case KensaRoller:
		return printer.Sprintf("Kensa Splat Roller")
	case HeroRoller:
		return printer.Sprintf("Hero Roller Replica")
	case dynamo:
		return printer.Sprintf("Dynamo Roller")
	case GoldDynamo:
		return printer.Sprintf("Gold Dynamo Roller")
	case KensaDynamo:
		return printer.Sprintf("Kensa Dynamo Roller")
	case flingza:
		return printer.Sprintf("Flingza Roller")
	case FoilFlingza:
		return printer.Sprintf("Foil Flingza Roller")
	case inkbrush:
		return printer.Sprintf("Inkbrush")
	case InkbrushNouveau:
		return printer.Sprintf("Inkbrush Nouveau")
	case PermanentInkbrush:
		return printer.Sprintf("Permanent Inkbrush")
	case octobrush:
		return printer.Sprintf("Octobrush")
	case octobrushNouveau:
		return printer.Sprintf("Octobrush Nouveau")
	case KensaOctobrush:
		return printer.Sprintf("Kensa Octobrush")
	case herobrush:
		return printer.Sprintf("Herobrush Replica")
	case squiffer:
		return printer.Sprintf("Classic Squiffer")
	case NewSquiffer:
		return printer.Sprintf("New Squiffer")
	case FreshSquiffer:
		return printer.Sprintf("Fresh Squiffer")
	case charger:
		return printer.Sprintf("Splat Charger")
	case FirefinCharger:
		return printer.Sprintf("Firefin Splat Charger")
	case KensaCharger:
		return printer.Sprintf("Kensa Charger")
	case HeroCharger:
		return printer.Sprintf("Hero Charger Replica")
	case scope:
		return printer.Sprintf("Splatterscope")
	case FirefinScope:
		return printer.Sprintf("Firefin Splatterscope")
	case KensaScope:
		return printer.Sprintf("Kensa Splatterscope")
	case eliter:
		return printer.Sprintf("E-liter 4K")
	case CustomEliter:
		return printer.Sprintf("Custom E-liter 4K")
	case EliterScope:
		return printer.Sprintf("E-liter 4K Scope")
	case CustomEliterScope:
		return printer.Sprintf("Custom E-liter 4k Scope")
	case bamboozler:
		return printer.Sprintf("Bamboozler 14 Mk I")
	case Bamboozler2:
		return printer.Sprintf("Bamboozler 14 Mk II")
	case Bamboozler3:
		return printer.Sprintf("Bamboozler 14 Mk III")
	case goo:
		return printer.Sprintf("Goo Tuber")
	case CustomGoo:
		return printer.Sprintf("Custom Goo Tuber")
	case slosher:
		return printer.Sprintf("Slosher")
	case SlosherDeco:
		return printer.Sprintf("Slosher Deco")
	case SodaSlosher:
		return printer.Sprintf("Soda Slosher")
	case HeroSlosher:
		return printer.Sprintf("Hero Slosher Replica")
	case tri:
		return printer.Sprintf("Tri-Slosher")
	case TriNouveau:
		return printer.Sprintf("Tri-Slosher Nouveau")
	case machine:
		return printer.Sprintf("Sloshing Machine")
	case MachineNeo:
		return printer.Sprintf("Sloshing Machine Neo")
	case KensaMachine:
		return printer.Sprintf("Kensa Sloshing Machine")
	case blob:
		return printer.Sprintf("Bloblobber")
	case BlobDeco:
		return printer.Sprintf("Bloblobber Deco")
	case explosher:
		return printer.Sprintf("Explosher")
	case CustomExplosher:
		return printer.Sprintf("Custom Explosher")
	case mini:
		return printer.Sprintf("Mini Splatling")
	case ZinkMini:
		return printer.Sprintf("Zink Mini Splatling")
	case KensaMini:
		return printer.Sprintf("Kensa Mini Splatling")
	case heavy:
		return printer.Sprintf("Heavy Splatling")
	case HeavyDeco:
		return printer.Sprintf("Heavy Splatling Deco")
	case HeavyRemix:
		return printer.Sprintf("Heavy Splatling Remix")
	case HeroSplatling:
		return printer.Sprintf("Hero Splatling Replica")
	case hydra:
		return printer.Sprintf("Hydra Splatling")
	case CustomHydra:
		return printer.Sprintf("Custom Hydra Splatling")
	case ballpoint:
		return printer.Sprintf("Ballpoint Splatling")
	case BallpointNouveau:
		return printer.Sprintf("Ballpoint Splatling Nouveau")
	case Naut47:
		return printer.Sprintf("Nautilus 47")
	case Naut79:
		return printer.Sprintf("Nautilus 79")
	case dapple:
		return printer.Sprintf("Dapple Dualies")
	case DappleNouveau:
		return printer.Sprintf("Dapple Dualies Nouveau")
	case ClearDapple:
		return printer.Sprintf("Clear Dapple Dualies")
	case dualies:
		return printer.Sprintf("Splat Dualies")
	case EnperryDualies:
		return printer.Sprintf("Enperry Splat Dualies")
	case KensaDualies:
		return printer.Sprintf("Kensa Splat Dualies")
	case HeroDualies:
		return printer.Sprintf("Hero Dualie Replicas")
	case glooga:
		return printer.Sprintf("Glooga Dualies")
	case GloogaDeco:
		return printer.Sprintf("Glooga Dualies Deco")
	case KensaGlooga:
		return printer.Sprintf("Kensa Glooga Dualies")
	case squelchers:
		return printer.Sprintf("Dualie Squelchers")
	case CustomSquelchers:
		return printer.Sprintf("Custom Dualie Squelchers")
	case tetra:
		return printer.Sprintf("Dark Tetra Dualies")
	case LightTetra:
		return printer.Sprintf("Light Tetra Dualies")
	case brella:
		return printer.Sprintf("Splat Brella")
	case SorrellaBrella:
		return printer.Sprintf("Sorella Brella")
	case HeroBrella:
		return printer.Sprintf("Hero Brella Replica")
	case tenta:
		return printer.Sprintf("Tenta Brella")
	case TentaSorella:
		return printer.Sprintf("Tenta Sorella Brella")
	case TentaCamo:
		return printer.Sprintf("Tenta Camo Brella")
	case Undercover:
		return printer.Sprintf("Undercover Brella")
	case UndercoverSorella:
		return printer.Sprintf("Undercover Sorella Brella")
	case KensaUndercover:
		return printer.Sprintf("Kensa Undercover Brella")
	}
	return ""
}

func GetBattleWeaponEnum() []BattleWeaponEnum {
	return []BattleWeaponEnum{
		AnyWeapon, Sploosh, NeoSploosh, Sploosh7, Jr, CustomJr, KensaJr, Splash, NeoSplash, AeroMg, AeroRg, AeroPg,
		Splattershot, TtekSplattershot, KensaSplattershot, HeroShot, OctoShot, Point52Gal, Point52GalDeco,
		KensaPoint52Gal, Nzap85, Nzap89, Nzap83, Pro, ForgePro, KensaPro, Point96Gal, Point96GalDeco,
		Jet, CustomJet, Luna, LunaNeo, KensaLuna, Blaster, CustomBlaster, HeroBlaster, Range, CustomRange,
		GrimRange, Clash, ClashNeo, Rapid, RapidDeco, KensaRapid, RapidPro, RapidProDeco, L3, L3D, KensaL3,
		H3, H3D, CherryH3, Squeezer, FoilSqueezer, Carbon, CarbonDeco, roller, KrakOnRoller, KensaRoller,
		HeroRoller, dynamo, GoldDynamo, KensaDynamo, flingza, FoilFlingza, inkbrush, InkbrushNouveau,
		PermanentInkbrush, octobrush, octobrushNouveau, KensaOctobrush, herobrush, squiffer, NewSquiffer,
		FreshSquiffer, charger, FirefinCharger, KensaCharger, HeroCharger, scope, FirefinScope, KensaScope,
		eliter, CustomEliter, EliterScope, CustomEliterScope, bamboozler, Bamboozler2, Bamboozler3, goo,
		CustomGoo, slosher, SlosherDeco, SodaSlosher, HeroSlosher, tri, TriNouveau, machine, MachineNeo,
		KensaMachine, blob, BlobDeco, explosher, CustomExplosher, mini, ZinkMini, KensaMini, heavy, HeavyDeco,
		HeavyRemix, HeroSplatling, hydra, CustomHydra, ballpoint, BallpointNouveau, Naut47, Naut79, dapple,
		DappleNouveau, ClearDapple, dualies, EnperryDualies, KensaDualies, HeroDualies, glooga, GloogaDeco,
		KensaGlooga, squelchers, CustomSquelchers, tetra, LightTetra, brella, SorrellaBrella, HeroBrella, tenta,
		TentaSorella, TentaCamo, Undercover, UndercoverSorella, KensaUndercover,
	}
}

type BattleStatinkWeaponEnum string

const (
	statInkSploosh           BattleStatinkWeaponEnum = "bold"
	statInkNeoSploosh        BattleStatinkWeaponEnum = "bold_neo"
	statInkSploosh7          BattleStatinkWeaponEnum = "bold_7"
	statInkJr                BattleStatinkWeaponEnum = "wakaba"
	statInkCustomJr          BattleStatinkWeaponEnum = "momiji"
	statInkKensaJr           BattleStatinkWeaponEnum = "ochiba"
	statInkSplash            BattleStatinkWeaponEnum = "sharp"
	statInkNeoSplash         BattleStatinkWeaponEnum = "sharp_neo"
	statInkAeroMG            BattleStatinkWeaponEnum = "promodeler_mg"
	statInkAeroRG            BattleStatinkWeaponEnum = "promodeler_rg"
	statInkAeroPG            BattleStatinkWeaponEnum = "promodeler_pg"
	statInkSplattershot      BattleStatinkWeaponEnum = "sshooter"
	statInkTtekSplattershot  BattleStatinkWeaponEnum = "sshooter_collabo"
	statInkKensaSplattershot BattleStatinkWeaponEnum = "sshooter_becchu"
	statInkHeroShot          BattleStatinkWeaponEnum = "heroshooter_replica"
	statInkOctoShot          BattleStatinkWeaponEnum = "octoshooter_replica"
	statInkPoint52Gal        BattleStatinkWeaponEnum = "52gal"
	statInkPoint52GalDeco    BattleStatinkWeaponEnum = "52gal_deco"
	statInkKensaPoint52Gal   BattleStatinkWeaponEnum = "52gal_becchu"
	statInkNZap85            BattleStatinkWeaponEnum = "nzap85"
	statInkNZap89            BattleStatinkWeaponEnum = "nzap89"
	statInkNZap83            BattleStatinkWeaponEnum = "nzap83"
	statInkPro               BattleStatinkWeaponEnum = "prime"
	statInkForgePro          BattleStatinkWeaponEnum = "prime_collabo"
	statInkKensaPro          BattleStatinkWeaponEnum = "prime_becchu"
	statInkPoint96Gal        BattleStatinkWeaponEnum = "96gal"
	statInkPoint96GalDeco    BattleStatinkWeaponEnum = "96gal_deco"
	statInkJet               BattleStatinkWeaponEnum = "jetsweeper"
	statInkCustomJet         BattleStatinkWeaponEnum = "jetsweeper_custom"
	statInkLuna              BattleStatinkWeaponEnum = "nova"
	statInkLunaNeo           BattleStatinkWeaponEnum = "nova_neo"
	statInkKensaLuna         BattleStatinkWeaponEnum = "nova_becchu"
	statInkBlaster           BattleStatinkWeaponEnum = "hotblaster"
	statInkCustomBlaster     BattleStatinkWeaponEnum = "hotblaster_custom"
	statInkHeroBlaster       BattleStatinkWeaponEnum = "heroblaster_replica"
	statInkRange             BattleStatinkWeaponEnum = "longblaster"
	statInkCustomRange       BattleStatinkWeaponEnum = "longblaster_custom"
	statInkGrimRange         BattleStatinkWeaponEnum = "longblaster_necro"
	statInkClash             BattleStatinkWeaponEnum = "clashblaster"
	statInkClashNeo          BattleStatinkWeaponEnum = "clashblaster_neo"
	statInkRapid             BattleStatinkWeaponEnum = "rapid"
	statInkRapidDeco         BattleStatinkWeaponEnum = "rapid_deco"
	statInkKensaRapid        BattleStatinkWeaponEnum = "rapid_becchu"
	statInkRapidPro          BattleStatinkWeaponEnum = "rapid_elite"
	statInkRapidProDeco      BattleStatinkWeaponEnum = "rapid_elite_deco"
	statInkL3                BattleStatinkWeaponEnum = "l3reelgun"
	statInkL3D               BattleStatinkWeaponEnum = "l3reelgun_d"
	statInkKensaL3           BattleStatinkWeaponEnum = "l3reelgun_becchu"
	statInkH3                BattleStatinkWeaponEnum = "h3reelgun"
	statInkH3D               BattleStatinkWeaponEnum = "h3reelgun_d"
	statInkCherryH3          BattleStatinkWeaponEnum = "h3reelgun_cherry"
	statInkSqueezer          BattleStatinkWeaponEnum = "bottlegeyser"
	statInkFoilSqueezer      BattleStatinkWeaponEnum = "bottlegeyser_foil"
	statInkCarbon            BattleStatinkWeaponEnum = "carbon"
	statInkCarbonDeco        BattleStatinkWeaponEnum = "carbon_deco"
	statInkRoller            BattleStatinkWeaponEnum = "splatroller"
	statInkKrakOnRoller      BattleStatinkWeaponEnum = "splatroller_collabo"
	statInkKensaRoller       BattleStatinkWeaponEnum = "splatroller_becchu"
	statInkHeroRoller        BattleStatinkWeaponEnum = "heroroller_replica"
	statInkDynamo            BattleStatinkWeaponEnum = "dynamo"
	statInkGoldDynamo        BattleStatinkWeaponEnum = "dynamo_tesla"
	statInkKensaDynamo       BattleStatinkWeaponEnum = "dynamo_becchu"
	statInkFlingza           BattleStatinkWeaponEnum = "variableroller"
	statInkFoilFlingza       BattleStatinkWeaponEnum = "variableroller_foil"
	statInkInkbrush          BattleStatinkWeaponEnum = "pablo"
	statInkInkbrushNouveau   BattleStatinkWeaponEnum = "pablo_hue"
	statInkPermanentInkbrush BattleStatinkWeaponEnum = "pablo_permanent"
	statInkOctobrush         BattleStatinkWeaponEnum = "hokusai"
	statInkOctobrushNoveau   BattleStatinkWeaponEnum = "hokusai_hue"
	statInkKensaOctobrush    BattleStatinkWeaponEnum = "hokusai_becchu"
	statInkHeroBrush         BattleStatinkWeaponEnum = "herobrush_replica"
	statInkSquiffer          BattleStatinkWeaponEnum = "squiclean_a"
	statInkNewSquiffer       BattleStatinkWeaponEnum = "squiclean_b"
	statInkFreshSquiffer     BattleStatinkWeaponEnum = "squiclean_g"
	statInkCharger           BattleStatinkWeaponEnum = "splatcharger"
	statInkFirefinCharger    BattleStatinkWeaponEnum = "splatcharger_collabo"
	statInkKensaCharger      BattleStatinkWeaponEnum = "splatcharger_becchu"
	statInkHeroCharger       BattleStatinkWeaponEnum = "herocharger_replica"
	statInkScope             BattleStatinkWeaponEnum = "splatscope"
	statInkFirefinScope      BattleStatinkWeaponEnum = "splatscope_collabo"
	statInkKensaScope        BattleStatinkWeaponEnum = "splatscope_becchu"
	statInkEliter            BattleStatinkWeaponEnum = "liter4k"
	statInkCustomEliter      BattleStatinkWeaponEnum = "liter4k_custom"
	statInkEliterScope       BattleStatinkWeaponEnum = "liter4k_scope"
	statInkCustomEliterScope BattleStatinkWeaponEnum = "liter4k_scope_custom"
	statInkBamboozler        BattleStatinkWeaponEnum = "bamboo14mk1"
	statInkBamboozler2       BattleStatinkWeaponEnum = "bamboo14mk2"
	statInkBamboozler3       BattleStatinkWeaponEnum = "bamboo14mk3"
	statInkGoo               BattleStatinkWeaponEnum = "soytuber"
	statInkCustomGoo         BattleStatinkWeaponEnum = "soytuber_custom"
	statInkSlosher           BattleStatinkWeaponEnum = "bucketslosher"
	statInkSlosherDeco       BattleStatinkWeaponEnum = "bucketslosher_deco"
	statInkSodaSlosher       BattleStatinkWeaponEnum = "bucketslosher_soda"
	statInkHeroSlosher       BattleStatinkWeaponEnum = "heroslosher_replica"
	statInkTri               BattleStatinkWeaponEnum = "hissen"
	statInkTriNouveau        BattleStatinkWeaponEnum = "hissen_hue"
	statInkMachine           BattleStatinkWeaponEnum = "screwslosher"
	statInkMachineNeo        BattleStatinkWeaponEnum = "screwslosher_neo"
	statInkKensaMachine      BattleStatinkWeaponEnum = "screwslosher_becchu"
	statInkBlob              BattleStatinkWeaponEnum = "furo"
	statInkBlobDeco          BattleStatinkWeaponEnum = "furo_deco"
	statInkExplosher         BattleStatinkWeaponEnum = "explosher"
	statInkCustomExplosher   BattleStatinkWeaponEnum = "explosher_custom"
	statInkMini              BattleStatinkWeaponEnum = "splatspinner"
	statInkZinkMini          BattleStatinkWeaponEnum = "splatspinner_collabo"
	statInkKensaMini         BattleStatinkWeaponEnum = "splatspinner_becchu"
	statInkHeavy             BattleStatinkWeaponEnum = "barrelspinner"
	statInkHeavyDeco         BattleStatinkWeaponEnum = "barrelspinner_deco"
	statInkHeavyRemix        BattleStatinkWeaponEnum = "barrelspinner_remix"
	statInkHeroSplatling     BattleStatinkWeaponEnum = "herospinner_replica"
	statInkHydra             BattleStatinkWeaponEnum = "hydra"
	statInkCustomHydra       BattleStatinkWeaponEnum = "hydra_custom"
	statInkBallpoint         BattleStatinkWeaponEnum = "kugelschreiber"
	statInkBallpointNouveau  BattleStatinkWeaponEnum = "kugelschreiber_hue"
	statInkNaut47            BattleStatinkWeaponEnum = "nautilus47"
	statInkNaut79            BattleStatinkWeaponEnum = "nautilus79"
	statInkDapple            BattleStatinkWeaponEnum = "sputtery"
	statInkDappleNouveau     BattleStatinkWeaponEnum = "sputtery_hue"
	statInkClearDapple       BattleStatinkWeaponEnum = "sputtery_clear"
	statInkDualies           BattleStatinkWeaponEnum = "maneuver"
	statInkEnperryDualies    BattleStatinkWeaponEnum = "maneuver_collabo"
	statInkKensaDualies      BattleStatinkWeaponEnum = "maneuver_becchu"
	statInkHeroDualies       BattleStatinkWeaponEnum = "heromaneuver_replica"
	statInkGlooga            BattleStatinkWeaponEnum = "kelvin525"
	statInkGloogaDeco        BattleStatinkWeaponEnum = "kelvin525_deco"
	statInkKensaGlooga       BattleStatinkWeaponEnum = "kelvin525_becchu"
	statInkSquelchers        BattleStatinkWeaponEnum = "dualsweeper"
	statInkCustomSquelchers  BattleStatinkWeaponEnum = "dualsweeper_custom"
	statInkTetra             BattleStatinkWeaponEnum = "quadhopper_black"
	statInkLightTetra        BattleStatinkWeaponEnum = "quadhopper_white"
	statInkBrella            BattleStatinkWeaponEnum = "parashelter"
	statInkSorrellaBrella    BattleStatinkWeaponEnum = "parashelter_sorella"
	statInkHeroBrella        BattleStatinkWeaponEnum = "heroshelter_replica"
	statInkTenta             BattleStatinkWeaponEnum = "campingshelter"
	statInkTentaSorella      BattleStatinkWeaponEnum = "campingshelter_sorella"
	statInkTentaCamo         BattleStatinkWeaponEnum = "campingshelter_camo"
	statInkUndercover        BattleStatinkWeaponEnum = "spygadget"
	statInkUndercoverSorella BattleStatinkWeaponEnum = "spygadget_sorella"
	statInkKensaUndercover   BattleStatinkWeaponEnum = "spygadget_becchu"
	statInkNoneWeapon        BattleStatinkWeaponEnum = ""
)

func (bswe *BattleStatinkWeaponEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type BSWE BattleStatinkWeaponEnum
	r := (*BSWE)(bswe)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *bswe {
	case statInkSploosh, statInkNeoSploosh, statInkSploosh7, statInkJr, statInkCustomJr, statInkKensaJr, statInkSplash,
		statInkNeoSplash, statInkAeroMG, statInkAeroRG, statInkAeroPG, statInkSplattershot, statInkTtekSplattershot,
		statInkKensaSplattershot, statInkHeroShot, statInkOctoShot, statInkPoint52Gal, statInkPoint52GalDeco,
		statInkKensaPoint52Gal, statInkNZap85, statInkNZap89, statInkNZap83, statInkPro, statInkForgePro,
		statInkKensaPro, statInkPoint96Gal, statInkPoint96GalDeco, statInkJet, statInkCustomJet, statInkLuna,
		statInkLunaNeo, statInkKensaLuna, statInkBlaster, statInkCustomBlaster, statInkHeroBlaster, statInkRange,
		statInkCustomRange, statInkGrimRange, statInkClash, statInkClashNeo, statInkRapid, statInkRapidDeco,
		statInkKensaRapid, statInkRapidPro, statInkRapidProDeco, statInkL3, statInkL3D, statInkKensaL3, statInkH3,
		statInkH3D, statInkCherryH3, statInkSqueezer, statInkFoilSqueezer, statInkCarbon, statInkCarbonDeco,
		statInkRoller, statInkKrakOnRoller, statInkKensaRoller, statInkHeroRoller, statInkDynamo, statInkGoldDynamo,
		statInkKensaDynamo, statInkFlingza, statInkFoilFlingza, statInkInkbrush, statInkInkbrushNouveau,
		statInkPermanentInkbrush, statInkOctobrush, statInkOctobrushNoveau, statInkKensaOctobrush, statInkHeroBrush,
		statInkSquiffer, statInkNewSquiffer, statInkFreshSquiffer, statInkCharger, statInkFirefinCharger,
		statInkKensaCharger, statInkHeroCharger, statInkScope, statInkFirefinScope, statInkKensaScope, statInkEliter,
		statInkCustomEliter, statInkEliterScope, statInkCustomEliterScope, statInkBamboozler, statInkBamboozler2,
		statInkBamboozler3, statInkGoo, statInkCustomGoo, statInkSlosher, statInkSlosherDeco, statInkSodaSlosher,
		statInkHeroSlosher, statInkTri, statInkTriNouveau, statInkMachine, statInkMachineNeo, statInkKensaMachine,
		statInkBlob, statInkBlobDeco, statInkExplosher, statInkCustomExplosher, statInkMini, statInkZinkMini,
		statInkKensaMini, statInkHeavy, statInkHeavyDeco, statInkHeavyRemix, statInkHeroSplatling, statInkHydra,
		statInkCustomHydra, statInkBallpoint, statInkBallpointNouveau, statInkNaut47, statInkNaut79, statInkDapple,
		statInkDappleNouveau, statInkClearDapple, statInkDualies, statInkEnperryDualies, statInkKensaDualies,
		statInkHeroDualies, statInkGlooga, statInkGloogaDeco, statInkKensaGlooga, statInkSquelchers,
		statInkCustomSquelchers, statInkTetra, statInkLightTetra, statInkBrella, statInkSorrellaBrella,
		statInkHeroBrella, statInkTenta, statInkTentaSorella, statInkTentaCamo, statInkUndercover,
		statInkUndercoverSorella, statInkKensaUndercover, statInkNoneWeapon:
		return nil
	}
	return errors.New("Invalid BattleStatinkWeaponEnum. Got: " + fmt.Sprint(*bswe))
}
