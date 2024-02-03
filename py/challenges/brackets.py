import unittest
import random


# A string S consisting of N characters is considered to be properly nested 
# if any of the following conditions is true:
# S is empty;
# S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
# S has the form "VW" where V and W are properly nested strings.
# For example, the string "{[()()]}" is properly nested but "([)()]" is not.
# 
# Write a function:
# 
# def solution(S)
# 
# that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.
# 
# For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", 
# the function should return 0, as explained above.
# 
# Write an efficient algorithm for the following assumptions:
# 
# N is an integer within the range [0..200,000];
# string S is made only of the following characters: '(', '{', '[', ']', '}' and/or ')'.

import logging
LOGGER = logging.getLogger()

OPENING_BRACKET_CHARS = ["(", "{", "["]
CLOSING_BRACKET_CHARS = [")", "}", "]"]

def solution(bracketStr: str) -> bool:
    LOGGER.info("Brackets challenge - input " + bracketStr)

    stack = []
    for _, v in enumerate(bracketStr):
        if v in OPENING_BRACKET_CHARS:
            stack.append(v)
            LOGGER.info("append v " + str(v))
        else:
            if len(stack) == 0:
                LOGGER.info("stack is empty")
                return False
            prev_v = stack.pop()
            if (v == CLOSING_BRACKET_CHARS[0] and prev_v != OPENING_BRACKET_CHARS[0]) or \
                (v == CLOSING_BRACKET_CHARS[1] and prev_v != OPENING_BRACKET_CHARS[1]) or \
                (v == CLOSING_BRACKET_CHARS[2] and prev_v != OPENING_BRACKET_CHARS[2]):
                LOGGER.info("v " + str(v) + " not == prev bracket " + str(prev_v))                        
                return False
    if len(stack) != 0:
        LOGGER.info("trailling opening brackes - stack not empty")
        return False

    return True



class BracketsTest(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution("{[()()]}"), 1)
        self.assertEqual(solution("([)()]"), 0)

    def test_min(self):
        self.assertEqual(solution("("), 0)
        self.assertEqual(solution(")"), 0)

    def test_max(self):
        max = 3000
        brackets = '(' * max
        closing = ')' * max
        self.assertEqual(solution(str(brackets + closing)), 1)
