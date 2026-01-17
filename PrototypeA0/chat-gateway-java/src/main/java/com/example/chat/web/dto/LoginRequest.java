package com.example.chat.web.dto;

public record LoginRequest(
    String email,
    String password
) {
    
}
