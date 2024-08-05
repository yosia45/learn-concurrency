## Task 5: Select With Channels

- 5.1. Create two channels: one for sending even numbers and another for sending odd numbers from 1 to 20.
- 5.2. Use a single goroutine that sends numbers to these channels based on whether they're even or odd.
- 5.3. In your main function, use the select statement to read from these channels, printing whether an even or odd number was received. Stop the operation after all numbers have been printed.