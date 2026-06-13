# E-Voting — Digital Ballot System

A console-based e-voting application built with Go. It manages candidates and voters, records votes, tallies election results, and demonstrates classic **searching** and **sorting** algorithms in a real use case.

**Member Group:**
- **Irly Fizaharis - 103012540025**
    - Feature: Searching, Min and Max, Selection Sorting for Candidate, Insertition Sorting for Voter, CRUD Voter, Intregation between Voter Controller and Voter View.
    - Menu CLI: Home, Manage Voter, Setting, and Election
    - Role: Lead the flow of development, Review the code from Pull Request
- **Muhamad Hammam Misbah - 103012500383**
    - Feature: Insertion Sorting for Candidate, Selection Sorting for Voter, CRUD Candidate, Intregation between Voter Controller and Voter View
    - Menu CLI: Manage Candidate and Casting Voter
    - Role: Review the code before pull request, Testing the app to make sure run smoothly
---

## Description

E-Voting is a data-driven application for managing election data and counting votes. It stores candidate profiles, voter records, and vote choices in memory, then processes them using configurable algorithms.

**Target users:**
- **Election committee** — manage candidates, voters, and view election results
- **Organization members** — cast a vote for their chosen candidate

---

## Features

| # | Feature | Description |
|---|---------|-------------|
| 1 | Candidate Management | View, add, delete, search, and update candidate data (name & vision mission) |
| 2 | Voter Management | View, add, delete, search, and update voter data |
| 3 | Voting | Assign a candidate choice to a voter |
| 4 | Election Management | View vote results (percentage & total) and reset the election |
| 5 | Settings | Switch sorting algorithm (selection / insertion) and search algorithm (sequential / binary) |

---

## Getting Started

### Requirements

- Go 1.25.5 or later

### Run the application

```bash
git clone https://github.com/ItsLyy/e-voting-golang-app.git
cd e-voting-golang-app
go run main.go
```

### Main menu

```
1. Candidate Management
2. Voter Management
3. Voting
4. Election Management
s. Settings
q. Quit App
```

On startup, the app loads **10 sample candidates** and **20 sample voters**, then opens the home menu with these default settings:

| Setting | Default |
|---------|---------|
| Sort algorithm | Selection sort |
| Search algorithm | Binary search |

---

## Project Structure

The app follows an **MVC (Model–View–Controller)** pattern:

```
e-voting-golang-app/
├── main.go                 # Entry point — loads data and starts the app
├── model/                  # Data structures and sample data
│   ├── candidate.go
│   ├── voter.go
│   ├── election.go
│   └── main.go             # DataSetting struct
├── controller/             # Business logic and algorithms
│   ├── candidate.go        # CRUD + optimized search for candidates
│   ├── voter.go            # CRUD, voting, + optimized search for voters
│   ├── election.go         # Vote tallying and percentage calculation
│   ├── searching.go        # Sequential & binary search
│   ├── sorting.go          # Selection & insertion sort
│   ├── min.go              # Min index & less-than comparisons
│   └── max.go              # Max index & greater-than comparisons
└── view/                   # Console UI and menu navigation
    ├── home.go
    ├── candidate.go
    ├── voter.go
    ├── election.go
    └── setting.go
```

---

## Data Structures

### Candidate

```go
CandidateData {
    CandidateNumber : int
    Name            : string
    VisionMission   : [999]string   // word-by-word input, ended with "STOP"
    VMLength        : int
}

Candidates {
    Data   : [100]CandidateData
    Length : int
}
```

### Voter

```go
VoterData {
    VoterId          : int
    Name             : string
    CandidateNumber  : int   // 0 = has not voted yet
}

Voters {
    Data   : [999]VoterData
    Length : int
}
```

### Election

```go
ElectionData {
    CandidateIndex : int
    TotalVote      : int
}

Elections {
    Data : [10]ElectionData
}
```

### Settings

```go
DataSetting {
    SortBy        : string   // "id", "name", "candidate"
    SortOrder     : string   // "asc", "desc"
    SortSetting   : string   // "selection", "insertion"
    SearchSetting : string   // "sequential", "binary"
}
```

---

## Algorithms

### Searching

