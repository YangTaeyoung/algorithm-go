def solution(s):
    stack = []
    for ch in s:
        if ch == "(":
            stack.append(ch)
        else:
            if len(stack) == 0:
                return False
            popped = stack.pop()


    return len(stack) == 0

