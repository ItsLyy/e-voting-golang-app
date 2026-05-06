package model

const NMAXCANDIDATE = 10

type Canditate struct {
	CandidateNumber int
	Name, VisionMission string
}

type Canditates [NMAXCANDIDATE]Canditate

func GenerateCandidatesData () Canditates {
	var data Canditates = Canditates{
		{CandidateNumber: 1001, Name: "Budi", VisionMission: "MembangunekonomiIndonesialebihkuat"},
		{CandidateNumber: 1002, Name: "Siti", VisionMission: "Meningkatkankualitaspendidikannasional"},
		{CandidateNumber: 1003, Name: "Ahmad", VisionMission: "Menciptakanlapangankerjaseluasluasnya"},
		{CandidateNumber: 1004, Name: "Dewi", VisionMission: "MewujudkanIndonesiayangadildanmakmur"},
		{CandidateNumber: 1005, Name: "Rizky", VisionMission: "Mempercepatpembangunaninfrastrukturnasional"},
		{CandidateNumber: 1006, Name: "Anisa", VisionMission: "MeningkatkankesehatanmasyarakatIndonesia"},
		{CandidateNumber: 1007, Name: "Fajar", VisionMission: "MenguatkandayasaingbangsaditingkatglobalL"},
		{CandidateNumber: 1008, Name: "Putri", VisionMission: "Menciptakanpemerintahanyangbersihdantransparan"},
		{CandidateNumber: 1009, Name: "Hendra", VisionMission: "MemberdayakanUMKMdanekonomirakyat"},
		{CandidateNumber: 1010, Name: "Lestari", VisionMission: "Menjagalingkunganhidupdankelestarianalam"},
	}
	return data
}