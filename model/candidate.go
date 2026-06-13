package model

const NMAXCANDIDATE = 100

type CandidateData struct {
	CandidateNumber int
	Name            string
	VisionMission   [999]string
	VMLength        int
}

type Candidates struct {
	Data   [NMAXCANDIDATE]CandidateData
	Length int
}

var Candidate Candidates

/*
 * GenerateCandidatesData creates the initial candidate dataset.
 * Purpose: Provide default candidate records when the app starts.
 * Flow: Build a Candidates struct with 10 sample entries -> return it.
 */
func GenerateCandidatesData() Candidates {
	var data Candidates = Candidates{
		Data: [NMAXCANDIDATE]CandidateData{
			{CandidateNumber: 1, Name: "Budi", VisionMission: [999]string{"Membangun", "ekonomi", "Indonesia", "lebih", "kuat"}, VMLength: 5},
			{CandidateNumber: 2, Name: "Siti", VisionMission: [999]string{"Meningkatkan", "kualitas", "pendidikan", "nasional"}, VMLength: 4},
			{CandidateNumber: 3, Name: "Ahmad", VisionMission: [999]string{"Menciptakan", "lapangan", "kerja", "seluas-luasnya"}, VMLength: 4},
			{CandidateNumber: 4, Name: "Dewi", VisionMission: [999]string{"Mewujudkan", "Indonesia", "yang", "adil", "dan", "makmur"}, VMLength: 6},
			{CandidateNumber: 5, Name: "Rizky", VisionMission: [999]string{"Mempercepat", "pembangunan", "infrastruktur", "nasional"}, VMLength: 4},
			{CandidateNumber: 6, Name: "Anisa", VisionMission: [999]string{"Meningkatkan", "kesehatan", "masyarakat", "Indonesia"}, VMLength: 4},
			{CandidateNumber: 7, Name: "Fajar", VisionMission: [999]string{"Menguatkan", "daya", "saing", "bangsa", "di", "tingkat", "global"}, VMLength: 7},
			{CandidateNumber: 8, Name: "Putri", VisionMission: [999]string{"Menciptakan", "pemerintahan", "yang", "bersih", "dan", "transparan"}, VMLength: 6},
			{CandidateNumber: 9, Name: "Hendra", VisionMission: [999]string{"Memberdayakan", "UMKM", "dan", "ekonomi", "rakyat"}, VMLength: 5},
			{CandidateNumber: 10, Name: "Lestari", VisionMission: [999]string{"Menjaga", "lingkungan", "hidup", "dan", "kelestarian", "alam"}, VMLength: 6},
		},
		Length: 10,
	}
	return data
}
