package com.example.chat.web;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;

@RestController
public class UserConversationController {
    private final ChatGrpcClient grpcClient;

    public UserConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @GetMapping("/users/{userId}/conversations")
    public List<String> myConversations(@PathVariable String userId) {
        return grpcClient.listConversationsIds(userId);
    }
}
