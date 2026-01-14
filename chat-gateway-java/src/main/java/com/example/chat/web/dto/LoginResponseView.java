package com.example.chat.web.dto;

public record LoginResponseView(
    String userId, 
    String name, 
    String token
) {
    
}
