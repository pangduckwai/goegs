package node

//////////////
// Constants
//////////////

// ExpandRetries the number of times to retry during Expand(). When a node has fewer child node than its total possible legal moves,
// it is not fully expanded. Playout() will try to add a child node to it with a randomly chosen move. If this move already exists
// in the sibling children, another move will be chosen randomly until either a new move is found, or number of retries exceed this
// value.
// Consideration: This value has to be sufficiently high to all the tree to be properly expanded, but not too high to avoid a
// fully expanded tree, which seems to defeat the purpose of a Monte Carlo tree...
const ExpandRetries = 500

// Exploration the exploration parameter in the Upper Confidence Bound (UCB) algorithm.
// Theoretically it equals to √2, but in practice usually chosen empirically.
const Exploration = 1.4142135623730950488016887242097 // √2
