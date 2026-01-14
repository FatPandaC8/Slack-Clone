package com.example.chat.web.controller.conversation;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.web.dto.ListPerUserConversationView;

import jakarta.servlet.http.HttpServletRequest;

@RestController
public class UserConversationController {
    private final ChatGrpcClient grpcClient;

    public UserConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @GetMapping("/users/me/conversations")
    public List<ListPerUserConversationView> myConversations(
        @RequestHeader("Authorization") String authHeader,
        HttpServletRequest httpRequest
    ) {
        String token = authHeader.replace("Bearer ", "");
        String userId = (String) httpRequest.getAttribute("userId");
        return grpcClient.listConversations(token, userId);
    }
}
