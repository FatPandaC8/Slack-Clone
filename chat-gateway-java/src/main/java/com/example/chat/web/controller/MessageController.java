package com.example.chat.web.controller;

import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.application.SendMessageService;
import com.example.chat.infra.grpc.MessageView;
import com.example.chat.web.dto.SendMessageHttpRequest;

import jakarta.servlet.http.HttpServletRequest;

@RestController
@RequestMapping("/conversations/{conversationId}/messages")
public class MessageController {
    private final SendMessageService sendMessageService;
    private final SimpMessagingTemplate ws;

    public MessageController(
        SendMessageService sendMessageService,
        SimpMessagingTemplate ws
    ) {
        this.sendMessageService = sendMessageService;
        this.ws = ws;
    }

    @PostMapping
    public void send(
        @PathVariable String conversationId,
        @RequestBody SendMessageHttpRequest req,
        @RequestHeader("Authorization") String authHeader,
        HttpServletRequest httpRequest
    ) {
        String token = authHeader.replace("Bearer ", "");
        String senderId = (String) httpRequest.getAttribute("userId");
        sendMessageService.send(
            token,
            conversationId,
            senderId,
            req.text()
        );

        ws.convertAndSend(
            "/topic/conversations/" + conversationId,
            new MessageView(
                req.messageId(),
                senderId,
                req.text()
            )
        );
    }
}
