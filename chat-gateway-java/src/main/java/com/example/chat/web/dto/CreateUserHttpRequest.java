package com.example.chat.web.dto;

public record CreateUserHttpRequest(
    String name, 
    String email,
    String password
) {
    
}
