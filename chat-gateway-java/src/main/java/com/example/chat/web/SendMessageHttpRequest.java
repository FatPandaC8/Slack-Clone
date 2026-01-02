package com.example.chat.web;

public record SendMessageHttpRequest(
    String conversationId,
    String senderId,
    String text
) {
    
}
