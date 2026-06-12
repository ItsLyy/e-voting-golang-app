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
			{CandidateNumber: 1, Name: "Budi", VisionMission: "MembangunekonomiIndonesialebihkuat"},
			{CandidateNumber: 2, Name: "Siti", VisionMission: "Meningkatkankualitaspendidikannasional"},
			{CandidateNumber: 3, Name: "Ahmad", VisionMission: "Menciptakanlapangankerjaseluasluasnya"},
			{CandidateNumber: 4, Name: "Dewi", VisionMission: "MewujudkanIndonesiayangadildanmakmur"},
			{CandidateNumber: 5, Name: "Rizky", VisionMission: "Mempercepatpembangunaninfrastrukturnasional"},
			{CandidateNumber: 6, Name: "Anisa", VisionMission: "MeningkatkankesehatanmasyarakatIndonesia"},
			{CandidateNumber: 7, Name: "Fajar", VisionMission: "MenguatkandayasaingbangsaditingkatglobalL"},
			{CandidateNumber: 8, Name: "Putri", VisionMission: "Menciptakanpemerintahanyangbersihdantransparan"},
			{CandidateNumber: 9, Name: "Hendra", VisionMission: "MemberdayakanUMKMdanekonomirakyat"},
			{CandidateNumber: 10, Name: "Lestari", VisionMission: "Menjagalingkunganhidupdankelestarianalam"},
		},
		Length: 10,
	}
	return data
}
