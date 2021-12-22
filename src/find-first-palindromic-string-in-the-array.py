#!/usr/bin/env python3
# https://leetcode-cn.com/problems/find-first-palindromic-string-in-the-array/
class Solution:
    def isPalindrome(self, str) -> bool:
        if not len(str):
            return False
        return str == str[::-1]

    def firstPalindrome(self, words: List[str]) -> str:
        if not len(words):
            return ""
        if self.isPalindrome(words[0]):
            return words[0]
        return self.firstPalindrome(words[1:])
