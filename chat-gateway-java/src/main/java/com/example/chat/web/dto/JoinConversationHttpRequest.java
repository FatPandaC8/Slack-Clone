package com.example.chat.web.dto;

public record JoinConversationHttpRequest(
    String inviteCode, 
    String userId
) {
    
}
