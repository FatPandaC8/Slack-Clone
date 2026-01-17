package com.example.chat.application;

import java.util.UUID;

import org.springframework.stereotype.Service;

import com.example.chat.infra.grpc.ChatGrpcClient;

@Service
public class SendMessageService {
    private final ChatGrpcClient grpcClient;

    public SendMessageService(ChatGrpcClient client) {
        this.grpcClient = client;
    }

    public void send(
        String token,
        String conversationId, 
        String senderId,
        String text
    ) {
        String messageId = UUID.randomUUID().toString();

        grpcClient.sendMessage(token, messageId, conversationId, senderId, text);
    }
}
