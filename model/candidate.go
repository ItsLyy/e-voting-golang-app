package model

const NMAXCANDIDATE = 100

type CandidateData struct {
	CandidateNumber     int
	Name, VisionMission string
}

type Candidates struct {
	Data   [NMAXCANDIDATE]CandidateData
	Length int
}

var Candidate Candidates

func GenerateCandidatesData() Candidates {
	var data Candidates = Candidates{
		Data: [NMAXCANDIDATE]CandidateData{
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
		},
		Length: 10,
	}
	return data
}
