package com.example.chat.web.dto;

public record SendMessageHttpRequest(
    String conversationId,
    String senderId,
    String name,
    String text
) {
    
}
