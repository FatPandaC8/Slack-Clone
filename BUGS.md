# Tracking
Record of all the obstacles encountered during dev time

(03/01/2026)
Q: How can multiple users find the conversations they're invited ?
For simple fix, get the userId, then find all the conversation ID including that user ID
But can turn into a graph problem OR use Union Find for faster grouping ? (Ideas)
But now it can talk to each other 

B: Only show the text when click the conversationId button => Need to use WebSocket

(04/01/2026)
Tips: Always check the inspect for network => check status code -> can resend if needed + check console
      Use Ctrl + Shift + R for top to bottom reload
A: Today I separate the js files into their respective folders => need to import with the extension of .js & export needed functions
   Or else, the file will return 404
BUGS: cannot see the new messages 
