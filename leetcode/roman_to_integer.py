# Problem: https://leetcode.com/problems/roman-to-integer/

class RomanToInt():
    def __init__(self):
        self.roman_numerals = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }

    def initial(self, s: str) -> int:
        # Initial solution: runtime 122 ms, memory usage 13.9 MB
        res = 0
        for i in range(len(s)):
            if i + 1 < len(s) and self.roman_numerals[s[i]] < self.roman_numerals[s[i + 1]]:
                res -= self.roman_numerals[s[i]]
            else:
                res += self.roman_numerals[s[i]]
        return res

    def improved(self, s: str) -> int:
        # Improved solution: runtime 74 ms, memory usage 13.8 MB
        res = 0
        res += self.roman_numerals[s[-1]]
        for i in reversed(range(0, len(s)-1)):
            if self.roman_numerals[s[i]] >= self.roman_numerals[s[i+1]]:
                res += self.roman_numerals[s[i]]
            else:
                res -= self.roman_numerals[s[i]]
        return res
