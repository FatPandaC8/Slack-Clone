grpcurl -plaintext -d '{"userID": "1", "name":"General"}' localhost:50051 chat.ChatService/CreateRoom
grpcurl -plaintext -d '{"userID": "2", "invite_code": "7092787b"}' localhost:50051 chat.ChatService/JoinRoom

curl -X POST localhost:8888/login -H "Content-Type: application/json" -d '{"email":"a@a.com","password":"123"}'
grpcurl -plaintext -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXNlcnZpY2UiLCJzdWIiOiI5M2NiMDExZC01OTcxLTRkNDktOTI4MS1mM2Y5MGNjOGEzODMiLCJleHAiOjE3Njg4MDAwODMsImlhdCI6MTc2ODc5OTE4M30.28qmPbgcPN2T_rS-oFABE3wmW7MRNI7aEp2-9usUXh4" -d '{"roomID":"room1","content":"hello"}' localhost:50051 chat.ChatService/SendMessage