import copy


class MinHeap:
    def __init__(self, size):
        self._heap = [0] * (size + 1)
        self.end = 1

    @property
    def heap(self):
        return copy.copy(self._heap[1:])

    def insert(self, value):
        self._heap[self.end] = value

        node = self.end
        parent = node // 2
        while parent > 0 and self._heap[parent] > self._heap[node]:
            self._swap(parent, node)
            node = parent
            parent = node // 2

        self.end += 1

    def pop_min(self):
        if self.empty():
            return None

        self.end -= 1

        minimum = self._heap[1]
        self._heap[1] = self._heap[self.end]

        node = 1
        left = node * 2
        while left < self.end:
            right = left + 1
            if right < self.end and self._heap[right] < self._heap[left]:
                mi = right
            else:
                mi = left

            if self._heap[node] > self._heap[mi]:
                self._swap(node, mi)
                node = mi
            else:
                node = self.end
            left = node * 2

        return minimum

    def _swap(self, i, j):
        l = self._heap
        l[i], l[j] = l[j], l[i]

    def empty(self):
        return self.end == 1


min_size = 10
min_heap = MinHeap(min_size)
for i in list(range(1, min_size+1))[::-1]:
    min_heap.insert(i)

print(min_heap.heap)

while not min_heap.empty():
    print(min_heap.pop_min())