| Algorithm | File | Used when |
|-----------|------|-----------|
| **Sequential Search** | `controller/searching.go` | Data is unsorted, or binary search is disabled in settings |
| **Binary Search** | `controller/searching.go` | Data is sorted by the searched field and binary search is enabled |

Search functions:

| Function | Searches by |
|----------|-------------|
| `SearchSequentialCandidateByName` | Candidate name (linear) |
| `SearchSequentialCandidateByCandidateNumber` | Candidate number (linear) |
| `SearchSequentialVoterByName` | Voter name (linear) |
| `SearchSequentialVoterById` | Voter ID (linear) |
| `SearchBinaryCandidateByName` | Candidate name (binary, requires sorted data) |
| `SearchBinaryCandidateByCandidateNumber` | Candidate number (binary, requires sorted data) |
| `SearchBinaryVoterByName` | Voter name (binary, requires sorted data) |
| `SearchBinaryVoterById` | Voter ID (binary, requires sorted data) |

Optimized search helpers automatically pick the right algorithm based on current settings:
- `optimizationCandidateSearchIndex` / `optimizationCandidateNameSearchIndex`
- `optimizationVoterIdSearchIndex` / `optimizationVoterNameSearchIndex`

### Sorting

| Algorithm | File | Used for |
|-----------|------|----------|
| **Selection Sort** | `controller/sorting.go` | Default sort for candidates and voters |
| **Insertion Sort** | `controller/sorting.go` | Optional sort for voters (configurable in settings) |

Sort fields:
- **Candidates** — by ID or name
- **Voters** — by ID, name, or chosen candidate number

Helper comparison functions in `controller/min.go` and `controller/max.go` support both sorting and sorted insertion when adding new records.

---

## Application Flow

```
main
 └── Home
      ├── Candidate Management
      │    ├── View All (search / sort)
      │    ├── Add Candidate
      │    └── Delete Candidate
      ├── Voter Management
      │    ├── View All (search / sort)
      │    ├── Add Voter
      │    └── Delete Voter
      ├── Voting
      │    └── Select voter → select candidate → record vote
      ├── Election Management
      │    ├── View Results (percentage & vote count)
      │    └── Reset Election
      └── Settings
           ├── Candidate / Voter Sorting (selection / insertion)
           └── Candidate / Voter Searching (sequential / binary)
```

### Election result flow

1. `InputElection` — tally votes by matching each voter's `CandidateNumber` to a candidate
2. `SumVotes` — count total votes across all candidates
3. `CalculatePercentage` — compute each candidate's vote share as a percentage
4. `displayTableElectionResults` — print the results table

### Reset election flow

1. `ResetElection` — set every voter's `CandidateNumber` back to `0`
2. All recorded votes are cleared; candidates and voter records remain

---

## Controller Functions

### Candidate (`controller/candidate.go`)

| Function | Purpose |
|----------|---------|
| `CreateCandidate` | Add a new candidate with validation and sorted insertion |
| `DeleteCandidate` | Remove a candidate by number |
| `UpdateCandidate` | Update an existing candidate's data |
| `ShowCandidateById` | Retrieve a candidate by number |
| `ShowCandidateByName` | Retrieve a candidate by name |

### Voter (`controller/voter.go`)

| Function | Purpose |
|----------|---------|
| `CreateVoter` | Add a new voter with validation and sorted insertion |
| `DeleteVoter` | Remove a voter by ID |
| `UpdateVoter` | Update an existing voter's data |
| `CastingVoteVoter` | Record a voter's candidate choice |
| `ShowVoterById` | Retrieve a voter by ID |
| `ShowVoterByName` | Retrieve a voter by name |

### Election (`controller/election.go`)

| Function | Purpose |
|----------|---------|
| `InputElection` | Tally votes from voter data into election results |
| `SumVotes` | Count total votes across all candidates |
| `CalculatePercentage` | Compute vote share as a percentage |
| `ResetElection` | Clear all voter choices |

---

## Code Documentation

Every function in the codebase is documented using block comments in this format:

```go
/*
 * FunctionName — short description of what it is.
 * Purpose: Why this function exists.
 * Flow: Step-by-step logic of what it does.
 */
```

Browse the source files for full documentation on each function's behavior and flow.

---

## License

This project was created for academic purposes. See [LICENSE](LICENSE) for more details.
