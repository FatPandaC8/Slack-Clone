package com.example.chat.web.dto;

public record CreateConversationResponseView(
    String id, 
    String name,
    String inviteCode
) {
    
}
