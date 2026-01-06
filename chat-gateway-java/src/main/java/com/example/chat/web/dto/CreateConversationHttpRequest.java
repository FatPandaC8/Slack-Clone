package com.example.chat.web.dto;

import java.util.List;

public record CreateConversationHttpRequest(
    String conversationId,
    List<String> memberIds
) {
    
}
