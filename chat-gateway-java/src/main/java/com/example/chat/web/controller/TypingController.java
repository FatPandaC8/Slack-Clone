package com.example.chat.web.controller;

import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

import com.example.chat.infra.grpc.Typing;
import com.example.chat.infra.grpc.TypingUsersView;
import com.example.chat.web.dto.TypingEvent;

@Controller
public class TypingController {
    private final Typing tracker;
    private final SimpMessagingTemplate ws;

    public TypingController(Typing tracker, SimpMessagingTemplate ws) {
        this.tracker = tracker;
        this.ws = ws;
    }

    @MessageMapping("/typing.start")
    public void startTyping(TypingEvent event) {
        tracker.startTyping(event.conversationId(), event.userId());

        ws.convertAndSend(
            "/topic/conversations/" + event.conversationId() + "/typing",
            new TypingUsersView(
                event.conversationId(),
                tracker.getTypingUsers(event.conversationId())
            )
        );
    }

    @MessageMapping("/typing.stop")
    public void stopTyping(TypingEvent event) {
        tracker.stopTyping(event.conversationId(), event.userId());

        ws.convertAndSend(
            "/topic/conversations/" + event.conversationId() + "/typing",
            new TypingUsersView(
                event.conversationId(),
                tracker.getTypingUsers(event.conversationId())
            )
        );
    }
}
