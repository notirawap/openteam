def encode(s: str) -> str:
    length = len(s)
    # Edge case: empty input string
    if length == 0:
        return ""
    encoded = [] # Store each run-length encoding
    left, right = 0, 1 # Initialize two pointers
    while left < length:
        # Expand the window
        while right < length and s[left] == s[right]:
            right += 1
        # Append character and count
        count = right - left
        encoded.append(s[left] + str(count))
        # Move to the next run of duplicate characters
        left = right
        right = left + 1
    return "".join(encoded)