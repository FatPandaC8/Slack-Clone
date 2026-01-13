package com.example.chat.web.controller.conversation;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.web.dto.ListPerUserConversationView;

@RestController
public class UserConversationController {
    private final ChatGrpcClient grpcClient;

    public UserConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @GetMapping("/users/{userId}/conversations")
    public List<ListPerUserConversationView> myConversations(
        @PathVariable String userId,
        @RequestHeader("Authorization") String authHeader
    ) {
        String token = authHeader.replace("Bearer", "");
        return grpcClient.listConversations(token, userId);
    }
}
