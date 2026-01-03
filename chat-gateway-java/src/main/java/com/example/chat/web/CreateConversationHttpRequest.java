package com.example.chat.web;

import java.util.List;

public record CreateConversationHttpRequest(
    String conversationId,
    List<String> memberIds
) {
    
}
