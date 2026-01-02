package com.example.chat.web;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.application.SendMessageService;

@RestController
@RequestMapping("/chat")
public class ChatController {
    private final SendMessageService sendMessageService;

    public ChatController(SendMessageService sendMessageService) {
        this.sendMessageService = sendMessageService;
    }

    @PostMapping("/send")
    public void send(@RequestBody SendMessageHttpRequest req) {
        sendMessageService.send(req.conversationId(), req.senderId(), req.text());
    }
}
