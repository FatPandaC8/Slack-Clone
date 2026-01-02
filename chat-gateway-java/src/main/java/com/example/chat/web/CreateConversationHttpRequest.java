package com.example.chat.web;

import java.util.List;

public record CreateConversationHttpRequest(
    String channelId,
    List<String> memberIds
) {
    
}
