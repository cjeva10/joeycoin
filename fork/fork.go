package fork

/*
module for handling fork choice algorithm

Given PoW algorithm, if we see a chain N blocks ahead, the probability of a chain
reorg is 0.49^N (i.e. the max probability that the longest chain is produced by a minority of hash power)

after a 10 block difference, that means the probability that the best chain being 
that far ahead is 10^-5, which is a reasonable pruning heuristic
i.e. we consider forks that are within 10 blocks 
*/

/*
We keep working on our chain unless we hear about another chain that is longer than ours.

Procedure:

Peer sends us block N+1 from propagate_block RPC call.

We verify N+1 is valid
    - either it is valid on our chain
    - or it's predecessor is different from our current tip
        - in this case decrementally request block get_block_by_number until:
            1) we have a match (in which case we need to reorg)
            2) a block is invalid (reject the fork)
    If we rejected the reorg, continue working on our chain until otherwise
    Else reorg our chain and rebuild state from the peer's blocks

    When we reorg, we need to update the tx pool because some of the transactions
    we had in our chain might now not be included, and others that we thought
    were pending might be executed

    To handle this, we search for tx's that are 
    1) in our chain but not in the reorg
    2) or in the reorg but not in our chain

    For 1)
        for txs in affected blocks on our chain:
            search for hash in reorg chain
            if not included:
                move this tx back to tx pool
            else:
                do nothing

    For 2)
        check tx hashes of reorg:
            if in our tx pool:      
                remove from pool
            else: 
                do nothing
*/
