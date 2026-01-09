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

(05/01/2026)
[UI - FIXED]            Invisible messages <= Lost import
[Logic - FIXED]         TypeError: conversationId is read-only => fixed by using a helper setter
[UI - FIXED]            The new messages are not aligned but it's fixed after a web restart
[Logic - FIXED]         Send a request to java that a person is typing, then after a 2 second interval, send a not typing request
                        The one keeping track of all the users typing is in the java
                        Done: use a single flag to reduce 95% of sending request.

(08/01/2026)
[Logic - FIXED]         At runtime, loadConversation create new msg element, but real time, it does not have any css applied to it.
                        Websocket was the one changing the content without applying the css => reuse the same renderMessage func from 
                        loadConversation.js

[UI - ONGOING]          Right now, the sidebar is jamming with user creation, conversation creation and conversation list => reduce UX 
                        flow => MAKE IT SEPARATE LIKE SLACK DOES.

                        chatapp - superchatapp