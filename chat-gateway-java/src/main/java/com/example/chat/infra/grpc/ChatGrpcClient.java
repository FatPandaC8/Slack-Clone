package com.example.chat.infra.grpc;

import java.util.List;
import java.util.stream.Collectors;

import org.springframework.stereotype.Component;

import com.example.chat.grpc.ChatServiceGrpc;
import com.example.chat.grpc.CreateConversationRequest;
import com.example.chat.grpc.CreateConversationResponse;
import com.example.chat.grpc.GetConversationRequest;
import com.example.chat.grpc.GetConversationResponse;
import com.example.chat.grpc.SendMessageRequest;
import com.example.chat.grpc.SendMessageResponse;

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

    public void createConversation(String channelId, List<String> members) {
        CreateConversationRequest req = CreateConversationRequest.newBuilder()
                                                        .setChannelId(channelId)
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
}
