class MyQueue:
    slice = []

    def __init__(self):
        self.slice = []

    def push(self, x: int) -> None:
        self.slice.append(x)

    def pop(self) -> int:
        n = self.slice[0]
        self.slice = self.slice[1:]
        return n

    def peek(self) -> int:
        return self.slice[0]

    def empty(self) -> bool:
        return len(self.slice) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
