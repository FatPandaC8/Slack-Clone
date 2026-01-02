package com.example.chat.web;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.infra.grpc.ConversationView;

@RestController
@RequestMapping("/channels")
public class ConversationController {
    private final ChatGrpcClient grpcClient;
    
    public ConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @PostMapping
    public void create(@RequestBody CreateConversationHttpRequest req) {
        grpcClient.createConversation(req.channelId(), req.memberIds());
    }

    @GetMapping("/{id}")
    public ConversationView get(@PathVariable String id) {
        return grpcClient.getConversation(id);
    }
}
