class MyHashSet:

    nums = []
    def __init__(self):
        self.nums = [None] * 10000
        i = 0
        while i < len(self.nums):
            self.nums[i] = [-1] * 1000
            i = i + 1


    def add(self, key: int) -> None:
        n = int(key / 1000)
        c = key % 1000
        self.nums[n][c] = 1



    def remove(self, key: int) -> None:
        n = int(key / 1000)
        c = key % 1000
        self.nums[n][c] = -1


    def contains(self, key: int) -> bool:
        n = int(key / 1000)
        c = key % 1000
        if self.nums[n][c] > 0:
            return True
        return False



# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)

