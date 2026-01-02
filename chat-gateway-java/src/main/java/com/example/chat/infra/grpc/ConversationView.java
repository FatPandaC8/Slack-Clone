package com.example.chat.infra.grpc;

import java.util.List;

public record ConversationView(
    String conversationId,
    List<MessageView> messages
) {}
