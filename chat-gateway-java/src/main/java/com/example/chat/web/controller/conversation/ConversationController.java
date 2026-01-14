package com.example.chat.web.controller.conversation;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.infra.grpc.ConversationView;
import com.example.chat.web.dto.CreateConversationHttpRequest;
import com.example.chat.web.dto.CreateConversationResponseView;
import com.example.chat.web.dto.JoinConversationHttpRequest;

import jakarta.servlet.http.HttpServletRequest;

@RestController
@RequestMapping("/conversations")
public class ConversationController {
    private final ChatGrpcClient grpcClient;
    
    public ConversationController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CreateConversationResponseView create(
        @RequestBody CreateConversationHttpRequest req,
        @RequestHeader("Authorization") String authHeader,
        HttpServletRequest httpRequest
    ) {
        String creatorId = (String) httpRequest.getAttribute("userId");
        String token = authHeader.replace("Bearer ", ""); 
        return grpcClient.createConversation(token, req.name(), creatorId);
    }

    @PostMapping("/join")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void join(
        @RequestBody JoinConversationHttpRequest req,
        @RequestHeader("Authorization") String authHeader
    ) {
        String token = authHeader.replace("Bearer ", "");
        grpcClient.joinConversation(token, req.inviteCode(), req.userId());
    }

    @GetMapping("/{conversationId}")
    public ConversationView get(
        @PathVariable String conversationId,
        @RequestHeader("Authorization") String authHeader
    ) {
        String token = authHeader.replace("Bearer ", "");
        return grpcClient.getConversation(token, conversationId);
    }
}
