package electionday

import "fmt"

type ElectionResult struct {
	// Name of the candidate
	Name string
	// Number of votes the candidate had
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var p *int
	p = &initialVotes
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	} else {
		return *counter
	}
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var elec ElectionResult
	var p *ElectionResult
	elec = ElectionResult{Name: candidateName, Votes: votes}
	p = &elec
	return p
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	str := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return str
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] = results[candidate] - 1 //Map ja conta como pointeiro
}
