# hierarchy-datastore
This is a test task

#### Brief description:

Service maintains a hierarchy of nodes in memory. The API
supports adding and deleting nodes, moving them to a new position in
the tree and searching.

Each node is identified by a name and an
ID (both strings). The name of a node must be unique among its
siblings (i.e., the children of the node's parent); the ID must be
unique among all nodes in the tree.

The interface uses JSON-encoded messages. Program receives
requests on standard input and write responses to standard output.
