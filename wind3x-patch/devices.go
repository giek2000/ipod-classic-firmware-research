package devices

type Kind string

const (
	Nano3      Kind = "n3g"
	Nano4      Kind = "n4g"
	Classic6G  Kind = "classic6g"
	Classic6GA Kind = "classic6g-a"
	Classic6GB Kind = "classic6g-b"
	Classic6GC Kind = "classic6g-c"
	Nano5      Kind = "n5g"
	Nano6      Kind = "n6g"
	Nano7      Kind = "n7g"
	Nano7Late  Kind = "n7g-late"
)

type InterfaceKind string

const (
	DFU  InterfaceKind = "dfu"
	WTF  InterfaceKind = "wtf"
	Disk InterfaceKind = "diskmode"
)

type DFUProtoVersion int

const (
	// DFUProtoVersion1 is implemented by Nano3G and all iPod Classic revisions.
	DFUProtoVersion1 DFUProtoVersion = 1
	// DFUProtoVersion2 is implemented by Nano4G+.
	DFUProtoVersion2 DFUProtoVersion = 2
)

func (k Kind) String() string {
	switch k {
	case Nano3:
		return "Nano 3G"
	case Nano4:
		return "Nano 4G"
	case Classic6G:
		return "Classic 6G (Initial)"
	case Classic6GA:
		return "Classic 6G (Rev A)"
	case Classic6GB:
		return "Classic 6G (Rev B)"
	case Classic6GC:
		return "Classic 6G (Rev C)"
	case Nano5:
		return "Nano 5G"
	case Nano6:
		return "Nano 6G"
	case Nano7:
		return "Nano 7G"
	case Nano7Late:
		return "Nano 7G (Mid-2015)"
	}
	return "UNKNOWN"
}

func (k Kind) SoCCode() string {
	switch k {
	case Nano3, Classic6G, Classic6GA, Classic6GB, Classic6GC:
		return "8702"
	case Nano4:
		return "8720"
	case Nano5:
		return "8730"
	case Nano6:
		return "8723"
	case Nano7, Nano7Late:
		return "8740"
	}
	return "INVL"
}

func (k Kind) DFUVersion() DFUProtoVersion {
	switch k {
	case Nano3, Classic6G, Classic6GA, Classic6GB, Classic6GC:
		return DFUProtoVersion1
	default:
		return DFUProtoVersion2
	}
}

func (k Kind) Description() Description {
	for _, d := range Descriptions {
		if d.Kind == k {
			return d
		}
	}
	panic("unreachable")
}

type Description struct {
	VID             int16
	PIDs            map[InterfaceKind]int16
	UpdaterFamilyID int
	Kind            Kind
}

var Descriptions = []Description{
	// === iPod Nano ===
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1223,
			WTF:  0x1242,
			Disk: 0x1262,
		},
		UpdaterFamilyID: 26,
		Kind:            Nano3,
	},
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1225,
			WTF:  0x1243,
			Disk: 0x1263,
		},
		UpdaterFamilyID: 31,
		Kind:            Nano4,
	},
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1231,
			WTF:  0x1246,
			Disk: 0x1265,
		},
		UpdaterFamilyID: 34,
		Kind:            Nano5,
	},
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1232,
			WTF:  0x1248,
			Disk: 0x1266,
		},
		UpdaterFamilyID: 36,
		Kind:            Nano6,
	},
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1234,
			WTF:  0x1249,
			Disk: 0x1267,
		},
		UpdaterFamilyID: 37,
		Kind:            Nano7,
	},
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			WTF: 0x124a,
		},
		UpdaterFamilyID: 37,
		Kind:            Nano7Late,
	},
	// === iPod Classic (all revisions) ===
	// Initial (80GB/160GB, Sep 2007) - MB029, MB147, MB145, MB150
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1223,
			WTF:  0x1241,
			Disk: 0x1261,
		},
		UpdaterFamilyID: 35,
		Kind:            Classic6G,
	},
	// Rev A (120GB, Sep 2008) - MB562, MB565
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1223,
			WTF:  0x1245,
			Disk: 0x1261,
		},
		UpdaterFamilyID: 35,
		Kind:            Classic6GA,
	},
	// Rev B (160GB, Sep 2009) - MC293, MC297
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1223,
			WTF:  0x1247,
			Disk: 0x1261,
		},
		UpdaterFamilyID: 35,
		Kind:            Classic6GB,
	},
	// Rev C (160GB, Oct 2012) - MD717, MD718
	{
		VID: 0x05ac,
		PIDs: map[InterfaceKind]int16{
			DFU:  0x1250,
			WTF:  0x1250,
			Disk: 0x1261,
		},
		UpdaterFamilyID: 35,
		Kind:            Classic6GC,
	},
}
