## Solution notes

### Task 01 – Run‑Length Encoder
- Language: Python
- Approach: Using dynamic sliding window algorithm, first, we initialize two pointers: left and right pointers, starting at the beginning of the string. If we encounter the duplicate, expand the window by moving the right pointer. When encountering a new character, shrink the window by moving the left pointer to current position of the right pointer pointing to a new unique character. After each run of duplicate characters, append the character along with the length in <char><count> format.
- Why: This method uses two pointers approach minimizies time complexity to O(n) and space complexity to O(n) eliminating the use of nesting loops in brute-force approaches which search through all possible substrings, resulting in worse time complexity.
- What to refine with more time: Although I personally found that using a while loop offers more intuitive control over the pointers, a more concise code using a for loop would improve readability.
- Time spent: ~10 min
- AI tools used: Claude (for double-checking)

### Task 02 – Fix‑the‑Bug
- Language: Go
- Approach: A Mutex (sync.Mutex) synchronizes access to the shared variable by locking it via the Lock() method at the start of the critical section and unlocking it with Unlock() at the end. This prevents multiple goroutines from concurrently accessing the critical section, which avoids race conditions.
- Why: Using a mutex ensures that only one goroutine can enter the critical section at a time. In this case, ID generation is thread‑safe.
- What to refine with more time: Consider exploring other higher-level synchronization alternatives for higher performance
- Time spent: ~5 min
- AI tools used: Claude (for double-checking)