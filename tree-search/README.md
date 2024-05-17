# Take-Home Coding Challenge

The goal is to create a function capable of identifying the shallowest duplicate node within a given tree structure.

## Algorithm description
Breadth-First Search (BFS): This method uses a queue to explore nodes level by level. I start from the root and explore all its children, then move on to the next level.

Tracking Duplicates: I use a map to keep track of each node’s value and its level when it is first encountered. If a node with the same value is encountered again, I compare its level to the previously stored level for that value.

Finding the Shallowest Duplicate: During the BFS traversal, as soon as a duplicate is found, we check if it is shallower (i.e., at a lower level) than any previously found duplicates.

## Time and Space Complexity:
Time Complexity: O(n), where n is the number of nodes in the tree. Each node is processed exactly once.

Space Complexity: O(n) for storing the nodes in the BFS queue and the map for tracking nodes.