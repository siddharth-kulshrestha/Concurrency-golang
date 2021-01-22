# Concurrency-golang


## Why is concurrency Hard ? 

### Race Conditions


### Atomicity


### Memory Access Synchronization


### Deadlocks

A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.

There are few conditions that must be present for deadlock to arise, also known as COFFMAN CONDITIONS.

- Mutual exclusion: A concurrent process holds exclusive rights to a resource at any one time.
- Wait for condition: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
- No Preemption: A resource is held by a concurrent process and can be released only by that process, so it fulfills that condition.
- Circular Wait: A concurrent process P1, must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.

If we ensure that at least one of the above conditions is not true, we can prevent deadlocks from occuring.


### Livelocks

Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.



### Starvation


