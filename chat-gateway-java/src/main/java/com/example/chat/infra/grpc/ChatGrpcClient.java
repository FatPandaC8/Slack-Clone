package com.example.chat.infra.grpc;

import java.util.List;

import org.springframework.stereotype.Component;

import com.example.chat.grpc.ChatServiceGrpc;
import com.example.chat.grpc.CreateChannelRequest;
import com.example.chat.grpc.CreateChannelResponse;
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

    public void createChannel(String channelId, List<String> members) {
        CreateChannelRequest req = CreateChannelRequest.newBuilder()
                                                        .setChannelId(channelId)
                                                        .addAllMemberIds(members)
                                                        .build();
        
        CreateChannelResponse res = stub.createChannel(req);
        
        if (!res.getOk()) {
            throw new RuntimeException(res.getError());
        }
    }
}
