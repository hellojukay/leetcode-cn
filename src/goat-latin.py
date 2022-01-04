class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        def convert(word: str, i: int) -> str:
            suffix = "a" * i
            if re.match("^[aoeiuAEOIU]", word):
                word = word + "ma"
            else:
                ch = word[0]
                word = word[1:] + ch + "ma"
            return word + suffix
        i = 0
        words = sentence.split(" ")
        while i < len(words):
            words[i] = convert(words[i], i+1)
            i = i + 1
        return " ".join(words)
