You can view the LeetCode submission [here]("https://leetcode.com/problems/add-two-numbers/solutions/4949824/add-two-number-solution-in-go")

**Intuition**

Think of how you'd manually add two large numbers written on paper. You'd start with the ones column, add the digits, and then carry over any extra to the next column. This problem uses linked lists to represent the numbers in reverse, so we apply the same "add and carry" principle.

**Approach**

Here's a breakdown of how we'll tackle it:

1. **Setting up:** We'll need a way to track the result (a new linked list) and a variable to handle the 'carry-over' part of addition. 

2. **Add and Build:**  We'll walk through both input lists at the same time.
   * For each position, we'll add the digits from the lists (handling cases where one list is shorter).
   * We'll also add in any carry-over from the previous addition.
   * The ones-place digit of this sum  becomes a new node in our result list.
   * Any value above 9 means we have a 'carry' for the next position.

3. **Wrap it up:** After the loop is done, we return the head of our result list.

**Complexity**

* **Time complexity:** Since we have to touch at most each node once in the lists, this is O(max(N1, N2)), with N1 and N2 being the lengths of the lists.
* **Space Complexity:** The worst case is when we have to make a result list as long as the larger input list, also O(max(N1, N2)).

