package com.example.chat.web.dto;

public record TypingEvent(
    String conversationId, 
    String userId
) {
    
}
