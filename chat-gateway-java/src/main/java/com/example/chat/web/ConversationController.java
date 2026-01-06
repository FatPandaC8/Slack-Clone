package com.example.chat.web;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.infra.grpc.ConversationView;
import com.example.chat.web.dto.CreateConversationHttpRequest;

@RestController
@RequestMapping("/conversations")
public class ConversationController {
    private final ChatGrpcClient grpcClient;
    
    public ConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public ConversationView create(@RequestBody CreateConversationHttpRequest req) {
        grpcClient.createConversation(req.conversationId(), req.memberIds());
        return grpcClient.getConversation(req.conversationId());
    }

    @GetMapping("/{conversationId}")
    public ConversationView get(@PathVariable String conversationId) {
        return grpcClient.getConversation(conversationId);
    }
}
