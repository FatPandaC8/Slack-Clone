package com.example.chat.infra.grpc;

public record MessageView(
    String messageId,
    String senderId,
    String name,
    String text
) {}