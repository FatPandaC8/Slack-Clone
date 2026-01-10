package com.example.chat.web.dto;

public record CreateConversationHttpRequest(
        String name,
        String creatorId
) {}
