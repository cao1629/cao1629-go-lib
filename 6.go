package main

import "time"

type Raft struct {
    state      State
    elected    chan struct{}
    toFollower chan struct{}
    done       chan struct{}
}

type State int

const (
    Leader    State = iota // 0
    Follower               // 1
    Candidate              // 2
)

func NewRaft() *Raft {
    rf := &Raft{
        state: Follower,
    }

    go func() {
        heartbeatTicker := time.NewTicker(time.Second)
        defer heartbeatTicker.Stop()

        electionTicker := time.NewTicker(time.Second * 5)
        defer electionTicker.Stop()

        for {
            select {
            case <-heartbeatTicker.C:
                // I'm a leader.
                go rf.InitiateAppendEntries()
            case <-electionTicker.C:
                // I just became a candidate.
                go rf.InitiateElection()
            case <-rf.elected:
                // I just became a leader.
                // 在 InitializeElection()中 rf.elected <- struct{}{}
                electionTicker.Stop()
                heartbeatTicker = time.NewTicker(time.Second)
            case <-rf.toFollower:
                // I just became a follower.
                // 在maybeUpdateTerm()中 rf.toFollower <- struct{}{}
                heartbeatTicker.Stop()
                electionTicker = time.NewTicker(time.Second * 5)
            case <-rf.done:
                return
            }
        }
    }()

    return rf
}

func (rf *Raft) InitiateElection() {}

func (rf *Raft) InitiateAppendEntries() {}
