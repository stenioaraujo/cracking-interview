class Node:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


def in_order(root, result):
    if root:
        in_order(root.left, result)
        result.append(root.value)
        in_order(root.right, result)


def pre_order(root, result):
    if root:
        result.append(root.value)
        pre_order(root.left, result)
        pre_order(root.right, result)


def post_order(root, result):
    if root:
        post_order(root.left, result)
        post_order(root.right, result)
        result.append(root.value)


# Filling tree
nodes = [Node(i) for i in range(0, 6)]
nodes[1].left = nodes[0]
nodes[1].right = nodes[2]
nodes[5].left = nodes[4]
nodes[3].left = nodes[1]
nodes[3].right = nodes[5]

root = nodes[3]


out = []
in_order(root, out)
print("In Order:   ", out)

out = []
pre_order(root, out)
print("Pre Order:  ", out)

out = []
post_order(root, out)
print("Post Order: ", out)
