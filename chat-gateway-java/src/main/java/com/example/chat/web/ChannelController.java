package com.example.chat.web;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;

@RestController
@RequestMapping("/channels")
public class ChannelController {
    private final ChatGrpcClient grpcClient;
    
    public ChannelController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @PostMapping
    public void create(@RequestBody CreateChannelHttpRequest req) {
        grpcClient.createChannel(req.channelId(), req.memberIds());
    }
}
