package com.example.chat.infra.grpc;

import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

import org.springframework.stereotype.Component;

@Component
public class Typing {
    private final Map<String, Set<String>> typingMap = new ConcurrentHashMap<>();

    public void startTyping(String conversationId, String userId) {
        typingMap.computeIfAbsent(conversationId, k -> ConcurrentHashMap.newKeySet()).add(userId);
    }

    public void stopTyping(String conversationId, String userId) {
        Set<String> users = typingMap.get(conversationId);
        if (users == null) return;

        users.remove(userId);
        if (users.isEmpty()) {
            typingMap.remove(conversationId);
        }
    }

    public Set<String> getTypingUsers(String conversationId) {
        return typingMap.getOrDefault(conversationId, Set.of());
    }
}
