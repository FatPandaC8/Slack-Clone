package com.example.chat.web.dto;

public record SendMessageHttpRequest(
    String messageId,
    String text
) {}
