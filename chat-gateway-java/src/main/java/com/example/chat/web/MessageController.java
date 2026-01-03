package com.example.chat.web;

import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.application.SendMessageService;

@RestController
@RequestMapping("/conversations/{conversationId}/messages")
public class MessageController {
    private final SendMessageService sendMessageService;

    public MessageController(SendMessageService sendMessageService) {
        this.sendMessageService = sendMessageService;
    }

    @PostMapping
    public void send(
        @PathVariable String conversationId,
        @RequestBody SendMessageHttpRequest req
    ) {
        sendMessageService.send(
            conversationId,
            req.senderId(),
            req.text()
        );
    }
}
