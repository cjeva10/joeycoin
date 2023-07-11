package txpool

/* 
* data structures for tracking the pending tx pool
* algorithm for constructing blocks from the pool
*/

/* 
for simplicity every transaction has a fixed 1% fee on top of the tx amount
Also we have a default block reward of 1 JOEY on each block 
(to incentivize mining even when there are no transactions)
therefore miners should include the transactions with the largest amounts first

simple algorithm for generating a valid block from the pool
maintain a sorted list of pending txs that are valid on current state

    init empty block
    set temp state = current state
    for each tx:
        if tx is valid on temp state:
            push tx onto block
            update temp state
        else: 
            go to next tx in list

    if reach end of pool, or reach block size limit:
        return block
    (possible to check invalid tx's to see if any of them are valid now)

    generate a proof of work for the block

    propagate block to other miners
    
    if our block is accepted on the chain, reorg the pending tx list
*/


/* rules for rejecting transactions from entering the pool

    Sometimes we can keep a transaction in the pool even if it is currently invalid
    For example, a balance that is too low.
    Maybe we haven't seen a block from somewhere on the network where the wallet's
    balance is increased. 
    We can also keep a table of nonce gapped transactions for the same reason.
    In this case we need to reorg these possible but currently invalid transactions
    after each new block is added to the chain, either by us or another miner

    Other transactions are always invalid 
    (we've already seen a nonce, amount > max total token supply)
*/

/* 
    Algorithm for reorging tx pool after a new block

    Check all valid tx's to see if any are now invalid after the latest block
    Check all invalid tx's to see if they are valid on current state
    if so, insert them by amount into the sorted valid list
*/
