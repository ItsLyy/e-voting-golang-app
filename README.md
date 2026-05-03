# 🗳️ E-Voting — Digital Ballot System

A data-driven application for managing elections digitally: registering candidates, recording votes, and displaying real-time election statistics.

---

## 📋 Description

E-Voting is an application for managing election data and counting votes quickly. The core data used includes candidate data, voter data, and vote tallies. The intended users are **election committees** and **organization members**.

---

## ✨ Features

| #   | Feature              | Description                                                                            |
| --- | -------------------- | -------------------------------------------------------------------------------------- |
| a   | Candidate Management | Add, edit, and delete candidate profile data                                           |
| b   | Vote Recording       | Record candidate number, vision-mission, and accumulated incoming votes                |
| c   | Candidate Search     | Search by candidate number using **Sequential Search** and **Binary Search**           |
| d   | Data Sorting         | Sort by most votes or candidate number using **Selection Sort** and **Insertion Sort** |
| e   | Vote Statistics      | Display vote percentage per candidate and total votes cast                             |

---

## 🚀 Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/ItsLyy/e-voting-golang-app.git
   cd e-voting-golang-app
   ```

2. Run the program:

   ```bash
   go run main.go
   ```

---

## 🗂️ Data Structure

```
candidate {
    candidate_number : int
    name             : string
    vision_mission   : string
}

voter {
    voter_id                : int
    name                    : string
    chosen_candidate_number : int
}
```

---

## 🧠 Algorithms Used

### Search

- **Sequential Search** — linear traversal through all candidates, no sorting required
- **Binary Search** — fast lookup by candidate number on sorted data

### Sort

- **Selection Sort** — sorts candidates by highest vote count
- **Insertion Sort** — sorts candidates by candidate number

---

## 👥 Users

- **Election Committee** — manages candidate and voter data, monitors vote results
- **Organization Members** — cast a vote for their chosen candidate

---

## 📄 License

This project was created for academic purposes. See [LICENSE](https://www.notion.so/LICENSE) for more details.
