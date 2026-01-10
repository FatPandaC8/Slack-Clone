package com.example.chat.infra.grpc;

import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import com.example.chat.grpc.ChatServiceGrpc;
import com.example.chat.grpc.CreateConversationRequest;
import com.example.chat.grpc.CreateConversationResponse;
import com.example.chat.grpc.CreateUserRequest;
import com.example.chat.grpc.CreateUserResponse;
import com.example.chat.grpc.GetConversationRequest;
import com.example.chat.grpc.GetConversationResponse;
import com.example.chat.grpc.ListConversationsRequest;
import com.example.chat.grpc.ListConversationsResponse;
import com.example.chat.grpc.ListUserRequest;
import com.example.chat.grpc.ListUserResponse;
import com.example.chat.grpc.SendMessageRequest;
import com.example.chat.grpc.SendMessageResponse;
import com.example.chat.web.dto.CreateUserResponseView;
import com.example.chat.web.dto.UserView;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

@Component
public class ChatGrpcClient {
    private final ChatServiceGrpc.ChatServiceBlockingStub stub;

    public ChatGrpcClient() {
        ManagedChannel channel = ManagedChannelBuilder
                                        .forAddress("localhost", 50051)
                                        .usePlaintext()
                                        .build();
        this.stub = ChatServiceGrpc.newBlockingStub(channel);
    }

    public void sendMessage(
        String messageId,
        String conversationId,
        String senderId,
        String text
    ) {
        SendMessageRequest request = SendMessageRequest.newBuilder()
                                                    .setMessageId(messageId)
                                                    .setConversationId(conversationId)
                                                    .setSenderId(senderId)
                                                    .setText(text)
                                                    .build();
        
        SendMessageResponse response = stub.sendMessage(request);
        
        if (!response.getOk()) {
            throw new RuntimeException(response.getError());
        }
    }

    public void createConversation(String conversationId, List<String> members) {
        CreateConversationRequest req = CreateConversationRequest.newBuilder()
                                                        .setConversationId(conversationId)
                                                        .addAllMemberIds(members)
                                                        .build();
        
        CreateConversationResponse res = stub.createConversation(req);
        
        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }
    }

    public ConversationView getConversation(String conversationId) {
        GetConversationRequest req = GetConversationRequest.newBuilder()
                                                            .setConversationId(conversationId)
                                                            .build();

        GetConversationResponse res = stub.getConversation(req);
        
        List<MessageView> messageViews = res.getMessagesList()
                                            .stream()
                                            .map(
                                                m -> new MessageView(
                                                m.getMessageId(),
                                                m.getSenderId(),
                                                m.getText()
                                            ))
                                            .collect(Collectors.toList());

        return new ConversationView(res.getConversationId(), messageViews);
    }

    public List<String> listConversationsIds(String userId) {
        ListConversationsRequest req = ListConversationsRequest.newBuilder()
                                                                .setUserId(userId)
                                                                .build();

        ListConversationsResponse res = stub.listConversations(req);
        
        return res.getConversationIdList();
    }

    public CreateUserResponseView createUser(String name, String email, String password) {
        CreateUserRequest req = CreateUserRequest.newBuilder()
                                                .setName(name)
                                                .setEmail(email)
                                                .setPassword(password)
                                                .build();

        CreateUserResponse res = stub.createUser(req);

        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }

        return new CreateUserResponseView(
            res.getUserId(),
            res.getName()
        );
    }

    public List<UserView> listUsers(String conversationId) {
        ListUserRequest req = ListUserRequest.newBuilder()
                                            .setConversationId(conversationId)
                                            .build();

        ListUserResponse res = stub.listUsers(req);
        return res.getUsersList()
                        .stream()
                        .map(u -> new UserView(u.getUserId(), u.getName()))
                        .collect(Collectors.toList());
    }
}
