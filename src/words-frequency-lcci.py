class WordsFrequency:

    def __init__(self, book: List[str]):
        self.dic = {}
        for s in book:
            count = self.dic.get(s,0)
            count = count + 1
            self.dic[s] = count

    def get(self, word: str) -> int:
        return self.dic.get(word,0)


# Your WordsFrequency object will be instantiated and called as such:
# obj = WordsFrequency(book)
# param_1 = obj.get(word)
