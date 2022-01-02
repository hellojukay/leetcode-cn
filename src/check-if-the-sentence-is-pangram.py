class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        if len(sentence) < 26:
            return False
        tmp = "abcdefghijklmnopqrstuvwxyz"
        for c in tmp:
            if c not in sentence:
                return False
        return True
