package com.example.chat.infra.grpc;

import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import com.example.chat.grpc.ChatServiceGrpc;
import com.example.chat.grpc.CreateConversationRequest;
import com.example.chat.grpc.CreateConversationResponse;
import com.example.chat.grpc.GetConversationRequest;
import com.example.chat.grpc.GetConversationResponse;
import com.example.chat.grpc.JoinConversationRequest;
import com.example.chat.grpc.JoinConversationResponse;
import com.example.chat.grpc.ListConversationsRequest;
import com.example.chat.grpc.ListConversationsResponse;
import com.example.chat.grpc.ListUserRequest;
import com.example.chat.grpc.ListUserResponse;
import com.example.chat.grpc.LoginUserRequest;
import com.example.chat.grpc.LoginUserResponse;
import com.example.chat.grpc.RegisterUserRequest;
import com.example.chat.grpc.RegisterUserResponse;
import com.example.chat.grpc.SendMessageRequest;
import com.example.chat.grpc.SendMessageResponse;
import com.example.chat.web.dto.CreateConversationResponseView;
import com.example.chat.web.dto.ListPerUserConversationView;
import com.example.chat.web.dto.LoginResponseView;
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

    public CreateConversationResponseView createConversation(String name, String creatorId) {
        CreateConversationRequest req = CreateConversationRequest.newBuilder()
                                                        .setName(name)
                                                        .setCreatorId(creatorId)
                                                        .build();
        
        CreateConversationResponse res = stub.createConversation(req);
        
        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }

        return new CreateConversationResponseView(
            res.getConversationId(),
            res.getName(),
            res.getInviteCode()
        );
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
                                                m.getName(),
                                                m.getText()
                                            ))
                                            .collect(Collectors.toList());

        return new ConversationView(res.getConversationId(), messageViews);
    }

    public List<ListPerUserConversationView> listConversations(String userId) {
        ListConversationsRequest req = ListConversationsRequest.newBuilder()
                                                                .setUserId(userId)
                                                                .build();

        ListConversationsResponse res = stub.listConversations(req);
        
        return res.getConversationsList()
                                .stream()
                                .map(c -> new ListPerUserConversationView(
                                        c.getConversationId(),
                                        c.getName()
                                ))
                                .collect(Collectors.toList());
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

    public void joinConversation(String inviteCode, String userId) {
        JoinConversationRequest req = JoinConversationRequest.newBuilder()
                                                            .setInviteCode(inviteCode)
                                                            .setUserId(userId)
                                                            .build();

        JoinConversationResponse res = stub.joinConversation(req);

        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }
    }

    public UserView register(String name, String email, String password) {
        RegisterUserRequest req = RegisterUserRequest.newBuilder()
            .setName(name)
            .setEmail(email)
            .setPassword(password)
            .build();

        RegisterUserResponse res = stub.registerUser(req);

        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }

        return new UserView(
            res.getUserId(),
            res.getName()
        );
    }

    public LoginResponseView login(String email, String password) {
        LoginUserRequest req = LoginUserRequest.newBuilder()
            .setEmail(email)
            .setPassword(password)
            .build();

        LoginUserResponse res = stub.loginUser(req);

        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }

        return new LoginResponseView(
            res.getUserId(),
            res.getName(),
            res.getToken()
        );
    }

}
