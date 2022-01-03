class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        result = []
        for word in sentence.split(" "):
            for r in dictionary:
                if word.startswith(r):
                    word = r
            result.append(word)
        return " ".join(result)
