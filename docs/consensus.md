# Abstract Consensus
This is the consensus model for the Abstract Network. We take inspiration
from XRPL's validator list and Proof-of-Stake to create our own consensus
model that we call ACE (Abstract-Consensus-Engine). The ACE requires
validators to stake tokens in order to become a validator. The nodes
are then reward a portion of the gas fee as their role in being a validator.
If a node wishes to become a validator it must first stake 0.1% of the total XAB
in circulation (100000000 XAB). There is a maximum amount of 100 validators,
meaning only a total of 10% of the total circulating supply will ever be staked.
Once a node becomes a validator they join a state of contract with the network,
vowing to be a validator for at least 1000000 blocks (The average block takes 3 seconds
to validate taking a total of 3000000 seconds). If the node exits as a validator before
the end of contract, the node will loose 10% of the value staked. The tokens lost will
be distributed to the other validators. A node that wants to become a validator must
be running an AVM that is as preformant as AVMx86 or better. Once a validator's
contract runs out they are placed in limbo, in which they still validate transactions
but once a new validator is found they are replaced. If a node wants to be first in
priority to become a validator they must outbid the highest bidder, to place a bid
for a validator spot it will cost 100x the normal transaction fee. Nodes who can't
afford the staking amount can take loans from users of the token without imposing
risk to the loaners. The people who give the node tokens will receive a percentage of
the gas fees accured if the node is elected, if the node is not elected then all funds
will be returned. This brings decentralization to the network and avoids the fears of
decentralization brought by Proof-of-Stake and validators. The validator nodes will be
in-charge of choosing which nodes will run which programs and which transactions are
valid or not. Any node on the network can challenge the validity of a transaction or a
block, but it will cost the total amount of gas times 100 used in the block/tx to 
challenge a block or transaction. This allows for a second layer of security but avoids
the spamming of second-validations on the network. Once a block has been challenged and
proven it then won't be able to be challenged again. Transactions cannot be sent to
validators directly, instead the use of a gateway will be needed. Although creating a
gateway is permissionless, although it does cost an amount of 10 XAB to create one.
A gateway cannot be deleted, a gateway is a way to allow a node to submit transactions
to the validators but in a method in which the spam of transactions is avoided.

For a block to be validated on the ACE we only need a 100% verification. This means
that it would require all 100 signatures from the validators.

Since we are using validators for the consensus, we can queue a list of blocks to be
validated. We can delegate the block creation process to certain validators to automatically
distribute the load. Theoretically we can queue up thousands of blocks to be validated. If the
blocks that are queued happen to be lost then no loss would be incurred since the blocks were
never added.

