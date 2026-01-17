package com.example.chat.infra.grpc;

import java.util.Set;

public record TypingUsersView(
    String conversationId, 
    Set<String> usersTyping
) {
    
}
