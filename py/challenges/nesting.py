import unittest


# A string S consisting of N characters is called properly nested if:

# S is empty;
# S has the form "(U)" where U is a properly nested string;
# S has the form "VW" where V and W are properly nested strings.
# For example, string "(()(())())" is properly nested but string "())" isn't.

# Write a function:

# def solution(S)

# that, given a string S consisting of N characters, returns 1 if string S is 
# properly nested and 0 otherwise.

# For example, given S = "(()(())())", the function should return 1 and given S = "())", 
# the function should return 0, as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [0..1,000,000];
# string S is made only of the characters '(' and/or ')'.


def solution(S):
    strArr = S        
    if len(strArr)%2 != 0:
        return 0

    balanced = 0
    for c in strArr[0:]:
        if c == "(":
            balanced+=1
        else:
            if balanced == 0:
                return 0
            balanced-=1

    if balanced > 0:
        return 0

    return 1 



class TestNested(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution("(()(())())"), 1)
        self.assertEqual(solution("())"), 0)

    def test_min(self):
        self.assertEqual(solution(""), 1)
        self.assertEqual(solution(")"), 0)
        self.assertEqual(solution("("), 0)

    def test_max(self):
        max = 100000
        strArr = ['(' for n in range(0, max)]
        self.assertEqual(solution(strArr), 0)
        strArr_1 = ['(' for n in range(0, int(max/2))]
        strArr_1 += [')' for n in range(int(max/2), max)]
        self.assertEqual(solution(strArr_1), 1)

